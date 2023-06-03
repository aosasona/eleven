package eleven

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	// Path to the endpoint without the URL or version, eg: /voices
	Path string

	// Headers to be sent with the request (xi-api-key is already set)
	Headers map[string]string

	// Additional data to be sent with the request
	Data any
}

type RequestMethod string

const (
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"
)

func (e *Eleven) sendRequest(method RequestMethod, args Request) (*http.Response, error) {
	args.Path = strings.Trim(args.Path, "/")
	url := fmt.Sprintf("%s/%s", e.baseURL, args.Path)

	if args.Headers == nil {
		args.Headers = make(map[string]string)
	}

	args.Headers["xi-api-key"] = e.secret

	data := &bytes.Buffer{}

	if args.Data != nil {
		body, err := json.Marshal(args.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}

		data = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(string(method), url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range args.Headers {
		req.Header.Set(k, v)
	}

	return e.client.Do(req)
}

func (e *Eleven) get(args Request) (*http.Response, error) {
	return e.sendRequest(GET, args)
}

func (e *Eleven) post(args Request) (*http.Response, error) {
	return e.sendRequest(POST, args)
}

func decodeResponse[T any](res *http.Response, out *T) error {
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(out)
}
