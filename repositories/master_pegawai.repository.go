package repositories

import (
	"master-data-service-go/models"

	"gorm.io/gorm"
)

/*
Interface per-domain
*/
type MasterPegawaiRepository interface {
	FindAll() ([]models.MasterPegawai, error)
	FindByID(id string) (models.MasterPegawai, error)
	Create(data models.MasterPegawai) (models.MasterPegawai, error)
	Update(data models.MasterPegawai) (models.MasterPegawai, error)
	Delete(id string) error
	SoftDelete(id string) error
}

/*
Implementasi
*/
type ImplMasterPegawaiRepo struct {
	base *BaseRepository
}

/*
Constructor
*/
func NewMasterPegawaiRepository(db *gorm.DB) MasterPegawaiRepository {
	return &ImplMasterPegawaiRepo{
		base: NewBaseRepository(db),
	}
}

/*
	CRUD IMPLEMENTATION
*/

func (r *ImplMasterPegawaiRepo) FindAll() ([]models.MasterPegawai, error) {
	var datas []models.MasterPegawai
	err := r.base.FindAll(&datas)
	return datas, err
}

func (r *ImplMasterPegawaiRepo) FindByID(id string) (models.MasterPegawai, error) {
	var data models.MasterPegawai
	err := r.base.FindByID(&data, id)
	return data, err
}

func (r *ImplMasterPegawaiRepo) Create(data models.MasterPegawai) (models.MasterPegawai, error) {
	err := r.base.Create(&data)
	return data, err
}

func (r *ImplMasterPegawaiRepo) Update(data models.MasterPegawai) (models.MasterPegawai, error) {
	err := r.base.Update(&data)
	return data, err
}

func (r *ImplMasterPegawaiRepo) Delete(id string) error {
	return r.base.Delete(&models.MasterPegawai{}, id)
}

func (r *ImplMasterPegawaiRepo) SoftDelete(id string) error {
	return r.base.SoftDelete(&models.MasterPegawai{}, id)
}

/*
	CUSTOM QUERY (domain-specific)
*/
