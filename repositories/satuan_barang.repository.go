package repositories

import (
	"master-data-service-go/models"

	"gorm.io/gorm"
)

/*
Interface per-domain
*/
type MasterSatuanBarangRepository interface {
	FindAll() ([]models.MasterSatuanBarang, error)
	FindByID(id string) (models.MasterSatuanBarang, error)
	Create(data models.MasterSatuanBarang) (models.MasterSatuanBarang, error)
	Update(data models.MasterSatuanBarang) (models.MasterSatuanBarang, error)
	Delete(id string) error
	SoftDelete(id string) error
}

/*
Implementasi
*/
type ImplMasterSatuanBarangRepo struct {
	base *BaseRepository
}

/*
Constructor
*/
func NewMasterSatuanBarangRepository(db *gorm.DB) MasterSatuanBarangRepository {
	return &ImplMasterSatuanBarangRepo{
		base: NewBaseRepository(db),
	}
}

/*
	CRUD IMPLEMENTATION
*/

func (r *ImplMasterSatuanBarangRepo) FindAll() ([]models.MasterSatuanBarang, error) {
	var datas []models.MasterSatuanBarang
	err := r.base.FindAll(&datas)
	return datas, err
}

func (r *ImplMasterSatuanBarangRepo) FindByID(id string) (models.MasterSatuanBarang, error) {
	var data models.MasterSatuanBarang
	err := r.base.FindByID(&data, id)
	return data, err
}

func (r *ImplMasterSatuanBarangRepo) Create(data models.MasterSatuanBarang) (models.MasterSatuanBarang, error) {
	err := r.base.Create(&data)
	return data, err
}

func (r *ImplMasterSatuanBarangRepo) Update(data models.MasterSatuanBarang) (models.MasterSatuanBarang, error) {
	err := r.base.Update(&data)
	return data, err
}

func (r *ImplMasterSatuanBarangRepo) Delete(id string) error {
	return r.base.Delete(&models.MasterSatuanBarang{}, id)
}

func (r *ImplMasterSatuanBarangRepo) SoftDelete(id string) error {
	return r.base.SoftDelete(&models.MasterSatuanBarang{}, id)
}

/*
	CUSTOM QUERY (domain-specific)
*/
