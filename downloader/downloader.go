package downloader

// Content type for the downloaded content
//type Content string

// Downloader struct to download a url
type Downloader interface {
	Download() []byte
}
