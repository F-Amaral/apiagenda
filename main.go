package main

import (
	"github.com/F-Amaral/apiagenda/agenda/agendahttp"
	"github.com/F-Amaral/apiagenda/agenda/agendaservices"
	"github.com/F-Amaral/apiagenda/api/apiencodes"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	agendaService := agendaservices.New()

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(apiencodes.EncodeError),
	}

	getContactsEndpoint := httptransport.NewServer(agendahttp.MakeGetContacts(agendaService), agendahttp.DecodeSearchContacts, apiencodes.EncodeResponse, serverOptions...)
	getContactByIdEndpoint := httptransport.NewServer(agendahttp.MakeGetContactById(agendaService), agendahttp.DecodeContactById, apiencodes.EncodeResponse, serverOptions...)
	getContactByNameEndpoint := httptransport.NewServer(agendahttp.MakeGetContactByName(agendaService), agendahttp.DecodeGetContactByName, apiencodes.EncodeResponse, serverOptions...)
	addContactEndpoint := httptransport.NewServer(agendahttp.MakeAddContact(agendaService), agendahttp.DecodeContact, apiencodes.EncodeResponse, serverOptions...)
	updateContactEndpoint := httptransport.NewServer(agendahttp.MakeUpdateContact(agendaService), agendahttp.DecodeContact, apiencodes.EncodeResponse, serverOptions...)
	deleteContactEndpoint := httptransport.NewServer(agendahttp.MakeDeleteContact(agendaService), agendahttp.DecodeContactById, apiencodes.EncodeResponse, serverOptions...)

	router := mux.NewRouter()
	router.Handle("/contacts/search", getContactsEndpoint)
	router.Handle("/contacts/{id}", getContactByIdEndpoint).Methods(http.MethodGet)
	router.Handle("/contacts", getContactByNameEndpoint).Methods(http.MethodGet)
	router.Handle("/contacts", addContactEndpoint).Methods(http.MethodPost)
	router.Handle("/contacts", updateContactEndpoint).Methods(http.MethodPut)
	router.Handle("/contacts/{id}", deleteContactEndpoint).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		print(err.Error())
		return
	}
}
