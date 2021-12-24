package httptransport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/truescotian/golang-rest-projstructure/internal"
)

type handler struct {
	internal.Service
}

var r = mux.NewRouter()

func Routes() http.Handler {
	return r
}

func UserHandler(s internal.Service) {
	h := &handler{s}
	r.HandleFunc("/api/users", h.createUser).Methods("POST")
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var in internal.CreateUserIn
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	out, err := h.CreateUser(ctx, in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
