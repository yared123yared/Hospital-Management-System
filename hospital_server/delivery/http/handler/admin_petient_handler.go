package handler

import (
	"encoding/json"
	"fmt"
	"github.com/getach1/web1/web1_group_project/hospital_server/entity"
	"github.com/getach1/web1/web1_group_project/hospital_server/petient"

	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// AdminPetientHandler handles petient related http requests
type AdminPetientHandler struct {
	petientService petient.PetientService
}

// NewAdminPetientHandler returns new AdminPetientHandler object
func NewAdminPetientHandler(cmntService petient.PetientService) *AdminPetientHandler {
	return &AdminPetientHandler{petientService: cmntService}
}

// GetPetients handles GET /v1/admin/petients request
func (aph *AdminPetientHandler) GetPetients(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	petients, errs := aph.petientService.Petients()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(petients, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetSinglePetient handles GET /v1/admin/petients/:id request
func (aph *AdminPetientHandler) GetSinglePetient(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	petient, errs := aph.petientService.Petient(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(petient, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostPetient handles POST /v1/admin/petients request
func (aph *AdminPetientHandler) PostPetient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	petient := &entity.Petient{}

	err := json.Unmarshal(body, petient)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	petient, errs := aph.petientService.StorePetient(petient)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/petients/%d", petient.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutPetient handles PUT /v1/admin/petients/:id request
func (aph *AdminPetientHandler) PutPetient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	petient, errs := aph.petientService.Petient(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &petient)
	petient, errs = aph.petientService.UpdatePetient(petient)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(petient, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeletePetient handles DELETE /v1/admin/petients/:id request
func (aph *AdminPetientHandler) DeletePetient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := aph.petientService.DeletePetient(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
