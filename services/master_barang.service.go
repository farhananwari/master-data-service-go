package services

import (
	"errors"
	"strings"

	"master-data-service-go/dto"
	"master-data-service-go/models"
	"master-data-service-go/repositories"
)

type MasterBarangService interface {
	GetAll() ([]models.MasterBarang, error)
	GetByID(KodeBarang string) (models.MasterBarang, error)
	Create(req dto.CreateMasterBarangRequest) (models.MasterBarang, error)
	Update(KodeBarang string, req dto.UpdateMasterBarangRequest) (models.MasterBarang, error)
	Delete(KodeBarang string) error
	SoftDelete(KodeBarang string) error
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

func (s *ImplMasterBarangService) GetByID(KodeBarang string) (models.MasterBarang, error) {
	if KodeBarang == "" {
		return models.MasterBarang{}, errors.New("kode barang tidak boleh kosong")
	}
	return s.repo.FindByKodeBarang(KodeBarang)
}

func (s *ImplMasterBarangService) Create(
	req dto.CreateMasterBarangRequest,
) (models.MasterBarang, error) {

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

func (s *ImplMasterBarangService) Update(
	KodeBarang string,
	req dto.UpdateMasterBarangRequest,
) (models.MasterBarang, error) {

	if KodeBarang == "" {
		return models.MasterBarang{}, errors.New("Kode barang tidak boleh kosong")
	}

	existing, err := s.repo.FindByKodeBarang(KodeBarang)
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

	return s.repo.Update(existing)
}

func (s *ImplMasterBarangService) Delete(KodeBarang string) error {
	if KodeBarang == "" {
		return errors.New("Kode barang tidak boleh kosong")
	}
	return s.repo.Delete(KodeBarang)
}

func (s *ImplMasterBarangService) SoftDelete(KodeBarang string) error {
	if KodeBarang == "" {
		return errors.New("Kode barang tidak boleh kosong")
	}
	return s.repo.SoftDeleteByKodeBarang(KodeBarang)
}
