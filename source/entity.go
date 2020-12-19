package source

// Client es la entidad de "cliente"
type Client struct {
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	CardType  int    `gorm:"column:card_type"`
	CardID    string `gorm:"column:id_card"`
}

// TableName regresa el nombre de la tabla para la entidad "cliente"
func (c Client) TableName() string {
	return "players_player"
}
