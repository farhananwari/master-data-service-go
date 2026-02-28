package services

import (
	"errors"
	"strings"

	"master-data-service-go/dto"
	"master-data-service-go/models"
	"master-data-service-go/repositories"
)

type MasterSatuanBarangService interface {
	GetAll() ([]models.MasterSatuanBarang, error)
	GetByID(id string) (models.MasterSatuanBarang, error)
	Create(req dto.CreateMasterDataRequest) (models.MasterSatuanBarang, error)
	Update(id string, req dto.UpdateMasterDataRequest) (models.MasterSatuanBarang, error)
	Delete(id string) error
	SoftDelete(id string) error
}

type ImplMasterSatuanBarangService struct {
	repo repositories.MasterSatuanBarangRepository
}

func NewMasterSatuanBarangService(repo repositories.MasterSatuanBarangRepository) MasterSatuanBarangService {
	return &ImplMasterSatuanBarangService{repo: repo}
}

func (s *ImplMasterSatuanBarangService) GetAll() ([]models.MasterSatuanBarang, error) {
	return s.repo.FindAll()
}

func (s *ImplMasterSatuanBarangService) GetByID(id string) (models.MasterSatuanBarang, error) {
	if id == "" {
		return models.MasterSatuanBarang{}, errors.New("id tidak boleh kosong")
	}
	return s.repo.FindByID(id)
}

func (s *ImplMasterSatuanBarangService) Create(
	req dto.CreateMasterDataRequest,
) (models.MasterSatuanBarang, error) {

	data := models.MasterSatuanBarang{
		Nama:      strings.TrimSpace(req.Nama),
		IsActive:  true,
		CreatedBy: "system",
	}

	return s.repo.Create(data)
}

func (s *ImplMasterSatuanBarangService) Update(
	id string,
	req dto.UpdateMasterDataRequest,
) (models.MasterSatuanBarang, error) {

	existingData, err := s.repo.FindByID(id)
	if err != nil {
		return models.MasterSatuanBarang{}, err
	}

	existingData.Nama = strings.TrimSpace(*req.Nama)
	existingData.CreatedBy = "system"

	return s.repo.Update(existingData)
}

func (s *ImplMasterSatuanBarangService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *ImplMasterSatuanBarangService) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}
