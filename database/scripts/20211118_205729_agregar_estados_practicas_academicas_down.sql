DELETE FROM solicitud.estado WHERE tipo_solicitud_id = (SELECT id FROM solicitud.tipo_solicitud WHERE codigo_abreviacion = 'SoPA');

DELETE FROM solicitud.estado WHERE codigo_abreviacion = 'PenRPC';
DELETE FROM solicitud.estado WHERE codigo_abreviacion = 'EJEC';

DELETE FROM solicitud.tipo_solicitud WHERE codigo_abreviacion = 'SoPA';