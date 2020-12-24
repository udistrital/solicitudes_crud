package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type TrPaquete struct {
	Paquete            *Paquete
	SolicitudesPaquete *[]TrSolicitudPaquete
}

type TrSolicitudPaquete struct {
	PaqueteSolicitud *PaqueteSolicitud
	//EstadoSolicitudPaquete *EstadoTipoSolicitud
	Observaciones     *[]Observacion
	EvolucionesEstado *[]SolicitudEvolucionEstado
}

func AddNuevoPaquete(m *TrPaquete) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	m.Paquete.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.Paquete.FechaModificacion = time_bogota.TiempoBogotaFormato()

	if idPaquete, errTr := o.Insert(m.Paquete); errTr == nil {

		fmt.Println(idPaquete)

		for _, v := range *m.SolicitudesPaquete {
			v.PaqueteSolicitud.PaqueteId = m.Paquete

			v.PaqueteSolicitud.FechaCreacion = time_bogota.TiempoBogotaFormato()
			v.PaqueteSolicitud.FechaModificacion = time_bogota.TiempoBogotaFormato()

			if _, errTr = o.Insert(v.PaqueteSolicitud); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}

			v.PaqueteSolicitud.SolicitudId.FechaCreacion = time_bogota.TiempoBogotaFormato()
			v.PaqueteSolicitud.SolicitudId.FechaModificacion = time_bogota.TiempoBogotaFormato()
			v.PaqueteSolicitud.SolicitudId.FechaRadicacion = time_bogota.TiempoBogotaFormato()

			if _, errTr = o.Update(v.PaqueteSolicitud.SolicitudId, "EstadoTipoSolicitudId", "Referencia", "Resultado", "FechaRadicacion", "FechaModificacion"); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}

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

					// v2.FechaLimite = time_bogota.TiempoBogotaFormato()
					v2.FechaCreacion = time_bogota.TiempoBogotaFormato()
					v2.FechaModificacion = time_bogota.TiempoBogotaFormato()

					// evolucionEstado.FechaLimite = time_bogota.TiempoBogotaFormato()
					evolucionEstado.FechaCreacion = time_bogota.TiempoBogotaFormato()
					evolucionEstado.FechaModificacion = time_bogota.TiempoBogotaFormato()

					if evolucionEstado.FechaLimite != v2.FechaLimite {
						evolucionEstado.FechaLimite = v2.FechaLimite
					}

					if v2.Id != 0 {
						if _, errTr = o.Update(&evolucionEstado, "TerceroId", "EstadoTipoSolicitudIdAnterior", "EstadoTipoSolicitudId", "FechaLimite", "FechaModificacion"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
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
						if v2.Id != 0 {
							observacion.FechaModificacion = time_bogota.TiempoBogotaFormato()
							if _, errTr = o.Update(&observacion, "TipoObservacionId", "TerceroId", "Valor", "FechaModificacion"); errTr != nil {
								err = errTr
								fmt.Println(err)
								_ = o.Rollback()
								return
							}
						} else {
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

func UpdatePaquete(m *TrPaquete) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	v := Paquete{Id: m.Paquete.Id}
	// ascertain id exists in the database
	if errTr := o.Read(&v); errTr == nil {
		var num int64
		if num, errTr = o.Update(m.Paquete, "Nombre", "FechaRadicacion", "FechaModificacion"); errTr == nil {
			fmt.Println("Number of records updated in database:", num)

			for _, v2 := range *m.SolicitudesPaquete {
				var paqueteSolicitud PaqueteSolicitud
				if errTr = o.QueryTable(new(PaqueteSolicitud)).RelatedSel().Filter("Id", v2.PaqueteSolicitud.Id).One(&paqueteSolicitud); err == nil {
					// if paqueteSolicitud.EstadoTipoSolicitudId.Id != v2.EstadoSolicitudPaquete.Id {
					// 	paqueteSolicitud.EstadoTipoSolicitudId.Id = v2.EstadoSolicitudPaquete.Id
					// }

					if v2.PaqueteSolicitud.Id != 0 {
						if _, errTr = o.Update(&paqueteSolicitud, "EstadoTipoSolicitudId", "FechaModificacion"); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					} else {
						v2.PaqueteSolicitud.PaqueteId.Id = int(v.Id)
						if _, errTr = o.Insert(&v2.PaqueteSolicitud); errTr != nil {
							err = errTr
							fmt.Println(err)
							_ = o.Rollback()
							return
						}
					}
				}
				/*
					for _, v3 := range *v2.EvolucionesEstado {
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

							if evolucionEstado.FechaLimite != v3.FechaLimite {
								evolucionEstado.FechaLimite = v3.FechaLimite
							}

							if v3.Id != 0 {
								if _, errTr = o.Update(&evolucionEstado, "TerceroId", "EstadoTipoSolicitudIdAnterior", "EstadoTipoSolicitudId", "FechaLimite", "FechaModificacion"); errTr != nil {
									err = errTr
									fmt.Println(err)
									_ = o.Rollback()
									return
								}
							} else {
								v3.SolicitudId = v2.PaqueteSolicitud.SolicitudId
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

					for _, v3 := range *v2.Observaciones{
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
				*/
			}
		}
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

func GetAllPaquetes() (v []interface{}, err error) {
	o := orm.NewOrm()

	var paquetes []*Paquete
	if _, err := o.QueryTable(new(Paquete)).RelatedSel().Filter("Activo", true).All(&paquetes); err == nil {
		for _, paquete := range paquetes {
			v = append(v, map[string]interface{}{
				"Id":           paquete.Id,
				"Nombre":       paquete.Nombre,
				"NumeroComite": paquete.NumeroComite,
				"FechaComite":  paquete.FechaComite,
			})
		}
		return v, nil
	}
	return nil, err
}
