# solicitudes_crud
El api proveé los métodos para la gestión de solicitudes.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# parametros de api
SOLICITUD_CRUD_HTTP_PORT=[Puerto de exposición del API]
SOLICITUD_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd
SOLICITUD_CRUD_PGUSER=[Usuario de BD]
SOLICITUD_CRUD_PGPASS=[Contraseña del usaurio de BD]
SOLICITUD_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
SOLICITUD_CRUD_PGPORT=[Puerto de la BD]
SOLICITUD_CRUD_PGDB=[Nombre de Base de Datos]
SOLICITUD_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con SOLICITUD_CRUD__...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/solicitudes_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/solicitudes_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
SOLICITUD_CRUD_HTTP_PORT=8080 SOLICITUD_CRUD_PGURL=127.0.0.1 SOLICITUD_CRUD_SOME_VARIABLE=some_value bee run
```
### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/entradas_crud

#2. Moverse a la carpeta del repositorio
cd solicitudes_crud

#3. Crear un fichero con el nombre **custom.env**
touch .env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```
### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/solicitudes_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/solicitudes_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/solicitudes_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/solicitudes_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/solicitudes_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/solicitudes_crud) |

## Modelo de Datos

[Modelo de datos API CRUD solicitudes](https://user-images.githubusercontent.com/44048060/99477113-d9d0d900-291f-11eb-999b-e187cdd823f7.png)


## Licencia

This file is part of solicitudes_crud.

solicitudes_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

solicitudes_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with solicitudes_crud. If not, see https://www.gnu.org/licenses/.
