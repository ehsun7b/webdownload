package downloader

import (
	"io/ioutil"
	"net/http"
	"path"
)

// URL string type for all Url fields
//type URL string

// SimpleDownloader simple downloader, implements Downloader interface
type SimpleDownloader struct {
	URL string
}

// Download method
func (sd *SimpleDownloader) Download() ([]byte, string, error) {
	//resp, err := http.Get(sd.URL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", sd.URL, nil)

	if err != nil {
		return nil, "", err
	}

	file := path.Base(req.URL.Path)

	if file == "." {
		file = req.URL.Host
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, file, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, file, err
	}

	return body, file, nil
}
