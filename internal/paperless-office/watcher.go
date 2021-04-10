package paperlessoffice

import (
	"os"
	"path/filepath"

	"github.com/dietsche/rfsnotify"
	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/ocr"
	"gopkg.in/fsnotify.v1"
)

func Watch() {
	var (
		watcher   *rfsnotify.RWatcher
		err       error
		event     fsnotify.Event
		ok        bool
		file      os.FileInfo
		ocrClient *ocr.OCR
	)

	if watcher, err = rfsnotify.NewWatcher(); err != nil {
		log.Fatal("Failed to initialize file watcher")
	}
	watcher.AddRecursive(globalConfig.RawPDFDir)
	for {
		select {
		case event, ok = <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				go func() {
					if file, err = os.Stat(event.Name); err != nil {
						log.Error(err)
						return
					}
					if ocrClient, err = ocr.NewOCR(filepath.Join(globalConfig.RawPDFDir, file.Name())); err != nil {
						log.Errorf("OCRClientErr: %s", err.Error())
						return
					}
					if ocrClient == nil {
						return
					}
					log.Infof("Enqueueing file %s", filepath.Base(ocrClient.RawPDFFile))
					pdfChan <- ocrClient
				}()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Errorf("Watcher error: %s", err.Error())
		}
	}
}
