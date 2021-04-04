package config

import (
	"time"

	"github.com/go-git/go-git/v5"
	"gorm.io/gorm"
)

type logging struct {
	Level  string
	Format string
}

type listen struct {
	Port      int
	Host      string
	TLSConfig string
}

type server struct {
	Listen listen
}

type author struct {
	Name  string
	Email string
}

type gitConfig struct {
	Author  author
	gitRepo *git.Repository
}

type Configuration struct {
	Server          server
	Workers         int
	Logging         logging
	ServeDir        string
	RawPDFDir       string
	ProcessedPDFDir string
	Git             gitConfig
	DB              *gorm.DB
	GitPushChan     chan CommitMessage
}

type Date time.Time
type CommitMessage string

type PDFEntry struct {
	gorm.Model
	Filename       string `json:"filename"`
	RescanFilename string `json:"rescanFilename"`
	Content        string `json:"content"`
	Title          string `json:"title"`
	EMailAdresses  string `json:"emailAdresses"`
	URLs           string `json:"urls"`
	PhoneNumbers   string `json:"phoneNumbers"`
	Date           Date   `json:"date"`
	Confirmed      bool   `json:"confirmed"`
	InTrash        bool   `json:"inTrash" gorm:"inTrash"`
	Tags           []Tag  `json:"tags" gorm:"many2many:pdf_tags;"`
}

type Tag struct {
	gorm.Model
	Title string `json:"title" gorm:"not null;unique"`
	Color string `json:"color"`
}
