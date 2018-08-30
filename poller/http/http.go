package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type HTTP struct {
	Method string
	URL    *url.URL
	Body   string

	client http.Client
}

func New(method string, fullURL string, body string) (HTTP, error) {
	u, err := url.Parse(fullURL)
	if err != nil {
		return HTTP{}, err
	}
	return HTTP{
		Method: method,
		URL:    u,
		Body:   body,

		client: http.Client{},
	}, nil
}

func isValidResponseCode(code int) bool {
	return 200 <= code && code < 300
}

func (h HTTP) CheckReady() (ready bool, err error) {
	req, err := http.NewRequest(h.Method, h.URL.String(), strings.NewReader(h.Body))
	if err != nil {
		return false, err
	}
	resp, err := h.client.Do(req)

	log.WithFields(log.Fields{
		"module":      "poller.HTTP",
		"status_code": resp.StatusCode,
	}).Debug("http_response")

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if !isValidResponseCode(resp.StatusCode) {
		return false, fmt.Errorf("expected 2XX, received: %d", resp.StatusCode)
	}

	return true, nil
}
