package dto

type CreateMasterPegawaiRequest struct {
	Nama       string `json:"nama" binding:"required"`
	NoTelp     string `json:"no_telp" binding:"required"`
	Alamat     string `json:"alamat" binding:"required"`
	IsActive   bool   `json:"is_active"`
	CreataedBy string `json:"created_by" binding:"required"`
}

type UpdateMasterPegawaiRequest struct {
	Nama      *string `json:"nama" binding:"required"`
	NoTelp    *string `json:"no_telp" binding:"required"`
	Alamat    *string `json:"alamat" binding:"required"`
	IsActive  *bool   `json:"is_active"`
	UpdatedBy *string `json:"updated_by" binding:"required"`
}
