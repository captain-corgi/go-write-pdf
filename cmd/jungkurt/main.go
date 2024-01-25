package main

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func generatePDF(outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font for the title
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(0, 0, 0)

	// Set up header function
	header := func() {
		pdf.SetY(10) // Adjust the Y position as needed
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(0, 10, "My Header", "0", 0, "C", false, 0, "")
	}

	// Set the header function
	pdf.SetHeaderFunc(func() {
		header()
	})

	// Add a big title at the top of the first page
	title := "My Big Title"
	pdf.CellFormat(0, 10, title, "", 0, "C", false, 0, "")
	pdf.Ln(20)

	// Set font for the table
	pdf.SetFont("Arial", "", 12)

	// Get page width
	pageWidth, _ := pdf.GetPageSize()

	// Add margins
	leftMargin := 20
	rightMargin := 20

	// Add table headers
	pdf.SetFillColor(200, 220, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(0.1)

	headers := []string{"Date", "Column 2", "Column 3", "Column 4"}
	colWidths := calculateColumnWidths(pdf, headers)

	// Calculate the total width of the table
	totalWidth := sumFloat64Slice(colWidths)

	// Dynamically adjust the width of the second column to occupy remaining space
	remainingWidth := pageWidth - totalWidth - float64(leftMargin+rightMargin)
	colWidths[1] = remainingWidth

	// Calculate left margin to center the table
	leftMargin = int((pageWidth - totalWidth) / 2)

	pdf.SetX(float64(leftMargin))

	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 7, header, "1", 0, "C", true, 0, "")
	}

	pdf.Ln(-1)

	// Add table data with 100 rows
	data := make([][]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = []string{
			fmt.Sprintf("Data %d", i+1),
			fmt.Sprintf("Data %d", i+2),
			fmt.Sprintf("Data %d", i+3),
			fmt.Sprintf("Data %d", i+4),
		}
	}

	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(0, 0, 0)

	for _, row := range data {
		pdf.SetX(float64(leftMargin))
		for i, cell := range row {
			pdf.CellFormat(colWidths[i], 10, cell, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
	}

	// Add a signing row at the bottom of the page
	pdf.Ln(10) // Add space between the table and signing row
	signingText := "Signed by: ____________________"
	pdf.CellFormat(0, 10, signingText, "", 0, "L", false, 0, "")

	// Save the PDF to the specified output path
	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func calculateColumnWidths(pdf *gofpdf.Fpdf, headers []string) []float64 {
	colWidths := make([]float64, len(headers))
	for i, header := range headers {
		colWidths[i] = pdf.GetStringWidth(header) + 6 // Add padding
	}
	return colWidths
}

func sumFloat64Slice(slice []float64) float64 {
	sum := 0.0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func main() {
	outputPath := "./output/output_gofpdf.pdf"

	err := generatePDF(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF created successfully at:", outputPath)
}
