package agendahttp

import (
	"context"
	"github.com/F-Amaral/apiagenda/entities"
	"github.com/F-Amaral/apiagenda/services"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetContacts(agendaService services.AgendaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return agendaService.GetAll(ctx)
	}
}

func MakeGetContactById(agendaService services.AgendaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(string)
		return agendaService.GetById(ctx, id)
	}
}

func MakeGetContactByName(agendaService services.AgendaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		name := request.(string)
		return agendaService.GetByName(ctx, name)
	}
}

func MakeAddContact(agendaService services.AgendaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		contact := request.(*entities.Contact)
		return agendaService.Add(ctx, contact)
	}
}

func MakeUpdateContact(agendaService services.AgendaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		contact := request.(*entities.Contact)
		return agendaService.Update(ctx, contact)
	}
}

func MakeDeleteContact(agendaService services.AgendaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(string)
		return agendaService.Remove(ctx, id)
	}
}
