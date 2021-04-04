package config

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/parsers"
	"gorm.io/gorm"
)

func (entry *PDFEntry) AfterDelete(tx *gorm.DB) (err error) {

	if len(entry.Filename) == 0 {
		return errors.New("Couldn't remove files")
	}
	if err = os.Remove(filepath.Join(instance.ProcessedPDFDir, entry.Filename)); err != nil {
		log.Error(err)
	}
	if err = os.Remove(filepath.Join(instance.RawPDFDir, entry.Filename)); err != nil {
		log.Error(err)
	}
	if err = os.Remove(filepath.Join(instance.ProcessedPDFDir, strings.Replace(entry.Filename, ".pdf", ".png", 1))); err != nil {
		log.Error(err)
	}
	instance.GitPushChan <- CommitMessage(fmt.Sprintf("Deleted %s", entry.Filename))
	return nil
}

func (entry *PDFEntry) AfterUpdate(tx *gorm.DB) (err error) {
	if entry.Confirmed {
		if _, err = os.Stat(filepath.Join(instance.RawPDFDir, entry.Filename)); !os.IsNotExist(err) {
			if err = os.Remove(filepath.Join(instance.RawPDFDir, entry.Filename)); err != nil {
				log.Error(err)
				return err
			}
		}
	}
	if err = tx.Find(&entry).Error; err != nil {
		log.Error(err)
		return err
	}
	instance.GitPushChan <- CommitMessage(fmt.Sprintf("Updated %s", entry.Filename))
	return nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var (
		parsedDate time.Time
		err        error
	)
	b = bytes.ReplaceAll(b, []byte(`"`), []byte(""))
	if parsedDate, err = parsers.ParseDate(string(b)); err != nil {
		return err
	}

	*d = Date(parsedDate)
	return nil
}

func (d *Date) Scan(value interface{}) error {
	var (
		ok bool
		t  time.Time
	)

	t, ok = value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal date value:", value))
	}
	*d = Date(t)
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}
