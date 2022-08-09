package impl

import (
	"HelloWorld/types"
	"net/http"
)

type ApiImpl struct{}

// Get /api/volunteers
func (api *ApiImpl) GetVolunteers(w http.ResponseWriter, r *http.Request, params types.GetVolunteersParams) {
	// Implement your logic here

	w.WriteHeader(200)
}
