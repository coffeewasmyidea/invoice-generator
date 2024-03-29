package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var (
	ver string
	sha string
)

func version() string {
	if ver == "" {
		ver = "dev"
		sha = "unknown"
	}

	return fmt.Sprintf("invoice-generator version %s (built from sha %s)", ver, sha)
}

func main() {
	// Declare InvoiceData struct
	var invoice_data InvoiceData

	// Read TOML file
	toml_file, err := os.ReadFile("invoice-generator.toml")
	if err != nil {
		ok := YnPrompt("The file invoice-generator.toml was not found in the current directory. Want to create it?", true)
		if ok {
			createTOML()
		} else {
			fmt.Println("Ok, exiting..")
			os.Exit(0)
		}
	}

	toml_file, err = os.ReadFile("invoice-generator.toml")

	if err != nil {
		log.Fatal(err)
	}

	toml.Unmarshal(toml_file, &invoice_data)

	// Run cli
	CommandLineTool(invoice_data)
}
