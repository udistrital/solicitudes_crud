package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EstadoSolicitudTransferencias_20220308_211742 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EstadoSolicitudTransferencias_20220308_211742{}
	m.Created = "20220308_211742"

	migration.Register("EstadoSolicitudTransferencias_20220308_211742", m)
}

// Run the migrations
func (m *EstadoSolicitudTransferencias_20220308_211742) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220308_211742_estado_solicitud_transferencias_up.sql")

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
func (m *EstadoSolicitudTransferencias_20220308_211742) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220308_211742_estado_solicitud_transferencias_down.sql")

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

