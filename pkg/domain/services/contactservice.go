package services

import (
	"context"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
)

type ContactService interface {
	Add(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError)
	Remove(ctx context.Context, id string) (*entities.Contact, apierror.ApiError)
	GetByName(ctx context.Context, name string) (*entities.Contact, apierror.ApiError)
	GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError)
	GetAll(ctx context.Context) ([]*entities.Contact, apierror.ApiError)
	Update(ctx context.Context, contact *entities.Contact) (*entities.Contact, apierror.ApiError)
}
