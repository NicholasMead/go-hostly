package ghoster

import "net/http"

type Handler struct {
}

func (h Handler) ServeHTTP(_ http.ResponseWriter,_ *http.Request) {
}

