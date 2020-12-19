package source

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func setupTestCase(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func(t *testing.T)) {
	t.Log("setup test case")
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gdb, err := gorm.Open(sqlserver.New(sqlserver.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	return gdb, mock, func(t *testing.T) {
		t.Log("teardown test case")
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		assert.NoError(t, err)
	}
}

func TestRepositoryListAll__should_return_clients(t *testing.T) {
	db, mock, teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	rows := mock.NewRows([]string{"first_name", "last_name", "card_type", "id_card"}).
		AddRow("Jorge", "Tolentino", 1, "12346578").
		AddRow("Diego", "Conga", 2, "87654321")

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players_player`)).
		WillReturnRows(rows)

	repository := CreateRepository(db)
	list, err := repository.ListAll()

	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestRepositoryListAll__should_return_empty(t *testing.T) {
	db, mock, teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players_player`)).
		WillReturnRows(sqlmock.NewRows(nil))

	repository := CreateRepository(db)
	list, err := repository.ListAll()

	assert.NoError(t, err)
	assert.Len(t, list, 0)
}
