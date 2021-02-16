package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTipoSolicitudActualizacionDatos_20210216_104420 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTipoSolicitudActualizacionDatos_20210216_104420{}
	m.Created = "20210216_104420"

	migration.Register("CrearTipoSolicitudActualizacionDatos_20210216_104420", m)
}

// Run the migrations
func (m *CrearTipoSolicitudActualizacionDatos_20210216_104420) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210216_104420_crear_tipo_solicitud_actualizacion_datos_up.sql")

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
func (m *CrearTipoSolicitudActualizacionDatos_20210216_104420) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210216_104420_crear_tipo_solicitud_actualizacion_datos_down.sql")

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
