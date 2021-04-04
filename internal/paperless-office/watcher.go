package paperlessoffice

import (
	"os"

	"github.com/dietsche/rfsnotify"
	log "github.com/sirupsen/logrus"
	"gopkg.in/fsnotify.v1"
)

func Watch() {
	var (
		watcher *rfsnotify.RWatcher
		err     error
		event   fsnotify.Event
		ok      bool
		file    os.FileInfo
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
					log.Infof("Enqueueing file %s", file.Name())
					pdfChan <- file
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
