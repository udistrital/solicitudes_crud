package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/udistrital/solicitudes_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// TrPaqueteController ...
type TrPaqueteController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrPaqueteController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title PostTrPaquete
// @Description create the TrPaquete
// @Param	body		body 	models.TrPaquete	true	"body for TrPaquete content"
// @Success 201 {int} models.TrPaquete
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrPaqueteController) Post() {
	var v models.TrPaquete
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddNuevoPaquete(&v); err == nil {
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
// @Description update the TrPaquete
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrPaquete	true		"body for TrPaquete content"
// @Success 200 {object} models.TrPaquete
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TrPaqueteController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("Id es: " + idStr)
	id, _ := strconv.Atoi(idStr)
	var v models.TrPaquete
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Paquete.Id = id
		if err := models.UpdatePaquete(&v); err == nil {
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
// @Description get TrPaqueteController
// @Success 200 {object} models.TrPaqueteController
// @Failure 404 not found resource
// @router / [get]
func (c *TrPaqueteController) GetAll() {
	l, err := models.GetAllPaquetes()
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