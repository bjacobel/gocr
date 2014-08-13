###gocr
---

Speedy OCR of weird shit people put on Pinterest.

Requirements:
    - [tesseract-ocr](https://code.google.com/p/tesseract-ocr/)
    - [imagemagick](http://www.imagemagick.org/)
    - go 1.3



Speed without goroutines:

        real    0m14.225s
        user    0m6.422s
        sys     0m0.975s

With goroutines:

        real    0m6.240s
        user    0m10.584s
        sys     0m1.321s