package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterBarang struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	KodeBarang     string    `gorm:"type:varchar(100);unique;not null;" json:"kode_barang"`
	Nama           string    `gorm:"type:varchar(100);not null;" json:"nama"`
	SatuanBarangID string    `gorm:"type:varchar(100);not null;" json:"satuan_barang_id"`
	JenisBarangID  string    `gorm:"type:varchar(100);not null;" json:"jenis_barang_id"`
	IsActive       bool      `gorm:"type:boolean;default:true;" json:"is_active"`
	CreatedBy      string    `gorm:"type:varchar(100);not null;" json:"created_by"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (MasterBarang) TableName() string {
	return "master_barang"
}
