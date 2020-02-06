package main

import (
	"flag"

	"github.com/tobbbles/muzip/archive"
	p "github.com/tobbbles/muzip/print"
)

var (
	file = flag.String("file", "", "file to read from")
	dir  = flag.String("dir", "", "directory to recursively read from zip files")
)

func main() {
	flag.Parse()

	if len(*file) != 0 {
		a, err := Index(*file)
		if err != nil {
			panic(err)
		}

		p.Pretty([]*archive.Archive{a})
		return
	}

	if len(*dir) != 0 {
		as, err := Walk(*dir)
		if err != nil {
			panic(err)
		}

		p.Pretty(as)
		return
	}

}
