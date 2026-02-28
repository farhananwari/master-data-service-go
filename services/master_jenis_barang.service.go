package services

import (
	"errors"
	"strings"

	"master-data-service-go/dto"
	"master-data-service-go/models"
	"master-data-service-go/repositories"
)

type MasterJenisBarangService interface {
	GetAll() ([]models.MasterJenisBarang, error)
	GetByID(id string) (models.MasterJenisBarang, error)
	Create(req dto.CreateMasterDataRequest) (models.MasterJenisBarang, error)
	Update(id string, req dto.UpdateMasterDataRequest) (models.MasterJenisBarang, error)
	Delete(id string) error
	SoftDelete(id string) error
}

type ImplMasterJenisBarangService struct {
	repo repositories.MasterJenisBarangRepository
}

func NewMasterJenisBarangService(repo repositories.MasterJenisBarangRepository) MasterJenisBarangService {
	return &ImplMasterJenisBarangService{repo: repo}
}

func (s *ImplMasterJenisBarangService) GetAll() ([]models.MasterJenisBarang, error) {
	return s.repo.FindAll()
}

func (s *ImplMasterJenisBarangService) GetByID(id string) (models.MasterJenisBarang, error) {
	if id == "" {
		return models.MasterJenisBarang{}, errors.New("id tidak boleh kosong")
	}
	return s.repo.FindByID(id)
}

func (s *ImplMasterJenisBarangService) Create(
	req dto.CreateMasterDataRequest,
) (models.MasterJenisBarang, error) {

	data := models.MasterJenisBarang{
		Nama:      strings.TrimSpace(req.Nama),
		IsActive:  true,
		CreatedBy: "system",
	}

	return s.repo.Create(data)
}

func (s *ImplMasterJenisBarangService) Update(
	id string,
	req dto.UpdateMasterDataRequest,
) (models.MasterJenisBarang, error) {

	existingData, err := s.repo.FindByID(id)
	if err != nil {
		return models.MasterJenisBarang{}, err
	}

	existingData.Nama = strings.TrimSpace(*req.Nama)
	existingData.CreatedBy = "system"

	return s.repo.Update(existingData)
}

func (s *ImplMasterJenisBarangService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *ImplMasterJenisBarangService) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}
