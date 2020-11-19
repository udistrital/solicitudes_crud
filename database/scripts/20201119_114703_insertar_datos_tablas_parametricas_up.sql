--Inserta los estados para las solicitudes de producción docente para la oficina de docencia
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (1, 'Radicada', '', 'rad', 1, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (2, 'Evaluación', '', 'eval', 2, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (3, 'Completa', '', 'comp', 3, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (4, 'Evaluación nueva', '', 'neval', 4, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (5, 'Espera de par', '', 'esppar', 5, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (6, 'En evaluación', '', 'pteeval', 6, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (7, 'Verificada', '', 'ver', 7, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (8, 'Agrupada', '', 'agru', 8, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (9, 'Revisión', '', 'rev', 9, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (10, 'Preparada', '', 'pre', 10, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (11, 'Negada', '', 'neg', 11, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (12, 'Aprobada', '', 'apr', 12, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (13, 'Comentario', '', 'com', 13, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (14, 'Devuelta', '', 'dev', 14, now(), now(), true);

--Inserta el tipo de solicitud de producción docente para la oficina de docencia
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (1, 'Producción docentes oficina docencia', '', 'proddocen', 1, now(), now(), true);

--Inserta los tipos de observación para las solicitudes de producción docente para la oficina de docencia
INSERT INTO solicitud.tipo_observacion(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (1, 'Comentario', 'Indica que la observación es un comentario', 'com', 1, now(), now(), true);
INSERT INTO solicitud.tipo_observacion(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (2, 'Alerta', 'Indica que la observación es una alerta', 'alt', 2, now(), now(), true);

--Inserta los estados para las solicitudes de producción docente para la oficina de docencia
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (1, 1, 1, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (2, 1, 2, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (3, 1, 3, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (4, 1, 4, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (5, 1, 5, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (6, 1, 6, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (7, 1, 7, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (8, 1, 8, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (9, 1, 9, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (10, 1, 10, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (11, 1, 11, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (12, 1, 12, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (13, 1, 13, 1, 3, now(), now(), true);
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (14, 1, 14, 1, 3, now(), now(), true);