package main

import (
	"fmt"

	"github.com/ehsun7b/webdownload/downloader"
	//"io/ioutil"
	//"net/http"
)

func main() {
	downloader := downloader.SimpleDownloader{URL: "http://www.google.com"}
	content, err := downloader.Download()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	} else {
		s := string(content[:])
		fmt.Println(s)
	}
}
