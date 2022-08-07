package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/over55/workery-invoicebuilder/internal/config"
	"github.com/over55/workery-invoicebuilder/pkg/dtos"
	"github.com/over55/workery-invoicebuilder/pkg/rpc"
)

func init() {
	rootCmd.AddCommand(sampleCmd)
}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Print the sample number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate sample PDF")

		dto := &dtos.WorkOrderInvoiceRequestDTO{
			Id:                       "999",
			Uuid:                     "xxxx-xxxx-xxxx-xxxx-xxxx",
			TenantID:                 "1",
			OrderId:                  "30069",
			InvoiceID:                "12345678901234",
			InvoiceDate:              "04/28/2022",
			AssociateName:            "Bartlomiej Piotr Mika XXXXXXX",
			AssociateTelephone:       "(123) 456-7898 XXXXXXXXXX",
			ClientName:               "Frank Herbert XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			ClientAddress:            "1000 - 1234 World Avenue XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			ClientTelephone:          "(987) 654-3210 XXXXX",
			ClientEmail:              "fherbert@dune.com XXXXXXXXXXXXXXXXXXXXXXX",
			Line01Qty:                "1 XXXXX",
			Line01Desc:               "Proof Reading Services XXXXXXXXXXXXXXXX",
			Line01Price:              "$15 XXX",
			Line01Amount:             "$500.00 XX",
			Line02Qty:                "2 XXXXX",
			Line02Desc:               "2x Proof Reading Services XXXXXXXXXXXXXX",
			Line02Price:              "$15 XXX",
			Line02Amount:             "$500.00 XX",
			Line03Qty:                "3 XXXXX",
			Line03Desc:               "3x Proof Reading Services XXXXXXXXXXXXXX",
			Line03Price:              "$15 XXX",
			Line03Amount:             "$500.00 XX",
			Line04Qty:                "4 XXXXX",
			Line04Desc:               "4x Proof Reading Services XXXXXXXXXXXXXX",
			Line04Price:              "$15 XXX",
			Line04Amount:             "$500.00 XX",
			Line05Qty:                "5 XXXXX",
			Line05Desc:               "5x Proof Reading Services XXXXXXXXXXXXXX",
			Line05Price:              "$15 XXX",
			Line05Amount:             "$500.00 XX",
			Line06Qty:                "6 XXXXX",
			Line06Desc:               "6x Proof Reading Services XXXXXXXXXXXXXX",
			Line06Price:              "$15 XXX",
			Line06Amount:             "$500.00 XX",
			Line07Qty:                "7 XXXXX",
			Line07Desc:               "7x Proof Reading Services XXXXXXXXXXXXXX",
			Line07Price:              "$15 XXX",
			Line07Amount:             "$500.00 XX",
			Line08Qty:                "8 XXXXX",
			Line08Desc:               "8x Proof Reading Services XXXXXXXXXXXXXX",
			Line08Price:              "$15 XXX",
			Line08Amount:             "$500.00 XX",
			Line09Qty:                "9 XXXXX",
			Line09Desc:               "9x Proof Reading Services XXXXXXXXXXXXXX",
			Line09Price:              "$15 XXX",
			Line09Amount:             "$500.00 XX",
			Line10Qty:                "10 XXXX",
			Line10Desc:               "10x Proof Reading Services XXXXXXXXXXXXX",
			Line10Price:              "$15 XXX",
			Line10Amount:             "$500.00 XX",
			Line11Qty:                "11 XXXX",
			Line11Desc:               "11x Proof Reading Services XXXXXXXXXXXXX",
			Line11Price:              "$15 XXX",
			Line11Amount:             "$500.00 XX",
			Line12Qty:                "12 XXXX",
			Line12Desc:               "12x Proof Reading Services XXXXXXXXXXXXX",
			Line12Price:              "$15 XXX",
			Line12Amount:             "$500.00 XX",
			Line13Qty:                "13 XXXX",
			Line13Desc:               "13x Proof Reading Services XXXXXXXXXXXXX",
			Line13Price:              "$15 XXX",
			Line13Amount:             "$500.00 XX",
			Line14Qty:                "14 XXXX",
			Line14Desc:               "14x Proof Reading Services XXXXXXXXXXXXX",
			Line14Price:              "$15 XXX",
			Line14Amount:             "$500.00 XX",
			Line15Qty:                "15 XXXX",
			Line15Desc:               "15x Proof Reading Services XXXXXXXXXXXXX",
			Line15Price:              "$15 XXX",
			Line15Amount:             "$500.00 XX",
			InvoiceQuoteDays:         "90",
			TotalLabour:              "$1200.00 XXXXXXXXXX",
			TotalMaterials:           "$1200.00 XXXXXXXXXX",
			OtherCosts:               "$1200.00 XXXXXXXXXX",
			SubTotal:                 "$1200.00 XXXXXXXXXX",
			Tax:                      "$1200.00 XXXXXXXXXX",
			Total:                    "$1200.00 XXXXXXXXXX",
			InvoiceAssociateTax:      "#1234567890 XXX",
			InvoiceQuoteDate:         "04/28/2022",
			InvoiceCustomersApproval: "FHerbert",
			Line01Notes:              "This is a comment. XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Line02Notes:              "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			PaymentDate:              "04/28/2022 XXXX",
			PaymentAmount:            "$1200.00 XXXXXX",
			Deposit:                  "$1,000.00 XXXXX",
			IsCash:                   "X",
			IsCheque:                 "X",
			IsDebit:                  "X",
			IsCredit:                 "X",
			IsOther:                  "X",
			ClientSignature:          "Frank Herbert XXXXXXXXXXXXXXX",
			AssociateSignDate:        "04/28/2022",
			AssociateSignature:       "Bart Mika XXXXXXX",
			WorkOrderId:              "12345 XXXX",
		}

		// Load up all the environment variables.
		appConf := config.AppConfig()

		// Connect to a running client.
		applicationAddress := fmt.Sprintf("%s:%s", appConf.Server.IP, appConf.Server.Port)
		client := rpc.NewClient(applicationAddress)

		// Execute the remote call.
		res, err := client.GeneratePDF(dto)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)

	},
}
