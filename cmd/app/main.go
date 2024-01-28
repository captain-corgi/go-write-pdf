package report

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/signintech/gopdf"
)

// var NO_DATA_ERR_MSG string = "No Data"
func Generate(req model.ReportRequest, reportConfig model.ReportInfo, _ time.Time, generatorModel model.ReportGeneraterModel) error {
	requestDate, err := time.Parse("2006-01-02", req.Date)
	businessDate := requestDate.Add(-24 * time.Hour).Format("012006")
	reportInfoLst, err := r.repo.WhtRepository.GetWhtReportInfo(businessDate)
	filePath := filepath.Join(reportConfig.Path, generatorModel.FileName)
	pdf := initPDF()
	for _, reportInfo := range reportInfoLst {
		err = r.addPage(pdf, reportConfig)
		if err != nil {
			return err
		}
		err = processWhtReport(pdf, reportConfig, reportInfo, requestDate)
		if err != nil {
			return err
		}
	}
	err = pdf.WritePdf(filePath)
	if err != nil {
		return err
	}

	return nil
}

const (
	FontFreesiaUPCNormal = "FreesiaUPC"

	A4FullWidth  = 210 //mm
	A4FullHeight = 297 //mm

	MarginRight  = 25.4 //mm
	MarginTop    = 25.4 //mm
	MarginBottom = 25.4 //mm

	FontSizeNormal = 14 //pt

	TabSpace              = 12.5              //mm
	LineHeight            = 10                //pt
	ParagraphSpacingAfter = 1.15 * LineHeight //pt
	MLeft                 = 10.0
	MRight                = 10.0
	TableWidth            = float64(A4FullWidth - (MLeft + MRight)) //float64(A4FullWidth - (MLeft + MarginRight))
)

func initPDF() *gopdf.GoPdf {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: gopdf.Rect{W: A4FullWidth, H: A4FullHeight},
		Unit:     gopdf.UnitMM,
	})
	return pdf
}

func addPage(pdf *gopdf.GoPdf, reportConfig model.ReportInfo) error {
	pdf.AddPage()

	// Add the font from source: https://www.wfonts.com/
	err := pdf.AddTTFFont(FontFreesiaUPCNormal, reportConfig.FontSource)
	if err != nil {
		return err
	}

	err = pdf.SetFont(FontFreesiaUPCNormal, "", FontSizeNormal)
	if err != nil {
		return err
	}

	pdf.SetMargins(MLeft, MarginTop, MarginRight, MarginBottom)
	pdf.SetX(0)
	pdf.SetY(0)

	return nil
}

