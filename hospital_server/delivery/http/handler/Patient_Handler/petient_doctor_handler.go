package Patient_Handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/web1_group_project/hospital_server/petient"
	"net/http"
	"strconv"
)

// PetientDoctorHandler handles doctor related http doctors
type PetientDoctorHandler struct {
	doctorService petient.DoctorService
}

// NewPetientDoctorHandler returns new PetientDoctorHandler object
func NewPetientDoctorHandler(cmntService petient.DoctorService) *PetientDoctorHandler {
	return &PetientDoctorHandler{doctorService: cmntService}
}

// GetDoctors handles GET /v1/admin/doctors doctor
func (aph *PetientDoctorHandler) GetDoctors(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {
	fmt.Println("Getinng doctotrs ......... in handler")

	doctors, errs := aph.doctorService.Doctors()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(doctors, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetSingleDoctor handles GET /v1/admin/doctors/:id doctor
func (aph *PetientDoctorHandler) GetSingleDoctor(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	doctor, errs := aph.doctorService.Doctor(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(doctor, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
