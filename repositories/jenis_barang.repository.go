package repositories

import (
	"master-data-service-go/models"

	"gorm.io/gorm"
)

/*
Interface per-domain
*/
type MasterJenisBarangRepository interface {
	FindAll() ([]models.MasterJenisBarang, error)
	FindByID(id string) (models.MasterJenisBarang, error)
	Create(data models.MasterJenisBarang) (models.MasterJenisBarang, error)
	Update(data models.MasterJenisBarang) (models.MasterJenisBarang, error)
	Delete(id string) error
	SoftDelete(id string) error
}

/*
Implementasi
*/
type ImplMasterJenisBarangRepo struct {
	base *BaseRepository
}

/*
Constructor
*/
func NewMasterJenisBarangRepository(db *gorm.DB) MasterJenisBarangRepository {
	return &ImplMasterJenisBarangRepo{
		base: NewBaseRepository(db),
	}
}

/*
	CRUD IMPLEMENTATION
*/

func (r *ImplMasterJenisBarangRepo) FindAll() ([]models.MasterJenisBarang, error) {
	var datas []models.MasterJenisBarang
	err := r.base.FindAll(&datas)
	return datas, err
}

func (r *ImplMasterJenisBarangRepo) FindByID(id string) (models.MasterJenisBarang, error) {
	var data models.MasterJenisBarang
	err := r.base.FindByID(&data, id)
	return data, err
}

func (r *ImplMasterJenisBarangRepo) Create(data models.MasterJenisBarang) (models.MasterJenisBarang, error) {
	err := r.base.Create(&data)
	return data, err
}

func (r *ImplMasterJenisBarangRepo) Update(data models.MasterJenisBarang) (models.MasterJenisBarang, error) {
	err := r.base.Update(&data)
	return data, err
}

func (r *ImplMasterJenisBarangRepo) Delete(id string) error {
	return r.base.Delete(&models.MasterJenisBarang{}, id)
}

func (r *ImplMasterJenisBarangRepo) SoftDelete(id string) error {
	return r.base.SoftDelete(&models.MasterJenisBarang{}, id)
}

/*
	CUSTOM QUERY (domain-specific)
*/
