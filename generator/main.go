package main

import (
	"github.com/mathieubrun/go-google-photos-api/generator/codegenerator"
	"go/format"
	"log"
	"os"
	"strings"
)

func main() {
	document, err := os.Open("../photoslibrary/v1/photoslibrary-api.json")
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

	code, err := os.Create("../photoslibrary/v1/photoslibrary-gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer code.Close()

	code.Write(fmt)
}
