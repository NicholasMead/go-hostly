package ghoster

import (
	"net/http"
	"strings"
	"testing"

	"github.com/NicholasMead/go-hostly/internal/target"
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

	t.Run("ForwardsRequest", func(t *testing.T) {
		mockClient := &MockClient{
			Resp: &http.Response{
				Status:     "Ok",
				StatusCode: 200,
				Header:     map[string][]string{},
				Body:       nil,
			},
		}
		client = mockClient

		target := target.Target{
			Host: "http://target-host.com",
		}
		request, _ := http.NewRequest(
			"POST",
			"http://gohostly.io",
			strings.NewReader(""))
		responceWriter := &stubRespWriter{}

		Handler{target}.ServeHTTP(responceWriter, request)

		if mockClient.Req == nil {
			t.Fatalf("No forwarded request")
		}

	})
}

