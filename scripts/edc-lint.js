#!/usr/bin/env node
// Runs the full EDC documentation lint suite against a project root and
// prints one consolidated report — see documentation-structure-and-
// governance.md's "Verification tooling" section.
//
// Usage: node edc-lint.js <project-root>

const checks = [
  { name: 'links', label: 'Internal links', mod: require('./check-links') },
  { name: 'metadata', label: 'Metadata headers', mod: require('./check-metadata') },
  { name: 'related', label: 'related: consistency', mod: require('./check-related-consistency') },
  { name: 'governance', label: 'Governance structure', mod: require('./check-governance') },
];

function run(root) {
  const results = checks.map(({ name, label, mod }) => ({
    name,
    label,
    ...mod.run(root),
  }));
  const ok = results.every((r) => r.ok);
  return { ok, results };
}

module.exports = { run };

if (require.main === module) {
  const root = process.argv[2];
  if (!root) {
    console.error('usage: node edc-lint.js <project-root>');
    process.exit(1);
  }

  const { ok, results } = run(root);

  for (const r of results) {
    console.log(`\n== ${r.label} ==`);
    console.log(r.summary);
    r.details.forEach((d) => console.log(`  ${d}`));
    (r.advisories || []).forEach((a) => console.log(`  (advisory) ${a}`));
  }

  // Advisories are reported but never gate: they're the findings a script
  // can surface but only a human can rule on (see check-related-
  // consistency.js's note on section 1.6's one-off-example exception).
  const totalErrors = results.reduce((sum, r) => sum + r.details.length, 0);
  const totalAdvisories = results.reduce((sum, r) => sum + (r.advisories || []).length, 0);
  const advisoryNote = totalAdvisories ? `, ${totalAdvisories} advisory` : '';
  console.log(
    `\n${ok ? 'PASS' : 'FAIL'} — ${totalErrors} error(s)${advisoryNote} across ${results.length} checks.`
  );
  process.exit(ok ? 0 : 1);
}
