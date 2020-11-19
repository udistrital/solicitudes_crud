-- object: solicitud | type: SCHEMA --
-- DROP SCHEMA IF EXISTS solicitud CASCADE;
CREATE SCHEMA solicitud;
-- ddl-end --
COMMENT ON SCHEMA solicitud IS 'Esquema de la base de datos que gestionará las solicitudes.';


-- object: solicitud.solicitud | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.solicitud CASCADE;
CREATE TABLE solicitud.solicitud(
	id serial NOT NULL,
	estado_tipo_solicitud_id integer NOT NULL,
	referencia json NOT NULL,
	resultado json,
	fecha_radicacion timestamp,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	solicitud_padre_id integer,
	CONSTRAINT pk_solicitud PRIMARY KEY (id)

);

-- object: solicitud.tipo_observacion | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.tipo_observacion CASCADE;
CREATE TABLE solicitud.tipo_observacion(
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_tipo_observacion PRIMARY KEY (id)

);

-- object: solicitud.observacion | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.observacion CASCADE;
CREATE TABLE solicitud.observacion(
	id bigserial NOT NULL,
	tipo_observacion_id integer NOT NULL,
	solicitud_id integer NOT NULL,
	tercero_id integer NOT NULL,
	valor text NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_observacion PRIMARY KEY (id)

);

-- object: solicitud.paquete | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.paquete CASCADE;
CREATE TABLE solicitud.paquete(
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_paquete PRIMARY KEY (id)

);

-- object: solicitud.paquete_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.paquete_solicitud CASCADE;
CREATE TABLE solicitud.paquete_solicitud(
	id bigserial NOT NULL,
	paquete_id integer NOT NULL,
	solicitud_id integer NOT NULL,
	estado_tipo_solicitud_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_paquete_solicitud PRIMARY KEY (id)

);

-- object: solicitud.estado | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.estado CASCADE;
CREATE TABLE solicitud.estado(
	id serial NOT NULL,
	nombre varchar(100),
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_estado PRIMARY KEY (id)

);

-- object: solicitud.tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.tipo_solicitud CASCADE;
CREATE TABLE solicitud.tipo_solicitud(
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_tipo_solicitud PRIMARY KEY (id)

);

-- object: solicitud.soporte_paquete | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.soporte_paquete CASCADE;
CREATE TABLE solicitud.soporte_paquete(
	id serial NOT NULL,
	paquete_id integer NOT NULL,
	documento_id integer NOT NULL,
	descripcion character varying(250),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_soporte_paquete PRIMARY KEY (id)

);

-- object: solicitud.estado_tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.estado_tipo_solicitud CASCADE;
CREATE TABLE solicitud.estado_tipo_solicitud(
	id serial NOT NULL,
	tipo_solicitud integer NOT NULL,
	estado_id integer NOT NULL,
	dependencia_id integer NOT NULL,
	numero_dias integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_tipo_solicitud PRIMARY KEY (id)

);

-- object: solicitud.solicitante | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.solicitante CASCADE;
CREATE TABLE solicitud.solicitante(
	id bigserial NOT NULL,
	tercero_id integer NOT NULL,
	solicitud_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_solicitante PRIMARY KEY (id)

);

-- object: solicitud.solicitud_evolucion_estado | type: TABLE --
-- DROP TABLE IF EXISTS solicitud.solicitud_evolucion_estado CASCADE;
CREATE TABLE solicitud.solicitud_evolucion_estado(
	id bigserial NOT NULL,
	tercero_id integer NOT NULL,
	solicitud_id integer NOT NULL,
	estado_tipo_solicitud_id_anterior integer,
	estado_tipo_solicitud_id integer NOT NULL,
	fecha_limite timestamp,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_solicitud_evolucion_estado PRIMARY KEY (id)

);
-- ddl-end --



