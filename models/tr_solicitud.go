package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type TrSolicitud struct {
	Solicitud         *Solicitud
	Solicitantes      *[]Solicitante
	Observaciones     *[]Observacion
	EvolucionesEstado *[]SolicitudEvolucionEstado
}

func AddNuevaSolicitud(m *TrSolicitud) (err error) {
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

		for _, v := range *m.EvolucionesEstado {
			v.SolicitudId.Id = int(idSolicitud)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		// for _, v := range *m.Observaciones {
		// 	v.SolicitudId.Id = int(idSolicitud)
		// 	if _, errTr = o.Insert(&v); errTr != nil {
		// 		err = errTr
		// 		fmt.Println(err)
		// 		_ = o.Rollback()
		// 		return
		// 	}
		// }

		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

func UpdateSolicitud(m *TrSolicitud) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	v := Solicitud{Id: m.Solicitud.Id}
	// ascertain id exists in the database
	if errTr := o.Read(&v); errTr == nil {
		var num int64
		if num, errTr = o.Update(m.Solicitud, "EstadoTipoSolicitudId", "Referencia", "Resultado", "FechaRadicacion", "FechaModificacion"); errTr == nil {
			fmt.Println("Number of records updated in database:", num)

			for _, v := range *m.EvolucionesEstado {
				var evolucionEstado SolicitudEvolucionEstado
				if errTr = o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", m.Solicitud.Id).One(&evolucionEstado); err == nil {

					if evolucionEstado.TerceroId != v.TerceroId {
						evolucionEstado.TerceroId = v.TerceroId
					}

					if evolucionEstado.EstadoTipoSolicitudIdAnterior != v.EstadoTipoSolicitudIdAnterior {
						evolucionEstado.EstadoTipoSolicitudIdAnterior = v.EstadoTipoSolicitudIdAnterior
					}

					if evolucionEstado.EstadoTipoSolicitudId != v.EstadoTipoSolicitudId {
						evolucionEstado.EstadoTipoSolicitudId = v.EstadoTipoSolicitudId
					}

					if evolucionEstado.FechaLimite != v.FechaLimite {
						evolucionEstado.FechaLimite = v.FechaLimite
					}

					if v.Id != 0 {
						if _, errTr = o.Update(&evolucionEstado, "TerceroId", "EstadoTipoSolicitudIdAnterior", "EstadoTipoSolicitudId", "FechaLimite", "FechaModificacion"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
						v.SolicitudId = m.Solicitud
						if _, errTr = o.Insert(&v); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					}
				} else {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			}

			for _, v := range *m.Observaciones {
				if v.Activo {
					var observacion Observacion
					if errTr = o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", m.Solicitud.Id).One(&observacion); err == nil {

						if observacion.TipoObservacionId != v.TipoObservacionId {
							observacion.TipoObservacionId = v.TipoObservacionId
						}

						if observacion.TerceroId != v.TerceroId {
							observacion.TerceroId = v.TerceroId
						}

						if observacion.Valor != v.Valor {
							observacion.Valor = v.Valor
						}
						if v.Id != 0 {
							observacion.FechaModificacion = time_bogota.TiempoBogotaFormato()
							if _, errTr = o.Update(&observacion, "TipoObservacionId", "TerceroId", "Valor", "FechaModificacion"); errTr != nil {
								err = errTr
								fmt.Println(err)
								_ = o.Rollback()
								return
							}
						} else {
							v.SolicitudId = m.Solicitud
							if _, errTr = o.Insert(&v); errTr != nil {
								err = errTr
								fmt.Println(err)
								_ = o.Rollback()
								return
							}
						}
					} else {
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					}
				}
			}

			_ = o.Commit()
		} else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}
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

			var evolucionEstado []SolicitudEvolucionEstado
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

			var evolucionEstado []SolicitudEvolucionEstado
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

/*
func GetSolicitudesByFiltros(tercero int, estadoTipo int)(v []interface{}, err error){

	orm.Debug = true
	var w io.Writer

	o := orm.NewOrm()
	var solicitudes []*Solicitud

	fmt.Println("tercero:", tercero)
	fmt.Println("estadoTipo:", estadoTipo)

	var qs orm.QuerySeter

	qs = o.QueryTable(new(Solicitud))
	qs = qs.Filter("Activo", true)

	if (estadoTipo > 0){
		qs = qs.Filter("estado_tipo_solicitud_id", estadoTipo)
	}

	if (tercero > 0){
		qs = qs.RelatedSel(new(Solicitante))
		qs = qs.Filter("Solicitante__tercero_id", tercero)
	}

	orm.DebugLog = orm.NewLog(w)
	fmt.Println("qs:", qs)

	if _, err := qs.All(&solicitudes); err == nil{
		fmt.Println("solicitudes:", solicitudes)
		for _, solicitud := range solicitudes {


			var solicitanteSolicitud []Solicitante
			if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Id", solicitud.Id).All(&solicitanteSolicitud); err != nil {
				return nil, err
			}

			v = append(v, map[string]interface{}{
				"Id":                    solicitud.Id,
				"EstadoTipoSolicitudId": solicitud.EstadoTipoSolicitudId,
				"Referencia":            solicitud.Referencia,
				"Resultado":             solicitud.Resultado,
				"FechaRadicacion":       solicitud.FechaRadicacion,
				// "EvolucionEstado":       &evolucionEstado,
				"Solicitantes":          &solicitanteSolicitud,
				// "Observaciones":         &observaciones,
			})

		}
		return v, nil
	}
	return nil, err
}
*/
