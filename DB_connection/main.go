package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Autor struct {
	ID     int64
	nombre string
}

type Libro struct {
	ID       int64
	Titulo   string
	AuthorID string
	Fecha    time.Time
	Autor    Autor
}

func main() {
	db, err := crearConexiones()
	ctx := context.Background()

	if err != nil {
		panic(err)
	}

	// err = insertarLibro(ctx, db, "Gabriel García que no es Marquez", 1, time.Now())

	// if err != nil {
	// 	panic(err)
	// }
	err = selectQuery(ctx, db)
	if err != nil {
		panic(err)
	}

	db.Close()

}

func crearConexiones() (*sql.DB, error) {
	var connection = "root:admin@tcp(localhost:3306)/test?parseTime=True"

	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	//Golang ya trae por defecto un pool de conexiones
	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func selectQuery(ctx context.Context, db *sql.DB) error {

	qry := `SELECT l.id, l.titulo, l.author_id, a.nombre as 'Autores' , l.fecha_publicacion
	FROM libros l
	INNER JOIN autores a 
	ON a.id = l.author_id`

	rows, err := db.QueryContext(ctx, qry)

	if err != nil {
		return err
	}

	libros := []Libro{}

	for rows.Next() {
		b := Libro{}

		err := rows.Scan(&b.ID, &b.Titulo, &b.AuthorID, &b.Autor.nombre, &b.Fecha)
		if err != nil {
			return err
		}

		libros = append(libros, b)
	}

	fmt.Print(libros)

	return nil
}

func insertarLibro(ctx context.Context, db *sql.DB, titulo string, authorID int64, fecha time.Time) error {

	qryInsertar := `INSERT INTO libros (titulo, author_id, fecha_publicacion)
 					VALUES(?,?,?)`
	//Este ExecContext nos devuelve un resultado y un error
	resultado, err := db.ExecContext(ctx, qryInsertar, titulo, authorID, fecha)
	if err != nil {
		return err
	}

	id, err := resultado.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Println("INSERT ID: ", id)

	return nil
}
