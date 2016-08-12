package http_helper

import (
	"io/ioutil"
	"net/http"
)

type DefaultHttpHelper struct{}

func (DefaultHttpHelper) Fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return getResponseBody(response)
}

func getResponseBody(response *http.Response) ([]byte, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
