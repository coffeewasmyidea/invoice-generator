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

	// Run cli
	CommandLineTool(invoice_data)

	// Read TOML file
	toml_file, err := os.ReadFile("invoice-generator.toml")
	if err != nil {
		log.Fatal(err, ". You need to download the invoice-generator.toml from here https://raw.githubusercontent.com/coffeewasmyidea/invoice-generator/main/invoice-generator.toml.example")
	}

	toml.Unmarshal(toml_file, &invoice_data)
}
