package labor_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fasikawkn/web1_group_project/hospital_server/Laboratorist"
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
	"github.com/julienschmidt/httprouter"
)

// PharmMedicineHandler handles comment related http requests
type LaborDiagnosisHandler struct {
	diagnosisService Laboratorist.DiagnosisService
}

// NewPharmMedicineHandler returns new AdminCommentHandler object
func NewLaborDiagnosisHandler(diagService Laboratorist.DiagnosisService) *LaborDiagnosisHandler {
	return &LaborDiagnosisHandler{diagnosisService: diagService}
}

// GetMedicines handles GET /v1/admin/comments request
func (ach *LaborDiagnosisHandler) GetDiagnosiss(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	medicines, errs := ach.diagnosisService.Diagnosiss()

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

// GetMedicines handles GET /v1/admin/comments request
func (ach *LaborDiagnosisHandler) GetMultiDiagnosis(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	medicines, errs := ach.diagnosisService.GetDiagnosiss()

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
func (ach *LaborDiagnosisHandler) GetSingleDiagnosis(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.diagnosisService.Diagnosis(uint(id))

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
func (ach *LaborDiagnosisHandler) PostDiagnosis(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	medicine := &entity.Diagnosis{}

	err := json.Unmarshal(body, medicine)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.diagnosisService.AddDiagnosis(medicine)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/labor/diag/%d", medicine.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutMedicine handles PUT /v1/admin/comments/:id request
func (ach *LaborDiagnosisHandler) PutDiagnosis(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := ach.diagnosisService.Diagnosis(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &medicine)

	medicine, errs = ach.diagnosisService.UpdateDiagnosis(medicine)

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
func (ach *LaborDiagnosisHandler) DeleteDiagnosis(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.diagnosisService.DeleteDiagnosis(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
