package agendaservices

import (
	"context"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/contracts"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"github.com/F-Amaral/apiagenda/pkg/domain/repositories"
	"net/http"
)

type contactService struct {
	contactRepository repositories.ContactRepository
}

func NewContactService(repository repositories.ContactRepository) *contactService {
	return &contactService{
		contactRepository: repository,
	}
}

func (self *contactService) Add(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError) {
	err := self.contactRepository.Add(ctx, contact)
	if err != nil {
		return nil, err
	}
	addedContact, err := self.GetById(ctx, contact.Id)
	if err != nil {
		return nil, err
	}
	return addedContact, nil
}

func (self *contactService) Remove(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {

	err := self.contactRepository.Remove(ctx, id)
	if err != nil {
		return nil, err
	}

	contact, err := self.contactRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (self *contactService) GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {
	return self.contactRepository.GetById(ctx, id)
}

func (self *contactService) Search(ctx context.Context, searchRequest *contracts.SearchRequest) ([]*entities.Contact, apierror.ApiError) {
	return self.contactRepository.Search(ctx, searchRequest)
}

func (self *contactService) Update(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError) {

	currentContact, err := self.contactRepository.GetById(ctx, contact.Id)
	if err != nil {
		return nil, err
	}

	if currentContact.Deleted {
		return nil, apierror.New(http.StatusNotFound, "contact not found")
	}

	err = self.contactRepository.Update(ctx, contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
