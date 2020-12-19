package source

import "gorm.io/gorm"

// Repository es la capa de datos del paquete
type Repository interface {
	ListAll() ([]*Client, error)
}

type repository struct {
	db *gorm.DB
}

// ListAll regresa toda la lista de clientes
func (r *repository) ListAll() (list []*Client, err error) {
	err = r.db.Find(&list).Error
	return
}

// CreateRepository crea una instancia de repositorio
func CreateRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
