package Doctor_Handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
	"net/http"
	"strconv"
)

type DoctorPatientHandler struct {
	patientService Doctor.PatientService
}

// NewAdminCommentHandler returns new AdminCommentHandler object
func NewDoctorPatientHandler(ptService Doctor.PatientService) *DoctorPatientHandler {
	return &DoctorPatientHandler{patientService: ptService}
}

// GetUsers handles GET /v1/admin/users request
func (aph *DoctorPatientHandler) GetPatients(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	patientes, errs := aph.patientService.Patientes()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(patientes, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleUsers handles GET /v1/admin/users/:id request
func (aph *DoctorPatientHandler) GetSinglePatient(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am about to get single value")

	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	patient, errs := aph.patientService.Patient(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(patient, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostUser handles POST /v1/admin/users request
func (aph *DoctorPatientHandler) PostPatient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am at the post method")

	l := r.ContentLength
	body := make([]byte, l)
	fmt.Println(" ia have changed the data to byte")
	fmt.Println(string(body))
	r.Body.Read(body)
	patient := &entity.Petient{}

	err := json.Unmarshal(body, patient)
	fmt.Println("thise is the unmarchal jeson")
	fmt.Println(patient)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	patient, errs := aph.patientService.StorePatient(patient)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/users/%d")
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutComment handles PUT /v1/admin/comments/:id request
func (aph *DoctorPatientHandler) PutPatient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	patient, errs := aph.patientService.Patient(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &patient)

	patient, errs = aph.patientService.UpdatePatient(patient)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(patient, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteComment handles DELETE /v1/admin/comments/:id request
func (aph *DoctorPatientHandler) DeletePatient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i amabout to delete ")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := aph.patientService.DeletePatient(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
