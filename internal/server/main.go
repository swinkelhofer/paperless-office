package server

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/swinkelhofer/paperless-office/internal/config"
	"gitlab.com/swinkelhofer/paperless-office/internal/fileserver"
	paperlessoffice "gitlab.com/swinkelhofer/paperless-office/internal/paperless-office"
)

func NewServer() *http.Server {
	var (
		server = &http.Server{
			ReadTimeout:       300 * time.Second,
			ReadHeaderTimeout: 100 * time.Second,
			WriteTimeout:      300 * time.Second,
			IdleTimeout:       300 * time.Second,
		}
		fileServer http.Handler
	)
	globalConfig = config.NewConfig()

	fileServer = fileserver.NewFileServer(http.Dir(globalConfig.ServeDir))

	http.HandleFunc("/api/documents", paperlessoffice.HandleDocuments)
	http.HandleFunc("/api/tags", paperlessoffice.HandleTags)
	http.HandleFunc("/api/search", paperlessoffice.HandleSearch)
	http.HandleFunc("/api/rescan", paperlessoffice.HandleRescan)
	http.HandleFunc("/api/rescan/finalize", paperlessoffice.HandleRescanFinalize)
	http.HandleFunc("/api/rescan/abort", paperlessoffice.HandleRescanAbort)
	http.HandleFunc("/api/upload", paperlessoffice.HandleUpload)
	http.HandleFunc("/", fileServer.ServeHTTP)

	server.Addr = fmt.Sprintf("%s:%d", globalConfig.Server.Listen.Host, globalConfig.Server.Listen.Port)
	return server
}
