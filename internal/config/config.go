package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server serverConf
}

type serverConf struct {
	Port                string `env:"WORKERY_INVOICEBUILDER_PORT,required"`
	IP                  string `env:"WORKERY_INVOICEBUILDER_IP,required"`
	PDFTemplateFilePath string `env:"WORKERY_INVOICEBUILDER_PDF_TEMPLATE_FILE_PATH,required"`
	DataDirectoryPath   string `env:"WORKERY_INVOICEBUILDER_DATA,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
