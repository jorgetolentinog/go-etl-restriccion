package target

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockServiceRepository struct {
	mock.Mock
}

// Exists is a mock method
func (m *MockServiceRepository) Exists(arg1 string) (bool, error) {
	args := m.Called(arg1)
	return args.Bool(0), args.Error(1)
}

// Exists is a mock method
func (m *MockServiceRepository) Save(arg1 string, arg2 string, arg3 int, arg4 string) (Client, error) {
	args := m.Called(arg1, arg2, arg3, arg4)
	return args.Get(0).(Client), args.Error(1)
}

func TestSaveClientIfNotExists(t *testing.T) {
	repo := new(MockServiceRepository)
	// setup expectations with a placeholder in the argument list
	repo.On("Exists", mock.Anything).Return(true, nil)

	service := CreateService(repo)
	firstName := "jorge"
	lastName := "tolentino"
	cardType := 1
	cardID := "12345678"

	service.SaveClientIfNotExists(firstName, lastName, cardType, cardID)

	// assert that the expectations were met
	repo.AssertExpectations(t)
}
