package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/leekchan/accounting"
)

type InvoiceData struct {
	Fullname           string
	MyAddress          string
	MyEmail            string
	TaxId              string
	InvoiceTo          string
	CompanyAddress     string
	ServiceDescription string
	Currency           string
	Price              int
	FormattedPrice     string
	Quantity           int
	VAT                int
	PaymentDetails     string
	BeneficiaryName    string
	BeneficiaryAddress string
	BankName           string
	BankAddress        string
	IBAN               string
	SwiftBIC           string
}

func (invoice *InvoiceData) InvoiceNumber() string {
	now := time.Now()
	year, month, _ := now.Date()
	location := now.Location()
	current_month := (time.Date(year, month, 1, 0, 0, 0, 0, location)).Format("02.01.06")
	s := strings.Replace(current_month, ".", "", 2)
	return fmt.Sprintf("SE-%s", s)
}

func (invoice *InvoiceData) InvoiceDate() string {
	now := time.Now()
	year, month, day := now.Date()
	location := now.Location()
	current_date := (time.Date(year, month, day, 0, 0, 0, 0, location)).Format("02.01.2006")
	return fmt.Sprintf("%s", current_date)
}

func (invoice *InvoiceData) InvoiceDueDate() string {
	now := time.Now()
	year, month, day := now.Date()
	location := now.Location()
	current_date := (time.Date(year, month, day+10, 0, 0, 0, 0, location)).Format("02.01.2006")
	return fmt.Sprintf("%s", current_date)
}

func (invoice *InvoiceData) ServicePeriod() string {
	now := time.Now()
	year, month, day := now.Date()
	location := now.Location()
	current_date := time.Date(year, month, day, 0, 0, 0, 0, location)
	return fmt.Sprintf("%s, %d", current_date.Month(), current_date.Year())
}

func (invoice *InvoiceData) getServiceDescription() string {
	now := time.Now()
	year, month, _ := now.Date()
	location := now.Location()
	start := time.Date(year, month, 1, 0, 0, 0, 0, location)
	end := start.AddDate(0, 1, -1)
	start_date := start.Format("02.01.2006")
	end_date := end.Format("02.01.2006")
	s := strings.Replace(invoice.ServiceDescription, "{xx.xx.xxxx}", start_date, 1)
	return strings.Replace(s, "{yy.yy.yyyy}", end_date, 1)
}

func (invoice *InvoiceData) getPrice() string {
	if invoice.Currency != "EUR" && invoice.Currency != "USD" {
		panic("works only with EUR(€) or USD($)")
	}
	ac := accounting.Accounting{Symbol: invoice.Currency + " ", Precision: 2}
	return ac.FormatMoney(invoice.Price)

}

func (invoice *InvoiceData) getTotalVAT() string {
	total := invoice.Price * invoice.Quantity
	s := total / 100 * invoice.VAT
	ac := accounting.Accounting{Symbol: invoice.Currency + " ", Precision: 2}
	return ac.FormatMoney(s)

}

func (invoice *InvoiceData) getTotal() string {
	total := invoice.Price * invoice.Quantity
	ac := accounting.Accounting{Symbol: invoice.Currency + " ", Precision: 2}
	return ac.FormatMoney(total)

}

func (invoice *InvoiceData) getTotalVATAmount() string {
	total := invoice.Price * invoice.Quantity
	total_vat := total / 100 * invoice.VAT
	ac := accounting.Accounting{Symbol: invoice.Currency + " ", Precision: 2}
	return ac.FormatMoney(total - total_vat)

}
