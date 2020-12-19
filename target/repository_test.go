package target

import (
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

func TestRepositoryExists__should_return_true(t *testing.T) {
	db, mock, teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cardID := "12345678"
	mock.
		ExpectQuery(`SELECT count\(1\) FROM "clientes_identidadcliente"`).WithArgs(cardID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	repository := CreateRepository(db)
	exists, err := repository.Exists(cardID)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)
}

func TestRepositoryExists__should_return_false(t *testing.T) {
	db, mock, teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	cardID := "12345678"
	mock.
		ExpectQuery(`SELECT count\(1\) FROM "clientes_identidadcliente"`).
		WithArgs(cardID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	repository := CreateRepository(db)
	exists, err := repository.Exists(cardID)

	assert.NoError(t, err)
	assert.Equal(t, exists, false)
}

func TestRepositorySave__should_register_client(t *testing.T) {
	db, mock, teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	newID := 1
	firstName := "jorge"
	lastName := "tolentino"
	cardType := 1
	cardID := "12345678"

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "clientes_cliente"`).
		WithArgs(firstName, lastName, true).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newID))

	mock.ExpectExec(`INSERT INTO "clientes_identidadcliente"`).
		WithArgs(newID, cardType, cardID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`INSERT INTO "clientes_restriccioncliente"`).
		WithArgs(newID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repository := CreateRepository(db)
	client, err := repository.Save(firstName, lastName, cardType, cardID)

	assert.NoError(t, err)
	assert.Equal(t, client.ID, newID)
}
