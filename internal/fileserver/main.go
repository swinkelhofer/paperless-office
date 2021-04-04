package fileserver

import (
	"net/http"
	"strings"
)

func NewFileServer(dir http.FileSystem) *FileServer {
	return &FileServer{
		FileServer: http.FileServer(dir),
	}
}

func (handler *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		buf = NewResponseBuffer(w)
	)

	if strings.HasSuffix(r.URL.Path, ".pdf") {
		w.Header().Add("Cache-Control", "no-cache")
	} else {
		w.Header().Add("Cache-Control", "max-age=3600")
	}

	handler.FileServer.ServeHTTP(&buf, r)
	createLog(buf, r)
}
