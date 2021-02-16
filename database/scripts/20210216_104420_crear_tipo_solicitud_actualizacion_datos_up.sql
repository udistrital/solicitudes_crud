--Insert de actualización de identificación en la tabla tipo de solicitud
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (3, 'Actualizar ID', 'Solicitud de actualizacion de la identificacion del estudiante', 'ActID', 3, now(), now(), true);
--Insert de actualización de nombres en la tabla tipo de solicitud
INSERT INTO solicitud.tipo_solicitud(id, nombre, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)VALUES (4, 'Actualizar Nombre', 'Solicitud de actualizacion de nombre del estudiante', 'ActNombre', 4, now(), now(), true);
