package api

import (
	"HelloWorld/types"
	"net/http"
)

// ServerInterface represents all server handlers.
type ApiStub interface {
	GetVolunteers(w http.ResponseWriter, r *http.Request, params types.GetVolunteersParams)
}

type ApiWrapper struct {
	ApiDelegate ApiStub
}

func (stub *ApiWrapper) GetVolunteers(w http.ResponseWriter, r *http.Request) {

	getvolunteersParams := types.GetVolunteersParams{
		Email: r.URL.Query().Get("Email"),
	}

	stub.ApiDelegate.GetVolunteers(w, r, getvolunteersParams)
}
