package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarSolicitudTransferenciaReingreso_20220213_204724 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarSolicitudTransferenciaReingreso_20220213_204724{}
	m.Created = "20220213_204724"

	migration.Register("AgregarSolicitudTransferenciaReingreso_20220213_204724", m)
}

// Run the migrations
func (m *AgregarSolicitudTransferenciaReingreso_20220213_204724) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220213_204724_agregar_solicitud_transferencia_reingreso_up.sql")

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
func (m *AgregarSolicitudTransferenciaReingreso_20220213_204724) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220213_204724_agregar_solicitud_transferencia_reingreso_down.sql")

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
