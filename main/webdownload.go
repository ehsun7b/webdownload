package main

import (
	"fmt"
	
	//"io/ioutil"
	//"net/http"
)

func main() {
	downloader := Downloader{url: "http://www.google.com"}
	fmt.Println(downloader)
}

