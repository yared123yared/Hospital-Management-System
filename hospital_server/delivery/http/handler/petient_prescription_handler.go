package handler

import (
	"encoding/json"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/petient"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// PetientPrescriptionHandler handles prescription related http requests
type PetientPrescriptionHandler struct {
	prescriptionService petient.PrescriptionService
}

// NewPetientPrescriptionHandler returns new PetientPrescriptionHandler object
func NewPetientPrescriptionHandler(cmntService petient.PrescriptionService) *PetientPrescriptionHandler {
	return &PetientPrescriptionHandler{prescriptionService: cmntService}
}

// GetPrescriptions handles GET /v1/admin/prescriptions request
func (aph *PetientPrescriptionHandler) GetPrescriptions(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	prescriptions, errs := aph.prescriptionService.Prescriptions()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(prescriptions, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetSinglePrescription handles GET /v1/admin/prescriptions/:id request
func (aph *PetientPrescriptionHandler) GetSinglePrescription(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	prescription, errs := aph.prescriptionService.Prescription(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(prescription, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
