package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server serverConf
}

type serverConf struct {
	// Port        string `env:"WORKERY_INVOICEBUILDER_PORT,required"`
	// IP          string `env:"WORKERY_INVOICEBUILDER_IP,required"`
	// SecretKey   []byte `env:"WORKERY_INVOICEBUILDER_SECRET_KEY,required"`
	PDFFilePath string `env:"WORKERY_INVOICEBUILDER_PDF_FILE_PATH,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
