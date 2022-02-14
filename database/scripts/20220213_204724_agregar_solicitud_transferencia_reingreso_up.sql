--Insert de solicitudes de transferencias
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (25, 'Transferencias y reintegros', 'Solicitudes de reintegros, transferencias externas e internas', 'TrnRe', 1, now(), now(), true);

--Insert de actualización de nombres en la tabla tipo de solicitud
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (24, 'Aprobada', 'Se aprueba la solicitud por parte de la persona responsable', 'SAPR', 1, now(), now(), true);

-- Radicada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (25, 1, 4, 0, now(), now(), true);
-- Rechazada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (25, 11, 4, 0, now(), now(), true);
-- Aprobada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (25, 24, 4, 0, now(), now(), true);