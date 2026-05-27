---
description: Converts a PDF from output/claude-design-pdf/ into PNG slides saved in output/images/<slug>/. Run after exporting a Canva design as PDF. Usage: /pdf-to-images YYYY-MM-DD
argument-hint: <YYYY-MM-DD>
disable-model-invocation: false
---

Convert the PDF for the given date slug into PNG slides.

1. Verify `output/claude-design-pdf/$ARGUMENTS.pdf` exists. If not, tell the user the file is missing and stop.

2. Run the conversion via uv (handles pymupdf install automatically):
   ```
   uv run --with pymupdf python scripts/pdf_to_images.py $ARGUMENTS
   ```

3. Report back: how many slides were converted, the output folder path, and the pixel dimensions.
