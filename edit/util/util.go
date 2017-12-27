package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Get(url string) ([]byte, error) {
	reader, err := GetReader(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetReader(url string) (io.Reader, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		return nil, os.ErrNotExist
	}
	if response.StatusCode != 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http error %d: %s", response.StatusCode, body)
	}
	return response.Body, nil
}
