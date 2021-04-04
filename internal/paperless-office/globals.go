package paperlessoffice

import (
	"io/fs"

	"gitlab.com/swinkelhofer/paperless-office/internal/config"
)

var (
	globalConfig config.Configuration
	pdfChan      chan fs.FileInfo
)
