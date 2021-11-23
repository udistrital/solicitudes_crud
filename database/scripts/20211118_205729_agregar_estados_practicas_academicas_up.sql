--Insert de actualización de identificación en la tabla tipo de solicitud
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (23, 'Prácticas académicas', 'Solicitud de prácticas académica', 'SoPA', 1, now(), now(), true);
--Insert de actualización de nombres en la tabla tipo de solicitud
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (22, 'En consejo curricular', 'Pendiente de respuesta por parte del proyecto curricular', 'PenRPC', 2, now(), now(), true);
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (23, 'Ejecutada', 'Ejecutada y documentadad devidamente', 'EJEC', 2, now(), now(), true);

-- Radicada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 1, 4, 0, now(), now(), true);
-- Requiere modificación
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 2, 4, 0, now(), now(), true);
-- Rectificar
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 17, 4, 0, now(), now(), true);
-- En consejo curricular
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 22, 4, 0, now(), now(), true);
-- Rechazada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 11, 4, 0, now(), now(), true);
-- Acta aprobada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 9, 4, 0, now(), now(), true);
-- Ejecutada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (23, 23, 4, 0, now(), now(), true);
