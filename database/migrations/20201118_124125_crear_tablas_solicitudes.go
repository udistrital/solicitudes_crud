package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasSolicitudes_20201118_124125 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasSolicitudes_20201118_124125{}
	m.Created = "20201118_124125"

	migration.Register("CrearTablasSolicitudes_20201118_124125", m)
}

// Run the migrations
func (m *CrearTablasSolicitudes_20201118_124125) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201118_124125_crear_tablas_solicitudes_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearTablasSolicitudes_20201118_124125) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201118_124125_crear_tablas_solicitudes_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
