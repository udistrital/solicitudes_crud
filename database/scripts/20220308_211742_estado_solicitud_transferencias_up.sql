--Insert de actualización de nombres en la tabla tipo de solicitud
INSERT INTO solicitud.estado(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (26, 'Prueba especifica', 'Cuando un aspirante desea ingresar a un programa de facultad de artes', 'PrEs', 1, now(), now(), true);

-- Solicitado
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)
SELECT * FROM (SELECT 47, 25, 15, 4, 0, now(), now(), true) AS tmp
WHERE NOT EXISTS (
    SELECT id FROM solicitud.estado_tipo_solicitud WHERE id = 47
) LIMIT 1;

-- Requiere modificación
INSERT INTO solicitud.estado_tipo_solicitud(id, tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo)
SELECT * FROM (SELECT 48, 25, 2, 4, 0, now(), now(), true) AS tmp
WHERE NOT EXISTS (
    SELECT id FROM solicitud.estado_tipo_solicitud WHERE id = 48
) LIMIT 1;

-- Prueba especifica
INSERT INTO solicitud.estado_tipo_solicitud(tipo_solicitud_id, estado_id, dependencia_id, numero_dias, fecha_creacion, fecha_modificacion, activo) VALUES (25, 26, 4, 0, now(), now(), true);