package comandos

import (
	"CommandLineInterfaceCLI/costos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lector *bufio.Reader = bufio.NewReader(os.Stdin)

func GetEntrada() (string, error) {

	fmt.Print("Indica un valor-> ")
	str, err := lector.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\r\n", "", 1)

	return str, nil
}

func MostrarConsola(arreglos []float32) {

	fmt.Println(contenerString(arreglos))

}

func contenerString(arreglos []float32) string {
	builder := strings.Builder{}
	maximo, minimo, promedio, total := costosDetalle(arreglos)
	for i, index := range arreglos {
		builder.WriteString(fmt.Sprintf("Valor: %6.2f\n", index))

		if i == len(arreglos)-1 {
			builder.WriteString("=========================\n")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Maximo: %6.2f\n", maximo))
			builder.WriteString(fmt.Sprintf("Minimo: %6.2f\n", minimo))
			builder.WriteString(fmt.Sprintf("Promedio: %6.2f\n", promedio))
			builder.WriteString("=========================")
		}
	}

	return builder.String()
}

func costosDetalle(arreglos []float32) (maximo, minimo, promedio, total float32) {

	if len(arreglos) == 0 {
		return
	}

	minimo = costos.Minimo(arreglos...)
	maximo = costos.Maximo(arreglos...)
	total = costos.Suma(arreglos...)
	promedio = costos.Promedio(arreglos...)

	return
}

func Export(archivoNombre string, lista []float32) error {

	f, err := os.Create(archivoNombre)
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(contenerString(lista))
	if err != nil {
		return err
	}

	return w.Flush()
}
