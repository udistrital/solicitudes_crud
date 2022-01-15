package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarSolicitudDerechosPecuniarios_20220115_025328 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarSolicitudDerechosPecuniarios_20220115_025328{}
	m.Created = "20220115_025328"

	migration.Register("AgregarSolicitudDerechosPecuniarios_20220115_025328", m)
}

// Run the migrations
func (m *AgregarSolicitudDerechosPecuniarios_20220115_025328) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220115_025328_agregar_solicitud_derechos_pecuniarios_up.sql")

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
func (m *AgregarSolicitudDerechosPecuniarios_20220115_025328) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220115_025328_agregar_solicitud_derechos_pecuniarios_down.sql")

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
