#!/usr/bin/env node
// Verifies that files distributed from `_template/` into each project still
// match their canonical copy — see documentation-structure-and-governance.md's
// "Verification tooling" section.
//
// The EDC deliberately duplicates rather than references: the manifesto, the
// metadata spec, the governance guide and these very scripts physically exist
// inside every project, which is what lets their links stay relative and
// survive a clone. The cost of that choice is drift, and nothing used to
// notice it. This does.
//
// Comparison ignores the fields a project is *supposed* to own — `owner:` and
// `last_updated:` — and a short table of documented exceptions. Anything else
// that differs is drift to fix, not a local variation to keep: per the
// governance doc, a project-specific application of a shared document belongs
// in PRODUCT_PHILOSOPHY.md, not in a fork of the shared document.
//
// The canonical copies live outside any repository, so this check is only
// meaningful on the host that holds `_template/`. From a clone it reports that
// it could not run and passes, rather than failing work it cannot judge.
//
// Usage: node check-distribution.js <project-root>

const fs = require('fs');
const path = require('path');

// Files every project receives from the template.
const DISTRIBUTED = [
  'ENGINEERING_MANIFESTO.md',
  'METADATA_HEADER.md',
  'docs/00-project/documentation-structure-and-governance.md',
  'docs/10-decisions/ADR-TEMPLATE.md',
  'docs/11-rfc/RFC-TEMPLATE.md',
  'scripts/check-links.js',
  'scripts/check-metadata.js',
  'scripts/check-related-consistency.js',
  'scripts/check-governance.js',
  'scripts/check-distribution.js',
  'scripts/edc-lint.js',
  'scripts/lib/edc-fs.js',
  'scripts/lib/frontmatter.js',
];

// Differences that are correct rather than drift. Each needs a reason, so an
// exception can be argued with later instead of just inherited.
const EXCEPTIONS = [
  {
    file: 'METADATA_HEADER.md',
    contains: '00_HOW_TO_USE_THIS_TEMPLATE.md',
    reason: 'that walkthrough exists only in _template, so copies drop the reference',
  },
];

// Whole files a named project legitimately owns rather than receives.
// `retroflag-power` is the project this standard was generalized *from*, so
// several of its documents are the ancestors of the template's, not copies of
// them — reconciling those would mean rewriting the original to match the
// abstraction drawn out of it, which is backwards.
const PROJECT_EXEMPT = {
  'retroflag-power': {
    'docs/00-project/documentation-structure-and-governance.md': 'the original this template generalized',
    'docs/10-decisions/ADR-TEMPLATE.md': 'predates the template and carries its own filled-in header',
    'docs/11-rfc/RFC-TEMPLATE.md': 'predates the template and carries its own filled-in header',
  },
};

// Fields a project legitimately sets for itself.
const OWNED = /^(owner|last_updated):.*$/gm;

function normalise(text, relPath) {
  let out = text.replace(OWNED, '');
  for (const e of EXCEPTIONS) {
    if (e.file !== relPath) continue;
    out = out
      .split('\n')
      .filter((line) => !line.includes(e.contains))
      .join('\n');
  }
  return out.replace(/\s+/g, ' ').trim();
}

// Walk up looking for a sibling `_template/`. Absent from a clone, by design.
function findTemplate(start) {
  let dir = path.resolve(start);
  for (let i = 0; i < 6; i++) {
    const candidate = path.join(dir, '_template');
    if (fs.existsSync(path.join(candidate, 'METADATA_HEADER.md'))) return candidate;
    const parent = path.dirname(dir);
    if (parent === dir) break;
    dir = parent;
  }
  return null;
}

function run(root) {
  const abs = path.resolve(root);
  const template = findTemplate(abs);

  if (!template) {
    return {
      ok: true,
      summary: 'canonical _template/ not reachable from here — distribution check skipped.',
      details: [],
      advisories: ['distribution not verified: this looks like a clone rather than the host that holds _template/'],
    };
  }
  if (path.resolve(template) === abs) {
    return { ok: true, summary: 'this is the canonical template; nothing to compare.', details: [], advisories: [] };
  }

  const issues = [];
  const advisories = [];
  let compared = 0;
  const exempt = PROJECT_EXEMPT[path.basename(abs)] || {};

  for (const rel of DISTRIBUTED) {
    const mine = path.join(abs, rel);
    const canonical = path.join(template, rel);
    if (!fs.existsSync(canonical)) continue;
    if (exempt[rel]) {
      advisories.push(`${rel}: exempt — ${exempt[rel]}`);
      continue;
    }
    if (!fs.existsSync(mine)) {
      advisories.push(`${rel}: not present in this project (never distributed, or deliberately dropped)`);
      continue;
    }
    compared++;
    const a = normalise(fs.readFileSync(canonical, 'utf8'), rel);
    const b = normalise(fs.readFileSync(mine, 'utf8'), rel);
    if (a !== b) {
      issues.push(`${rel}: has drifted from _template/${rel} — reconcile them, don't fork the shared document`);
    }
  }

  const ok = issues.length === 0;
  const summary = `${compared} distributed file(s) compared against _template, ${issues.length} drifted.`;
  return { ok, summary, details: issues, advisories };
}

module.exports = { run };

if (require.main === module) {
  const root = process.argv[2];
  if (!root) {
    console.error('usage: node check-distribution.js <project-root>');
    process.exit(1);
  }
  const { ok, summary, details, advisories } = run(root);
  console.log(summary);
  details.forEach((d) => console.log(`  ${d}`));
  advisories.forEach((a) => console.log(`  (advisory) ${a}`));
  process.exit(ok ? 0 : 1);
}
