package ghoster

import (
	"net/http"
	"strings"
)

type stubRespWriter struct {
	StatusCode int
	Body       string
	Headers    http.Header
}

// Header implements http.ResponseWriter.
func (s *stubRespWriter) Header() http.Header {
	if s.Headers == nil {
		s.Headers = make(http.Header)
	}

	return s.Headers
}

// Write implements http.ResponseWriter.
func (s *stubRespWriter) Write(text []byte) (int, error) {
	builder := strings.Builder{}
	builder.WriteString(s.Body)
	i, e := builder.Write(text)

	s.Body = builder.String()
	return i, e
}

// WriteHeader implements http.ResponseWriter.
func (s *stubRespWriter) WriteHeader(statusCode int) {
	s.StatusCode = statusCode
}
