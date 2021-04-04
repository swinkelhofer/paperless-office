package fileserver

import "net/http"

func NewResponseBuffer(w http.ResponseWriter) ResponseBuffer {
	return ResponseBuffer{
		w: w,
	}
}

func (rb *ResponseBuffer) Header() http.Header {
	return rb.w.Header()
}

func (rb *ResponseBuffer) Write(b []byte) (int, error) {
	return rb.w.Write(b)
}

func (rb *ResponseBuffer) WriteHeader(statusCode int) {
	rb.StatusCode = statusCode
	rb.Status = http.StatusText(statusCode)
	rb.w.WriteHeader(statusCode)

}
