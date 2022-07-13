DELETE FROM solicitud.estado_tipo_solicitud WHERE tipo_solicitud_id = (SELECT id FROM solicitud.estado WHERE codigo_abreviacion = 'PrEs');

DELETE FROM solicitud.estado WHERE codigo_abreviacion = 'PrEs';