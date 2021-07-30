package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err            error
	db             *sql.DB
	tablaBD        = "Membresia"
	membresia_tipo = "Itzel Cabrera Sanchez"
)

func main() {
	nuevaConexion()
	fmt.Println("Practica #1")
	separador()
	comprobarTipoMembresia(tablaBD, membresia_tipo)
}

func nuevaConexion() {
	db, err = sql.Open("mysql", string("Bienhechor:Bienhechor_1234;@tcp(189.130.27.82:3306)/Bienhechor"))

	revisarError(err)

	err = db.Ping()
	revisarError(err)
}

func separador() {
	fmt.Println("________________________________")
}

func mostrarTabla(tabla string) {
	query, _ := db.Query("SELECT * FROM " + tabla)

	for query.Next() {
		var id_membresia, tipo_membresia string
		err = query.Scan(&id_membresia, &tipo_membresia)
		revisarError(err)

		fmt.Println(tipo_membresia + ". ID: " + id_membresia)
	}
}

func comprobarTipoMembresia(tabla, membresia string) {
	var id, tipo_membresia string
	var estado = false
	query, _ := db.Query("SELECT * FROM "+tabla+" where TipoMembresia = ?", membresia)
	revisarError(err)

	for query.Next() {
		err = query.Scan(&id, &tipo_membresia)
		revisarError(err)

		if id != "" {
			fmt.Println(tipo_membresia + ", usted ya ha sido registrado. ID: " + id)
			estado = true
			separador()
			//mostrarTabla("Membresia")
		}
	}
	if estado != true {
		fmt.Println("0")
		agregarDatosBD(membresia_tipo)
	}
	CerrarBD()
}

func agregarDatosBD(tipo_membresia string) {
	agregar, err := db.Exec("insert into Membresia (TipoMembresia) values (?)", tipo_membresia)
	revisarError(err)

	estatus, err := agregar.LastInsertId()
	revisarError(err)

	if err != nil {
		fmt.Println(false)
		fmt.Println(estatus)
		fmt.Println("Registro Exitoso!")
		separador()
	} else {
		fmt.Println(true)
	}
	CerrarBD()
}

func revisarError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func CerrarBD() {
	defer db.Close()
}
