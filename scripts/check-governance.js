#!/usr/bin/env node
// Checks structural rules from documentation-structure-and-governance.md
// that aren't about any single document's own content: folder placement,
// the one-README rule, the MILESTONES/HOST-EVENTS anchor-before-heading
// pattern, and the small set of document types with one fixed correct
// `status` value.
//
// Usage: node check-governance.js <project-root>

const fs = require('fs');
const path = require('path');
const { walkMarkdown } = require('./lib/edc-fs');
const { parseFrontmatter } = require('./lib/frontmatter');

// relative path (from project root) -> the one status value section 5 of
// the governance doc allows for it. Only doc types with a single fixed
// value belong here — e.g. ROADMAP.md is deliberately left out, since its
// own recommended flow is "Draft -> Active or Accepted."
const FIXED_STATUS = {
  'PROJECT_MEMORY.md': 'Active',
  'docs/00-project/MILESTONES.md': 'Active',
  'docs/14-infrastructure/HOST-EVENTS.md': 'Active',
  'docs/10-decisions/ADR-TEMPLATE.md': 'Accepted',
  'docs/11-rfc/RFC-TEMPLATE.md': 'Accepted',
};

const ANCHORED_LOG_FILES = {
  'docs/00-project/MILESTONES.md': { headingRe: /^##\s+(M-\d{4})\b/gm, idPrefix: 'm-' },
  'docs/14-infrastructure/HOST-EVENTS.md': { headingRe: /^##\s+(INFRA-\d{4})\b/gm, idPrefix: 'infra-' },
};

function checkOneReadmeRule(root, issues) {
  const rootReadme = path.join(root, 'README.md');
  const readmes = walkMarkdown(root).filter((f) => path.basename(f) === 'README.md');
  for (const f of readmes) {
    if (f !== rootReadme) {
      issues.push(`${path.relative(root, f)}: stray README.md — only the project-root README.md is allowed`);
    }
  }
}

function checkTemplateOnlyFolders(root, issues) {
  const rules = [
    { dir: 'docs/10-decisions', allow: 'ADR-TEMPLATE.md' },
    { dir: 'docs/11-rfc', allow: 'RFC-TEMPLATE.md' },
  ];
  for (const { dir, allow } of rules) {
    const full = path.join(root, dir);
    if (!fs.existsSync(full)) continue;
    for (const entry of fs.readdirSync(full, { withFileTypes: true })) {
      if (entry.isFile() && entry.name !== allow) {
        issues.push(`${dir}/${entry.name}: only ${allow} belongs in ${dir}/ — real docs live in docs/${dir.endsWith('rfc') ? 'rfc' : 'adr'}/`);
      }
    }
  }
}

function checkAdrRfcNaming(root, issues) {
  const patterns = [
    { re: /^ADR-\d+.*\.md$/i, home: 'docs/adr' },
    { re: /^RFC-\d+.*\.md$/i, home: 'docs/rfc' },
  ];
  const files = walkMarkdown(root);
  for (const file of files) {
    const rel = path.relative(root, file);
    const base = path.basename(file);
    const dir = path.dirname(rel).replace(/\\/g, '/');
    for (const { re, home } of patterns) {
      if (re.test(base) && !base.endsWith('-TEMPLATE.md') && dir !== home) {
        issues.push(`${rel}: named like a real ${home === 'docs/adr' ? 'ADR' : 'RFC'} but lives outside ${home}/`);
      }
    }
  }
}

function checkNoStrayFiles(root, issues) {
  const banned = ['STATUS.md', 'HANDOFF.md'];
  const files = walkMarkdown(root, {});
  for (const file of files) {
    if (banned.includes(path.basename(file))) {
      issues.push(`${path.relative(root, file)}: retired file type still present (see documentation-structure-and-governance.md section 7.1)`);
    }
  }
  if (fs.existsSync(path.join(root, 'ai'))) {
    issues.push(`ai/: retired folder still present — condensed AI docs were folded into docs/05-development/AI_GUIDANCE.md`);
  }
}

function checkAnchoredLogs(root, issues) {
  for (const [rel, { headingRe, idPrefix }] of Object.entries(ANCHORED_LOG_FILES)) {
    const full = path.join(root, rel);
    if (!fs.existsSync(full)) continue;
    const content = fs.readFileSync(full, 'utf8');
    const lines = content.split(/\r?\n/);
    for (let i = 0; i < lines.length; i++) {
      const m = lines[i].match(/^##\s+(M-\d{4}|INFRA-\d{4})\b/);
      if (!m) continue;
      const expectedAnchor = `<a id="${idPrefix}${m[1].split('-')[1]}"></a>`;
      const prevLine = i > 0 ? lines[i - 1].trim() : '';
      if (prevLine !== expectedAnchor) {
        issues.push(`${rel}: heading '${m[1]}' missing '${expectedAnchor}' on the line immediately above it`);
      }
    }
  }
}

function checkFixedStatuses(root, issues) {
  for (const [rel, expected] of Object.entries(FIXED_STATUS)) {
    const full = path.join(root, rel);
    if (!fs.existsSync(full)) continue;
    const parsed = parseFrontmatter(fs.readFileSync(full, 'utf8'));
    if (!parsed) continue;
    if (parsed.fields.status !== expected) {
      issues.push(`${rel}: status is '${parsed.fields.status}', expected '${expected}' per documentation-structure-and-governance.md section 5`);
    }
  }
}

function run(root) {
  const issues = [];
  checkOneReadmeRule(root, issues);
  checkTemplateOnlyFolders(root, issues);
  checkAdrRfcNaming(root, issues);
  checkNoStrayFiles(root, issues);
  checkAnchoredLogs(root, issues);
  checkFixedStatuses(root, issues);

  const ok = issues.length === 0;
  const summary = `${issues.length} governance-structure issues.`;
  return { ok, summary, details: issues };
}

module.exports = { run };

if (require.main === module) {
  const root = process.argv[2];
  if (!root) {
    console.error('usage: node check-governance.js <project-root>');
    process.exit(1);
  }
  const { ok, summary, details } = run(root);
  console.log(summary);
  details.forEach((d) => console.log(`  ${d}`));
  process.exit(ok ? 0 : 1);
}
