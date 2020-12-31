package models

import (
	"fmt"
	"strconv"

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
			if v.Activo {
				v.SolicitudId.Id = int(idSolicitud)
				if _, errTr = o.Insert(&v); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
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
func GetAllSolicitudes() (ml []interface{}, err error) {

	fmt.Println("GetAllSolicitudes")

	var l []Solicitud
	o := orm.NewOrm()
	
	num, err := o.Raw(`SELECT 
							sol.* 
						FROM solicitud.solicitud sol
						INNER JOIN solicitud.estado_tipo_solicitud ets
							ON sol.estado_tipo_solicitud_id = ets.id
						WHERE ets.tipo_solicitud_id = 1`).QueryRows(&l)
	if err == nil {
		fmt.Println("num sols: ", num)
		for _, v := range l {
			var estadoTipoSolicitud EstadoTipoSolicitud
			if _, err := o.QueryTable(new(EstadoTipoSolicitud)).RelatedSel().Filter("Id", v.EstadoTipoSolicitudId).All(&estadoTipoSolicitud); err != nil {
				return nil, err
			}
			
			var solicitantesSolicitud []Solicitante
			if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&solicitantesSolicitud); err != nil {
				return nil, err
			}

			var evolucionEstado []SolicitudEvolucionEstado
			if _, err := o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&evolucionEstado); err != nil {
				return nil, err
			}

			var observaciones []Observacion
			if _, err := o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&observaciones); err != nil {
				return nil, err
			}

			ml = append(ml, map[string]interface{}{
				"Id":                    v.Id,
				"EstadoTipoSolicitudId": estadoTipoSolicitud,
				"Referencia":            v.Referencia,
				"Resultado":             v.Resultado,
				"FechaRadicacion":       v.FechaRadicacion,
				"EvolucionEstado":       &evolucionEstado,
				"Solicitantes":          &solicitantesSolicitud,
				"Observaciones":         &observaciones,
			})
		}
		return ml, nil
	}

	return nil, err
}

// GetSolicitudesByPersona Transacción para consultar todas las solicitudes con toda la información de las mismas
func GetSolicitudesByPersona(persona int) (ml []interface{}, err error) {
	fmt.Println("GetSolicitudesByPersona: ", strconv.Itoa(persona))

	var l []Solicitud
	o := orm.NewOrm()
	
	num, err := o.Raw(`SELECT 
							sol.*
						FROM solicitud.solicitud sol
						INNER JOIN solicitud.estado_tipo_solicitud ets
							ON sol.estado_tipo_solicitud_id = ets.id
						INNER JOIN solicitud.solicitante sol2
							ON sol.id = sol2.solicitud_id
						WHERE ets.tipo_solicitud_id = 1
						AND sol2.tercero_id = ?`, persona).QueryRows(&l)
	if err == nil {
		fmt.Println("num sols: ", num)
		for _, v := range l {
			var estadoTipoSolicitud EstadoTipoSolicitud
			if _, err := o.QueryTable(new(EstadoTipoSolicitud)).RelatedSel().Filter("Id", v.EstadoTipoSolicitudId).All(&estadoTipoSolicitud); err != nil {
				return nil, err
			}
			
			var solicitantesSolicitud []Solicitante
			if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&solicitantesSolicitud); err != nil {
				return nil, err
			}

			var evolucionEstado []SolicitudEvolucionEstado
			if _, err := o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&evolucionEstado); err != nil {
				return nil, err
			}

			var observaciones []Observacion
			if _, err := o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&observaciones); err != nil {
				return nil, err
			}

			ml = append(ml, map[string]interface{}{
				"Id":                    v.Id,
				"EstadoTipoSolicitudId": estadoTipoSolicitud,
				"Referencia":            v.Referencia,
				"Resultado":             v.Resultado,
				"FechaRadicacion":       v.FechaRadicacion,
				"EvolucionEstado":       &evolucionEstado,
				"Solicitantes":          &solicitantesSolicitud,
				"Observaciones":         &observaciones,
			})
		}
		return ml, nil
	}

	return nil, err
}
