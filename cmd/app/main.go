package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/signintech/gopdf"
)

func generatePDFFromHTMLTemplate(templatePath, outputPath string, data map[string]interface{}) error {
	// Read the HTML template file
	htmlTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Execute the template and write the result to a string
	htmlContent := new(strings.Builder)
	err = htmlTemplate.Execute(htmlContent, data)
	if err != nil {
		return err
	}

	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to a data URL with the HTML content
	url := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(htmlContent.String()))
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Set the viewport size to capture the entire content
			err := emulation.SetDeviceMetricsOverride(1200, 800, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	)
	if err != nil {
		return err
	}

	// Capture a screenshot of the rendered content
	var screenshot []byte
	err = chromedp.Run(ctx, chromedp.CaptureScreenshot(&screenshot))
	if err != nil {
		return err
	}

	// Create a new PDF document
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.276, H: 841.890}}) // A4 dimensions in points
	pdf.AddPage()

	// Add the screenshot image to the PDF
	image, err := gopdf.ImageHolderByBytes(screenshot)
	if err != nil {
		return err
	}
	pdf.ImageByHolder(image, 0, 0, nil)

	// Save the PDF to the specified output path
	err = pdf.WritePdf(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Path to your HTML template file
	templatePath := "./template/template.html"

	// Path to the output PDF file
	outputPath := "./output/output_pdf.pdf"

	// Data to replace placeholders in the HTML template
	data := map[string]interface{}{
		"Name":  "Le Tuan Anh",
		"Email": "letuananh@captain.corgi.com",
		// Add more data as needed
	}

	err := generatePDFFromHTMLTemplate(templatePath, outputPath, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF created successfully at:", outputPath)
}
