package agendahttp

import (
	"context"
	"encoding/json"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/contracts"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func DecodeSearchContacts(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func DecodeContactById(ctx context.Context, r *http.Request) (request interface{}, err error) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	if id == "" {
		return nil, apierror.New(http.StatusBadRequest, "route parameter id is required")
	}
	return id, nil

}

func DecodeContact(ctx context.Context, r *http.Request) (request interface{}, err error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, err.Error())
	}
	contact := &entities.Contact{}
	err = json.Unmarshal(bytes, contact)
	if err != nil {
		return nil, apierror.New(http.StatusBadRequest, err.Error())
	}

	return contact, nil
}

func DecodeSearchRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	query := r.URL.Query()
	searchRequest := &contracts.SearchRequest{}

	id := query.Get("id")
	if id != "" {
		searchRequest.Id = &id
	}

	name := query.Get("name")
	if name != "" {
		searchRequest.Name = &name
	}

	email := query.Get("email")
	if email != "" {
		searchRequest.Email = &email
	}

	deletedString := query.Get("deleted")
	if deletedString != "" {
		deleted, err := strconv.ParseBool(deletedString)
		if err == nil {
			searchRequest.Deleted = &deleted
		}
	}

	return searchRequest, nil
}
