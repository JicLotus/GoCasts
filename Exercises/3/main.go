package main

import (
	"io"
	"log"
	"os"
)

func main() {

	fn := os.Args[1]
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
}
