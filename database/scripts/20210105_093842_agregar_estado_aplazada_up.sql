INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (14, 'Aplazada', '', 'apla', 10, now(), now(), true);

INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)VALUES (14, 1, 14, 46, 0, now(), now(), true);