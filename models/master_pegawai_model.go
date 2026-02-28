package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterPegawai struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;" json:"nama"`
	NoTelp    string    `gorm:"type:varchar(15);not null;" json:"no_telp"`
	Alamat    string    `gorm:"type:text;not null;" json:"alamat"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;" json:"user_id"`
	IsActive  bool      `gorm:"type:boolean;default:true;" json:"is_active"`
	CreatedBy string    `gorm:"type:varchar(50);not null;" json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
