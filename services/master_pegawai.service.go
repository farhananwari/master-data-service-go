package services

import (
	"errors"
	"strings"

	"master-data-service-go/dto"
	"master-data-service-go/models"
	"master-data-service-go/repositories"
)

type MasterPegawaiService interface {
	GetAll() ([]models.MasterPegawai, error)
	GetByID(id string) (models.MasterPegawai, error)
	Create(req dto.CreateMasterPegawaiRequest) (models.MasterPegawai, error)
	Update(id string, req dto.UpdateMasterPegawaiRequest) (models.MasterPegawai, error)
	Delete(id string) error
	SoftDelete(id string) error
}

type ImplMasterPegawaiService struct {
	repo repositories.MasterPegawaiRepository
}

func NewMasterPegawaiService(repo repositories.MasterPegawaiRepository) MasterPegawaiService {
	return &ImplMasterPegawaiService{repo: repo}
}

func (s *ImplMasterPegawaiService) GetAll() ([]models.MasterPegawai, error) {
	return s.repo.FindAll()
}

func (s *ImplMasterPegawaiService) GetByID(id string) (models.MasterPegawai, error) {
	if id == "" {
		return models.MasterPegawai{}, errors.New("id tidak boleh kosong")
	}
	return s.repo.FindByID(id)
}

func (s *ImplMasterPegawaiService) Create(
	req dto.CreateMasterPegawaiRequest,
) (models.MasterPegawai, error) {

	// mapping DTO -> Model
	data := models.MasterPegawai{
		Nama:      strings.TrimSpace(req.Nama),
		NoTelp:    strings.TrimSpace(req.NoTelp),
		Alamat:    strings.TrimSpace(req.Alamat),
		IsActive:  req.IsActive,
		CreatedBy: "system",
	}

	return s.repo.Create(data)
}

func (s *ImplMasterPegawaiService) Update(
	id string,
	req dto.UpdateMasterPegawaiRequest,
) (models.MasterPegawai, error) {

	// ambil data lama dari database
	existingData, err := s.repo.FindByID(id)
	if err != nil {
		return models.MasterPegawai{}, err
	}

	// mapping DTO -> Model
	if req.Nama != nil {
		existingData.Nama = strings.TrimSpace(*req.Nama)
	}
	if req.NoTelp != nil {
		existingData.NoTelp = strings.TrimSpace(*req.NoTelp)
	}
	if req.Alamat != nil {
		existingData.Alamat = strings.TrimSpace(*req.Alamat)
	}
	if req.IsActive != nil {
		existingData.IsActive = *req.IsActive
	}

	return s.repo.Update(existingData)
}

func (s *ImplMasterPegawaiService) Delete(id string) error {
	if id == "" {
		return errors.New("id tidak boleh kosong")
	}
	return s.repo.Delete(id)
}

func (s *ImplMasterPegawaiService) SoftDelete(id string) error {
	if id == "" {
		return errors.New("id tidak boleh kosong")
	}
	return s.repo.SoftDelete(id)
}
