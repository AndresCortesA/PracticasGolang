package handlers

import (
	conectiondb "Crud_go/conectionDB"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conect := conectiondb.ConectionDB()
		insert, err := conect.Prepare("INSERT INTO employee(nameEmployee, mailEmployee)VALUES('Andriu', 'andri@correo.com')")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec()

		tmpl.ExecuteTemplate(w, "index", nil)
	})
}

func Create() {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "create", nil)
	})
}

func Insert() {
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			mail := r.FormValue("mail")
			conect := conectiondb.ConectionDB()
			insert, err := conect.Prepare("INSERT INTO employee(nameEmployee, mailEmployee)VALUES(?, ?)")
			if err != nil {
				panic(err.Error())
			}
			insert.Exec(name, mail)
			http.Redirect(w, r, "/", http.StatusFound)

		}
	})
}
