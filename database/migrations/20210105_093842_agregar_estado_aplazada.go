package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarEstadoAplazada_20210105_093842 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarEstadoAplazada_20210105_093842{}
	m.Created = "20210105_093842"

	migration.Register("AgregarEstadoAplazada_20210105_093842", m)
}

// Run the migrations
func (m *AgregarEstadoAplazada_20210105_093842) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210105_093842_agregar_estado_aplazada_up.sql")

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
func (m *AgregarEstadoAplazada_20210105_093842) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210105_093842_agregar_estado_aplazada_down.sql")

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
