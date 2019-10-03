package main

import (
	"fmt"
	"github.com/rroble/modest-go/arrogant"
)

func main() {
	src := "<html><head></head><body><div>Hello World</div></body></html>"
	parser := arrogant.New()
	defer parser.Release()
	tree, err := parser.Parse(src)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	defer tree.Release()

	div, _ := tree.ByTagName("div")
	div.First().InnerText("Hello Go!")
	fmt.Printf("first div: '%s'\n", div.First())

	fmt.Printf("Document: '%s'\n", tree.Document())
}
