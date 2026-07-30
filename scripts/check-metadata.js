#!/usr/bin/env node
// Validates the EDC metadata header (see METADATA_HEADER.md) on every
// Markdown file in a project: header presence, required fields, a
// canonical `status` value, unique `id`s, a real `last_updated` date,
// and no unresolved template placeholders (`[OWNER]`, `[DATE]`, etc.)
// left behind in a document that isn't itself a template.
//
// Usage: node check-metadata.js <project-root>

const fs = require('fs');
const path = require('path');
const { walkMarkdown } = require('./lib/edc-fs');
const { parseFrontmatter } = require('./lib/frontmatter');

const CANONICAL_STATUSES = [
  'Draft', 'Review', 'Accepted', 'Active', 'Implemented',
  'Verified', 'Superseded', 'Archived', 'Rejected',
];
const REQUIRED_FIELDS = [
  'id', 'title', 'version', 'status', 'owner',
  'audience', 'purpose', 'related', 'last_updated',
];
const ISO_DATE_RE = /^\d{4}-\d{2}-\d{2}$/;
const PLACEHOLDER_RE = /^\[.*\]$/;

// Files that ship with intentionally-unfilled placeholder values — they
// exist to be copied and filled in per project (ADR/RFC/QUEST templates,
// the how-to-use guide) or per epoch transition (the generic, not-yet-
// copied EPOCH_READINESS.md — see documentation-structure-and-governance.md
// section 8's "copy per transition" rule).
function isTemplateFile(relPath) {
  const base = path.basename(relPath);
  return (
    base.endsWith('-TEMPLATE.md') ||
    base === '00_HOW_TO_USE_THIS_TEMPLATE.md' ||
    base === 'EPOCH_READINESS.md'
  );
}

function isValidDate(s) {
  if (!ISO_DATE_RE.test(s)) return false;
  const d = new Date(`${s}T00:00:00Z`);
  return !Number.isNaN(d.getTime()) && d.toISOString().slice(0, 10) === s;
}

function run(root) {
  const projectName = path.basename(path.resolve(root));
  const wholeProjectIsTemplate = projectName === '_template';
  const files = walkMarkdown(root);
  const issues = [];
  const idsSeen = new Map();

  for (const file of files) {
    const rel = path.relative(root, file);
    const content = fs.readFileSync(file, 'utf8');
    const parsed = parseFrontmatter(content);

    if (!parsed) {
      issues.push(`${rel}: missing metadata header (no --- frontmatter block)`);
      continue;
    }

    const { fields } = parsed;
    const exempt = wholeProjectIsTemplate || isTemplateFile(rel);

    for (const key of REQUIRED_FIELDS) {
      if (!(key in fields)) {
        issues.push(`${rel}: missing required field '${key}'`);
      }
    }

    if (typeof fields.id === 'string' && fields.id.trim()) {
      const list = idsSeen.get(fields.id) || [];
      list.push(rel);
      idsSeen.set(fields.id, list);
    }

    if (fields.status !== undefined && !CANONICAL_STATUSES.includes(fields.status)) {
      issues.push(
        `${rel}: status '${fields.status}' is not one of ${CANONICAL_STATUSES.join('/')}`
      );
    }

    if (!exempt && fields.last_updated !== undefined && !isValidDate(fields.last_updated)) {
      issues.push(`${rel}: last_updated '${fields.last_updated}' is not a valid YYYY-MM-DD date`);
    }

    if (!exempt) {
      for (const key of ['id', 'title', 'owner']) {
        const v = fields[key];
        if (typeof v === 'string' && PLACEHOLDER_RE.test(v.trim())) {
          issues.push(`${rel}: field '${key}' still holds an unresolved placeholder value (${v})`);
        }
      }
    }
  }

  for (const [id, paths] of idsSeen) {
    if (paths.length > 1) {
      issues.push(`duplicate id '${id}' used in: ${paths.join(', ')}`);
    }
  }

  const ok = issues.length === 0;
  const summary = `${files.length} markdown files checked, ${issues.length} metadata issues.`;
  return { ok, summary, details: issues };
}

module.exports = { run };

if (require.main === module) {
  const root = process.argv[2];
  if (!root) {
    console.error('usage: node check-metadata.js <project-root>');
    process.exit(1);
  }
  const { ok, summary, details } = run(root);
  console.log(summary);
  details.forEach((d) => console.log(`  ${d}`));
  process.exit(ok ? 0 : 1);
}
