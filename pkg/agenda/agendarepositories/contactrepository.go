package agendarepositories

import (
	"context"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"net/http"
)

type contactRepository struct {
	contacts map[string]*entities.Contact
}

func NewContactRepository() *contactRepository {
	return &contactRepository{
		contacts: make(map[string]*entities.Contact),
	}
}

func (self contactRepository) Add(ctx context.Context, contact *entities.Contact) apierror.ApiError {
	self.contacts[contact.Id] = contact
	return nil
}

func (self contactRepository) Remove(ctx context.Context, id string) apierror.ApiError {
	delete(self.contacts, id)
	return nil
}

func (self contactRepository) GetByName(ctx context.Context, name string) (*entities.Contact, apierror.ApiError) {
	for _, contact := range self.contacts {
		if contact.Name == name {
			return contact, nil
		}
	}
	return nil, apierror.New(http.StatusNotFound, "contact not found")
}

func (self contactRepository) GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {
	contact := self.contacts[id]
	if contact == nil {
		return nil, apierror.New(http.StatusNotFound, "contact not found")
	}
	return contact, nil
}

func (self contactRepository) GetAll(ctx context.Context) ([]*entities.Contact, apierror.ApiError) {
	var contactArray []*entities.Contact
	for _, contact := range self.contacts {
		contactArray = append(contactArray, contact)
	}
	return contactArray, nil
}

func (self contactRepository) Update(ctx context.Context, contact *entities.Contact) apierror.ApiError {
	self.contacts[contact.Id] = contact
	return nil
}
