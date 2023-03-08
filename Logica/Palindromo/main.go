package main

import (
	"fmt"
	"strings"
)

/*
Quiero crear un programa que me detecte un palindromo, por lo que se

entonces:

Debo primero detectar la cadena de strings y recorrerla "I like 'racecar' that very fast"
ahí esta el palidromo, lo que deseo es obtener esa cadena como salida ya es la respuesta

entonces lo que debo omitir por ahora, es mayusculas y espacios
*/

//primero definamos una funcion que me ayude a leer una cadena y me indique si es un palindromo

func esPalindromo(s string) bool {
	//creo un for que me ayude a leer una cadena hasta la mitad y me indique si la otra mitad es igual
	for i := 0; i < len(s)/2; /*este me lee hasta la mitad*/ i++ {
		//entonces le defino que me lea la otra mitad de la cadena y me indique si no es igual
		if s[i] != s[len(s)-i-1 /*este me lee la otra mitad faltante*/] {
			return false
		}
	}

	return true
}

func main() {
	// entradaS := "racecar"
	entradaS := "I like racecar that very fast"
	entradaS = strings.ToLower(entradaS)
	//el -1 indica que los " " espaciados son eliminados
	entradaS = strings.Replace(entradaS, " ", "", -1)

	// fmt.Println(esPalindromo(entradaS))
	//ahora necesitamos almacenar el palindromo más largo

	largoPalindromo := ""
	sizePalindromo := 0

	for i := 0; i < len(entradaS); i++ {
		for j := i + 1; j <= len(entradaS); j++ {
			subcadenas := entradaS[i:j]
			/* lo que devuelve esta subcadena es lo siguite
			supongamos que recibe la palabra "hello"
			"h"
			"he"
			"hel"
			"hell"
			"hello"
			"e"
			"el"
			"ell"
			"ello"
			"l"
			"ll"
			"llo"
			"l"
			"lo"
			"o"
			esto es lo que devuelve por cada iteracion hasta recorrer la cadena
			para ser más claro supongamos que el i := 2 es dos empieza en "l" y termina en "llo"
			y asi sucesivamente para la cadena
			*/

			//hora vamos a definir el tamaño del palindromo y que realmente sea un palindromo

			if esPalindromo(subcadenas) && len(subcadenas) > sizePalindromo {
				largoPalindromo = subcadenas
				sizePalindromo = len(subcadenas)
			}
		}
	}

	fmt.Println(largoPalindromo)
}
