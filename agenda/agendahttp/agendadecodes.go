package agendahttp

import (
	"context"
	"encoding/json"
	"github.com/F-Amaral/apiagenda/api/apierror"
	"github.com/F-Amaral/apiagenda/entities"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
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

func DecodeGetContactByName(ctx context.Context, r *http.Request) (request interface{}, err error) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		return nil, apierror.New(http.StatusBadRequest, "query parameter name is required")
	}
	return name, nil
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
