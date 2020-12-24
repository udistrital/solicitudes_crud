package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizacionNombresEstados_20201223_202004 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizacionNombresEstados_20201223_202004{}
	m.Created = "20201223_202004"

	migration.Register("ActualizacionNombresEstados_20201223_202004", m)
}

// Run the migrations
func (m *ActualizacionNombresEstados_20201223_202004) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201223_202004_actualizacion_nombres_estados_up.sql")

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
func (m *ActualizacionNombresEstados_20201223_202004) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201223_202004_actualizacion_nombres_estados_down.sql")

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
