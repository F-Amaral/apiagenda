package repositories

import (
	"context"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
)

type ContactRepository interface {
	Add(ctx context.Context, contact *entities.Contact) apierror.ApiError
	Remove(ctx context.Context, id string) apierror.ApiError
	GetByName(ctx context.Context, name string) (*entities.Contact, apierror.ApiError)
	GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError)
	GetAll(ctx context.Context) ([]*entities.Contact, apierror.ApiError)
	Update(ctx context.Context, contact *entities.Contact) apierror.ApiError
}
