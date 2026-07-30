// Shared filesystem walking for the EDC lint scripts.
const fs = require('fs');
const path = require('path');

// .github/ holds GitHub's own issue/PR templates and workflows, not EDC
// documentation — METADATA_HEADER.md only governs docs/ and the Phase A
// root files. docs-store/ (driverworks-tooling only) holds vendored
// external DriverWorks API reference docs, not authored under EDC
// conventions — see CONTROL4_FAMILY.md.
const ALWAYS_SKIP_DIRS = new Set(['.git', 'node_modules', '.github', 'docs-store']);

// Walks a project root for .md files. `skipDirNames` lets a caller opt out
// of vendored/non-EDC content (e.g. driverworks-tooling's docs-store/,
// which holds external DriverWorks API reference docs, not EDC-governed
// project documentation).
function walkMarkdown(root, { skipDirNames = [] } = {}) {
  const skip = new Set([...ALWAYS_SKIP_DIRS, ...skipDirNames]);
  const files = [];

  function walk(dir) {
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
      if (skip.has(entry.name)) continue;
      const full = path.join(dir, entry.name);
      if (entry.isDirectory()) walk(full);
      else if (entry.name.endsWith('.md')) files.push(full);
    }
  }

  walk(root);
  return files;
}

module.exports = { walkMarkdown };
