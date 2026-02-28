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
	FindByID(id string) (models.MasterBarang, error)
	Create(data models.MasterBarang) (models.MasterBarang, error)
	Update(data models.MasterBarang) (models.MasterBarang, error)
	Delete(id string) error
	SoftDelete(id string) error
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

func (r *ImplMasterBarangRepo) FindByID(id string) (models.MasterBarang, error) {
	var data models.MasterBarang
	err := r.base.FindByID(&data, id)
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

func (r *ImplMasterBarangRepo) Delete(id string) error {
	return r.base.Delete(&models.MasterBarang{}, id)
}

func (r *ImplMasterBarangRepo) SoftDelete(id string) error {
	return r.base.SoftDelete(&models.MasterBarang{}, id)
}

/*
	CUSTOM QUERY (domain-specific)
*/
