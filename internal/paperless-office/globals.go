package paperlessoffice

import (
	"gitlab.com/swinkelhofer/paperless-office/internal/config"
	"gitlab.com/swinkelhofer/paperless-office/internal/ocr"
)

var (
	globalConfig config.Configuration
	pdfChan      chan *ocr.OCR
)
