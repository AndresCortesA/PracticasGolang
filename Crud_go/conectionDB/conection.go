package conectiondb

import "database/sql"

func ConectionDB() (conect *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := "admin"

	NameDB := "sistemago"

	conection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+NameDB)
	if err != nil {
		panic(err.Error())
	}
	return conection

}
