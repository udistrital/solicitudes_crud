package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposComite_20201223_202016 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposComite_20201223_202016{}
	m.Created = "20201223_202016"

	migration.Register("AgregarCamposComite_20201223_202016", m)
}

// Run the migrations
func (m *AgregarCamposComite_20201223_202016) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201223_202016_agregar_campos_comite_up.sql")

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
func (m *AgregarCamposComite_20201223_202016) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201223_202016_agregar_campos_comite_down.sql")

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
