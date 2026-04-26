package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

/*
	CRUD UMUM
*/

func (b *BaseRepository) FindAll(out interface{}) error {
	return b.db.Find(out).Error
}

func (b *BaseRepository) FindByID(out interface{}, id string) error {
	return b.db.First(out, "id = ?", id).Error
}

func (b *BaseRepository) FindByKodeBarang(out interface{}, KodeBarang string) error {
	return b.db.First(out, "kode_barang = ?", KodeBarang).Error
}

func (b *BaseRepository) Create(data interface{}) error {
	return b.db.Create(data).Error
}

func (b *BaseRepository) Update(data interface{}) error {
	return b.db.Save(data).Error
}

func (b *BaseRepository) Delete(model interface{}, id string) error {
	return b.db.Delete(model, "id = ?", id).Error
}

func (r *BaseRepository) SoftDelete(
	model interface{},
	id string,
) error {
	return r.db.
		Model(model).
		Where("id = ?", id).
		Update("is_active", false).
		Error
}

func (r *BaseRepository) SoftDeleteByKodeBarang(
	model interface{},
	KodeBarang string,
) error {
	return r.db.
		Model(model).
		Where("kode_barang = ?", KodeBarang).
		Update("is_active", false).
		Error
}

/*
	HELPER TAMBAHAN (opsional)
*/

func (b *BaseRepository) DB() *gorm.DB {
	return b.db
}
