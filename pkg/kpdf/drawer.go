package kpdf

import (
	"github.com/signintech/gopdf"
)

type (
	Drawer interface {
		NextCol()
		NextLn()

		DrawSingleCell(data string, fontSize int, hasBorder bool, width float64) (float64, float64, error)
		DrawSingleCellLn(data string, fontSize int, hasBorder bool, width float64) error
		DrawSingleCellCol(data string, fontSize int, hasBorder bool, width float64) error
	}
	DrawerImpl struct {
		pdf *gopdf.GoPdf

		x, y, w, h                              float64
		marginL, marginR, marginT, marginBT     float64
		paddingL, paddingR, paddingT, paddingBT float64

		textWidth, textHeight float64
		fontSize              int

		pageWidth, pageHeight float32
	}
)

func NewDrawer(pdf *gopdf.GoPdf) Drawer {
	return &DrawerImpl{
		pdf: pdf,

		marginL:  DefaultMarginLeft,
		marginR:  DefaultMarginLeft,
		marginT:  DefaultMarginLeft,
		marginBT: DefaultMarginLeft,

		paddingL:  DefaultPadding,
		paddingR:  DefaultPadding,
		paddingT:  DefaultPadding,
		paddingBT: DefaultPadding,

		fontSize: DefaultFontSize,
	}
}

func (r *DrawerImpl) NextCol() {
	r.x += r.w + r.paddingL + r.textWidth + r.paddingR
}

func (r *DrawerImpl) NextLn() {
	r.x = 0
	r.y += r.h + r.paddingT + r.textHeight + r.paddingBT
}

func (r *DrawerImpl) DrawSingleCell(data string, fontSize int, hasBorder bool, width float64) (float64, float64, error) {
	var (
		x, y, w, h   float64
		textX, textY float64
	)

	// Draw border
	var (
		textWidth, textHeight float64
		err                   error
	)
	if textWidth, err = r.pdf.MeasureTextWidth(data); err != nil {
		return w, h, err
	}
	if textHeight, err = r.pdf.MeasureCellHeightByText(data); err != nil {
		return w, h, err
	}
	x = r.x + r.marginL
	y = r.y + r.marginT
	if width == 0 {
		w = r.w + r.paddingL + textWidth + r.paddingR
	} else {
		// TODO: Support fill page width later
		w = width
	}
	h = r.h + r.paddingT + textHeight + r.paddingBT

	if hasBorder {
		r.pdf.RectFromUpperLeftWithStyle(x, y, w, h, DefaultTableFormat)
	}

	// Draw text
	textX = x + r.paddingL
	textY = y + r.paddingT
	if data != "" {
		// Fill text
		r.pdf.SetX(textX)
		r.pdf.SetY(textY)
		r.pdf.Cell(gopdf.PageSizeA4, data)
	}

	// Store existing text info
	r.textWidth = textWidth
	r.textHeight = textHeight

	return w, h, nil
}

func (r *DrawerImpl) DrawSingleCellCol(data string, fontSize int, hasBorder bool, width float64) error {
	w, _, err := r.DrawSingleCell(data, fontSize, hasBorder, width)
	if err != nil {
		return err
	}

	// Next Col
	r.x += w
	return nil
}

func (r *DrawerImpl) DrawSingleCellLn(data string, fontSize int, hasBorder bool, width float64) error {
	_, h, err := r.DrawSingleCell(data, fontSize, hasBorder, width)
	if err != nil {
		return err
	}

	// Next Ln
	r.x = 0
	r.y += h
	return nil
}
