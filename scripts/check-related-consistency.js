#!/usr/bin/env node
// Checks that a document's `related:` list actually matches what's linked
// in its body — see documentation-structure-and-governance.md section 1.6,
// "the metadata chain rabbit-holes, the body does not repeat it."
//
// Section 1.6 is asymmetric, and this script follows it exactly:
//
//   ERROR    a `related:` entry whose target doesn't exist. A dead
//            reference is never defensible — the chain just stops.
//   ERROR    a `related:` entry that exists but isn't linked anywhere in
//            the body ("every file listed in `related:` should actually be
//            linked somewhere in the body"). Confirmed as a hard rule
//            2026-07-29 rather than softened to advice.
//   ADVISORY a doc linked in the body that isn't in `related:`. Section 1.6
//            grants this one an explicit exception ("that isn't a one-off
//            example"), which no script can judge, so it's reported and
//            left to a human.
//
// `related:` paths are relative to the file's own directory, same as a
// normal Markdown link (see METADATA_HEADER.md) — so comparison resolves
// both sides to an absolute path rather than comparing raw strings.
//
// Usage: node check-related-consistency.js <project-root>

const fs = require('fs');
const path = require('path');
const { walkMarkdown } = require('./lib/edc-fs');
const { parseFrontmatter } = require('./lib/frontmatter');

const LINK_RE = /\[[^\]]*\]\(([^)]+)\)/g;

function isExternal(target) {
  return /^[a-z][a-z0-9+.-]*:/i.test(target);
}

// Every local target the body links, whatever its extension. `related:`
// legitimately points at non-Markdown things — LICENSE, a directory like
// docs/10-decisions/, .github/workflows/ci.yml.example — and restricting
// this side to .md made those entries impossible to satisfy.
function bodyLinkTargets(file, body) {
  const targets = new Set();
  let m;
  LINK_RE.lastIndex = 0;
  while ((m = LINK_RE.exec(body))) {
    const target = m[1].trim();
    if (isExternal(target)) continue;
    const [filePart] = target.split('#');
    if (!filePart) continue;
    targets.add(path.resolve(path.dirname(file), filePart));
  }
  return targets;
}

// The advisory direction stays Markdown-only on purpose: a body link to a
// source file or a config example is not a doc-graph edge worth nagging
// about.
function bodyLinkedMarkdownFiles(file, body) {
  const targets = new Set();
  for (const t of bodyLinkTargets(file, body)) {
    if (t.endsWith('.md')) targets.add(t);
  }
  return targets;
}

function relatedTargets(file, related) {
  const targets = new Map();
  for (const entry of related) {
    if (typeof entry !== 'string') continue;
    const [filePart] = entry.split('#');
    if (!filePart) continue;
    targets.set(path.resolve(path.dirname(file), filePart), entry);
  }
  return targets;
}

function run(root) {
  const files = walkMarkdown(root);
  const issues = [];
  const advisories = [];
  let checked = 0;

  for (const file of files) {
    const rel = path.relative(root, file);
    const content = fs.readFileSync(file, 'utf8');
    const parsed = parseFrontmatter(content);
    if (!parsed || !Array.isArray(parsed.fields.related)) continue;

    checked++;
    const related = relatedTargets(file, parsed.fields.related);
    const linkedAny = bodyLinkTargets(file, parsed.body);
    const linkedMd = bodyLinkedMarkdownFiles(file, parsed.body);

    for (const [target, entry] of related) {
      if (!fs.existsSync(target)) {
        issues.push(`${rel}: related: '${entry}' points at a file that does not exist`);
      } else if (!linkedAny.has(target)) {
        issues.push(
          `${rel}: related: lists ${path.relative(root, target)} but it's not linked in the body`
        );
      }
    }

    for (const target of linkedMd) {
      if (!related.has(target) && target !== file) {
        advisories.push(
          `${rel}: body links ${path.relative(root, target)} but it's not in related: (one-off example? use judgment)`
        );
      }
    }
  }

  const ok = issues.length === 0;
  const summary =
    `${checked} documents with a related: field checked, ` +
    `${issues.length} error(s), ${advisories.length} advisory.`;
  return { ok, summary, details: issues, advisories };
}

module.exports = { run };

if (require.main === module) {
  const root = process.argv[2];
  if (!root) {
    console.error('usage: node check-related-consistency.js <project-root>');
    process.exit(1);
  }
  const { ok, summary, details, advisories } = run(root);
  console.log(summary);
  details.forEach((d) => console.log(`  ${d}`));
  advisories.forEach((a) => console.log(`  (advisory) ${a}`));
  process.exit(ok ? 0 : 1);
}
