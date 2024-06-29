package MyAnimeListSDK

import (
	"encoding/json"
	"io"
	"net/http"
	URL "net/url"
	"strings"
)

var (
	baseURL, _ = URL.Parse("https://api.myanimelist.net")
	pathPrefix = "v2"
)

type Method string

var (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

func buildReqURL(path string, queryParams map[string]string) string {
	newUrl := baseURL.JoinPath(pathPrefix, path)

	qParams := URL.Values{}
	for key, value := range queryParams {
		qParams.Add(key, value)
	}

	newUrl.RawQuery = qParams.Encode()
	return newUrl.String()
}

func buildRequest(method Method, path string, queryParams map[string]string, reqBody string) (*http.Request, error) {
	url := buildReqURL(path, queryParams)

	var bodyReader io.Reader
	if reqBody != "" {
		bodyReader = strings.NewReader(reqBody)
	}

	req, err := http.NewRequest(string(method), url, bodyReader)
	if err == nil && bodyReader != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, err
}

func (c MyAnimeListClient) request(method Method, path string, queryParams map[string]string, reqBody string) (*map[string]any, int, error) {
	req, err := buildRequest(method, path, queryParams, reqBody)
	if err != nil {
		return nil, -1, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer func(resp *http.Response) {
		_ = resp.Body.Close()
	}(resp)

	body := &map[string]any{}
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
