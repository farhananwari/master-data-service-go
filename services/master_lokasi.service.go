package services

import (
	"errors"
	"strings"

	"master-data-service-go/dto"
	"master-data-service-go/models"
	"master-data-service-go/repositories"
)

type MasterLokasiService interface {
	GetAll() ([]models.MasterLokasi, error)
	GetByID(id string) (models.MasterLokasi, error)
	Create(req dto.CreateMasterLokasiRequest) (models.MasterLokasi, error)
	Update(id string, req dto.UpdateMasterLokasiRequest) (models.MasterLokasi, error)
	Delete(id string) error
	SoftDelete(id string) error
}

type ImplMasterLokasiService struct {
	repo repositories.MasterLokasiRepository
}

func NewMasterLokasiService(repo repositories.MasterLokasiRepository) MasterLokasiService {
	return &ImplMasterLokasiService{repo: repo}
}

func (s *ImplMasterLokasiService) GetAll() ([]models.MasterLokasi, error) {
	return s.repo.FindAll()
}

func (s *ImplMasterLokasiService) GetByID(id string) (models.MasterLokasi, error) {
	if id == "" {
		return models.MasterLokasi{}, errors.New("id tidak boleh kosong")
	}
	return s.repo.FindByID(id)
}

func (s *ImplMasterLokasiService) Create(
	req dto.CreateMasterLokasiRequest,
) (models.MasterLokasi, error) {

	data := models.MasterLokasi{
		Nama:        strings.TrimSpace(req.Nama),
		JenisLokasi: strings.TrimSpace(req.JenisLokasi),
		Alamat:      strings.TrimSpace(req.Alamat),
		NoTelp:      strings.TrimSpace(req.NoTelp),
		IsActive:    req.IsActive,
		CreatedBy:   "system",
	}

	return s.repo.Create(data)
}

func (s *ImplMasterLokasiService) Update(
	id string,
	req dto.UpdateMasterLokasiRequest,
) (models.MasterLokasi, error) {

	existingData, err := s.repo.FindByID(id)
	if err != nil {
		return models.MasterLokasi{}, err
	}

	if req.Nama != nil {
		existingData.Nama = strings.TrimSpace(*req.Nama)
	}
	if req.JenisLokasi != nil {
		existingData.JenisLokasi = strings.TrimSpace(*req.JenisLokasi)
	}
	if req.Alamat != nil {
		existingData.Alamat = strings.TrimSpace(*req.Alamat)
	}
	if req.NoTelp != nil {
		existingData.NoTelp = strings.TrimSpace(*req.NoTelp)
	}
	if req.IsActive != nil {
		existingData.IsActive = *req.IsActive
	}

	return s.repo.Update(existingData)
}

func (s *ImplMasterLokasiService) Delete(id string) error {
	if id == "" {
		return errors.New("id tidak boleh kosong")
	}
	return s.repo.Delete(id)
}

func (s *ImplMasterLokasiService) SoftDelete(id string) error {
	if id == "" {
		return errors.New("id tidak boleh kosong")
	}
	return s.repo.SoftDelete(id)
}
