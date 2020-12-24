package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarParametrizacionSolicitudEvaluacion_20201223_202030 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarParametrizacionSolicitudEvaluacion_20201223_202030{}
	m.Created = "20201223_202030"

	migration.Register("AgregarParametrizacionSolicitudEvaluacion_20201223_202030", m)
}

// Run the migrations
func (m *AgregarParametrizacionSolicitudEvaluacion_20201223_202030) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201223_202030_agregar_parametrizacion_solicitud_evaluacion_up.sql")

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
func (m *AgregarParametrizacionSolicitudEvaluacion_20201223_202030) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201223_202030_agregar_parametrizacion_solicitud_evaluacion_down.sql")

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
