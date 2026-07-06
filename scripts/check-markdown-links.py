#!/usr/bin/env python3
"""Check local Markdown file links and anchors without third-party packages."""

from __future__ import annotations

import re
import sys
from pathlib import Path
from urllib.parse import unquote, urlsplit


ROOT = Path(__file__).resolve().parent.parent
SKIP_PARTS = {".git", "node_modules", "vendor"}
LINK_RE = re.compile(r"!?\[[^\]]*\]\(\s*(<[^>]+>|[^)\s]+)(?:\s+[^)]*)?\)")
EXPLICIT_ANCHOR_RE = re.compile(
    r"""<a\s+[^>]*(?:id|name)\s*=\s*["']([^"']+)["'][^>]*>""",
    re.IGNORECASE,
)
HEADING_RE = re.compile(r"^ {0,3}#{1,6}\s+(.+?)\s*#*\s*$")
HTML_TAG_RE = re.compile(r"<[^>]+>")


def markdown_files() -> list[Path]:
    return sorted(
        path
        for path in ROOT.rglob("*.md")
        if not any(part in SKIP_PARTS for part in path.relative_to(ROOT).parts)
    )


def github_heading_slug(heading: str) -> str:
    """Approximate GitHub's heading slug for the repository's common headings."""
    heading = HTML_TAG_RE.sub("", heading)
    heading = re.sub(r"[^\w\s-]", "", heading.lower(), flags=re.UNICODE)
    return re.sub(r"\s", "-", heading)


def anchors_for(path: Path) -> set[str]:
    anchors: set[str] = set()
    slug_counts: dict[str, int] = {}
    in_fence = False

    for line in path.read_text(encoding="utf-8").splitlines():
        stripped = line.lstrip()
        if stripped.startswith(("```", "~~~")):
            in_fence = not in_fence
            continue
        if in_fence:
            continue

        anchors.update(EXPLICIT_ANCHOR_RE.findall(line))

        heading_match = HEADING_RE.match(line)
        if not heading_match:
            continue

        base_slug = github_heading_slug(heading_match.group(1))
        duplicate_number = slug_counts.get(base_slug, 0)
        slug_counts[base_slug] = duplicate_number + 1
        anchors.add(
            base_slug if duplicate_number == 0 else f"{base_slug}-{duplicate_number}"
        )

    return anchors


def internal_target(raw_target: str) -> tuple[str, str] | None:
    target = raw_target[1:-1] if raw_target.startswith("<") else raw_target
    parsed = urlsplit(target)
    if parsed.scheme or parsed.netloc:
        return None
    return unquote(parsed.path), unquote(parsed.fragment)


def main() -> int:
    files = markdown_files()
    anchor_cache: dict[Path, set[str]] = {}
    errors: list[str] = []
    checked_links = 0

    for source in files:
        in_fence = False
        for line_number, line in enumerate(
            source.read_text(encoding="utf-8").splitlines(), start=1
        ):
            if line.lstrip().startswith(("```", "~~~")):
                in_fence = not in_fence
                continue
            if in_fence:
                continue

            for match in LINK_RE.finditer(line):
                target = internal_target(match.group(1))
                if target is None:
                    continue

                checked_links += 1
                path_part, fragment = target
                destination = (
                    source if not path_part else (source.parent / path_part).resolve()
                )

                try:
                    destination.relative_to(ROOT)
                except ValueError:
                    errors.append(
                        f"{source.relative_to(ROOT)}:{line_number}: "
                        f"link escapes repository: {match.group(1)}"
                    )
                    continue

                if not destination.exists():
                    errors.append(
                        f"{source.relative_to(ROOT)}:{line_number}: "
                        f"missing file: {match.group(1)}"
                    )
                    continue

                if fragment:
                    if destination.suffix.lower() != ".md":
                        errors.append(
                            f"{source.relative_to(ROOT)}:{line_number}: "
                            f"cannot verify anchor in non-Markdown file: "
                            f"{match.group(1)}"
                        )
                        continue
                    anchors = anchor_cache.setdefault(
                        destination, anchors_for(destination)
                    )
                    if fragment not in anchors:
                        errors.append(
                            f"{source.relative_to(ROOT)}:{line_number}: "
                            f"missing anchor #{fragment} in "
                            f"{destination.relative_to(ROOT)}"
                        )

    if errors:
        print("Link Lantern found broken internal Markdown links:")
        for error in errors:
            print(f"  {error}")
        return 1

    print(
        f"Link Lantern checked {checked_links} internal links "
        f"across {len(files)} Markdown files."
    )
    return 0


if __name__ == "__main__":
    sys.exit(main())
