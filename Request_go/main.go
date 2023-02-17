package main

import (
	"Request_go/response"
	"fmt"
)

func main() {
	//obtener todos los post
	posts, err := response.GetPosts("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		panic(err.Error())
	}

	for _, p := range posts {
		//Para la respuesta obteniendo solo el post
		// if p.Id == 6 {
		// 	post, err := response.Response("https://jsonplaceholder.typicode.com/todos/", p.Id)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// 	fmt.Println(post)
		// }
		//Para guardar un post

		if p.UserId == 8 {
			post, err := response.SavePost("https://jsonplaceholder.typicode.com/posts", p.UserId, "test tittle",
				true)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Saved full post: %v", post)
		}
	}
}
