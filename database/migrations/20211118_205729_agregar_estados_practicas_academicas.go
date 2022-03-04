package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarEstadosPracticasAcademicas_20211118_205729 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarEstadosPracticasAcademicas_20211118_205729{}
	m.Created = "20211118_205729"

	migration.Register("AgregarEstadosPracticasAcademicas_20211118_205729", m)
}

// Run the migrations
func (m *AgregarEstadosPracticasAcademicas_20211118_205729) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20211118_205729_agregar_estados_practicas_academicas_up.sql")

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
func (m *AgregarEstadosPracticasAcademicas_20211118_205729) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20211118_205729_agregar_estados_practicas_academicas_down.sql")

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
