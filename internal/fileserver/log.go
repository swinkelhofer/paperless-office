package fileserver

import (
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func createLog(w ResponseBuffer, r *http.Request) {
	var (
		ip, port string
		fields   log.Fields
	)

	ip, port, _ = net.SplitHostPort(r.RemoteAddr)
	fields = log.Fields{
		"code":         fmt.Sprintf("%d - %s", w.StatusCode, w.Status),
		"responseCode": w.StatusCode,
		"method":       r.Method,
		"path":         r.RequestURI,
		"remoteHost":   ip,
		"remotePort":   port,
		"proto":        r.Proto,
	}
	log.WithFields(fields).Info("Handled request")
}
