package config

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConfig() Configuration {
	var (
		err error
	)

	once.Do(func() {
		instance = Configuration{}
		parseFlags()
		instance.InitGitWorker()
		if instance.DB, err = gorm.Open(sqlite.Open(filepath.Join(instance.ProcessedPDFDir, "db.sqlite")), &gorm.Config{}); err != nil {
			log.Fatal(err)
		}
		if err = instance.DB.AutoMigrate(PDFEntry{}); err != nil {
			log.Fatal(err)
		}
		if err = instance.DB.AutoMigrate(Tag{}); err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
