package paperlessoffice

import (
	"time"

	"gitlab.com/swinkelhofer/paperless-office/internal/config"
)

type Date struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type Search struct {
	ID          string       `json:"ID"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	Tags        []config.Tag `json:"tags"`
	EmailAdress string       `json:"emailAdress"`
	URL         string       `json:"url"`
	Date        Date         `json:"date"`
	Unconfirmed bool         `json:"unconfirmed"`
	InTrash     bool         `json:"inTrash"`
}

type Rescan struct {
	EntryID int64 `json:"entryID"`
}

type Worker int

type UploadRequest struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	FileData string `json:fileData`
}
