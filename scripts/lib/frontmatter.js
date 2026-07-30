// Minimal parser for this project family's metadata header dialect — see
// METADATA_HEADER.md. Intentionally not a general YAML parser: every real
// header in this repo family is flat scalars plus one level of `- item`
// lists (audience, related), so a hand-rolled parser avoids pulling in a
// new dependency for a format this constrained.

function parseFrontmatter(content) {
  if (!content.startsWith('---\n') && !content.startsWith('---\r\n')) return null;
  const end = content.indexOf('\n---', 3);
  if (end === -1) return null;

  const raw = content.slice(0, end).split(/\r?\n/).slice(1); // drop opening '---'
  const fields = {};
  let currentKey = null;

  for (const line of raw) {
    const kv = line.match(/^([A-Za-z_][A-Za-z0-9_]*):\s*(.*)$/);
    const item = line.match(/^\s+-\s+(.*)$/);
    if (kv) {
      currentKey = kv[1];
      const val = kv[2].trim();
      if (val === '' ) {
        fields[currentKey] = []; // a list is expected to follow
      } else if (val === '[]') {
        fields[currentKey] = [];
      } else {
        fields[currentKey] = val;
      }
    } else if (item && currentKey && Array.isArray(fields[currentKey])) {
      fields[currentKey].push(item[1].trim());
    }
  }

  const closingLineEnd = content.indexOf('\n', end + 1);
  const bodyStart = closingLineEnd === -1 ? content.length : closingLineEnd + 1;

  return { fields, body: content.slice(bodyStart) };
}

module.exports = { parseFrontmatter };
