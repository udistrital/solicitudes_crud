package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizacionDefinicionEstados_20201203_184105 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizacionDefinicionEstados_20201203_184105{}
	m.Created = "20201203_184105"

	migration.Register("ActualizacionDefinicionEstados_20201203_184105", m)
}

// Run the migrations
func (m *ActualizacionDefinicionEstados_20201203_184105) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201203_184105_actualizacion_definicion_estados_up.sql")

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
func (m *ActualizacionDefinicionEstados_20201203_184105) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201203_184105_actualizacion_definicion_estados_down.sql")

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
