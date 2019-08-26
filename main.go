package main

import (
	"github.com/mathieubrun/go-google-photos-api/codegenerator"
	"go/format"
	"log"
	"os"
	"strings"
)

func main() {
	document, err := os.Open("v1/photoslibrary-gen.json")
	if err != nil {
		log.Fatal(err)
	}
	defer document.Close()

	buf := &strings.Builder{}

	err = codegenerator.ProcessAll(document, buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt, err := format.Source([]byte(buf.String()))
	if err != nil {
		log.Fatal(err)
	}

	code, err := os.Create("v1/photoslibrary-gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer code.Close()

	code.Write(fmt)
}
