package target

import (
	"fmt"
)

// Service es la capa logica del paquete
type Service interface {
	SaveClientIfNotExists(string, string, int, string) (Client, error)
}

type service struct {
	repository Repository
}

// SaveClientIfNotExists registra el cliente si no existe
func (s service) SaveClientIfNotExists(firstName string, lastName string, cardType int, cardID string) (client Client, err error) {
	skip, err := s.repository.Exists(cardID)
	if err != nil || skip {
		fmt.Printf("- Skip CardID %s", cardID)
		return
	}
	client, err = s.repository.Save(firstName, lastName, cardType, cardID)
	fmt.Printf("- Target ID %d", client.ID)
	return
}

// CreateService crea una instancia de servicio
func CreateService(repository Repository) Service {
	return &service{repository: repository}
}
