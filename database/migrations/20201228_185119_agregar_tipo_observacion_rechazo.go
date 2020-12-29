package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarTipoObservacionRechazo_20201228_185119 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarTipoObservacionRechazo_20201228_185119{}
	m.Created = "20201228_185119"

	migration.Register("AgregarTipoObservacionRechazo_20201228_185119", m)
}

// Run the migrations
func (m *AgregarTipoObservacionRechazo_20201228_185119) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO solicitud.tipo_observacion(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (3, 'Comentario rechazo', 'Indica que la observación es un comentario de rechazo', 'comrec', 3, now(), now(), true);")
}

// Reverse the migrations
func (m *AgregarTipoObservacionRechazo_20201228_185119) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("delete from solicitud.tipo_observacion where id = 3;")
}
