package main

import (
	"CommandLineInterfaceCLI/comandos"
	"flag"
	"log"
	"strconv"
)

func main() {

	var valores []float32
	var export string

	flag.StringVar(&export, "export", "", "Exportando los detalles de .txt")
	flag.Parse()

	for {
		entrada, err := comandos.GetEntrada()
		if err != nil {
			log.Fatal(err)
		}
		if entrada == "cls" {
			break

		}

		valor, err := strconv.ParseFloat(entrada, 32)
		if err != nil {
			log.Fatal(err)
		}

		valores = append(valores, float32(valor))
	}

	if export == "" {
		comandos.MostrarConsola(valores)
	} else {
		comandos.Export(export, valores)
	}
}
