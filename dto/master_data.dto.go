package dto

type CreateMasterDataRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateMasterDataRequest struct {
	Nama *string `json:"nama"`
}
