--Se actualizan los nombre de los estado
update solicitud.estado set nombre = 'Radicada', codigo_abreviacion = 'RAD' where id = 1;
update solicitud.estado set nombre = 'Requiere modificación', codigo_abreviacion = 'RMO' where id = 2;
update solicitud.estado set nombre = 'En asignación de par', codigo_abreviacion = 'EEP' where id = 3;
update solicitud.estado set nombre = 'Preparada para presentar a comité', codigo_abreviacion = 'PPC' where id = 4;
update solicitud.estado set nombre = 'En evaluación de pares', codigo_abreviacion = 'EVP' where id = 5;
update solicitud.estado set nombre = 'Caso seleccionado', codigo_abreviacion = 'CSE' where id = 6;
update solicitud.estado set nombre = 'Revision previa al comité', codigo_abreviacion = 'RPC' where id = 7;
update solicitud.estado set nombre = 'Tratado en el comité', codigo_abreviacion = 'TRC' where id = 8;
update solicitud.estado set nombre = 'Acta aprobada', codigo_abreviacion = 'APR' where id = 9;

--Se actualiza el número de días para las alertas
update solicitud.estado_tipo_solicitud set numero_dias = 14 where estado_id = 1;
update solicitud.estado_tipo_solicitud set numero_dias = 45 where estado_id = 2;
update solicitud.estado_tipo_solicitud set numero_dias = 14 where estado_id = 3;
update solicitud.estado_tipo_solicitud set numero_dias = 35 where estado_id = 4;
update solicitud.estado_tipo_solicitud set numero_dias = 14 where estado_id = 5;
update solicitud.estado_tipo_solicitud set numero_dias = 7 where estado_id = 6;
update solicitud.estado_tipo_solicitud set numero_dias = 7 where estado_id = 7;
update solicitud.estado_tipo_solicitud set numero_dias = 0 where estado_id = 8;
update solicitud.estado_tipo_solicitud set numero_dias = 0 where estado_id = 9;