--Insert de actualización de identificación en la tabla tipo de solicitud
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (24, 'Derechos pecuniario', 'Solicitud de derechos pecuniario', 'SolDerPec', 12, now(), now(), true);

-- Solicitado
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (24, 15, 4, 0, now(), now(), true);
-- Ejecutada
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (24, 23, 4, 0, now(), now(), true);
