package main

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml/v2"
)

func main() {
	// Declare InvoiceData struct
	var invoice_data InvoiceData

	// Read TOML file
	toml_file, err := ioutil.ReadFile("invoice-generator.toml")
	if err != nil {
		log.Fatal(err)
	}

	toml.Unmarshal(toml_file, &invoice_data)

	// Run cli
	CommandLineTool(invoice_data)

}
