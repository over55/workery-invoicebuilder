package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"

	"github.com/over55/workery-invoicebuilder/internal/config"
	"github.com/over55/workery-invoicebuilder/pkg/dtos"
	"github.com/over55/workery-invoicebuilder/pkg/uuid"
)

type InvoiceBuilder interface {
	GeneratePDF(dto *dtos.WorkOrderInvoiceRequestDTO) (*dtos.WorkOrderInvoiceResponseDTO, error)
}

type invoiceBuilder struct {
	AppConfig    *config.Conf
	UUIDProvider uuid.Provider
}

func New(appConfig *config.Conf, uuidp uuid.Provider) (InvoiceBuilder, error) {
	// Defensive code: Make sure we have access to the file before proceeding any further with the code.
	log.Println("Opening up file at path:", appConfig.Server.PDFTemplateFilePath)
	_, err := os.Stat(appConfig.Server.PDFTemplateFilePath)
	if os.IsNotExist(err) {
		return nil, errors.New("file does not exist")
	}

	return &invoiceBuilder{
		AppConfig:    appConfig,
		UUIDProvider: uuidp,
	}, nil
}

func (bdr *invoiceBuilder) GeneratePDF(dto *dtos.WorkOrderInvoiceRequestDTO) (*dtos.WorkOrderInvoiceResponseDTO, error) {
	var err error

	// Open our PDF invoice template and create clone it for the PDF invoice we will be building with.
	pdf := gofpdf.New("P", "mm", "A4", "")
	tpl1 := gofpdi.ImportPage(pdf, bdr.AppConfig.Server.PDFTemplateFilePath, 1, "/MediaBox")

	pdf.AddPage()

	// Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 0, 210, 0)

	pdf.SetFont("Courier", "", 11)

	pdf.SetXY(159, 19)
	pdf.Cell(0, 0, dto.InvoiceID)

	pdf.SetFont("Helvetica", "", 11)

	//
	// Header
	//

	pdf.SetXY(116, 25)
	pdf.Cell(0, 0, dto.InvoiceDate)
	pdf.SetXY(136, 32)
	pdf.Cell(0, 0, dto.AssociateName)
	pdf.SetXY(136, 39)
	pdf.Cell(0, 0, dto.AssociateTelephone)
	pdf.SetXY(46, 46)
	pdf.Cell(0, 0, dto.ClientName)
	pdf.SetXY(36, 53)
	pdf.Cell(0, 0, dto.ClientAddress)
	pdf.SetXY(33, 60)
	pdf.Cell(0, 0, dto.ClientTelephone)
	pdf.SetXY(92, 60)
	pdf.Cell(0, 0, dto.ClientEmail)

	//
	// Line #1
	//

	pdf.SetXY(30, 72)
	pdf.Cell(0, 0, dto.Line01Qty)
	pdf.SetXY(49, 72)
	pdf.Cell(0, 0, dto.Line01Desc)
	pdf.SetXY(136, 72)
	pdf.Cell(0, 0, dto.Line01Price)
	pdf.SetXY(169, 72)
	pdf.Cell(0, 0, dto.Line01Amount)

	//
	// Line #2
	//

	pdf.SetXY(30, 78)
	pdf.Cell(0, 0, dto.Line02Qty)
	pdf.SetXY(49, 78)
	pdf.Cell(0, 0, dto.Line02Desc)
	pdf.SetXY(136, 78)
	pdf.Cell(0, 0, dto.Line02Price)
	pdf.SetXY(169, 78)
	pdf.Cell(0, 0, dto.Line02Amount)

	//
	// Line #3
	//

	pdf.SetXY(30, 83)
	pdf.Cell(0, 0, dto.Line03Qty)
	pdf.SetXY(49, 83)
	pdf.Cell(0, 0, dto.Line03Desc)
	pdf.SetXY(136, 83)
	pdf.Cell(0, 0, dto.Line03Price)
	pdf.SetXY(169, 83)
	pdf.Cell(0, 0, dto.Line03Amount)

	//
	// Line #4
	//

	pdf.SetXY(30, 88)
	pdf.Cell(0, 0, dto.Line04Qty)
	pdf.SetXY(49, 88)
	pdf.Cell(0, 0, dto.Line04Desc)
	pdf.SetXY(136, 88)
	pdf.Cell(0, 0, dto.Line04Price)
	pdf.SetXY(169, 88)
	pdf.Cell(0, 0, dto.Line04Amount)

	//
	// Line #5
	//

	pdf.SetXY(30, 93)
	pdf.Cell(0, 0, dto.Line05Qty)
	pdf.SetXY(49, 93)
	pdf.Cell(0, 0, dto.Line05Desc)
	pdf.SetXY(136, 93)
	pdf.Cell(0, 0, dto.Line05Price)
	pdf.SetXY(169, 93)
	pdf.Cell(0, 0, dto.Line05Amount)

	//
	// Line #6
	//

	pdf.SetXY(30, 98)
	pdf.Cell(0, 0, dto.Line06Qty)
	pdf.SetXY(49, 98)
	pdf.Cell(0, 0, dto.Line06Desc)
	pdf.SetXY(136, 98)
	pdf.Cell(0, 0, dto.Line06Price)
	pdf.SetXY(169, 98)
	pdf.Cell(0, 0, dto.Line06Amount)

	//
	// Line #7
	//

	pdf.SetXY(30, 103)
	pdf.Cell(0, 0, dto.Line07Qty)
	pdf.SetXY(49, 103)
	pdf.Cell(0, 0, dto.Line07Desc)
	pdf.SetXY(136, 103)
	pdf.Cell(0, 0, dto.Line07Price)
	pdf.SetXY(169, 103)
	pdf.Cell(0, 0, dto.Line07Amount)

	//
	// Line #8
	//

	pdf.SetXY(30, 109)
	pdf.Cell(0, 0, dto.Line08Qty)
	pdf.SetXY(49, 109)
	pdf.Cell(0, 0, dto.Line08Desc)
	pdf.SetXY(136, 109)
	pdf.Cell(0, 0, dto.Line08Price)
	pdf.SetXY(169, 109)
	pdf.Cell(0, 0, dto.Line08Amount)

	//
	// Line #9
	//

	pdf.SetXY(30, 114)
	pdf.Cell(0, 0, dto.Line09Qty)
	pdf.SetXY(49, 114)
	pdf.Cell(0, 0, dto.Line09Desc)
	pdf.SetXY(136, 114)
	pdf.Cell(0, 0, dto.Line09Price)
	pdf.SetXY(169, 114)
	pdf.Cell(0, 0, dto.Line09Amount)

	//
	// Line #10
	//

	pdf.SetXY(30, 119)
	pdf.Cell(0, 0, dto.Line10Qty)
	pdf.SetXY(49, 119)
	pdf.Cell(0, 0, dto.Line10Desc)
	pdf.SetXY(136, 119)
	pdf.Cell(0, 0, dto.Line10Price)
	pdf.SetXY(169, 119)
	pdf.Cell(0, 0, dto.Line10Amount)

	//
	// Line #11
	//

	pdf.SetXY(30, 124)
	pdf.Cell(0, 0, dto.Line11Qty)
	pdf.SetXY(49, 124)
	pdf.Cell(0, 0, dto.Line11Desc)
	pdf.SetXY(136, 124)
	pdf.Cell(0, 0, dto.Line11Price)
	pdf.SetXY(169, 124)
	pdf.Cell(0, 0, dto.Line11Amount)

	//
	// Line #12
	//

	pdf.SetXY(30, 129)
	pdf.Cell(0, 0, dto.Line12Qty)
	pdf.SetXY(49, 129)
	pdf.Cell(0, 0, dto.Line12Desc)
	pdf.SetXY(136, 129)
	pdf.Cell(0, 0, dto.Line12Price)
	pdf.SetXY(169, 129)
	pdf.Cell(0, 0, dto.Line12Amount)

	//
	// Line #13
	//

	pdf.SetXY(30, 134)
	pdf.Cell(0, 0, dto.Line13Qty)
	pdf.SetXY(49, 134)
	pdf.Cell(0, 0, dto.Line13Desc)
	pdf.SetXY(136, 134)
	pdf.Cell(0, 0, dto.Line13Price)
	pdf.SetXY(169, 134)
	pdf.Cell(0, 0, dto.Line13Amount)

	//
	// Line #14
	//

	pdf.SetXY(30, 140)
	pdf.Cell(0, 0, dto.Line14Qty)
	pdf.SetXY(49, 140)
	pdf.Cell(0, 0, dto.Line14Desc)
	pdf.SetXY(136, 140)
	pdf.Cell(0, 0, dto.Line14Price)
	pdf.SetXY(169, 140)
	pdf.Cell(0, 0, dto.Line14Amount)

	//
	// Line #15
	//

	pdf.SetXY(30, 145)
	pdf.Cell(0, 0, dto.Line15Qty)
	pdf.SetXY(49, 145)
	pdf.Cell(0, 0, dto.Line15Desc)
	pdf.SetXY(136, 145)
	pdf.Cell(0, 0, dto.Line15Price)
	pdf.SetXY(169, 145)
	pdf.Cell(0, 0, dto.Line15Amount)

	//
	// Footer
	//

	pdf.SetXY(66, 155)
	pdf.Cell(0, 0, dto.InvoiceQuoteDays)
	pdf.SetXY(146, 151)
	pdf.Cell(0, 0, dto.TotalLabour)
	pdf.SetXY(146, 160)
	pdf.Cell(0, 0, dto.TotalMaterials)
	pdf.SetXY(146, 167)
	pdf.Cell(0, 0, dto.OtherCosts)
	pdf.SetXY(146, 175)
	pdf.Cell(0, 0, dto.SubTotal)
	pdf.SetXY(146, 182)
	pdf.Cell(0, 0, dto.Tax)
	pdf.SetXY(146, 192)
	pdf.Cell(0, 0, dto.Total)
	pdf.SetXY(77, 170)
	pdf.Cell(0, 0, dto.InvoiceAssociateTax)
	pdf.SetXY(77, 179)
	pdf.Cell(0, 0, dto.InvoiceQuoteDate)
	pdf.SetXY(77, 187)
	pdf.Cell(0, 0, dto.InvoiceCustomersApproval)
	pdf.SetXY(19, 206)
	pdf.Cell(0, 0, dto.Line01Notes)
	pdf.SetXY(19, 216)
	pdf.Cell(0, 0, dto.Line02Notes)
	pdf.SetXY(55, 225)
	pdf.Cell(0, 0, dto.PaymentDate)
	pdf.SetXY(158, 225)
	pdf.Cell(0, 0, dto.PaymentAmount)
	pdf.SetXY(158, 217)
	pdf.Cell(0, 0, dto.Deposit)

	//
	// Checks
	//

	pdf.SetXY(27, 232)
	pdf.Cell(0, 0, dto.IsCash)
	pdf.SetXY(62, 232)
	pdf.Cell(0, 0, dto.IsCheque)
	pdf.SetXY(97, 232)
	pdf.Cell(0, 0, dto.IsDebit)
	pdf.SetXY(131, 232)
	pdf.Cell(0, 0, dto.IsCredit)
	pdf.SetXY(167, 232)
	pdf.Cell(0, 0, dto.IsOther)
	pdf.SetXY(42, 248)
	pdf.Cell(0, 0, dto.AssociateSignDate)
	pdf.SetXY(155, 248)
	pdf.Cell(0, 0, dto.AssociateSignature)
	pdf.SetXY(170, 262)
	pdf.Cell(0, 0, dto.WorkOrderId)

	//
	// Signature
	//

	pdf.SetXY(128, 239)
	pdf.Cell(0, 0, dto.ClientSignature)

	////
	//// Generate the file and save it to the file.
	////

	fileName := fmt.Sprintf("%s.pdf", bdr.UUIDProvider.NewUUID())
	filePath := fmt.Sprintf("%s/%s", bdr.AppConfig.Server.DataDirectoryPath, fileName)

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		return nil, err
	}

	////
	//// Open the file and read all the binary data.
	////

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bin, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	////
	//// Return the generate invoice.
	////

	return &dtos.WorkOrderInvoiceResponseDTO{
		FileName: fileName,
		Content:  bin,
	}, err
}
