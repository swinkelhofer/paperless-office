package paperlessoffice

import (
	"io/fs"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/config"
)

func Init() {
	globalConfig = config.NewConfig()
	InitWorkers()
	InitialRawDirectoryProcessing()
	Watch()
}

func InitialRawDirectoryProcessing() {
	var (
		files []fs.FileInfo
		file  fs.FileInfo
		err   error
	)

	if files, err = ioutil.ReadDir(globalConfig.RawPDFDir); err != nil {
		log.Error(err)
		return
	}

	for _, file = range files {
		log.Infof("Enqueueing file %s", file.Name())
		pdfChan <- file
	}
}
