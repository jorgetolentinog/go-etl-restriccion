package target

import (
	"time"

	"gorm.io/gorm"
)

// Repository es la capa de datos del paquete
type Repository interface {
	Exists(string) (bool, error)
	Save(string, string, int, string) (Client, error)
}

type repository struct {
	db *gorm.DB
}

// Exists revisa si un cliente existe en la base de datos por numero de documento
func (r *repository) Exists(cardID string) (exists bool, err error) {
	var count int64
	filter := &ClientIdentity{CardID: cardID}
	err = r.db.Table(filter.TableName()).Where(filter).Count(&count).Error
	exists = count > 0
	return
}

// Save guarda los datos del cliente en diferentes tablas
func (r *repository) Save(firstName string, lastName string, cardType int, cardID string) (client Client, err error) {
	tx := r.db.Begin()
	client = Client{FirstName: firstName, LastName: lastName, Verified: true}
	err = tx.Create(&client).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Create(&ClientIdentity{ClientID: client.ID, CardType: cardType, CardID: cardID}).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Create(&ClientRestriction{ClientID: client.ID, CreatedAt: time.Now()}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// CreateRepository crea una instancia de repositorio
func CreateRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
