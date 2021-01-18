package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearAtributoSolicitud_20210118_091314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearAtributoSolicitud_20210118_091314{}
	m.Created = "20210118_091314"

	migration.Register("CrearAtributoSolicitud_20210118_091314", m)
}

// Run the migrations
func (m *CrearAtributoSolicitud_20210118_091314) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210118_091314_crear_atributo_solicitud.up.sql")

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
func (m *CrearAtributoSolicitud_20210118_091314) Down() {
	file, err := ioutil.ReadFile("../scripts/20210118_091314_crear_atributo_solicitud.down.sql")

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
