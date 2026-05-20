package services

import (
	"errors"
	"master-data-service-go/dto"
	"master-data-service-go/models"
	"master-data-service-go/repositories"
	"strings"
)

type MasterBarangService interface {
	GetAll() ([]models.MasterBarang, error)
	GetByID(id string) (models.MasterBarang, error)
	GetByKodeBarang(kodeBarang string) (models.MasterBarang, error)
	Create(req dto.CreateMasterBarangRequest) (models.MasterBarang, error)
	Update(id string, req dto.UpdateMasterBarangRequest) (models.MasterBarang, error)
	Delete(id string) error
	SoftDelete(id string) error
}

type ImplMasterBarangService struct {
	repo             repositories.MasterBarangRepository
	jenisBarangRepo  repositories.MasterJenisBarangRepository
	satuanBarangRepo repositories.MasterSatuanBarangRepository
}

func NewMasterBarangService(repo repositories.MasterBarangRepository) MasterBarangService {
	return &ImplMasterBarangService{repo: repo}
}

func (s *ImplMasterBarangService) GetAll() ([]models.MasterBarang, error) {
	return s.repo.FindAll()
}

func (s *ImplMasterBarangService) GetByID(id string) (models.MasterBarang, error) {
	return s.repo.FindByID(id)
}

func (s *ImplMasterBarangService) GetByKodeBarang(kodeBarang string) (models.MasterBarang, error) {
	return s.repo.FindByKodeBarang(kodeBarang)
}
func (s *ImplMasterBarangService) Create(req dto.CreateMasterBarangRequest) (models.MasterBarang, error) {
	// mapping DTO -> Model
	data := models.MasterBarang{
		KodeBarang:     strings.TrimSpace(req.KodeBarang),
		Nama:           strings.TrimSpace(req.Nama),
		SatuanBarangID: req.SatuanBarangID,
		JenisBarangID:  req.JenisBarangID,
		IsActive:       true,
		CreatedBy:      "system", // nanti dari auth context
	}

	// validasi jenis barang
	if _, err := s.jenisBarangRepo.FindByID(req.JenisBarangID); err != nil {
		return models.MasterBarang{}, errors.New("jenis barang tidak ditemukan")
	}

	// validasi satuan barang
	if _, err := s.satuanBarangRepo.FindByID(req.SatuanBarangID); err != nil {
		return models.MasterBarang{}, errors.New("satuan barang tidak ditemukan")
	}

	// create
	return s.repo.Create(data)
}
func (s *ImplMasterBarangService) Update(id string, req dto.UpdateMasterBarangRequest) (models.MasterBarang, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return models.MasterBarang{}, err
	}

	if req.Nama != nil {
		existing.Nama = strings.TrimSpace(*req.Nama)
	}
	if req.SatuanBarangID != nil {
		existing.SatuanBarangID = *req.SatuanBarangID
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	return s.repo.Update(id, existing)
}
func (s *ImplMasterBarangService) Delete(id string) error {

	return s.repo.Delete(id)
}
func (s *ImplMasterBarangService) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}
