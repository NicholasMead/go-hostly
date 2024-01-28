package ghoster

import "net/http"

type MockClient struct {
	Req  *http.Request
	Resp *http.Response
}

// Do implements httpClient.
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.Req = req
	return m.Resp, nil
}
