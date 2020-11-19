# Solicitudes_crud
El api proveé los métodos para la gestión de solicitudes.

## Especificaciones Técnicas

### Variables de Entorno
```shell
SOLICITUD_CRUD__PGDB=[nombre de la base de datos]
SOLICITUD_CRUD__PGPASS=[password del usuario]
SOLICITUD_CRUD__PGURLS=[direccion de la base de datos]
SOLICITUD_CRUD__PGUSER=[usuario con acceso a la base de datos]
SOLICITUD_CRUD__PGSCHEMA=[esquema donde se ubican las tablas]
API_PRODUCCION_ACADEMICA_HTTP_PORT=[puerto de ejecucion] bee run
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con SOLICITUD_CRUD__...

## Modelo de Datos

[Modelo de datos API CRUD solicitudes](https://user-images.githubusercontent.com/44048060/99477113-d9d0d900-291f-11eb-999b-e187cdd823f7.png)