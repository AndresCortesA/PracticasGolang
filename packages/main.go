package main

import "packages/countries"

//para importar paquetes externos como lo es uuid de google ponemos go get github.com/google/uuid y lo podemos
//con el sgte. código
//id := uuid.New()
//fmt.Println("ID user:", id)

func main() {
	countries.Add("USA")
	countries.Add("Mexico")
	countries.Add("Japan")
	countries.Add("Spain")

	err := countries.SetMyCountry("Spain")

	if err != nil {
		panic(err)
	}

	countries.List()
}
