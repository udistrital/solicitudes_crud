package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:EstadoTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:PaqueteSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitanteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SolicitudEvolucionEstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:SoportePaqueteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoObservacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/solicitudes_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
