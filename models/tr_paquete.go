package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

//TrPaquete is...
type TrPaquete struct {
	Paquete            *Paquete
	SolicitudesPaquete *[]TrSolicitudPaquete
}

//TrSolicitudPaquete is...
type TrSolicitudPaquete struct {
	PaqueteSolicitud *PaqueteSolicitud
	//EstadoSolicitudPaquete *EstadoTipoSolicitud
	Observaciones     *[]Observacion
	EvolucionesEstado *[]SolicitudEvolucionEstado
}

//AddNuevoPaquete is...
func AddNuevoPaquete(m *TrPaquete) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	// fechas que se pueden comentar
	m.Paquete.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.Paquete.FechaModificacion = time_bogota.TiempoBogotaFormato()
	//fin

	m.Paquete.FechaComite = time_bogota.TiempoBogotaFormato()
	if idPaquete, errTr := o.Insert(m.Paquete); errTr == nil {

		fmt.Println(idPaquete)

		for _, v := range *m.SolicitudesPaquete {
			fmt.Println("procesando solicitud en paquete")

			v.PaqueteSolicitud.PaqueteId = m.Paquete

			v.PaqueteSolicitud.FechaCreacion = time_bogota.TiempoBogotaFormato()
			v.PaqueteSolicitud.FechaModificacion = time_bogota.TiempoBogotaFormato()
			fmt.Println("insertando solicitud en paquete")
			if _, errTr = o.Insert(v.PaqueteSolicitud); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}

			// fechas que se pueden comentar
			v.PaqueteSolicitud.SolicitudId.FechaRadicacion = time_bogota.TiempoBogotaFormato()
			v.PaqueteSolicitud.SolicitudId.FechaCreacion = time_bogota.TiempoBogotaFormato()
			//fin

			v.PaqueteSolicitud.SolicitudId.FechaModificacion = time_bogota.TiempoBogotaFormato()
			fmt.Println("actualizando solicitud en paquete")
			if _, errTr = o.Update(v.PaqueteSolicitud.SolicitudId, "EstadoTipoSolicitudId", "Referencia", "Resultado", "FechaModificacion"); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}

			fmt.Println("procesando estado")
			for _, v2 := range *v.EvolucionesEstado {
				var evolucionEstado SolicitudEvolucionEstado
				if errTr = o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", v.PaqueteSolicitud.SolicitudId.Id).One(&evolucionEstado); err == nil {

					if evolucionEstado.TerceroId != v2.TerceroId {
						evolucionEstado.TerceroId = v2.TerceroId
					}

					if evolucionEstado.EstadoTipoSolicitudIdAnterior != v2.EstadoTipoSolicitudIdAnterior {
						evolucionEstado.EstadoTipoSolicitudIdAnterior = v2.EstadoTipoSolicitudIdAnterior
					}

					if evolucionEstado.EstadoTipoSolicitudId != v2.EstadoTipoSolicitudId {
						evolucionEstado.EstadoTipoSolicitudId = v2.EstadoTipoSolicitudId
					}

					// fechas que se pueden comentar
					v2.FechaLimite = time_bogota.TiempoBogotaFormato()
					evolucionEstado.FechaCreacion = time_bogota.TiempoBogotaFormato()
					evolucionEstado.FechaLimite = time_bogota.TiempoBogotaFormato()
					//fin

					if evolucionEstado.FechaLimite != v2.FechaLimite {
						evolucionEstado.FechaLimite = v2.FechaLimite
					}

					if v2.Id != 0 {
						evolucionEstado.FechaModificacion = time_bogota.TiempoBogotaFormato()
						fmt.Println("actualizando estado[] solicitud en paquete")
						if _, errTr = o.Update(&evolucionEstado, "TerceroId", "EstadoTipoSolicitudIdAnterior", "EstadoTipoSolicitudId", "FechaLimite", "FechaModificacion"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
						v2.FechaCreacion = time_bogota.TiempoBogotaFormato()
						v2.FechaModificacion = time_bogota.TiempoBogotaFormato()
						v2.SolicitudId = v.PaqueteSolicitud.SolicitudId
						fmt.Println("insertando estado[] solicitud en paquete")
						if _, errTr = o.Insert(&v2); errTr != nil {
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

			fmt.Println("procesando observaciones")
			for _, v2 := range *v.Observaciones {
				if v2.Activo {
					var observacion Observacion

					if errTr = o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", v.PaqueteSolicitud.SolicitudId.Id).One(&observacion); err == nil {

						if observacion.TipoObservacionId != v2.TipoObservacionId {
							observacion.TipoObservacionId = v2.TipoObservacionId
						}

						if observacion.TerceroId != v2.TerceroId {
							observacion.TerceroId = v2.TerceroId
						}

						if observacion.Valor != v2.Valor {
							observacion.Valor = v2.Valor
						}

						// fechas que se pueden comentar
						observacion.FechaCreacion = time_bogota.TiempoBogotaFormato()
						//fin

						if v2.Id != 0 {
							observacion.FechaModificacion = time_bogota.TiempoBogotaFormato()
							if _, errTr = o.Update(&observacion, "TipoObservacionId", "TerceroId", "Valor", "FechaModificacion"); errTr != nil {
								err = errTr
								fmt.Println(err)
								_ = o.Rollback()
								return
							}
						} else {
							v2.FechaCreacion = time_bogota.TiempoBogotaFormato()
							v2.FechaModificacion = time_bogota.TiempoBogotaFormato()
							v2.SolicitudId = v.PaqueteSolicitud.SolicitudId
							if _, errTr = o.Insert(&v2); errTr != nil {
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
		}
		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println("err")
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

// UpdatePaquete is...
func UpdatePaquete(m *TrPaquete) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	v := Paquete{Id: m.Paquete.Id}
	// ascertain id exists in the database
	if errTr := o.Read(&v); errTr == nil {

		var num int64

		// fechas que se pueden comentar
		m.Paquete.FechaCreacion = time_bogota.TiempoBogotaFormato()
		m.Paquete.FechaModificacion = time_bogota.TiempoBogotaFormato()
		m.Paquete.FechaComite = time_bogota.TiempoCorreccionFormato(m.Paquete.FechaComite)
		// fin
		if num, errTr = o.Update(m.Paquete, "Nombre", "NumeroComite", "FechaComite", "FechaModificacion"); errTr == nil {
			fmt.Println("Number of records updated in database:", num)
			for _, v2 := range *m.SolicitudesPaquete {
				fmt.Println("procesando solicitud en paquete")
				var paqueteSolicitud PaqueteSolicitud
				if errTr = o.QueryTable(new(PaqueteSolicitud)).RelatedSel().Filter("Id", v2.PaqueteSolicitud.Id).One(&paqueteSolicitud); err == nil {

					if paqueteSolicitud.EstadoTipoSolicitudId != v2.PaqueteSolicitud.EstadoTipoSolicitudId {
						paqueteSolicitud.EstadoTipoSolicitudId = v2.PaqueteSolicitud.EstadoTipoSolicitudId
					}

					// fechas que se pueden comentar
					paqueteSolicitud.FechaCreacion = time_bogota.TiempoBogotaFormato()
					//fin

					if v2.PaqueteSolicitud.Id != 0 {
						paqueteSolicitud.FechaModificacion = time_bogota.TiempoBogotaFormato()
						fmt.Println("actualizando relación solicitud en paquete")
						if _, errTr = o.Update(&paqueteSolicitud, "EstadoTipoSolicitudId", "FechaModificacion"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
						v2.PaqueteSolicitud.FechaCreacion = time_bogota.TiempoBogotaFormato()
						v2.PaqueteSolicitud.FechaModificacion = time_bogota.TiempoBogotaFormato()
						v2.PaqueteSolicitud.PaqueteId.Id = int(v.Id)
						fmt.Println("insertando relación solicitud en paquete")
						if _, errTr = o.Insert(&v2.PaqueteSolicitud); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					}
				}

				// fechas que se pueden comentar
				v2.PaqueteSolicitud.SolicitudId.FechaCreacion = time_bogota.TiempoBogotaFormato()
				v2.PaqueteSolicitud.SolicitudId.FechaRadicacion = time_bogota.TiempoBogotaFormato()
				//fin

				v2.PaqueteSolicitud.SolicitudId.FechaModificacion = time_bogota.TiempoBogotaFormato()
				fmt.Println("actualizando solicitud en paquete")
				if _, errTr = o.Update(v2.PaqueteSolicitud.SolicitudId, "EstadoTipoSolicitudId", "Referencia", "Resultado", "SolicitudFinalizada", "FechaModificacion"); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}

				fmt.Println("procesando estado")
				for _, v3 := range *v2.EvolucionesEstado {
					if v3.Activo {
						var evolucionEstado SolicitudEvolucionEstado
						if errTr = o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", v2.PaqueteSolicitud.SolicitudId.Id).One(&evolucionEstado); err == nil {

							if evolucionEstado.TerceroId != v3.TerceroId {
								evolucionEstado.TerceroId = v3.TerceroId
							}

							if evolucionEstado.EstadoTipoSolicitudIdAnterior != v3.EstadoTipoSolicitudIdAnterior {
								evolucionEstado.EstadoTipoSolicitudIdAnterior = v3.EstadoTipoSolicitudIdAnterior
							}

							if evolucionEstado.EstadoTipoSolicitudId != v3.EstadoTipoSolicitudId {
								evolucionEstado.EstadoTipoSolicitudId = v3.EstadoTipoSolicitudId
							}

							// fechas que se pueden comentar
							// v3.FechaLimite = time_bogota.TiempoBogotaFormato()
							// evolucionEstado.FechaCreacion = time_bogota.TiempoBogotaFormato()
							// evolucionEstado.FechaLimite = time_bogota.TiempoBogotaFormato()
							//fin

							if evolucionEstado.FechaLimite != v3.FechaLimite {
								evolucionEstado.FechaLimite = v3.FechaLimite
							}

							if v3.Id != 0 {
								evolucionEstado.FechaModificacion = time_bogota.TiempoBogotaFormato()
								fmt.Println("actualizando estado[] solicitud en paquete")
								if _, errTr = o.Update(&evolucionEstado, "TerceroId", "EstadoTipoSolicitudIdAnterior", "EstadoTipoSolicitudId", "FechaLimite", "FechaModificacion"); errTr != nil {
									err = errTr
									fmt.Println(err)
									_ = o.Rollback()
									return
								}
							} else {
								v3.FechaCreacion = time_bogota.TiempoBogotaFormato()
								v3.FechaModificacion = time_bogota.TiempoBogotaFormato()
								v3.SolicitudId = v2.PaqueteSolicitud.SolicitudId
								fmt.Println("insertando estado[] solicitud en paquete")
								if _, errTr = o.Insert(&v3); errTr != nil {
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

				fmt.Println("procesando observaciones")
				for _, v3 := range *v2.Observaciones {
					if v3.Activo {
						var observacion Observacion

						if errTr = o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", v2.PaqueteSolicitud.SolicitudId.Id).One(&observacion); err == nil {

							if observacion.TipoObservacionId != v3.TipoObservacionId {
								observacion.TipoObservacionId = v3.TipoObservacionId
							}

							if observacion.TerceroId != v3.TerceroId {
								observacion.TerceroId = v3.TerceroId
							}

							if observacion.Valor != v3.Valor {
								observacion.Valor = v3.Valor
							}

							// fechas que se pueden comentar
							observacion.FechaCreacion = time_bogota.TiempoBogotaFormato()
							//fin

							if v3.Id != 0 {
								observacion.FechaModificacion = time_bogota.TiempoBogotaFormato()
								if _, errTr = o.Update(&observacion, "TipoObservacionId", "TerceroId", "Valor", "FechaModificacion"); errTr != nil {
									err = errTr
									fmt.Println(err)
									_ = o.Rollback()
									return
								}
							} else {
								v3.SolicitudId = v2.PaqueteSolicitud.SolicitudId
								v3.FechaCreacion = time_bogota.TiempoBogotaFormato()
								v3.FechaModificacion = time_bogota.TiempoBogotaFormato()
								if _, errTr = o.Insert(&v3); errTr != nil {
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
			}
		}
		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

// GetAllPaquetes is ...
func GetAllPaquetes() (v []interface{}, err error) {
	o := orm.NewOrm()

	var paquetes []*Paquete
	if _, err := o.QueryTable(new(Paquete)).RelatedSel().All(&paquetes); err == nil {
		for _, paquete := range paquetes {
			var paqueteSolicitudes []*PaqueteSolicitud
			if _, errS := o.QueryTable(new(PaqueteSolicitud)).Filter("PaqueteId", paquete.Id).All(&paqueteSolicitudes); errS == nil {
				v = append(v, map[string]interface{}{
					"Id":           paquete.Id,
					"Nombre":       paquete.Nombre,
					"NumeroComite": paquete.NumeroComite,
					"FechaComite":  paquete.FechaComite,
					"Activo":       paquete.Activo,
					"Solicitudes":  len(paqueteSolicitudes),
				})
			}
		}
		return v, nil
	}
	return nil, err
}
