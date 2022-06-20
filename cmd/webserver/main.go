package main

import (
	"fmt"
	gfservice "gofirst/src/service"
	gflog "gofirst/src/util/log"
)

func cleanup() {
	gflog.Close()
}

/*
TODO:
Add websocket support
Add database of movies and files
*/

func main() {
	fmt.Printf("*** Starting application ***\n")
	defer cleanup()

	gfservice.Start()
	// gfutil.CopyFile("/home/asn/Videos/movies/korean/2022/In_Our_Prime.mp4", "/home/asn/testvid.mp4")
}
