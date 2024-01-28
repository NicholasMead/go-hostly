package ghoster

import (
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("IsHTTPHandler", func(t *testing.T) {
		var h interface{} = Handler{}

		switch h.(type) {
		case http.Handler:
			break
		default:
			t.Fatalf("Not http handler")
		}
	})
}
