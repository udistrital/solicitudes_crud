package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "strconv"

	"github.com/udistrital/solicitudes_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// TrSolicitudController ...
type TrSolicitudController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrSolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAllByPersona", c.GetAllByPersona)
	// c.Mapping("Delete", c.Delete)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title PostTrSolicitud
// @Description create the TrSolicitud
// @Param	body		body 	models.TrSolicitud	true	"body for TrSolicitud content"
// @Success 201 {int} models.TrSolicitud
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrSolicitudController) Post() {
	var v models.TrSolicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddNuevaSolicitud(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TrSolicitud
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrSolicitud	true		"body for TrSolicitud content"
// @Success 200 {object} models.TrSolicitud
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TrSolicitudController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("Id es: " + idStr)
	id, _ := strconv.Atoi(idStr)
	var v models.TrSolicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Solicitud.Id = id
		if err := models.UpdateSolicitud(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TrSolicitudController
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router / [get]
func (c *TrSolicitudController) GetAll() {
	l, err := models.GetAllSolicitudes()
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetAllByPersona ...
// @Title Get All By Persona
// @Description get TrSolicitudController
// @Param	id		path 	string	true		"Persona"
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router /:persona [get]
func (c *TrSolicitudController) GetAllByPersona() {
	idPersonaStr := c.Ctx.Input.Param(":persona")
	idPersona, _ := strconv.Atoi(idPersonaStr)
	l, err := models.GetSolicitudesByPersona(idPersona)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

/*
// GetAllByFiltros ...
// @Title Get All By Filtros
// @Description get TrSolicitudController
// @Param	tercero_id		path 	string	true		"Tercero"
// @Param	estado_tipo_id		path 	string	true		"EstadoTipo"
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router /:tercero_id/:estado_tipo_id [get]
func (c *TrSolicitudController) GetAllByFiltros(){

	idTerceroStr := c.Ctx.Input.Param(":tercero_id")
	fmt.Println("idTerceroStr:", idTerceroStr)

	idTercero, _ := strconv.Atoi(idTerceroStr)
	fmt.Println("idTercero:", idTercero)

	idEstadoTipoStr := c.Ctx.Input.Param(":estado_tipo_id")
	idEstadoTipo, _ := strconv.Atoi(idEstadoTipoStr)

	l, err := models.GetSolicitudesByFiltros(idTercero, idEstadoTipo)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}
*/
