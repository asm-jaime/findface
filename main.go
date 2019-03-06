package main

import (
	"flag"
	"log"
)

type flags struct {
	start       *string
	pathtovideo *string
}

func main() {
	fs := flags{}
	fs.start = flag.String("start", "face", "start face detection")
	fs.pathtovideo = flag.String("pathtovideo", "docs/std.webm", "path to videofile")
	flag.Parse()

	switch *fs.start {
	case "face":
		getdataDlib(*fs.pathtovideo)
	default:
		log.Println("wrong command")
	}

}
