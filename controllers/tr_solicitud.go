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
	c.Mapping("GetAllByPersonaActive", c.GetAllByPersonaActive)
	c.Mapping("GetAllByPersonaInactive", c.GetAllByPersonaInactive)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllActive", c.GetAllActive)
	c.Mapping("GetAllInactive", c.GetAllInactive)
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
// @Param	persona		path 	string	true		"Persona"
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

// GetAllActive ...
// @Title Get All
// @Description get TrSolicitudController for active request
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router /active/ [get]
func (c *TrSolicitudController) GetAllActive() {
	var limit int64 = 0
	var offset int64 = 0

	// limit: 0 (default is 0)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	l, err := models.GetAllSolicitudesWithFilter(false, offset, limit)
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

// GetAllByPersonaActive ...
// @Title Get All By Persona
// @Description get TrSolicitudController
// @Param	persona		path 	string	true		"Persona"
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router /active/:persona [get]
func (c *TrSolicitudController) GetAllByPersonaActive() {
	idPersonaStr := c.Ctx.Input.Param(":persona")
	idPersona, _ := strconv.Atoi(idPersonaStr)
	l, err := models.GetSolicitudesByPersonaWithFilter(false, idPersona)
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

// GetAllInactive ...
// @Title Get All
// @Description get TrSolicitudController for active request
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router /inactive/ [get]
func (c *TrSolicitudController) GetAllInactive() {

	var limit int64 = 0
	var offset int64 = 0

	// limit: 0 (default is 0)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	l, err := models.GetAllSolicitudesWithFilter(true, offset, limit)
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

// GetAllByPersonaInactive ...
// @Title Get All By Persona
// @Description get TrSolicitudController
// @Param	persona		path 	string	true		"Persona"
// @Success 200 {object} models.TrSolicitudController
// @Failure 404 not found resource
// @router /inactive/:persona [get]
func (c *TrSolicitudController) GetAllByPersonaInactive() {
	idPersonaStr := c.Ctx.Input.Param(":persona")
	idPersona, _ := strconv.Atoi(idPersonaStr)
	l, err := models.GetSolicitudesByPersonaWithFilter(true, idPersona)
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
