package handler

import (
	"encoding/json"
	"fmt"
	appointment "github.com/getach1/web1/hospital/Appointment"
	"github.com/getach1/web1/hospital/entity"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// AdminAppointmentHandler handles appointment related http requests
type AdminAppointmentHandler struct {
	appointmentService appointment.AppointmentService
}

// NewAdminAppointmentHandler returns new AdminAppointmentHandler object
func NewAdminAppointmentHandler(cmntService appointment.AppointmentService) *AdminAppointmentHandler {
	return &AdminAppointmentHandler{appointmentService: cmntService}
}

// GetAppointments handles GET /v1/admin/appointments request
func (aph *AdminAppointmentHandler) GetAppointments(w http.ResponseWriter,
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
func (aph *AdminAppointmentHandler) GetSingleAppointment(w http.ResponseWriter,
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

// PostAppointment handles POST /v1/admin/appointments request
func (aph *AdminAppointmentHandler) PostAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	appointment := &entity.Appointment{}

	err := json.Unmarshal(body, appointment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	appointment, errs := aph.appointmentService.StoreAppointment(appointment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/appointments/%d", appointment.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutAppointment handles PUT /v1/admin/appointments/:id request
func (aph *AdminAppointmentHandler) PutAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &appointment)
	appointment, errs = aph.appointmentService.UpdateAppointment(appointment)
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

// DeleteAppointment handles DELETE /v1/admin/appointments/:id request
func (aph *AdminAppointmentHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := aph.appointmentService.DeleteAppointment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
