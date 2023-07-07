package main

import (
	"fmt"
	"io"
	"os"

	Link "github.com/tusharr-patil/html-link-parser"
	"golang.org/x/net/html"
)

func main() {
	testCase := 4
	for i := 1; i <= testCase; i++ {
		filePath := fmt.Sprintf("./Testcase/ex%v.html", i)
		fmt.Println(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		r := io.Reader(file)

		// parse the reader
		doc, _ := html.Parse(r)

		// check for the a tag
		list := Link.GetLinks(doc)

		fmt.Printf("output for test case %v \n", i)
		fmt.Println(list)
	}
}
