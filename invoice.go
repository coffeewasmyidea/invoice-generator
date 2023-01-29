package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-pdf/fpdf"
)

func InvoiceGenerator(invoice_data InvoiceData) {
	titleStr := fmt.Sprintf("INVOICE #%s", invoice_data.InvoiceNumber())
	pdf := fpdf.New("P", "mm", "A4", "") // 210 x 297
	pdf.SetMargins(20, 20, 20)
	pdf.SetCellMargin(3)
	pdf.AddPage()

	cellGap := 2.0
	pdf.SetFont("helvetica", "", 24)
	pdf.Text(20, 30, titleStr)
	pdf.Ln(20)

	// From
	pdf.SetFont("helvetica", "BU", 11)
	pdf.CellFormat(90-cellGap-cellGap, 45, "From:", "1", 0, "LT", false, 0, "")

	// Fullname
	y := pdf.GetY()
	pdf.SetFont("helvetica", "", 11)
	pdf.SetXY(20, y+6)
	pdf.Cell(40, 10, invoice_data.Fullname)

	// MyAddress
	y = pdf.GetY()
	pdf.SetXY(20, y+8)
	pdf.MultiCell(90-cellGap-cellGap, 6, invoice_data.MyAddress, "0", "0", false)

	// MyEmail
	y = pdf.GetY()
	pdf.SetXY(20, y-2)
	pdf.Cell(40, 10, invoice_data.MyEmail)

	// Tax ID
	y = pdf.GetY()
	pdf.SetXY(20, y+8)
	pdf.Cell(40, 10, "TAX ID: "+invoice_data.TaxId)

	// To
	pdf.SetXY(20, 85)
	pdf.SetFont("helvetica", "BU", 11)
	pdf.CellFormat(90-cellGap-cellGap, 40, "Invoice for:", "1", 0, "LT", false, 0, "")

	// Invoice to
	y = pdf.GetY()
	pdf.SetXY(20, y+6)
	pdf.SetFont("helvetica", "", 11)
	pdf.Cell(40, 10, invoice_data.InvoiceTo)

	// Company address
	y = pdf.GetY()
	pdf.SetXY(20, y+8)
	pdf.MultiCell(90-cellGap-cellGap, 6, invoice_data.CompanyAddress, "0", "0", false)

	pdf.SetXY(106, 40)
	pdf.CellFormat(90-cellGap-cellGap, 45, "", "1", 0, "LT", false, 0, "")

	// Invoice number
	y = pdf.GetY()
	pdf.SetXY(106, y+6)
	pdf.Cell(40, 10, "Invoice number: "+invoice_data.InvoiceNumber())

	// Service Period
	y = pdf.GetY()
	pdf.SetXY(106, y+8)
	pdf.Cell(40, 10, "Service period: "+invoice_data.ServicePeriod())

	// Invoice Date
	y = pdf.GetY()
	pdf.SetXY(106, y+8)
	pdf.Cell(40, 10, "Invoice date: "+invoice_data.InvoiceDate())

	// Invoice due date
	y = pdf.GetY()
	pdf.SetXY(106, y+8)
	pdf.Cell(40, 10, "Invoice due date: "+invoice_data.InvoiceDueDate())

	// Empty
	pdf.SetXY(106, 85)
	pdf.CellFormat(90-cellGap-cellGap, 40, "", "1", 0, "LT", false, 0, "")

	pdf.SetXY(20, 138)

	// Colors, line width and bold font
	pdf.SetFillColor(73, 77, 99)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(96, 96, 96)
	pdf.SetLineWidth(.3)
	pdf.SetFont("helvetica", "B", 12)

	pdf.CellFormat(82, 11, "Service description", "", 0, "LM", true, 0, "")
	pdf.CellFormat(35, 11, "Price/mo", "", 0, "LM", true, 0, "")
	pdf.CellFormat(20, 11, "Qty", "", 0, "LM", true, 0, "")
	pdf.CellFormat(35, 11, "Total", "", 0, "LM", true, 0, "")

	// The last printed cell
	pdf.Ln(-1)

	// Color and font restoration
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetDrawColor(96, 96, 96)
	pdf.SetFont("helvetica", "", 11)
	list := pdf.SplitLines([]byte(invoice_data.getServiceDescription()), 87-cellGap-cellGap)

	// Service description
	pdf.Rect(20, 138, 82, 30+cellGap+cellGap, "D")
	y = pdf.GetY()
	cell_ht := float64(len(list)) * 6
	cellY := y + cellGap + (6-cell_ht)/2
	for splitJ := 0; splitJ < len(list); splitJ++ {
		pdf.SetXY(20+cellGap, cellY)
		pdf.CellFormat(82-cellGap-cellGap, 20, string(list[splitJ]), "", 0, "L", false, 0, "")
		cellY += 6
	}

	pdf.Rect(20, 138+34, 82, 8+cellGap+cellGap, "D")
	pdf.Rect(20, 138+34+12, 82, 8+cellGap+cellGap, "D")

	// Price
	pdf.Rect(20+82, 138, 35, 30+cellGap+cellGap, "D")
	y = pdf.GetY()
	pdf.SetXY(20+82, y)
	pdf.CellFormat(35-cellGap-cellGap, 14, invoice_data.getPrice(), "", 0, "L", false, 0, "")

	pdf.Rect(20+82, 138+34, 35, 8+cellGap+cellGap, "D")
	pdf.Rect(20+82, 138+34+12, 35, 8+cellGap+cellGap, "D")

	// Quantity
	pdf.Rect(20+82+35, 138, 20, 30+cellGap+cellGap, "D")
	y = pdf.GetY()
	pdf.SetXY(20+82+35, y)
	pdf.CellFormat(20-cellGap-cellGap, 14, fmt.Sprint(invoice_data.Quantity, ",00"), "", 0, "L", false, 0, "")

	pdf.Rect(20+82+35, 138+34, 20, 8+cellGap+cellGap, "D")

	// VAT
	pdf.SetXY(20+82+35, y+34/2)
	s := fmt.Sprintf("VAT(%d%s)", invoice_data.VAT, "%")
	pdf.CellFormat(20-cellGap-cellGap, 14, s, "", 0, "L", false, 0, "")

	pdf.Rect(20+82+35, 138+34+12, 20, 8+cellGap+cellGap, "D")
	pdf.SetXY(20+82+35, y)
	pdf.SetFont("helvetica", "B", 11)
	pdf.SetXY(20+82+35, (y+34/2)+12)
	pdf.CellFormat(20-cellGap-cellGap, 14, "Total", "", 0, "L", false, 0, "")

	// Total
	pdf.Rect(20+82+35+20, 138, 35, 30+cellGap+cellGap, "D")
	pdf.SetXY(20+82+35+20, 138+33/2)
	pdf.SetFont("helvetica", "", 11)
	pdf.CellFormat(35-cellGap-cellGap, 14, invoice_data.getTotal(), "", 0, "L", false, 0, "")

	y = pdf.GetY()
	pdf.Rect(20+82+35+20, 138+34, 35, 8+cellGap+cellGap, "D")
	pdf.SetXY(20+82+35+20, y+34/2)
	pdf.CellFormat(35-cellGap-cellGap, 14, invoice_data.getTotalVAT(), "", 0, "L", false, 0, "")

	pdf.Rect(20+82+35+20, 138+34+12, 35, 8+cellGap+cellGap, "D")
	pdf.SetXY(20+82+35+20, (y+34/2)+12)
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(35-cellGap-cellGap, 14, invoice_data.getTotalVATAmount(), "", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "", 11)

	y = pdf.GetY()
	pdf.SetXY(18, y+10)
	pdf.SetFont("helvetica", "", 12)
	pdf.Cell(18, 25, invoice_data.PaymentDetails)

	y = pdf.GetY()
	pdf.SetXY(18, y+8)
	pdf.SetFont("helvetica", "BU", 12)
	pdf.Cell(18, 30, "Payment details")

	pdf.SetFont("helvetica", "", 11)

	y = pdf.GetY()
	pdf.SetXY(18, y+23)

	pdf.Ln(1)
	pdf.CellFormat(50, 8, "Beneficiary name:", "1", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(124, 8, invoice_data.BeneficiaryName, "1", 1, "L", false, 0, "")
	pdf.SetFont("helvetica", "", 11)

	pdf.CellFormat(50, 8, "Beneficiary address:", "1", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(124, 8, invoice_data.BeneficiaryAddress, "1", 1, "L", false, 0, "")
	pdf.SetFont("helvetica", "", 11)

	pdf.CellFormat(50, 8, "Bank name:", "1", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(124, 8, invoice_data.BankName, "1", 1, "L", false, 0, "")
	pdf.SetFont("helvetica", "", 11)

	pdf.CellFormat(50, 8, "Bank Address:", "1", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(124, 8, invoice_data.BankAddress, "1", 1, "L", false, 0, "")
	pdf.SetFont("helvetica", "", 11)

	pdf.CellFormat(50, 8, "IBAN:", "1", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(124, 8, invoice_data.IBAN, "1", 1, "L", false, 0, "")
	pdf.SetFont("helvetica", "", 11)

	pdf.CellFormat(50, 8, "SWIFT/BIC:", "1", 0, "L", false, 0, "")
	pdf.SetFont("helvetica", "B", 11)
	pdf.CellFormat(124, 8, invoice_data.SwiftBIC, "1", 0, "L", false, 0, "")

	now := time.Now()
	year, month, _ := now.Date()
	location := now.Location()
	current_month := (time.Date(year, month, 1, 0, 0, 0, 0, location)).Format("02.01.06")
	invoice_number := strings.Replace(current_month, ".", "", 2)
	document_name := fmt.Sprintf("SE-%s.pdf", invoice_number)
	pdf.OutputFileAndClose(document_name)
}
