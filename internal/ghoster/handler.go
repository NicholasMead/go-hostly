package ghoster

import (
	"net/http"
	"strings"

	"github.com/NicholasMead/go-hostly/internal/target"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

var client httpClient = &http.Client{}

type Handler struct {
	Target target.Target
}

func (h Handler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	req, _ := http.NewRequest("", "", strings.NewReader(""))

	client.Do(req)
}