func processWhtReport(
	pdf *gopdf.GoPdf,
	reportConfig model.ReportInfo,
	reportInfo model.WhtInfo,
	requestDate time.Time) error {

	// Page Header
	//pdf.Br(4 * LineHeight)
	pdf.SetX(90)
	pdf.SetY(MarginTop)

	pdf.Cell(nil, fmt.Sprintf("%s %s %s", "ประจำเดือน", nowThaiMonth, nowThaiYear))
	pdf.Br(ParagraphSpacingAfter)

	pdf.Cell(nil, fmt.Sprintf("%s%60s%s", "ส่วนบริการชำระเงิน", " ", "หมายเลขสังกัด 05531 Cash Management"))
	pdf.Br(ParagraphSpacingAfter)

	pdf.Cell(nil, "ฝ่ายการให้บริการชำระเงิน")
	pdf.Br(ParagraphSpacingAfter)

	// Table section
	// Table header
	tmpX := pdf.GetX()
	tmpY := pdf.GetY()
	colWidth := 38.2
	height := float64(LineHeight)
	x := MLeft
	err := writeTableHeader(pdf, tmpX, tmpY, colWidth, height, FontSizeNormal)
	if err != nil {
		return err
	}

	// Table row
	tmpX = MLeft
	tmpY = pdf.GetY()
	tableRow := 18
	tableCol := 4

	tableData := [18][4]string{ // 2 row, 4 col
		{"10/11/2566", "E-WITHHOLDING TAX (CROSS BORDER)", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"10/11/2566", "E-WITHHOLDING TAX (CROSS BORDER)", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"10/11/2566", "E-WITHHOLDING TAX (CROSS BORDER)", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"15/11/2566", "WITHHOLDING TAX REFUND (DIRECT DEBIT/BULK GATEWAY)", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"15/11/2566", "WITHHOLDING TAX REFUND (DIRECT DEBIT/BULK GATEWAY)", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"15/11/2566", "MITSUBISHI CHEMICAL (THAILAND) CO.,LTD.", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"15/11/2566", "THE MARKETPLACE CO.,LTD. (PPA)", "5,000,000,000,000.88", "7,000,000,000,000.12"},
		{"15/11/2566", "THE MARKETPLACE CO.,LTD. (PPA)", "7,935.12", "7,000,000,000,000.12"},
		{"15/11/2566", "BAANPLOY CONSTRUCTION AND ALUMINUM CO.,LTD.", "1,800.01", "7,000,000,000,000.12"},
		{"15/11/2566", "BAANPLOY CONSTRUCTION AND ALUMINUM CO.,LTD.", "7,935.12", "7,000,000,000,000.12"},

		{"31/11/2566", "MISS NAROEMOL KRATHINTHONG (BG)              (PPD)", "1,800.01", "7,000,000,000,000.12"},
		{"31/11/2566", "MISS NAROEMOL KRATHINTHONG (BG)              (PPD)", "7,935.12", "7,000,000,000,000.12"},
		{"31/11/2566", "KRUNGTHAI ZMICO SECURITIES COMPANY LIMITED", "1,800.01", "7,000,000,000,000.12"},
		{"31/11/2566", "KRUNGTHAI ZMICO SECURITIES COMPANY LIMITED", "7,935.12", "7,000,000,000,000.12"},
		{"31/11/2566", "MEESUPPAISARN ENGINEERING LIMITED PA", "1,800.01", "7,000,000,000,000.12"},
		{"31/11/2566", "LAYTEX (THAILAND) CO.,LTD.(BG)", "7,935.12", "7,000,000,000,000.12"},
		{"31/11/2566", "NETSOL TECHNOLOGES (THAILAND) CO.,LTD. (BG)", "1,800.01", "7,000,000,000,000.12"},
		{"31/11/2566", "NETWORK INTEGRATION SOLUTIONS CO.,LTD. (PPA)", "7,935.12", "7,000,000,000,000.12"},
	}

	fmt.Println("TableWidth: ", TableWidth)
	for i := 0; i < tableRow; i++ {
		for j := 0; j < tableCol; j++ {
			text := tableData[i][j]
			if j == 0 {
				// Column 1
				colWidth = 0.13 * TableWidth
				fmt.Println("colWidth 1: ", colWidth)
			} else if j == 1 {
				// Column 2
				colWidth = 0.47 * TableWidth
				fmt.Println("colWidth 2: ", colWidth)
			} else {
				// Column 3,4
				colWidth = 0.2 * TableWidth
			}
			if err := fillPdfRect(pdf, text, tmpX, tmpY, colWidth, LineHeight, FontSizeNormal); err != nil {
				fmt.Println(err)
			}
			tmpX += colWidth
		}
		tmpX = x
		tmpY += LineHeight
	}

	return nil
}

func fillPdfRect(pdf *gopdf.GoPdf, text string, x, y, w, h float64, fontSize int) (err error) {
	pdf.SetLineWidth(0.2)
	pdf.SetMarginTop(10.0)
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "D")
	textw, err := pdf.MeasureTextWidth(text)
	if err != nil {
		return
	}

	x = x + (w / 2) - (textw / 2)
	pdf.SetX(x)
	y = y + (h / 2) - (float64(fontSize) / 2)
	pdf.SetY(y)

	if err = pdf.CellWithOption(nil, text, gopdf.CellOption{CoefUnderlinePosition: 100}); err != nil {
		return
	}
	return
}

func writeTableHeader(pdf *gopdf.GoPdf, tmpX, tmpY, colWidth, height float64, fontSize int) (err error) {
	// 1st line
	text := "รายการที่ 3 รายละเอียดภาษีเงินได้นิติบุคคลหัก ณ ที่จ่าย ตามมาตรา 3 เตรส"
	if err = fillPdfRect(pdf, text, tmpX, tmpY, TableWidth, height, fontSize); err != nil {
		fmt.Println(err)
	}
	tmpX = MLeft
	tmpY += height

	// 2nd line
	tableRow := 2
	tableCol := 4
	tableData := [2][4]string{ // 2 row, 4 col
		{"ว/ด/ป", "ชื่อหน่วยงานผู้มีหน้าที่หักภาษี ณ ที่จ่าย", "  ", "ตามมาตรา 3 เตรส"},
		{"     ", "      ", "จำนวนเงิน", "จำนวนภาษีที่หัก"},
	}
	for i := 0; i < tableRow; i++ {
		for j := 0; j < tableCol; j++ {
			text := tableData[i][j]
			if j == 0 {
				// Column 1
				colWidth = 0.13 * TableWidth
			} else if j == 1 {
				// Column 2
				colWidth = 0.47 * TableWidth
			} else {
				// Column 3,4
				colWidth = 0.2 * TableWidth
			}
			if err = fillPdfRect(pdf, text, tmpX, tmpY, colWidth, height, fontSize); err != nil {
				fmt.Println(err)
			}
			tmpX += colWidth
		}
		tmpX = MLeft
		tmpY += height
	}
	pdf.SetX(tmpX)
	pdf.SetY(tmpY)
	return
}
