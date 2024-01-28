package main

import (
	"fmt"
	"log"

	"github.com/captain-corgi/go-write-pdf/pkg/kpdf"
	"github.com/signintech/gopdf"
)

func main() {
	outputPath := "./output/output_gopdf.pdf"
	outputPathChatGpt := "./output/output_gopdf_chatgpt.pdf"

	err := kpdf.GeneratePDFNew(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	err = generatePDF(outputPathChatGpt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF created successfully at:", outputPath)
}

func generatePDF(outputPath string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.276, H: 841.890}})
	pdf.AddPage()

	// Set font for the table
	pdf.AddTTFFont("Arial", "./font/arial.ttf")
	pdf.SetFont("Arial", "", 12)

	// Draw table headers
	drawCell(&pdf, 20, 60, 100, 30, "Header 1")
	drawCell(&pdf, 120, 60, 100, 30, "Header 2")
	drawCell(&pdf, 220, 60, 100, 30, "Header 3")

	// Draw regular cells
	drawCell(&pdf, 20, 90, 100, 30, "Data 1")
	drawCell(&pdf, 120, 90, 100, 30, "Data 2")

	// Merge down effect (draw the merged cell that spans down)
	drawMergedDownCell(&pdf, 220, 90, 100, 60, "Merged Down Cell", 2)

	// Draw additional regular cells
	drawCell(&pdf, 20, 120, 100, 30, "Data 3")
	drawCell(&pdf, 120, 120, 100, 30, "Data 4")
	drawCell(&pdf, 220, 150, 100, 30, "Data 5")

	// Draw three more rows
	drawCell(&pdf, 20, 180, 100, 30, "Data 6")
	drawCell(&pdf, 120, 180, 100, 30, "Data 7")
	drawCell(&pdf, 220, 180, 100, 30, "Data 8")

	// Merge cells in the middle of the table
	drawMergedCell(&pdf, 120, 210, 100, 60, "Merged Cell", 2)

	// Draw additional regular cells
	drawCell(&pdf, 20, 240, 100, 30, "Data 9")
	drawCell(&pdf, 220, 240, 100, 30, "Data 10")

	// Save the PDF to the specified output path
	err := pdf.WritePdf(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func drawMergedCell(pdf *gopdf.GoPdf, x, y, width, height float64, text string, colspan int) {
	// Draw the merged cell as a single rectangle
	pdf.RectFromUpperLeft(x, y, width, height)
	pdf.SetX(x + 5)
	pdf.SetY(y + 5)
	pdf.Cell(nil, text)

	// Draw vertical lines to separate the merged cells
	for i := 1; i < colspan; i++ {
		lineX := x + (float64(i) * width / float64(colspan))
		pdf.Line(lineX, y, lineX, y+height)
	}
}

func drawCell(pdf *gopdf.GoPdf, x, y, width, height float64, text string) {
	pdf.RectFromUpperLeft(x, y, width, height)
	pdf.SetX(x + 5)
	pdf.SetY(y + 5)
	pdf.Cell(nil, text)
}

func drawMergedDownCell(pdf *gopdf.GoPdf, x, y, width, height float64, text string, rowspan int) {
	// Draw the merged cell as a single rectangle
	pdf.RectFromUpperLeft(x, y, width, height)
	pdf.SetX(x + 5)
	pdf.SetY(y + 5)
	pdf.Cell(nil, text)

}
