--Se actualizan 14 -> 2
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 2
where estado_tipo_solicitud_id_anterior = 14;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 2
where estado_tipo_solicitud_id = 14;

update solicitud.solicitud
set estado_tipo_solicitud_id = 2
where estado_tipo_solicitud_id = 14;

--Se actualizan 7 -> 4
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 4
where estado_tipo_solicitud_id_anterior = 7;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 4
where estado_tipo_solicitud_id = 7;

update solicitud.solicitud
set estado_tipo_solicitud_id = 4
where estado_tipo_solicitud_id = 7;

--Se actualizan 8 -> 6
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 6
where estado_tipo_solicitud_id_anterior = 8;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 6
where estado_tipo_solicitud_id = 8;

update solicitud.solicitud
set estado_tipo_solicitud_id = 6
where estado_tipo_solicitud_id = 8;

--Se actualizan 9 -> 7
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 7
where estado_tipo_solicitud_id_anterior = 9;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 7
where estado_tipo_solicitud_id = 9;

update solicitud.solicitud
set estado_tipo_solicitud_id = 7
where estado_tipo_solicitud_id = 9;

--Se actualizan 11 -> 8
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 8
where estado_tipo_solicitud_id_anterior = 11;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 8
where estado_tipo_solicitud_id = 11;

update solicitud.solicitud
set estado_tipo_solicitud_id = 8
where estado_tipo_solicitud_id = 11;

--Se actualizan 12 -> 9
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 9
where estado_tipo_solicitud_id_anterior = 12;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 9
where estado_tipo_solicitud_id = 12;

update solicitud.solicitud
set estado_tipo_solicitud_id = 9
where estado_tipo_solicitud_id = 12;

--Se actualizan las que no cruzaron
update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id_anterior = 1
where estado_tipo_solicitud_id_anterior > 9;

update solicitud.solicitud_evolucion_estado
set estado_tipo_solicitud_id = 1
where estado_tipo_solicitud_id > 9;

update solicitud.solicitud
set estado_tipo_solicitud_id = 1
where estado_tipo_solicitud_id > 9;

--Se eliminan los estados adicionales
delete from solicitud.estado_tipo_solicitud
where estado_id > 9;

delete from solicitud.estado
where id > 9;

--Se actualizan los estados
update solicitud.estado set nombre = 'Radicada', codigo_abreviacion = 'RAD' where id = 1;
update solicitud.estado set nombre = 'Requiere modificación', codigo_abreviacion = 'RMO' where id = 2;
update solicitud.estado set nombre = 'Para evaluación', codigo_abreviacion = 'PEV' where id = 3;
update solicitud.estado set nombre = 'Verificada', codigo_abreviacion = 'VER' where id = 4;
update solicitud.estado set nombre = 'En espera par', codigo_abreviacion = 'EEP' where id = 5;
update solicitud.estado set nombre = 'Agrupada', codigo_abreviacion = 'AGR' where id = 6;
update solicitud.estado set nombre = 'Revision', codigo_abreviacion = 'REV' where id = 7;
update solicitud.estado set nombre = 'Negada', codigo_abreviacion = 'NEG' where id = 8;
update solicitud.estado set nombre = 'Aprobada', codigo_abreviacion = 'APR' where id = 9;

