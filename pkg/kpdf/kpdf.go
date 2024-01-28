package kpdf

import "github.com/signintech/gopdf"

type (
	Data struct {
		Text   string
		Length float64
	}
	Employee struct {
		ID          Data
		Name        Data
		PhoneNumber Data
		Email       Data
	}
)

func GeneratePDFNew(outputPath string) error {
	fontSize := DefaultFontSize

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.276, H: 841.890}}) // A4 dimensions in points
	pdf.AddPage()
	pdf.AddTTFFont("Arial", "./font/arial.ttf")
	pdf.SetFont("Arial", "", fontSize)

	// Create new drawer
	drawer := NewDrawer(&pdf)

	// Header
	if err := drawer.DrawSingleCellLn("This is a TITLE", fontSize, false, 0); err != nil {
		return err
	}

	// Table header
	header := []Data{
		{"No.", 50},
		{"Name", 150},
		{"Phone number", 150},
		{"Email", 200},
	}
	for _, h := range header {
		if err := drawer.DrawSingleCellCol(h.Text, fontSize, true, h.Length); err != nil {
			return err
		}
	}
	drawer.NextLn()

	// Table body
	body := []Employee{
		{
			Data{"1", header[0].Length},
			Data{"Nguyen Van A", header[1].Length},
			Data{"0123 456 789", header[2].Length},
			Data{"a.nguyen.van@mailmail.com", header[3].Length},
		},
		{
			Data{"2", header[0].Length},
			Data{"Nguyen Van B", header[1].Length},
			Data{"0123 456 789", header[2].Length},
			Data{"c.nguyen.van@mailmail.com", header[3].Length},
		},
		{
			Data{"3", header[0].Length},
			Data{"Nguyen Van C", header[1].Length},
			Data{"0123 456 789", header[2].Length},
			Data{"b.nguyen.van@mailmail.com", header[3].Length},
		},
		{
			Data{"4", header[0].Length},
			Data{"Nguyen Van D", header[1].Length},
			Data{"0123 456 789", header[2].Length},
			Data{"d.nguyen.van@mailmail.com", header[3].Length},
		},
	}
	for _, emp := range body {
		if err := drawer.DrawSingleCellCol(emp.ID.Text, fontSize, true, emp.ID.Length); err != nil {
			return err
		}
		if err := drawer.DrawSingleCellCol(emp.Name.Text, fontSize, true, emp.Name.Length); err != nil {
			return err
		}
		if err := drawer.DrawSingleCellCol(emp.PhoneNumber.Text, fontSize, true, emp.PhoneNumber.Length); err != nil {
			return err
		}
		if err := drawer.DrawSingleCellCol(emp.Email.Text, fontSize, true, emp.Email.Length); err != nil {
			return err
		}
		drawer.NextLn()
	}

	// Footer
	if err := drawer.DrawSingleCellLn("5. Some footer information here", fontSize, false, 595.276); err != nil {
		return err
	}

	// Save the PDF to the specified output path
	if err := pdf.WritePdf(outputPath); err != nil {
		return err
	}

	return nil
}
