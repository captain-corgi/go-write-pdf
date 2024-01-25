package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

func generatePDF(outputPath string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.276, H: 841.890}}) // A4 dimensions in points
	pdf.AddPage()

	// Set font for the title
	pdf.AddTTFFont("Arial", "./font/arial.ttf")
	pdf.SetFont("Arial", "B", 16)

	// Add a big title at the top of the page
	title := "My Big Title"
	pdf.SetX(20)
	pdf.SetY(20)
	pdf.Cell(gopdf.PageSizeA4, title)

	// Set font for the table
	pdf.AddTTFFont("Arial", "./font/arial.ttf")
	pdf.SetFont("Arial", "", 12)

	// Add table headers
	pdf.SetFillColor(200, 220, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetLineWidth(0.1)

	headers := []string{"Date", "Column 2", "Column 3", "Column 4"}
	colWidths := calculateColumnWidthsGopdf(&pdf, headers)

	// Calculate left margin to center the table
	leftMargin := 20

	for i, header := range headers {
		pdf.RectFromUpperLeftWithStyle(float64(leftMargin), 60, colWidths[i], 15, "fill:rgb(200,200,200);stroke:rgb(0,0,0)")
		pdf.SetX(float64(leftMargin + 5))
		pdf.SetY(65)
		pdf.Cell(nil, header)
		leftMargin += int(colWidths[i]) // Adjust for cell width
	}

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

	y := 80
	for _, row := range data {
		leftMargin = 20
		for i, value := range row {
			pdf.RectFromUpperLeftWithStyle(float64(leftMargin), float64(y), colWidths[i], 15, "fill:none;stroke:rgb(0,0,0)")
			pdf.SetX(float64(leftMargin + 5))
			pdf.SetY(float64(y + 5))
			pdf.Cell(nil, value)
			leftMargin += int(colWidths[i]) // Adjust for cell width
		}
		y += 15 // Adjust for line height
	}

	// Add a signing row at the bottom of the page
	y += 20 // Add space between the table and signing row
	signingText := "Signed by: ____________________"
	pdf.SetX(20)
	pdf.SetY(float64(y))
	pdf.Cell(nil, signingText)

	// Save the PDF to the specified output path
	err := pdf.WritePdf(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func calculateColumnWidthsGopdf(pdf *gopdf.GoPdf, headers []string) []float64 {
	colWidths := make([]float64, len(headers))
	for i, header := range headers {
		width, _ := pdf.MeasureTextWidth(header)
		colWidths[i] = width + 10 // Add padding
	}
	return colWidths
}

func main() {
	outputPath := "./output/output_gopdf.pdf"

	err := generatePDF(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF created successfully at:", outputPath)
}
