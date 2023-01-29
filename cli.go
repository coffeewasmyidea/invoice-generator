package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func CommandLineTool(invoice_data InvoiceData) {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "generate",
				Aliases:     []string{"g"},
				Usage:       "generate a new invoice based on invoice-generator.toml information and current date",
				UsageText:   "super easy to use, no extra work needed, just use `generate` or short `g`",
				Description: "generate a new invoice based on invoice-generator.toml information and current date",
				Action: func(cCtx *cli.Context) error {
					InvoiceGenerator(invoice_data)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
