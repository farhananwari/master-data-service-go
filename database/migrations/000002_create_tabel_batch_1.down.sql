-- =========================================
-- DROP TABLES (reverse order)
-- =========================================
DROP TABLE IF EXISTS master_supplier;

DROP TABLE IF EXISTS master_satuan_barang;

DROP TABLE IF EXISTS master_pegawai;

DROP TABLE IF EXISTS master_paket_barang_items;

DROP TABLE IF EXISTS master_paket_barang;

DROP TABLE IF EXISTS master_lokasi;

DROP TABLE IF EXISTS master_jenis_barang;

-- =========================================
-- DROP ENUM TYPES
-- =========================================
DROP TYPE IF EXISTS jenis_lokasi_enum;

-- =========================================
-- OPTIONAL: EXTENSION (usually kept)
-- =========================================
-- DROP EXTENSION IF EXISTS pgcrypto;