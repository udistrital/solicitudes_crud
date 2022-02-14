DELETE FROM solicitud.estado_tipo_solicitud WHERE tipo_solicitud_id = (SELECT id FROM solicitud.tipo_solicitud WHERE codigo_abreviacion = 'TrnRe');

DELETE FROM solicitud.estado WHERE codigo_abreviacion = 'SAPR';

DELETE FROM solicitud.tipo_solicitud WHERE codigo_abreviacion = 'TrnRe';