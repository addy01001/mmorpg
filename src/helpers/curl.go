package helpers

import (
	"io"
	"net/http"
	"time"
)

func Call(url, method string, body io.Reader, auth string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, body)
	var EmptyReader io.ReadCloser
	if err != nil {
		return EmptyReader, err
	}
	req.Header.Set("user-agent", "golang application")
	req.Header.Add("authorization", auth)

	response, err := client.Do(req)
	if err != nil {
		return response.Body, err
	}

	defer response.Body.Close()
	return response.Body, nil
}
