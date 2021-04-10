package paperlessoffice

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/config"
	"gitlab.com/swinkelhofer/paperless-office/internal/ocr"
)

func Init() {
	globalConfig = config.NewConfig()
	InitWorkers()
	InitialRawDirectoryProcessing()
	Watch()
}

func InitialRawDirectoryProcessing() {
	var (
		files     []fs.FileInfo
		file      fs.FileInfo
		err       error
		ocrClient *ocr.OCR
	)

	if files, err = ioutil.ReadDir(globalConfig.RawPDFDir); err != nil {
		log.Error(err)
		return
	}

	for _, file = range files {
		if ocrClient, err = ocr.NewOCR(filepath.Join(globalConfig.RawPDFDir, file.Name())); err != nil {
			log.Errorf("OCRClientErr: %s", err.Error())
			return
		}
		if ocrClient == nil {
			return
		}
		log.Infof("Enqueueing file %s", file.Name())
		pdfChan <- ocrClient
	}
}
