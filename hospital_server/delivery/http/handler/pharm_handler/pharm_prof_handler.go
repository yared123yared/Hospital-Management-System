package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fasikawkn/web1_group_project-1/hospital_server/Pharmacist"

	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
	"github.com/julienschmidt/httprouter"
)

// PharmProfileHandler handles comment related http requests
type PharmProfileHandler struct {
	profileService Pharmacist.PharmacistProfileService
}

// NewPharmMedicineHandler returns new AdminCommentHandler object
func NewPharmProfileHandler(profcService Pharmacist.PharmacistProfileService) *PharmProfileHandler {
	return &PharmProfileHandler{profileService: profcService}
}

// GetMedicines handles GET /v1/admin/comments request
func (ach *PharmProfileHandler) GetProfiles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	medicines, errs := ach.profileService.Profiles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(medicines, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleMedicine handles GET /v1/admin/comments/:id request
func (ach *PharmProfileHandler) GetSingleProfile(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.profileService.Profile(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(medicine, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostMedicine handles POST /v1/admin/comments request
func (ach *PharmProfileHandler) PostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("post  Pharmacist")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	medicine := &entity.Pharmacist{}

	err := json.Unmarshal(body, medicine)
	log.Println("Unmarshal1 done")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	log.Println("Unmarshal2 done")

	medicine, errs := ach.profileService.AddProfile(medicine)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/pharm/profile/%d", medicine.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutMedicine handles PUT /v1/admin/comments/:id request
func (ach *PharmProfileHandler) PutProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.profileService.Profile(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &medicine)

	medicine, errs = ach.profileService.UpdateProfile(medicine)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(medicine, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//DeleteMedicine handles DELETE /v1/admin/comments/:id request
func (ach *PharmProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.profileService.DeleteProfile(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
