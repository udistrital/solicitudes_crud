package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarTituloObservacion_20201130_215132 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarTituloObservacion_20201130_215132{}
	m.Created = "20201130_215132"

	migration.Register("AgregarTituloObservacion_20201130_215132", m)
}

// Run the migrations
func (m *AgregarTituloObservacion_20201130_215132) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201130_215132_agregar_titulo_observacion_up.sql")

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
func (m *AgregarTituloObservacion_20201130_215132) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201130_215132_agregar_titulo_observacion_down.sql")

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
