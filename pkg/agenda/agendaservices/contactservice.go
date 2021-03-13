package agendaservices

import (
	"context"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"github.com/F-Amaral/apiagenda/pkg/domain/repositories"
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
	contact, err := self.contactRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = self.contactRepository.Remove(ctx, id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (self *contactService) GetByName(ctx context.Context, name string) (*entities.Contact, apierror.ApiError) {
	return self.contactRepository.GetByName(ctx, name)
}

func (self *contactService) GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {
	return self.contactRepository.GetById(ctx, id)
}

func (self *contactService) GetAll(ctx context.Context) ([]*entities.Contact, apierror.ApiError) {
	return self.contactRepository.GetAll(ctx)

}

func (self *contactService) Update(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError) {
	err := self.contactRepository.Update(ctx, contact)
	if err != nil {
		return nil, err
	}
	updatedContact, err := self.contactRepository.GetById(ctx, contact.Id)
	if err != nil {
		return nil, err
	}

	return updatedContact, nil
}
