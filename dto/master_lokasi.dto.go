package dto

type CreateMasterLokasiRequest struct {
	Nama        string `json:"nama" binding:"required"`
	JenisLokasi string `json:"jenis_lokasi" `
	Alamat      string `json:"alamat"`
	NoTelp      string `json:"no_telp"`
	IsActive    bool   `json:"is_active"`
	CreatedBy   string `json:"created_by" binding:"required"`
}

type UpdateMasterLokasiRequest struct {
	Nama        *string `json:"nama"`
	JenisLokasi *string `json:"jenis_lokasi"`
	Alamat      *string `json:"alamat"`
	NoTelp      *string `json:"no_telp"`
	IsActive    *bool   `json:"is_active"`
}
