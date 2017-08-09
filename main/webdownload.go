package main

import (
	"fmt"

	"github.com/ehsun7b/webdownload/downloader"
	"github.com/ehsun7b/webdownload/save"
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
		//s := string(content[:])
		saver := new(save.FileSaver)
		err := saver.Save(content, "google.html")

		if err != nil {
			panic(err)
		}
	}
}

func test(saver save.Saver) {
	fmt.Println(saver)
}
