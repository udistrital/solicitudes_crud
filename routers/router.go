// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/solicitudes_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado",
			beego.NSInclude(
				&controllers.EstadoController{},
			),
		),

		beego.NSNamespace("/paquete",
			beego.NSInclude(
				&controllers.PaqueteController{},
			),
		),

		beego.NSNamespace("/soporte_paquete",
			beego.NSInclude(
				&controllers.SoportePaqueteController{},
			),
		),

		beego.NSNamespace("/estado_tipo_solicitud",
			beego.NSInclude(
				&controllers.EstadoTipoSolicitudController{},
			),
		),

		beego.NSNamespace("/paquete_solicitud",
			beego.NSInclude(
				&controllers.PaqueteSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_solicitud",
			beego.NSInclude(
				&controllers.TipoSolicitudController{},
			),
		),

		beego.NSNamespace("/solicitud_evolucion_estado",
			beego.NSInclude(
				&controllers.SolicitudEvolucionEstadoController{},
			),
		),

		beego.NSNamespace("/tipo_observacion",
			beego.NSInclude(
				&controllers.TipoObservacionController{},
			),
		),

		beego.NSNamespace("/solicitante",
			beego.NSInclude(
				&controllers.SolicitanteController{},
			),
		),

		beego.NSNamespace("/observacion",
			beego.NSInclude(
				&controllers.ObservacionController{},
			),
		),

		beego.NSNamespace("/solicitud",
			beego.NSInclude(
				&controllers.SolicitudController{},
			),
		),

		beego.NSNamespace("/tr_solicitud_crear",
			beego.NSInclude(
				&controllers.TrSolicitudCrearController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
