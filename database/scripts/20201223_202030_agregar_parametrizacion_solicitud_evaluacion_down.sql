--se eliminan los estados para el tipo de solicitud de evaluación
delete from solicitud.estado_tipo_solicitud
where tipo_solicitud_id = 2;

--Se eliminan los estados para evaluación
delete from solicitud.estado
where id = 10;
delete from solicitud.estado
where id = 11;
delete from solicitud.estado
where id = 12;
delete from solicitud.estado
where id = 13;

--se elimina el tipo de solicitud
delete from solicitud.tipo_solicitud
where id = 2;