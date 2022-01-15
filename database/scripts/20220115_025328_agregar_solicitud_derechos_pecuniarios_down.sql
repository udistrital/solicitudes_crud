DELETE FROM solicitud.estado WHERE tipo_solicitud_id = (SELECT id FROM solicitud.tipo_solicitud WHERE codigo_abreviacion = 'SolDerPec');

DELETE FROM solicitud.tipo_solicitud WHERE codigo_abreviacion = 'SolDerPec';