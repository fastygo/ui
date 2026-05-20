#!/usr/bin/env python3
"""Fix showcase Go: two-arg ui.Button/Label/… calls and }func spacing."""
from __future__ import annotations

import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SCAN = [ROOT / "internal" / "showcase" / "catalog", ROOT / "internal" / "views"]


def balanced_end(s: str, open_i: int) -> int:
    depth = 0
    i = open_i
    while i < len(s):
        c = s[i]
        if c == "(":
            depth += 1
        elif c == ")":
            depth -= 1
            if depth == 0:
                return i + 1
        elif c in "\"'`":
            q = c
            i += 1
            while i < len(s) and s[i] != q:
                if s[i] == "\\":
                    i += 1
                i += 1
        i += 1
    return -1


def split_call_args(inner: str) -> list[str] | None:
    args: list[str] = []
    start = 0
    depth = 0
    i = 0
    while i < len(inner):
        c = inner[i]
        if c in "({[":
            depth += 1
        elif c in ")}]":
            depth -= 1
        elif c == "," and depth == 0:
            args.append(inner[start:i].strip())
            start = i + 1
        elif c in "\"'`":
            q = c
            i += 1
            while i < len(inner) and inner[i] != q:
                if inner[i] == "\\":
                    i += 1
                i += 1
        i += 1
    args.append(inner[start:].strip())
    return args if len(args) == 2 else None


CONFIG = {
    "Button": ("showcaseutil.Button({p}, {l})", True),
    "Label": ("showcaseutil.RenderLabel(ctx, w, {p}, {l})", False),
    "Badge": ("showcaseutil.RenderBadge(ctx, w, {p}, {l})", False),
    "TableHeadCell": ("showcaseutil.RenderTableHeadCell(ctx, w, {p}, {l})", False),
    "TableCell": ("showcaseutil.RenderTableCell(ctx, w, {p}, {l})", False),
}


def fix_calls(content: str) -> str:
    for func, (tmpl, keep_render) in CONFIG.items():
        needle = f"ui.{func}("
        out: list[str] = []
        i = 0
        while True:
            idx = content.find(needle, i)
            if idx < 0:
                out.append(content[i:])
                break
            out.append(content[i:idx])
            open_paren = idx + len(needle) - 1
            end = balanced_end(content, open_paren)
            if end < 0:
                out.append(content[idx:])
                break
            inner = content[open_paren + 1 : end - 1]
            args = split_call_args(inner)
            suffix = ""
            render_end = end
            if keep_render and content[end : end + 14] == ".Render(ctx, w)":
                render_end = end + 14
            elif not keep_render and content[end : end + 14] == ".Render(ctx, w)":
                render_end = end + 14
                suffix = ""
            args_ok = args is not None
            if args_ok:
                props, label = args
                repl = tmpl.format(p=props, l=label)
                out.append(repl)
                if keep_render:
                    out.append(".Render(ctx, w)")
            else:
                out.append(content[idx:render_end])
            i = render_end
        content = "".join(out)
    return content


def ensure_import(content: str) -> str:
    if 'github.com/fastygo/ui/internal/showcase/showcaseutil' in content:
        return content
    if "showcaseutil." not in content:
        return content
    if "import (" not in content:
        return content
    return content.replace(
        "import (",
        'import (\n\t"github.com/fastygo/ui/internal/showcase/showcaseutil"',
        1,
    )


def process_file(path: Path) -> bool:
    text = path.read_text(encoding="utf-8")
    orig = text
    text = text.replace("}func ", "}\n\nfunc ")
    text = fix_calls(text)
    text = ensure_import(text)
    if text != orig:
        path.write_text(text, encoding="utf-8")
        return True
    return False


def main() -> None:
    n = 0
    for base in SCAN:
        if not base.exists():
            continue
        for path in base.rglob("*.go"):
            if path.name.endswith("_templ.go"):
                continue
            if process_file(path):
                print("fixed", path.relative_to(ROOT))
                n += 1
    print("done", n)


if __name__ == "__main__":
    main()
