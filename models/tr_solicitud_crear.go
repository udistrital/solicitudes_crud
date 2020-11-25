package models

import (
	// "errors"
	"fmt"
	// "reflect"
	// "strings"
	// "time"

	"github.com/astaxie/beego/orm"
	// "github.com/udistrital/utils_oas/time_bogota"
)
/*
type TrSolicitudCrear struct {
	FechaRadicacion     string
	Referencia          string
	Solicitantes        []int
	EstadoTipoSolicitud int
	SolicitudPadreId    int
}

func AddNuevaSolicitud(m *TrSolicitudCrear)(err error){
	o := orm.NewOrm()
	err = o.Begin()

	var solicitud = new(Solicitud)
	fmt.Println(solicitud)

	
	
	solicitud.FechaCreacion = time_bogota.TiempoBogotaFormato()
	solicitud.FechaModificacion = time_bogota.TiempoBogotaFormato()
	solicitud.FechaRadicacion = time_bogota.TiempoCorreccionFormato(m.FechaRadicacion)
	solicitud.Referencia = m.Referencia

	var estadoTipo *EstadoTipoSolicitud
	estadoTipo = &EstadoTipoSolicitud{Id: m.EstadoTipoSolicitud}
	if err = o.Read(estadoTipo); err != nil {
		solicitud.EstadoTipoSolicitudId = estadoTipo
		// return v, nil
	}
	fmt.Println(o.Read(estadoTipo))

	fmt.Println("===")
	fmt.Println(solicitud)
	fmt.Println("---")
	fmt.Println(solicitud.EstadoTipoSolicitudId)
	if idSolicitud, errTr := o.Insert(solicitud); errTr == nil {

		fmt.Println(idSolicitud)


		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

*/

type TrSolicitudCrear struct {
	Solicitud           *Solicitud
	Solicitantes        *[]Solicitante
	EvolucionEstado     *SolicitudEvolucionEstado
}

func AddNuevaSolicitud(m *TrSolicitudCrear)(err error){
	o := orm.NewOrm()
	err = o.Begin()

	if idSolicitud, errTr := o.Insert(m.Solicitud); errTr == nil {

		fmt.Println(idSolicitud)

		for _, v := range *m.Solicitantes {
			v.SolicitudId.Id = int(idSolicitud)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		m.EvolucionEstado.SolicitudId.Id = int(idSolicitud)
		if _, errTr = o.Insert(m.EvolucionEstado); errTr != nil {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}

		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}
