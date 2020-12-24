--Se inserta tipo de solicitud evaluación
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (2, 'Evaluación', '', 'evalproddoc', 2, now(), now(), true);

--Se insertan los estados
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (10, 'En espera', '', 'ESP', 1, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (11, 'Rechazada', '', 'REC', 2, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (12, 'En espera de evaluación', '', 'ESE', 3, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (13, 'Evaluada', '', 'EVA', 4, now(), now(), true);

--Se insertan los estados por tipo de solicitud
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (10, 2, 10, 46, 14, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (11, 2, 11, 46, 0, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (12, 2, 12, 46, 14, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (13, 2, 13, 46, 0, now(), now(), true);