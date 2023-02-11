package costos

func Promedio(costo ...float32) float32 {
	return Suma(costo...) / float32(len(costo))
}

func Suma(costo ...float32) float32 {
	var suma float32

	for _, cost := range costo {
		suma += cost
	}

	return suma
}

func Maximo(costo ...float32) float32 {
	var maximo float32

	for _, exp := range costo {
		if exp > maximo {
			maximo = exp
		}

	}
	return maximo
}

func Minimo(costo ...float32) float32 {
	if len(costo) == 0 {
		return 0
	}

	var minimo float32 = costo[0]

	for _, exp := range costo {
		if exp < minimo {
			minimo = exp
		}
	}

	return minimo
}
