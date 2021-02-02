package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	// "time"

	"github.com/astaxie/beego/orm"
)

type Solicitud struct {
	Id                    int                  `orm:"column(id);pk;auto"`
	EstadoTipoSolicitudId *EstadoTipoSolicitud `orm:"column(estado_tipo_solicitud_id);rel(fk)"`
	Referencia            string               `orm:"column(referencia);type(json)"`
	Resultado             string               `orm:"column(resultado);type(json);null"`
	FechaRadicacion       string               `orm:"column(fecha_radicacion);type(timestamp without time zone);null"`
	FechaCreacion         string               `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion     string               `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	SolicitudFinalizada   bool                 `orm:"column(solicitud_finalizada);null"`
	Activo                bool                 `orm:"column(activo)"`
	SolicitudPadreId      *Solicitud           `orm:"column(solicitud_padre_id);rel(fk);null"`
}

func (t *Solicitud) TableName() string {
	return "solicitud"
}

func init() {
	orm.RegisterModel(new(Solicitud))
}

// AddSolicitud insert a new Solicitud into database and returns
// last inserted Id on success.
func AddSolicitud(m *Solicitud) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudById retrieves Solicitud by Id. Returns error if
// Id doesn't exist
func GetSolicitudById(id int) (v *Solicitud, err error) {
	o := orm.NewOrm()
	v = &Solicitud{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitud retrieves all Solicitud matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitud(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Solicitud)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Solicitud
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSolicitud updates Solicitud by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudById(m *Solicitud) (err error) {
	o := orm.NewOrm()
	v := Solicitud{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitud deletes Solicitud by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitud(id int) (err error) {
	o := orm.NewOrm()
	v := Solicitud{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Solicitud{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetSolicitudesEvaluaciones(correo string) (ml []interface{}, err error) {
	fmt.Println("GetSolicitudesEvaluaciones")
	fmt.Println("correo: ", correo)

	var l []Solicitud
	o := orm.NewOrm()

	num, err := o.Raw(`SELECT 
							solHija.* 
						FROM solicitud.solicitud solHija
						INNER JOIN solicitud.estado_tipo_solicitud etsHija
							ON solHija.estado_tipo_solicitud_id = etsHija.id
						WHERE solHija.estado_tipo_solicitud_id IS NOT NULL
						AND etsHija.tipo_solicitud_id = 2
						AND solHija.referencia->>'Correo' = ?`, correo).QueryRows(&l)
	if err == nil {
		fmt.Println("num sols: ", num)
		for _, v := range l {
			var estadoTipoSolicitud EstadoTipoSolicitud
			if _, err := o.QueryTable(new(EstadoTipoSolicitud)).RelatedSel().Filter("Id", v.EstadoTipoSolicitudId).All(&estadoTipoSolicitud); err != nil {
				return nil, err
			}

			var solicitudPadre Solicitud
			if _, err := o.QueryTable(new(Solicitud)).RelatedSel().Filter("Id", v.SolicitudPadreId).All(&solicitudPadre); err != nil {
				return nil, err
			}

			var solicitanteSolicitud []Solicitante
			if _, err := o.QueryTable(new(Solicitante)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&solicitanteSolicitud); err != nil {
				return nil, err
			}

			var evolucionEstado []SolicitudEvolucionEstado
			if _, err := o.QueryTable(new(SolicitudEvolucionEstado)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&evolucionEstado); err != nil {
				return nil, err
			}

			// var observaciones []Observacion
			// if _, err := o.QueryTable(new(Observacion)).RelatedSel().Filter("SolicitudId__Id", v.Id).All(&observaciones); err != nil {
			// 	return nil, err
			// }

			ml = append(ml, map[string]interface{}{
				"Id":                    v.Id,
				"EstadoTipoSolicitudId": estadoTipoSolicitud,
				"Referencia":            v.Referencia,
				// "Resultado":             solicitudPadre.Resultado,
				"FechaRadicacion":  v.FechaRadicacion,
				"SolicitudPadreId": solicitudPadre,
				"EvolucionEstado":  &evolucionEstado,
				"Solicitantes":     &solicitanteSolicitud,
				// "Observaciones":         &observaciones,
			})
		}
		return ml, nil
	}
	fmt.Println(err)
	return nil, err
}
