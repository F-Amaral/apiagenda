package agendahttp

import (
	"context"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"github.com/F-Amaral/apiagenda/pkg/domain/services"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetContacts(service services.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.GetAll(ctx)
	}
}

func MakeGetContactById(service services.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(string)
		return service.GetById(ctx, id)
	}
}

func MakeGetContactByName(service services.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		name := request.(string)
		return service.GetByName(ctx, name)
	}
}

func MakeAddContact(service services.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		contact := request.(*entities.Contact)
		return service.Add(ctx, contact)
	}
}

func MakeUpdateContact(service services.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		contact := request.(*entities.Contact)
		return service.Update(ctx, contact)
	}
}

func MakeDeleteContact(service services.ContactService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(string)
		return service.Remove(ctx, id)
	}
}
