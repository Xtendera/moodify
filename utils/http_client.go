package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SendURLEncoded(method string, endpoint string, data url.Values, respObj interface{}, token *string) error {
	client := &http.Client{}
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if token != nil {
		req.Header.Set("Authorization", "Bearer "+*token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Stupid ahh bypass to make GoLand stop complaining.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &HTTPError{StatusCode: resp.StatusCode}
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(respBytes, respObj)
}

type HTTPError struct {
	StatusCode int
}

func (e *HTTPError) Error() string {
	return http.StatusText(e.StatusCode)
}
