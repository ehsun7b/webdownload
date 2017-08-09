package main

import (
	"fmt"
	"os"

	"github.com/ehsun7b/webdownload/downloader"
	"github.com/ehsun7b/webdownload/save"
	//"io/ioutil"
	//"net/http"
)

func main() {
	args := os.Args[1:]
	file := ""
	url := "S"

	if len(args) > 0 {
		url = args[0]
	} else {
		fmt.Println("Enter the URL as the first argument")
		os.Exit(0)
	}

	if len(args) > 1 {
		file = args[1]
	}

	downloader := downloader.SimpleDownloader{URL: url}
	content, file2, err := downloader.Download()

	if file == "" {
		file = file2
	}

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	} else {
		//s := string(content[:])
		saver := new(save.FileSaver)
		err := saver.Save(content, file)

		if err != nil {
			panic(err)
		}
	}
}

func test(saver save.Saver) {
	fmt.Println(saver)
}
