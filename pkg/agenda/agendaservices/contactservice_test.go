package agendaservices

import (
	"context"
	"fmt"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/mocks"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestNewContactService(t *testing.T) {
	// Arrange

	// Act
	sut := NewContactService(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*agendaservices.contactService", fmt.Sprintf("%T", sut))
}

func TestContactService_Add_WhenAddFails(t *testing.T) {
	// Arrange
	contactRepository := &mocks.ContactRepository{}
	errToReturn := apierror.New(http.StatusInternalServerError, "")
	contactRepository.On("Add", mock.Anything, mock.Anything).Return(errToReturn)
	sut := NewContactService(contactRepository)
	ctx := context.TODO()
	contact := &entities.Contact{}

	// Act
	response, err := sut.Add(ctx, contact)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, errToReturn, err)
	contactRepository.AssertNumberOfCalls(t, "Add", 1)
	contactRepository.AssertCalled(t, "Add", ctx, contact)
}

func TestContactService_Add_WhenGetFails(t *testing.T) {
	// Arrange
	contactRepository := &mocks.ContactRepository{}
	errToReturn := apierror.New(http.StatusInternalServerError, "")
	contactRepository.On("Add", mock.Anything, mock.Anything).Return(nil)
	contactRepository.On("GetById", mock.Anything, mock.Anything).Return(nil, errToReturn)
	sut := NewContactService(contactRepository)

	// Act
	response, err := sut.Add(context.TODO(), &entities.Contact{})

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, errToReturn, err)
}

func TestContactService_Add_Happy(t *testing.T) {
	// Arrange
	contactRepository := &mocks.ContactRepository{}
	contact := &entities.Contact{}
	contactRepository.On("Add", mock.Anything, mock.Anything).Return(nil)
	contactRepository.On("GetById", mock.Anything, mock.Anything).Return(contact, nil)
	sut := NewContactService(contactRepository)

	// Act
	response, err := sut.Add(context.TODO(), &entities.Contact{})

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, contact, response)
}
