package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterPaketBarang struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	KodeBarang string    `gorm:"type:varchar(100);unique;not null;" json:"kode_barang"`
	Nama       string    `gorm:"type:varchar(100);not null;" json:"nama"`
	BarangID   string    `gorm:"type:varchar(100);not null;" json:"barang_id"`
	IsActive   bool      `gorm:"type:boolean;default:true;" json:"is_active"`
	CreatedBy  string    `gorm:"type:varchar(100);not null;" json:"created_by"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MasterPaketBarangItems struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	MasterPaketBarangID string    `gorm:"type:varchar(100);not null;" json:"master_paket_barang_id"`
	BarangID            string    `gorm:"type:varchar(100);not null;" json:"barang_id"`
	Jumlah              int       `gorm:"type:int;not null;" json:"jumlah"`
	CreatedBy           string    `gorm:"type:varchar(100);not null;" json:"created_by"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (MasterPaketBarang) TableName() string {
	return "master_paket_barang"
}

func (MasterPaketBarangItems) TableName() string {
	return "master_paket_barang_items"
}
