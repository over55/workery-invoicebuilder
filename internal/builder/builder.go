package builder

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
)

type PDFBuilder struct {
}

func New(filePath string, m *WorkOrderInvoice) (*PDFBuilder, error) {
	var err error

	pdf := gofpdf.New("P", "mm", "A4", "")
	tpl1 := gofpdi.ImportPage(pdf, filePath, 1, "/MediaBox")

	pdf.AddPage()

	// Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 0, 210, 0)

	pdf.SetFont("Courier", "", 11)

	pdf.SetXY(159, 19)
	pdf.Cell(0, 0, m.InvoiceID)

	pdf.SetFont("Helvetica", "", 11)

	//
	// Header
	//

	pdf.SetXY(116, 25)
	pdf.Cell(0, 0, m.InvoiceDate)
	pdf.SetXY(136, 32)
	pdf.Cell(0, 0, m.AssociateName)
	pdf.SetXY(136, 39)
	pdf.Cell(0, 0, m.AssociateTelephone)
	pdf.SetXY(46, 46)
	pdf.Cell(0, 0, m.ClientName)
	pdf.SetXY(36, 53)
	pdf.Cell(0, 0, m.ClientAddress)
	pdf.SetXY(33, 60)
	pdf.Cell(0, 0, m.ClientTelephone)
	pdf.SetXY(92, 60)
	pdf.Cell(0, 0, m.ClientEmail)

	//
	// Line #1
	//

	pdf.SetXY(30, 72)
	pdf.Cell(0, 0, m.Line01Qty)
	pdf.SetXY(49, 72)
	pdf.Cell(0, 0, m.Line01Desc)
	pdf.SetXY(136, 72)
	pdf.Cell(0, 0, m.Line01Price)
	pdf.SetXY(169, 72)
	pdf.Cell(0, 0, m.Line01Amount)

	//
	// Line #2
	//

	pdf.SetXY(30, 78)
	pdf.Cell(0, 0, m.Line02Qty)
	pdf.SetXY(49, 78)
	pdf.Cell(0, 0, m.Line02Desc)
	pdf.SetXY(136, 78)
	pdf.Cell(0, 0, m.Line02Price)
	pdf.SetXY(169, 78)
	pdf.Cell(0, 0, m.Line02Amount)

	//
	// Line #3
	//

	pdf.SetXY(30, 83)
	pdf.Cell(0, 0, m.Line03Qty)
	pdf.SetXY(49, 83)
	pdf.Cell(0, 0, m.Line03Desc)
	pdf.SetXY(136, 83)
	pdf.Cell(0, 0, m.Line03Price)
	pdf.SetXY(169, 83)
	pdf.Cell(0, 0, m.Line03Amount)

	//
	// Line #4
	//

	pdf.SetXY(30, 88)
	pdf.Cell(0, 0, m.Line04Qty)
	pdf.SetXY(49, 88)
	pdf.Cell(0, 0, m.Line04Desc)
	pdf.SetXY(136, 88)
	pdf.Cell(0, 0, m.Line04Price)
	pdf.SetXY(169, 88)
	pdf.Cell(0, 0, m.Line04Amount)

	//
	// Line #5
	//

	pdf.SetXY(30, 93)
	pdf.Cell(0, 0, m.Line05Qty)
	pdf.SetXY(49, 93)
	pdf.Cell(0, 0, m.Line05Desc)
	pdf.SetXY(136, 93)
	pdf.Cell(0, 0, m.Line05Price)
	pdf.SetXY(169, 93)
	pdf.Cell(0, 0, m.Line05Amount)

	//
	// Line #6
	//

	pdf.SetXY(30, 98)
	pdf.Cell(0, 0, m.Line06Qty)
	pdf.SetXY(49, 98)
	pdf.Cell(0, 0, m.Line06Desc)
	pdf.SetXY(136, 98)
	pdf.Cell(0, 0, m.Line06Price)
	pdf.SetXY(169, 98)
	pdf.Cell(0, 0, m.Line06Amount)

	//
	// Line #7
	//

	pdf.SetXY(30, 103)
	pdf.Cell(0, 0, m.Line07Qty)
	pdf.SetXY(49, 103)
	pdf.Cell(0, 0, m.Line07Desc)
	pdf.SetXY(136, 103)
	pdf.Cell(0, 0, m.Line07Price)
	pdf.SetXY(169, 103)
	pdf.Cell(0, 0, m.Line07Amount)

	//
	// Line #8
	//

	pdf.SetXY(30, 109)
	pdf.Cell(0, 0, m.Line08Qty)
	pdf.SetXY(49, 109)
	pdf.Cell(0, 0, m.Line08Desc)
	pdf.SetXY(136, 109)
	pdf.Cell(0, 0, m.Line08Price)
	pdf.SetXY(169, 109)
	pdf.Cell(0, 0, m.Line08Amount)

	//
	// Line #9
	//

	pdf.SetXY(30, 114)
	pdf.Cell(0, 0, m.Line09Qty)
	pdf.SetXY(49, 114)
	pdf.Cell(0, 0, m.Line09Desc)
	pdf.SetXY(136, 114)
	pdf.Cell(0, 0, m.Line09Price)
	pdf.SetXY(169, 114)
	pdf.Cell(0, 0, m.Line09Amount)

	//
	// Line #10
	//

	pdf.SetXY(30, 119)
	pdf.Cell(0, 0, m.Line10Qty)
	pdf.SetXY(49, 119)
	pdf.Cell(0, 0, m.Line10Desc)
	pdf.SetXY(136, 119)
	pdf.Cell(0, 0, m.Line10Price)
	pdf.SetXY(169, 119)
	pdf.Cell(0, 0, m.Line10Amount)

	//
	// Line #11
	//

	pdf.SetXY(30, 124)
	pdf.Cell(0, 0, m.Line11Qty)
	pdf.SetXY(49, 124)
	pdf.Cell(0, 0, m.Line11Desc)
	pdf.SetXY(136, 124)
	pdf.Cell(0, 0, m.Line11Price)
	pdf.SetXY(169, 124)
	pdf.Cell(0, 0, m.Line11Amount)

	//
	// Line #12
	//

	pdf.SetXY(30, 129)
	pdf.Cell(0, 0, m.Line12Qty)
	pdf.SetXY(49, 129)
	pdf.Cell(0, 0, m.Line12Desc)
	pdf.SetXY(136, 129)
	pdf.Cell(0, 0, m.Line12Price)
	pdf.SetXY(169, 129)
	pdf.Cell(0, 0, m.Line12Amount)

	//
	// Line #13
	//

	pdf.SetXY(30, 134)
	pdf.Cell(0, 0, m.Line13Qty)
	pdf.SetXY(49, 134)
	pdf.Cell(0, 0, m.Line13Desc)
	pdf.SetXY(136, 134)
	pdf.Cell(0, 0, m.Line13Price)
	pdf.SetXY(169, 134)
	pdf.Cell(0, 0, m.Line13Amount)

	//
	// Line #14
	//

	pdf.SetXY(30, 140)
	pdf.Cell(0, 0, m.Line14Qty)
	pdf.SetXY(49, 140)
	pdf.Cell(0, 0, m.Line14Desc)
	pdf.SetXY(136, 140)
	pdf.Cell(0, 0, m.Line14Price)
	pdf.SetXY(169, 140)
	pdf.Cell(0, 0, m.Line14Amount)

	//
	// Line #15
	//

	pdf.SetXY(30, 145)
	pdf.Cell(0, 0, m.Line15Qty)
	pdf.SetXY(49, 145)
	pdf.Cell(0, 0, m.Line15Desc)
	pdf.SetXY(136, 145)
	pdf.Cell(0, 0, m.Line15Price)
	pdf.SetXY(169, 145)
	pdf.Cell(0, 0, m.Line15Amount)

	//
	// Footer
	//

	pdf.SetXY(66, 155)
	pdf.Cell(0, 0, m.InvoiceQuoteDays)
	pdf.SetXY(146, 151)
	pdf.Cell(0, 0, m.TotalLabour)
	pdf.SetXY(146, 160)
	pdf.Cell(0, 0, m.TotalMaterials)
	pdf.SetXY(146, 167)
	pdf.Cell(0, 0, m.OtherCosts)
	pdf.SetXY(146, 175)
	pdf.Cell(0, 0, m.SubTotal)
	pdf.SetXY(146, 182)
	pdf.Cell(0, 0, m.Tax)
	pdf.SetXY(146, 192)
	pdf.Cell(0, 0, m.Total)
	pdf.SetXY(77, 170)
	pdf.Cell(0, 0, m.InvoiceAssociateTax)
	pdf.SetXY(77, 179)
	pdf.Cell(0, 0, m.InvoiceQuoteDate)
	pdf.SetXY(77, 187)
	pdf.Cell(0, 0, m.InvoiceCustomersApproval)
	pdf.SetXY(19, 206)
	pdf.Cell(0, 0, m.Line01Notes)
	pdf.SetXY(19, 216)
	pdf.Cell(0, 0, m.Line02Notes)
	pdf.SetXY(55, 225)
	pdf.Cell(0, 0, m.PaymentDate)
	pdf.SetXY(158, 225)
	pdf.Cell(0, 0, m.PaymentAmount)
	pdf.SetXY(158, 217)
	pdf.Cell(0, 0, m.Deposit)

	//
	// Checks
	//

	pdf.SetXY(27, 232)
	pdf.Cell(0, 0, m.IsCash)
	pdf.SetXY(62, 232)
	pdf.Cell(0, 0, m.IsCheque)
	pdf.SetXY(97, 232)
	pdf.Cell(0, 0, m.IsDebit)
	pdf.SetXY(131, 232)
	pdf.Cell(0, 0, m.IsCredit)
	pdf.SetXY(167, 232)
	pdf.Cell(0, 0, m.IsOther)
	pdf.SetXY(42, 248)
	pdf.Cell(0, 0, m.AssociateSignDate)
	pdf.SetXY(155, 248)
	pdf.Cell(0, 0, m.AssociateSignature)
	pdf.SetXY(170, 262)
	pdf.Cell(0, 0, m.WorkOrderId)

	//
	// Signature
	//

	pdf.SetXY(128, 239)
	pdf.Cell(0, 0, m.ClientSignature)

	err = pdf.OutputFileAndClose("example.pdf")
	if err != nil {
		return nil, err
	}

	return &PDFBuilder{}, err
}
