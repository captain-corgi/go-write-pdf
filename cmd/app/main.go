package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/signintech/gopdf"
	"github.com/unidoc/unioffice/document"
)

func replacePlaceholder(doc *document.Document, placeholder, replacement string) {
	for _, p := range doc.Paragraphs() {
		for _, run := range p.Runs() {
			text := run.Text()
			if strings.Contains(text, placeholder) {
				run.Clear()
				run.AddText(strings.ReplaceAll(text, placeholder, replacement))
			}
		}
	}
}

func generatePDFFromWordTemplate(templatePath, outputPath string, replacements map[string]string) error {
	doc, err := document.Open(templatePath)
	if err != nil {
		return err
	}

	// Replace placeholders in the Word document
	for placeholder, replacement := range replacements {
		replacePlaceholder(doc, placeholder, replacement)
	}

	// Create a new PDF document
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.276, H: 841.890}})
	pdf.AddPage()

	// Set font for the PDF
	pdf.AddTTFFont("Arial", "./font/arial.ttf")
	pdf.SetFont("Arial", "", 12)

	// Convert Word document content to PDF
	for _, p := range doc.Paragraphs() {
		pdf.SetX(20)
		pdf.SetY(pdf.GetY() + 10) // Adjust spacing between paragraphs

		for _, run := range p.Runs() {
			text := run.Text()
			pdf.Cell(nil, text)
		}
	}

	// Save the PDF to the specified output path
	err = pdf.WritePdf(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	templatePath := "./templates/word_template.docx"
	outputPath := "./output/output_pdf.pdf"

	// Replace these placeholders with actual data
	replacements := map[string]string{
		"{Name}":    "John Doe",
		"{Address}": "123 Main St",
		// Add more replacements as needed
	}

	err := generatePDFFromWordTemplate(templatePath, outputPath, replacements)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF created successfully at:", outputPath)
}
