#!/usr/bin/env python3
"""Convert a PDF from output/claude-design-pdf/ into PNG slides in output/images/."""

import sys
from pathlib import Path


def main():
    if len(sys.argv) < 2:
        print("Usage: python pdf_to_images.py <date-slug>  (e.g. 2026-05-27)")
        sys.exit(1)

    slug = sys.argv[1]
    repo_root = Path(__file__).parent.parent
    pdf_path = repo_root / "output" / "claude-design-pdf" / f"{slug}.pdf"
    out_dir = repo_root / "output" / "images" / slug

    if not pdf_path.exists():
        print(f"PDF not found: {pdf_path}")
        sys.exit(1)

    out_dir.mkdir(parents=True, exist_ok=True)

    try:
        import fitz
    except ImportError:
        print("PyMuPDF not installed. Run: pip install pymupdf")
        sys.exit(1)

    doc = fitz.open(str(pdf_path))
    total = len(doc)

    for i, page in enumerate(doc):
        # Zoom 2x for high-quality output (~144–200 DPI depending on source)
        mat = fitz.Matrix(2, 2)
        pix = page.get_pixmap(matrix=mat)
        out_path = out_dir / f"slide-{i + 1:02d}.png"
        pix.save(str(out_path))
        print(f"  {out_path.name}  ({pix.width}×{pix.height}px)")

    doc.close()
    print(f"\nDone: {total} slides saved to output/images/{slug}/")


if __name__ == "__main__":
    main()
