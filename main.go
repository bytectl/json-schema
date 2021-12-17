package main

import (
	"flag"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

var (
	schema   = ""
	document = ""
)

func main() {
	flag.StringVar(&schema, "s", "file://./schema.json", "schema file or url")
	flag.StringVar(&document, "d", "file://./document.json", "document file or url")
	flag.Parse()

	schemaLoader := gojsonschema.NewReferenceLoader(schema)
	documentLoader := gojsonschema.NewReferenceLoader(document)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
