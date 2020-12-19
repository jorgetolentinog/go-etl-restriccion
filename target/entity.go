package target

import "time"

// Client es la entidad de "cliente"
type Client struct {
	ID        int    `gorm:"column:id"`
	FirstName string `gorm:"column:nombre"`
	LastName  string `gorm:"column:apellido"`
	Verified  bool   `gorm:"column:informacion_validada"`
}

// ClientIdentity es la entidad de "identidad de cliente"
type ClientIdentity struct {
	ClientID int    `gorm:"column:cliente_id"`
	CardType int    `gorm:"column:tipo_id"`
	CardID   string `gorm:"column:codigo"`
}

// ClientRestriction es la entidad de "restricción de cliente"
type ClientRestriction struct {
	ClientID  int       `gorm:"column:cliente_id"`
	CreatedAt time.Time `gorm:"column:fecha_creacion"`
}

// TableName regresa el nombre de la tabla para la entidad "cliente"
func (Client) TableName() string {
	return "clientes_cliente"
}

// TableName regresa el nombre de la tabla para la entidad "identidad de cliente"
func (ClientIdentity) TableName() string {
	return "clientes_identidadcliente"
}

// TableName regresa el nombre de la tabla para la entidad "restriccón de cliente"
func (ClientRestriction) TableName() string {
	return "clientes_restriccioncliente"
}
