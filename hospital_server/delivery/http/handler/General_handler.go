package handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/monday271/hospital_server/Doctor"
	//"github.com/yaredsolomon/webProgram1/hospital/request"
	"net/http"
	//"github.com/betsegawlemma/restaurant-rest/comment"
	"github.com/julienschmidt/httprouter"
	_ "github.com/yaredsolomon/webProgram1/hospital/entity"
	//"github.com/yaredsolomon/webProgram1/hospital/request"
)

//"github.com/yaredsolomon/webProgram1/sathurday18/entity"

// DoctorAppointmentHandler handles appointment related http requests
type GeneralHandler struct {
	generalService Doctor.GeneralService
}

// NewDoctorAppointmentHandler returns new DoctorAppointmentHandler object
func NewGeneralHandler(gnService Doctor.GeneralService) *GeneralHandler {
	return &GeneralHandler{generalService: gnService}
}

// GetAppointments handles GET /v1/doctor/appointments request
func (gh *GeneralHandler) GetPharmacists(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	pharmacists, errs := gh.generalService.Pharmacists()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(pharmacists, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (gh *GeneralHandler) GetLaboratorists(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {
	fmt.Println("i am at the labortorist method")
	laboratorist, errs := gh.generalService.Laboratorists()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(laboratorist, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
