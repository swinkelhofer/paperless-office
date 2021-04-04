package fileserver

import "net/http"

type FileServer struct {
	FileServer http.Handler
}

type ResponseBuffer struct {
	StatusCode int
	Status     string
	w          http.ResponseWriter
}
