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
				Name:    "generate",
				Aliases: []string{"g"},
				Usage: "Generate a new invoice based on invoice-generator.toml information and current date. \n" +
					"Alternatively, you can pass the date (month.year) of the desired service period as an" +
					"argument like this: invoice-generator g 10.2023",
				UsageText: "super easy to use, no extra work needed, just use `generate` or short `g` \n" +
					"Alternatively, you can pass the date (month.year) of the desired service period as an" +
					"argument like this: invoice-generator g 10.2023",
				Description: "generate a new invoice based on invoice-generator.toml information and current date",
				Action: func(cCtx *cli.Context) error {
					InvoiceGenerator(invoice_data, cCtx.Args())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
