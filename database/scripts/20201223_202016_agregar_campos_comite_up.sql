ALTER TABLE solicitud.paquete ADD COLUMN IF NOT EXISTS numero_comite varchar(100);
ALTER TABLE solicitud.paquete ADD COLUMN IF NOT EXISTS fecha_comite timestamp;