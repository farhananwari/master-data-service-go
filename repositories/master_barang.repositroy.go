package repositories

import (
	"master-data-service-go/models"

	"gorm.io/gorm"
)

/*
Interface per-domain
*/
type MasterBarangRepository interface {
	FindAll() ([]models.MasterBarang, error)
	FindByKodeBarang(kodeBarang string) (models.MasterBarang, error)
	Create(data models.MasterBarang) (models.MasterBarang, error)
	Update(data models.MasterBarang) (models.MasterBarang, error)
	Delete(kodeBarang string) error
	SoftDeleteByKodeBarang(kodeBarang string) error
}

/*
Implementasi
*/
type ImplMasterBarangRepo struct {
	db *gorm.DB
}

/*
Constructor
*/
func NewMasterBarangRepository(db *gorm.DB) MasterBarangRepository {
	return &ImplMasterBarangRepo{
		db: db,
	}
}

/*
CRUD IMPLEMENTATION
*/

func (r *ImplMasterBarangRepo) FindAll() ([]models.MasterBarang, error) {
	var datas []models.MasterBarang
	err := r.db.Find(&datas).Error
	return datas, err
}

func (r *ImplMasterBarangRepo) FindByKodeBarang(kodeBarang string) (models.MasterBarang, error) {
	var data models.MasterBarang
	err := r.db.Where("kode_barang = ?", kodeBarang).First(&data).Error
	return data, err
}

func (r *ImplMasterBarangRepo) Create(data models.MasterBarang) (models.MasterBarang, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *ImplMasterBarangRepo) Update(data models.MasterBarang) (models.MasterBarang, error) {
	err := r.db.Save(&data).Error
	return data, err
}

func (r *ImplMasterBarangRepo) Delete(kodeBarang string) error {
	return r.db.Where("kode_barang = ?", kodeBarang).
		Delete(&models.MasterBarang{}).Error
}

func (r *ImplMasterBarangRepo) SoftDeleteByKodeBarang(kodeBarang string) error {
	return r.db.Where("kode_barang = ?", kodeBarang).
		Delete(&models.MasterBarang{}).Error
}
