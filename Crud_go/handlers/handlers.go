package handlers

import (
	conectiondb "Crud_go/conectionDB"
	"html/template"
	"net/http"
)

type Employee struct {
	Id         int
	Name, Mail string
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conect := conectiondb.ConectionDB()
		selectRegister, err := conect.Query("SELECT * FROM employee")
		if err != nil {
			panic(err.Error())
		}
		employee := Employee{}
		arrayEmployee := []Employee{}

		for selectRegister.Next() {
			var id int
			var name, mail string
			err = selectRegister.Scan(&id, &name, &mail)
			if err != nil {
				panic(err.Error())
			}
			employee.Id = id
			employee.Name = name
			employee.Mail = mail

			arrayEmployee = append(arrayEmployee, employee)
		}
		tmpl.ExecuteTemplate(w, "index", arrayEmployee)
		// fmt.Println(arrayEmployee)

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

func Delete() {
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		idEmployee := r.URL.Query().Get("id")
		//fmt.Println(idEmployee)
		conect := conectiondb.ConectionDB()
		delete, err := conect.Prepare("DELETE FROM employee WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		delete.Exec(idEmployee)
		http.Redirect(w, r, "/", http.StatusFound)
	})
}

func Edit() {
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {

		idEmployee := r.URL.Query().Get("id")
		//fmt.Println(idEmployee)
		conect := conectiondb.ConectionDB()
		edit, err := conect.Query("SELECT *FROM employee WHERE id=?", idEmployee)
		if err != nil {
			panic(err.Error())
		}

		employee := Employee{}
		for edit.Next() {
			var id int
			var name, mail string
			err = edit.Scan(&id, &name, &mail)
			if err != nil {
				panic(err.Error())
			}
			employee.Id = id
			employee.Name = name
			employee.Mail = mail
		}

		tmpl.ExecuteTemplate(w, "edit", employee)

	})
}

func Update() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			mail := r.FormValue("mail")
			conect := conectiondb.ConectionDB()
			update, err := conect.Prepare("UPDATE employee SET nameEmployee=?,mailEmployee=? WHERE id=?")
			if err != nil {
				panic(err.Error())
			}
			update.Exec(name, mail, id)
			http.Redirect(w, r, "/", http.StatusFound)

		}
	})
}
