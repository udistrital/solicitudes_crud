package controllers

import (
	"encoding/json"

	// "strconv"

	"github.com/udistrital/solicitudes_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TrSolicitudCrearController struct {
	beego.Controller
}

func (c *TrSolicitudCrearController) URLMapping() {
	c.Mapping("Post", c.Post)
	// c.Mapping("GetAllByPersona", c.GetAllByPersona)
	// c.Mapping("Delete", c.Delete)
	// c.Mapping("Put", c.Put)
}

// @Title PostTrSolicitudCrear
// @Description create the TrSolicitudCrear
// @Param	body		body 	models.TrSolicitudCrear	true	"body for TrSolicitudCrear content"
// @Success 201 {int} models.TrSolicitudCrear
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrSolicitudCrearController) Post() {
	var v models.TrSolicitudCrear
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
