package downloader

import (
	"io/ioutil"
	"net/http"
)

// URL string type for all Url fields
//type URL string

// SimpleDownloader simple downloader, implements Downloader interface
type SimpleDownloader struct {
	URL string
}

// Download method
func (sd *SimpleDownloader) Download() ([]byte, error) {
	resp, err := http.Get(sd.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
