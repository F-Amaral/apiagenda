package agendaservices

import (
	"context"
	"github.com/F-Amaral/apiagenda/api/apierror"
	"github.com/F-Amaral/apiagenda/entities"
	"net/http"
)

type agendaService struct {
	contacts map[string]*entities.Contact
}

func New() *agendaService {
	return &agendaService{
		contacts: make(map[string]*entities.Contact),
	}
}

func (self *agendaService) Add(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError) {
	self.contacts[contact.Id] = contact
	return contact, nil
}

func (self *agendaService) Remove(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {
	contact := self.contacts[id]
	delete(self.contacts, id)
	return contact, nil
}

func (self *agendaService) GetByName(ctx context.Context, name string) (*entities.Contact, apierror.ApiError) {
	for _, contact := range self.contacts {
		if contact.Name == name {
			return contact, nil
		}
	}
	return nil, apierror.New(http.StatusNotFound, "contact not found")
}

func (self *agendaService) GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {
	contact := self.contacts[id]
	if contact == nil {
		return nil, apierror.New(http.StatusNotFound, "contact not found")
	}
	return contact, nil
}

func (self *agendaService) GetAll(ctx context.Context) ([]*entities.Contact, apierror.ApiError) {
	var contactArray []*entities.Contact
	for _, contact := range self.contacts {
		contactArray = append(contactArray, contact)
	}
	return contactArray, nil
}

func (self *agendaService) Update(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError) {
	self.contacts[contact.Id] = contact
	return contact, nil
}
