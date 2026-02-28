-- =========================
-- MASTER LOKASI
-- =========================
CREATE TABLE master_lokasis (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    nama VARCHAR(100) NOT NULL,
    jenis_lokasi VARCHAR(100) NOT NULL,
    alamat VARCHAR(255),
    no_telp VARCHAR(20),

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_master_lokasis_updated_at
BEFORE UPDATE ON master_lokasis
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- =========================
-- MASTER JENIS BARANG
-- =========================
CREATE TABLE master_jenis_barangs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    nama VARCHAR(100) NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_master_jenis_barangs_updated_at
BEFORE UPDATE ON master_jenis_barangs
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- =========================
-- MASTER SATUAN BARANG
-- =========================
CREATE TABLE master_satuan_barangs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    nama VARCHAR(100) NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_master_satuan_barangs_updated_at
BEFORE UPDATE ON master_satuan_barangs
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


-- =========================
-- MASTER PEGAWAI
-- =========================
CREATE TABLE master_pegawais (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    nama VARCHAR(100) NOT NULL,
    alamat VARCHAR(255),
    no_telp VARCHAR(20) NOT NULL DEFAULT 'NULL',

    user_id UUID
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER trg_master_pegawais_updated_at
BEFORE UPDATE ON master_pegawais
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();