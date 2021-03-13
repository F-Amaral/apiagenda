package apiencodes

import (
	"context"
	"encoding/json"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"net/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if apiError, ok := err.(apierror.ApiError); ok {
		w.WriteHeader(apiError.ErrorStatusCode())
		_ = json.NewEncoder(w).Encode(apiError)
		return
	}

	httpStatus := http.StatusInternalServerError

	unknownError := apierror.New(http.StatusInternalServerError, "unknown error")

	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(unknownError)

	return
}
