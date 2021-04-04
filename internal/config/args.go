package config

import (
	"flag"
	"path/filepath"

	"github.com/go-git/go-git/v5"

	log "github.com/sirupsen/logrus"
)

func initLogging() {
	switch instance.Logging.Format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{
			ForceColors: true,
		})
	}
	log.ParseLevel(instance.Logging.Level)

	log.Infof("Initialized %s logging on level %s", instance.Logging.Format, log.GetLevel().String())
}

func initGit() {
	var (
		err error
	)

	if instance.Git.gitRepo, err = git.PlainOpen(instance.ProcessedPDFDir); err != nil {
		log.Warnf("Directory target for processed PDFs [%s] is not a valid git repository", instance.ProcessedPDFDir)
		instance.Git.gitRepo = nil
	}
}

func parseFlags() {
	instance = Configuration{
		Server: server{
			Listen: listen{},
		},
		Logging: logging{},
		Git: gitConfig{
			Author: author{},
		},
	}
	flag.StringVar(&(instance.Server.Listen.Host), "listen-host", "0.0.0.0", "Server Listen Host")
	flag.IntVar(&(instance.Server.Listen.Port), "listen-port", 8000, "Server Listen Port")
	flag.IntVar(&(instance.Workers), "workers", 1, "Number of workers that can process PDFs in parallel")
	flag.StringVar(&(instance.Logging.Format), "log-format", "console", "Log format. One of [console, json]")
	flag.StringVar(&(instance.Logging.Level), "log-level", "info", "Log level. One of [info, debug, warn, error, trace]")
	flag.StringVar(&(instance.ServeDir), "serve-dir", "/data", "Directory holding files to serve")
	flag.StringVar(&(instance.Git.Author.Name), "git-author-name", "Paperless Office Bot", "Git author name to sign commits")
	flag.StringVar(&(instance.Git.Author.Email), "git-author-email", "paperless-office@bot.com", "Git author email address to sign commits")
	flag.Parse()
	instance.RawPDFDir = filepath.Join(instance.ServeDir, "data", "raw")
	instance.ProcessedPDFDir = filepath.Join(instance.ServeDir, "data", "processed")
	initLogging()
	initGit()
}
