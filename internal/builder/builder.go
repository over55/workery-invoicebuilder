package builder

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
)

type PDFBuilder struct {
}

func New(filePath string) (*PDFBuilder, error) {
	var err error

	pdf := gofpdf.New("P", "mm", "A4", "")
	tpl1 := gofpdi.ImportPage(pdf, filePath, 1, "/MediaBox")

	pdf.AddPage()

	// Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 0, 210, 0)

	pdf.SetFont("Helvetica", "", 20)
	pdf.Cell(111, 111, "Import existing PDF into gofpdf document with gofpdi")

	err = pdf.OutputFileAndClose("example.pdf")
	if err != nil {
		return nil, err
	}

	return &PDFBuilder{}, err
}
