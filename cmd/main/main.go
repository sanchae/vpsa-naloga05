package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	r "github.com/sanchae/vpsa-naloga05/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	r.DodajStudenta("63210001", "Ana", "Novak")
	r.DodajStudenta("63210002", "Boris", "Kralj")
	r.DodajStudenta("63210003", "Janez", "Novak")

	app := &cli.App{
		Name:  "Redovalnica CLI",
		Usage: "Upravljanje ocen študentov",

		Flags: []cli.Flag{
			&cli.IntFlag{Name: "stOcen", Value: 3},
			&cli.IntFlag{Name: "minOcena", Value: 1},
			&cli.IntFlag{Name: "maxOcena", Value: 10},
		},

		Commands: []*cli.Command{
			{
				Name:  "dodajOceno",
				Usage: "dodajOceno <vpisna> <ocena>",
				Action: func(c *cli.Context) error {
					vpisna := c.Args().Get(0)
					oc, _ := strconv.Atoi(c.Args().Get(1))
					r.DodajOceno(vpisna, oc)
					return nil
				},
			},
			{
				Name:  "izpisiVse",
				Usage: "izpisiVse <vpisna>",
				Action: func(c *cli.Context) error {
					vpisna := c.Args().Get(0)
					fmt.Println(r.IzpisVsehOcen(vpisna))
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpiše končni uspeh vseh študentov",
				Action: func(c *cli.Context) error {
					r.IzpisiKoncniUspeh(
						c.Int("minOcena"),
						c.Int("maxOcena"),
						c.Int("stOcen"),
					)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
