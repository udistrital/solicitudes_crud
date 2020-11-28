package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrSolicitudCrear struct {
	Solicitud       *Solicitud
	Solicitantes    *[]Solicitante
	Observaciones   *[]Observacion
	EvolucionEstado *SolicitudEvolucionEstado
}

func AddNuevaSolicitud(m *TrSolicitudCrear) (err error) {
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

// GetAllSolicitudes Transacción para consultar todas las producciones con toda la información de las mismas
func GetAllSolicitudes() (v []interface{}, err error) {
	o := orm.NewOrm()
	var solicitantes []*Solicitante
	if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Activo", true).All(&solicitantes); err == nil {
		for _, solicitante := range solicitantes {

			solicitud := solicitante.SolicitudId

			var solicitanteSolicitud []Solicitante
			if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&solicitanteSolicitud); err != nil {
				return nil, err
			}

			var evolucionEstado SolicitudEvolucionEstado
			if _, err := o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&evolucionEstado); err != nil {
				return nil, err
			}

			var observaciones []Observacion
			if _, err := o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&observaciones); err != nil {
				return nil, err
			}

			v = append(v, map[string]interface{}{
				"Id":                    solicitud.Id,
				"EstadoTipoSolicitudId": solicitud.EstadoTipoSolicitudId,
				"Referencia":            solicitud.Referencia,
				"Resultado":             solicitud.Resultado,
				"FechaRadicacion":       solicitud.FechaRadicacion,
				"EvolucionEstado":       &evolucionEstado,
				"Solicitantes":          &solicitanteSolicitud,
				"Observaciones":         &observaciones,
			})
		}

		return v, nil
	}
	return nil, err
}

// GetSolicitudesByPersona Transacción para consultar todas las solicitudes con toda la información de las mismas
func GetSolicitudesByPersona(persona int) (v []interface{}, err error) {
	o := orm.NewOrm()
	var solicitantes []*Solicitante
	if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("tercero_id", persona).Filter("SolicitudId__Activo", true).All(&solicitantes); err == nil {
		for _, solicitante := range solicitantes {

			solicitud := solicitante.SolicitudId

			var solicitantesSolicitud []Solicitante
			if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&solicitantesSolicitud); err != nil {
				return nil, err
			}

			var evolucionEstado SolicitudEvolucionEstado
			if _, err := o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&evolucionEstado); err != nil {
				return nil, err
			}

			var observaciones []Observacion
			if _, err := o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&observaciones); err != nil {
				return nil, err
			}

			v = append(v, map[string]interface{}{
				"Id":                    solicitud.Id,
				"EstadoTipoSolicitudId": solicitud.EstadoTipoSolicitudId,
				"Referencia":            solicitud.Referencia,
				"Resultado":             solicitud.Resultado,
				"FechaRadicacion":       solicitud.FechaRadicacion,
				"EvolucionEstado":       &evolucionEstado,
				"Solicitantes":          &solicitantesSolicitud,
				"Observaciones":         &observaciones,
			})
		}

		return v, nil
	}
	return nil, err
}
