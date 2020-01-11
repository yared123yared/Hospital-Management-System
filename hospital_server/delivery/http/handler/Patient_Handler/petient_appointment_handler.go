package Patient_Handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/web1_group_project/hospital_server/petient"
	"net/http"
	"strconv"
)

// PetientAppointmentHandler handles appointment related http requests
type PetientAppointmentHandler struct {
	appointmentService petient.AppointmentService
}

// NewPetientAppointmentHandler returns new PetientAppointmentHandler object
func NewPetientAppointmentHandler(cmntService petient.AppointmentService) *PetientAppointmentHandler {
	return &PetientAppointmentHandler{appointmentService: cmntService}
}

// GetAppointments handles GET /v1/admin/appointments request
func (aph *PetientAppointmentHandler) GetAppointments(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	appointments, errs := aph.appointmentService.Appointments()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(appointments, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetSingleAppointment handles GET /v1/admin/appointments/:id request
func (aph *PetientAppointmentHandler) GetSingleAppointment(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	appointment, errs := aph.appointmentService.Appointment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(appointment, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