-- object: fk_solicitud_estado_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.solicitud DROP CONSTRAINT IF EXISTS fk_solicitud_estado_tipo_solicitud CASCADE;
ALTER TABLE solicitud.solicitud ADD CONSTRAINT fk_solicitud_estado_tipo_solicitud FOREIGN KEY (estado_tipo_solicitud_id)
REFERENCES solicitud.estado_tipo_solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_solicitud_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.solicitud DROP CONSTRAINT IF EXISTS fk_solicitud_solicitud CASCADE;
ALTER TABLE solicitud.solicitud ADD CONSTRAINT fk_solicitud_solicitud FOREIGN KEY (solicitud_padre_id)
REFERENCES solicitud.solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_observacion_tipo_observacion | type: CONSTRAINT --
-- ALTER TABLE solicitud.observacion DROP CONSTRAINT IF EXISTS fk_observacion_tipo_observacion CASCADE;
ALTER TABLE solicitud.observacion ADD CONSTRAINT fk_observacion_tipo_observacion FOREIGN KEY (tipo_observacion_id)
REFERENCES solicitud.tipo_observacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_observacion_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.observacion DROP CONSTRAINT IF EXISTS fk_observacion_solicitud CASCADE;
ALTER TABLE solicitud.observacion ADD CONSTRAINT fk_observacion_solicitud FOREIGN KEY (solicitud_id)
REFERENCES solicitud.solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_paquete_solicitud_paquete | type: CONSTRAINT --
-- ALTER TABLE solicitud.paquete_solicitud DROP CONSTRAINT IF EXISTS fk_paquete_solicitud_paquete CASCADE;
ALTER TABLE solicitud.paquete_solicitud ADD CONSTRAINT fk_paquete_solicitud_paquete FOREIGN KEY (paquete_id)
REFERENCES solicitud.paquete (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_paquete_solicitud_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.paquete_solicitud DROP CONSTRAINT IF EXISTS fk_paquete_solicitud_solicitud CASCADE;
ALTER TABLE solicitud.paquete_solicitud ADD CONSTRAINT fk_paquete_solicitud_solicitud FOREIGN KEY (solicitud_id)
REFERENCES solicitud.solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_paquete_solicitud_estado_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.paquete_solicitud DROP CONSTRAINT IF EXISTS fk_paquete_solicitud_estado_tipo_solicitud CASCADE;
ALTER TABLE solicitud.paquete_solicitud ADD CONSTRAINT fk_paquete_solicitud_estado_tipo_solicitud FOREIGN KEY (estado_tipo_solicitud_id)
REFERENCES solicitud.estado_tipo_solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_soporte_paquete_paquete | type: CONSTRAINT --
-- ALTER TABLE solicitud.soporte_paquete DROP CONSTRAINT IF EXISTS fk_soporte_paquete_paquete CASCADE;
ALTER TABLE solicitud.soporte_paquete ADD CONSTRAINT fk_soporte_paquete_paquete FOREIGN KEY (paquete_id)
REFERENCES solicitud.paquete (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_estado_tipo_solicitud_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.estado_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_estado_tipo_solicitud_tipo_solicitud CASCADE;
ALTER TABLE solicitud.estado_tipo_solicitud ADD CONSTRAINT fk_estado_tipo_solicitud_tipo_solicitud FOREIGN KEY (tipo_solicitud)
REFERENCES solicitud.tipo_solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_estado_tipo_solicitud_estado | type: CONSTRAINT --
-- ALTER TABLE solicitud.estado_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_estado_tipo_solicitud_estado CASCADE;
ALTER TABLE solicitud.estado_tipo_solicitud ADD CONSTRAINT fk_estado_tipo_solicitud_estado FOREIGN KEY (estado_id)
REFERENCES solicitud.estado (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_solicitante_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.solicitante DROP CONSTRAINT IF EXISTS fk_solicitante_solicitud CASCADE;
ALTER TABLE solicitud.solicitante ADD CONSTRAINT fk_solicitante_solicitud FOREIGN KEY (solicitud_id)
REFERENCES solicitud.solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_solicitud_evolucion_estado_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.solicitud_evolucion_estado DROP CONSTRAINT IF EXISTS fk_solicitud_evolucion_estado_solicitud CASCADE;
ALTER TABLE solicitud.solicitud_evolucion_estado ADD CONSTRAINT fk_solicitud_evolucion_estado_solicitud FOREIGN KEY (solicitud_id)
REFERENCES solicitud.solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_solicitud_evolucion_estado_tipo_solicitud_anterior | type: CONSTRAINT --
-- ALTER TABLE solicitud.solicitud_evolucion_estado DROP CONSTRAINT IF EXISTS fk_solicitud_evolucion_estado_tipo_solicitud_anterior CASCADE;
ALTER TABLE solicitud.solicitud_evolucion_estado ADD CONSTRAINT fk_solicitud_evolucion_estado_tipo_solicitud_anterior FOREIGN KEY (estado_tipo_solicitud_id_anterior)
REFERENCES solicitud.estado_tipo_solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_solicitud_evolucion_estado_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE solicitud.solicitud_evolucion_estado DROP CONSTRAINT IF EXISTS fk_solicitud_evolucion_estado_tipo_solicitud CASCADE;
ALTER TABLE solicitud.solicitud_evolucion_estado ADD CONSTRAINT fk_solicitud_evolucion_estado_tipo_solicitud FOREIGN KEY (estado_tipo_solicitud_id)
REFERENCES solicitud.estado_tipo_solicitud (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
