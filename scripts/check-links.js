#!/usr/bin/env node
// Checks internal Markdown links (relative file paths + #anchors) across
// a project's docs. Requires the `github-slugger` package (same slugger
// GitHub itself uses) so anchor checks match GitHub's real rendering,
// not an approximation.
//
// Usage: node check-links.js <project-root>
//
// A heading can opt out of GitHub's auto-slug by placing an explicit
// stable anchor immediately before it:
//   <a id="m-0001"></a>
//   ## M-0001 — Some Title
// This script treats any such <a id="..."> as a valid anchor target in
// addition to auto-generated heading slugs — see docs/00-project/
// MILESTONES.md's entry template for why this project's numbered
// milestones use this pattern (auto-slugs are fragile: they change if a
// heading's title text changes, and GitHub's real slug algorithm produces
// non-obvious results for em dashes/periods, e.g. "Next.js" -> "nextjs").

const fs = require('fs');
const path = require('path');
const { walkMarkdown } = require('./lib/edc-fs');

let GithubSlugger;
try {
  GithubSlugger = require('github-slugger').default;
} catch {
  console.error(
    'Missing dependency: github-slugger. Run `npm install github-slugger` in this directory first.'
  );
  process.exit(2);
}

function anchorsFor(content) {
  const slugger = new GithubSlugger();
  const anchors = new Set();

  // Explicit stable anchors: <a id="..."></a> or <a name="...">
  for (const m of content.matchAll(/<a\s+(?:id|name)=["']([^"']+)["']/gi)) {
    anchors.add(m[1].toLowerCase());
  }

  // Auto-generated heading slugs, in document order (github-slugger
  // dedupes repeats the same way GitHub does, so order matters).
  for (const m of content.matchAll(/^#+\s+(.*)$/gm)) {
    anchors.add(slugger.slug(m[1]));
  }

  return anchors;
}

function run(root) {
  const mdFiles = walkMarkdown(root);
  const linkRe = /\[[^\]]*\]\(([^)]+)\)/g;
  let totalLinks = 0;
  const broken = [];

  for (const file of mdFiles) {
    const content = fs.readFileSync(file, 'utf8');
    let m;
    while ((m = linkRe.exec(content))) {
      const target = m[1].trim();
      if (/^https?:\/\//.test(target) || target.startsWith('mailto:')) continue;
      totalLinks++;

      const [filePart, anchor] = target.split('#');
      const resolved = filePart === '' ? file : path.resolve(path.dirname(file), filePart);

      // A link that leaves the repository resolves fine on the machine that
      // wrote it and nowhere else. It is broken for everyone who clones.
      const abs = path.resolve(root);
      const resolvedAbs = path.resolve(resolved);
      if (resolvedAbs !== abs && !resolvedAbs.startsWith(abs + path.sep)) {
        broken.push(`${path.relative(root, file)} -> ${target} (escapes the repository: ${resolvedAbs})`);
        continue;
      }

      if (!fs.existsSync(resolved)) {
        broken.push(`${path.relative(root, file)} -> ${target} (missing file: ${path.relative(root, resolved)})`);
        continue;
      }

      if (anchor && resolved.endsWith('.md')) {
        const targetContent = fs.readFileSync(resolved, 'utf8');
        const anchors = anchorsFor(targetContent);
        if (!anchors.has(anchor.toLowerCase())) {
          broken.push(`${path.relative(root, file)} -> ${target} (anchor #${anchor} not found in ${path.relative(root, resolved)})`);
        }
      }
    }
  }

  const ok = broken.length === 0;
  const summary = `${mdFiles.length} markdown files scanned, ${totalLinks} internal links checked.`;
  return { ok, summary, details: broken };
}

module.exports = { run };

if (require.main === module) {
  const root = process.argv[2];
  if (!root) {
    console.error('usage: node check-links.js <project-root>');
    process.exit(1);
  }
  const { ok, summary, details } = run(root);
  console.log(summary);
  if (details.length === 0) {
    console.log('0 broken links.');
  } else {
    console.log(`${details.length} broken links:`);
    details.forEach((b) => console.log(`  ${b}`));
  }
  process.exit(ok ? 0 : 1);
}
