package main

import (
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	filePath := "./Testcase/ex1.html"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := io.Reader(file)

	doc, _ := html.Parse(r)

	html.Render(os.Stdout, doc)

}
