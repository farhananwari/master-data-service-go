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
	FindByKodeBarang(KodeBarang string) (models.MasterBarang, error)
	Create(data models.MasterBarang) (models.MasterBarang, error)
	Update(data models.MasterBarang) (models.MasterBarang, error)
	Delete(KodeBarang string) error
	SoftDeleteByKodeBarang(KodeBarang string) error
}

/*
Implementasi
*/
type ImplMasterBarangRepo struct {
	base *BaseRepository
}

/*
Constructor
*/
func NewMasterBarangRepository(db *gorm.DB) MasterBarangRepository {
	return &ImplMasterBarangRepo{
		base: NewBaseRepository(db),
	}
}

/*
	CRUD IMPLEMENTATION
*/

func (r *ImplMasterBarangRepo) FindAll() ([]models.MasterBarang, error) {
	var datas []models.MasterBarang
	err := r.base.FindAll(&datas)
	return datas, err
}

func (r *ImplMasterBarangRepo) FindByKodeBarang(KodeBarang string) (models.MasterBarang, error) {
	var data models.MasterBarang
	err := r.base.FindByKodeBarang(&data, KodeBarang)
	return data, err
}

func (r *ImplMasterBarangRepo) Create(data models.MasterBarang) (models.MasterBarang, error) {
	err := r.base.Create(&data)
	return data, err
}

func (r *ImplMasterBarangRepo) Update(data models.MasterBarang) (models.MasterBarang, error) {
	err := r.base.Update(&data)
	return data, err
}

func (r *ImplMasterBarangRepo) Delete(KodeBarang string) error {
	return r.base.Delete(&models.MasterBarang{}, KodeBarang)
}

func (r *ImplMasterBarangRepo) SoftDeleteByKodeBarang(KodeBarang string) error {
	return r.base.SoftDeleteByKodeBarang(&models.MasterBarang{}, KodeBarang)
}

/*
	CUSTOM QUERY (domain-specific)
*/
