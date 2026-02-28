package dto

type CreateMasterBarangRequest struct {
	KodeBarang     string `json:"kode_barang" binding:"required"`
	Nama           string `json:"nama" binding:"required"`
	SatuanBarangID string `json:"satuan_barang_id" binding:"required"`
	JenisBarangID  string `json:"jenis_barang_id" binding:"required"`
}

type UpdateMasterBarangRequest struct {
	KodeBarang     *string `json:"kode_barang"`
	Nama           *string `json:"nama"`
	SatuanBarangID *string `json:"satuan_barang_id"`
	JenisBarangID  *string `json:"jenis_barang_id"`
	IsActive       *bool   `json:"is_active"`
}
