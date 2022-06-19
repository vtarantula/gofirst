package main

import (
	"fmt"
	gfservice "gofirst/pkg/service"
	gflog "gofirst/pkg/util/log"
)

func cleanup() {
	gflog.Close()
}

func main() {
	fmt.Printf("*** Starting application ***\n")
	defer cleanup()

	gfservice.Start()
	// gfutil.CopyFile("/home/asn/Videos/movies/korean/2022/In_Our_Prime.mp4", "/home/asn/testvid.mp4")
}
