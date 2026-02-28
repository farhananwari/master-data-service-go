package repositories

import (
	"master-data-service-go/models"

	"gorm.io/gorm"
)

/*
Interface per-domain
*/
type MasterLokasiRepository interface {
	FindAll() ([]models.MasterLokasi, error)
	FindByID(id string) (models.MasterLokasi, error)
	Create(data models.MasterLokasi) (models.MasterLokasi, error)
	Update(data models.MasterLokasi) (models.MasterLokasi, error)
	Delete(id string) error
	SoftDelete(id string) error
}

/*
Implementasi
*/
type ImplMasterLokasiRepo struct {
	base *BaseRepository
}

/*
Constructor
*/
func NewMasterLokasiRepository(db *gorm.DB) MasterLokasiRepository {
	return &ImplMasterLokasiRepo{
		base: NewBaseRepository(db),
	}
}

/*
	CRUD IMPLEMENTATION
*/

func (r *ImplMasterLokasiRepo) FindAll() ([]models.MasterLokasi, error) {
	var datas []models.MasterLokasi
	err := r.base.FindAll(&datas)
	return datas, err
}

func (r *ImplMasterLokasiRepo) FindByID(id string) (models.MasterLokasi, error) {
	var data models.MasterLokasi
	err := r.base.FindByID(&data, id)
	return data, err
}

func (r *ImplMasterLokasiRepo) Create(data models.MasterLokasi) (models.MasterLokasi, error) {
	err := r.base.Create(&data)
	return data, err
}

func (r *ImplMasterLokasiRepo) Update(data models.MasterLokasi) (models.MasterLokasi, error) {
	err := r.base.Update(&data)
	return data, err
}

func (r *ImplMasterLokasiRepo) Delete(id string) error {
	return r.base.Delete(&models.MasterLokasi{}, id)
}

func (r *ImplMasterLokasiRepo) SoftDelete(id string) error {
	return r.base.SoftDelete(&models.MasterLokasi{}, id)
}

/*
	CUSTOM QUERY (domain-specific)
*/
