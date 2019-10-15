package step

import (
	"strings"

	"github.com/iv-p/apid/common/http"
)

type executor interface {
	do(Request) (*http.Response, error)
}

type httpExecutor struct {
	client http.Client
}

// NewHTTPExecutor instantiates a new http executor
func NewHTTPExecutor(client http.Client) executor {
	return &httpExecutor{client: client}
}

func (e *httpExecutor) do(request Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Type, request.Endpoint, strings.NewReader(request.Body))
	if err != nil {
		return nil, err
	}
	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}
	return e.client.Do(req.Context(), req)
}
