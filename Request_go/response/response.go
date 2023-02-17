package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userid"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// obtiene el vinculo de la url con el id y opera conforme a los datos definidos del struct
func Response(url string, id int) (*Post, error) {
	vinculo := fmt.Sprintf(url+"%d", id)

	resp, err := http.Get(vinculo)

	if err != nil {
		return nil, err
	}
	//en caso de no existir algun error respondemos con el body cuerpo de la respuesta HTTP
	defer resp.Body.Close()

	/*Leeremos el contenido que esta dentro de la respuesta
	este nos devolvera un arreglo de bytes o un error, entonces lo declaramos para tenerlo en cuenta
	Usamos el package io y no el ioutil debido a que en la ultima version a quedado obsoleta
	aún se puede usar y no dara error pero si una advertencia, entonces es mejor evitar conflictos
	en versiones futuras y usamos lo que nos recomienda la documentacion oficial de golang

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}
	*/

	//Remplazamos la respuesta de io y utilizaremos los package de json para la lectura de datos

	post := &Post{}

	err = json.NewDecoder(resp.Body).Decode(post)
	if err != nil {
		return nil, err
	}
	// log.Println(post)
	return post, nil
}

// obtenemos los diferentes posts de prueba
func GetPosts(url string) ([]*Post, error) {

	posts := []*Post{}
	resp, err := http.Get(url)
	if err != nil {
		return posts, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return posts, err
	}

	return posts, nil

}

func SavePost(url string, userId int, title string, completed bool) (*Post, error) {
	// url https://jsonplaceholder.typicode.com/posts
	post := &Post{
		UserId:    userId,
		Title:     title,
		Completed: completed,
	}
	//Esto da respuesta un arreglo de bytes y un error
	content, err := json.Marshal(post)
	if err != nil {
		panic(err.Error())
	}
	//Esto devuelve una variable que llamaremos buffr
	buffr := bytes.NewBuffer(content)
	// este devolvera una respuesta y un error como las otras funciones
	resp, err := http.Post(url, "application/json", buffr)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(post)
	return post, err

}
