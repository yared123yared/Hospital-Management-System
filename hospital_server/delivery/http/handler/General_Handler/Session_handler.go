package General_Handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//_ "github.com/yaredsolomon/webProgram1/hospital/entity"

	"github.com/web1_group_project/hospital_server/General"
	"github.com/web1_group_project/hospital_server/entity"
)

type SessionHandler struct {
	sessionService General.SessionService
}

// NewGeneralHandler returns new GeneralHandler object
func NewSessionHandler(gnService General.SessionService) *SessionHandler {
	return &SessionHandler{sessionService: gnService}
}

// GetAppointments handles GET /v1/doctor/appointments request

func (sh *SessionHandler) Session(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println("thise is the session gate method")
	fmt.Println(" i am about to get single value")

	id := (ps.ByName("id"))
	fmt.Println(id)

	session, errs := sh.sessionService.Session((id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(session, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeletePatient handles DELETE /v1/admin/comments/:id request
func (sh *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am bout to delete ")
	id := (ps.ByName("id"))
	fmt.Println(id)

	_, errs := sh.sessionService.DeleteSession((id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (sh *SessionHandler) StoreSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am at the post method")

	l := r.ContentLength
	body := make([]byte, l)
	fmt.Println(" ia have changed the data to byte")
	fmt.Println(string(body))
	r.Body.Read(body)
	session := &entity.Session{}

	err := json.Unmarshal(body, session)
	fmt.Println("thise is the unmarchal jeson")
	fmt.Println(session)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	session, errs := sh.sessionService.StoreSession(session)
	fmt.Println("thise is the return value from the session")
	fmt.Println(session)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(session, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
