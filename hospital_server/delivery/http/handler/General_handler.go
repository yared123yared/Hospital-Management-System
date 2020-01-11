package handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"

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
func (dah *GeneralHandler) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" I am at the put method")
	//id, err := strconv.Atoi(ps.ByName("id"))
	//fmt.Println(id)
	//if err != nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	//	return
	//}

	//appointment, errs := dah.appointmentService.Appointment(uint(id))
	users := &entity.Profile{}
	//if len(errs) > 0 {
	//	w.Header().Set("Content-Type", "application/json")
	//	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	//	return
	//}
	fmt.Println(" i have get single value")

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &users)

	id := users.ID
	password := users.Password
	users, errs := dah.generalService.Users(int(id), password)
	fmt.Println("aaaaaaaaaaaaa")
	fmt.Println(users)
	if len(errs) > 0 {
		fmt.Println("thise is status not found")

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if users == nil {
		fmt.Println("thise is status not found")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		fmt.Println("thise is status not found")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
