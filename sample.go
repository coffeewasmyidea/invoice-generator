package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func YnPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

func createTOML() {
	f, err := os.Create("invoice-generator.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(TemplateTOML)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("invoice-generator.toml created!")

}

var TemplateTOML = `# Invoice data with SE-DDMMYY number and due date +10 days
Fullname = "Bob Sanders"
MyAddress = """Lorem ipsum amet, 221, 42 
Aenean ligula, 129104 City, Country."""

MyEmail = "hello@8am.dev"
TaxId = "400424221"
InvoiceTo = "World JSC"
CompanyAddress = """
c/o Alice Henderson
Nam quam nunc, blandit vel 11
Sed fringilla mauris, 68
409118, Middle-earth"""

ServiceDescription = "Backend development services for the period of {xx.xx.xxxx} to {yy.yy.yyyy}"
Price = 1500
Currency = "USD"
Quantity = 2
VAT = 10
PaymentDetails = "Please transfer the due amount according to the following payment details."
BeneficiaryName = "IE Bob Sanders"
BeneficiaryAddress = "FARFROMPOOPEN ROAD TURN 5, CITY, COUNTRY"
BankName = "Bank of the Middle-earth"
BankAddress = "42c GAGARIN STREET, Mars 0160, Middle-earth"
IBAN = "XXXXLX11100003168990YY"
SwiftBIC = "XXXXXX11"

`
