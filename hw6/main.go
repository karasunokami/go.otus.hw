package hw6

import (
	"flag"
	"log"
)

var from string
var to string
var limit int
var offset int

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.IntVar(&limit, "limit", 0, "limit of bytes to copy")
	flag.IntVar(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	err := Copy(from, to, limit, offset)
	if err != nil {
		log.Panicf("Failed to copy file: %s", err)
	}
}