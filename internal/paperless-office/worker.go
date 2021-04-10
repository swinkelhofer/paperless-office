package paperlessoffice

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/config"
	"gitlab.com/swinkelhofer/paperless-office/internal/ocr"
)

func InitWorkers() {
	var (
		i int
	)
	pdfChan = make(chan *ocr.OCR, 500)

	for i = 1; i <= globalConfig.Workers; i++ {
		go Worker(i).Work(pdfChan)
	}
}

func (w Worker) Work(pdfChan <-chan *ocr.OCR) {
	var (
		pdf *ocr.OCR
	)

	for {
		select {
		case pdf = <-pdfChan:
			w.ConvertRaw(pdf)
		}
	}
}

func (w Worker) ConvertRaw(ocrClient *ocr.OCR) {
	var (
		err           error
		dbEntry       *config.PDFEntry
		extractedText string
		copyBuffer    []byte
	)

	if ocrClient == nil {
		return
	}

	if _, err = os.Stat(ocrClient.ProcessedPDFFile); os.IsNotExist(err) {
		log.Infof("[worker #%d] Found %s. Trying to OCR", w, filepath.Base(ocrClient.RawPDFFile))
		if ocrClient.HasTextLayer() {
			log.Infof("[worker #%d] PDF already has text layer. Skipping OCR", w)
			if copyBuffer, err = ioutil.ReadFile(ocrClient.RawPDFFile); err != nil {
				return
			}
			if err = ioutil.WriteFile(ocrClient.ProcessedPDFFile, copyBuffer, 0644); err != nil {
				return
			}
		} else {
			if err = ocrClient.RunOCR(); err != nil {
				log.Errorf("[worker #%d] OCR error: %s", w, err.Error())
				return
			}
		}
		if err = ocrClient.GeneratePreview(); err != nil {
			log.Errorf("[worker #%d] Generate preview error: %s", w, err.Error())
			return
		}
		if extractedText, err = ocrClient.ExtractText(); err != nil {
			log.Errorf("[worker #%d] Extract text error: %s", w, err.Error())
			return
		}
		dbEntry = &config.PDFEntry{
			Filename:      filepath.Base(ocrClient.ProcessedPDFFile),
			Content:       extractedText,
			Date:          config.Date(ocrClient.ExtractDate(extractedText)),
			PhoneNumbers:  ocrClient.ExtractPhoneNumbers(extractedText),
			URLs:          ocrClient.ExtractURLs(extractedText),
			EMailAdresses: ocrClient.ExtractMailAdresses(extractedText),
		}

		if err = globalConfig.DB.Create(dbEntry).Error; err != nil {
			log.Errorf("[worker #%d] Write DB error: %s", w, err.Error())
			return
		}
		globalConfig.GitPushChan <- config.CommitMessage(fmt.Sprintf("Added %s", dbEntry.Filename))
	} else {
		log.Infof("[worker #%d] Enqueued file %s has been processed already. It's time to confirm it via the Web UI", w, filepath.Base(ocrClient.RawPDFFile))
	}
}
