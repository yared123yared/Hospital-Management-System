package handler

import (
	"encoding/json"
	"fmt"
	"github.com/getach1/web1/hospital/entity"
	presc "github.com/getach1/web1/hospital/prescribtion"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// AdminPrescriptionHandler handles prescription related http requests
type AdminPrescriptionHandler struct {
	prescriptionService presc.PrescriptionService
}

// NewAdminPrescriptionHandler returns new AdminPrescriptionHandler object
func NewAdminPrescriptionHandler(cmntService presc.PrescriptionService) *AdminPrescriptionHandler {
	return &AdminPrescriptionHandler{prescriptionService: cmntService}
}

// GetPrescriptions handles GET /v1/admin/prescriptions request
func (aph *AdminPrescriptionHandler) GetPrescriptions(w http.ResponseWriter,
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
func (aph *AdminPrescriptionHandler) GetSinglePrescription(w http.ResponseWriter,
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

// PostPrescription handles POST /v1/admin/prescriptions request
func (aph *AdminPrescriptionHandler) PostPrescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	prescription := &entity.Prescription{}

	err := json.Unmarshal(body, prescription)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	prescription, errs := aph.prescriptionService.StorePrescription(prescription)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/prescriptions/%d", prescription.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutPrescription handles PUT /v1/admin/prescriptions/:id request
func (aph *AdminPrescriptionHandler) PutPrescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &prescription)
	prescription, errs = aph.prescriptionService.UpdatePrescription(prescription)
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

// DeletePrescription handles DELETE /v1/admin/prescriptions/:id request
func (aph *AdminPrescriptionHandler) DeletePrescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := aph.prescriptionService.DeletePrescription(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
