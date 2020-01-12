package Patient_Handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/web1_group_project/hospital_server/petient"
	"net/http"
)
// AdminAdminHandler handles admin related http requests
type AdminAdminHandler struct {
	adminService petient.AdminService
}

// NewAdminAdminHandler returns new AdminAdminHandler object
func NewPatientAdminHandler(cmntService petient.AdminService) *AdminAdminHandler {
	return &AdminAdminHandler{adminService: cmntService}
}

// GetSingleAdmin handles GET /v1/admin/admins/:id request
func (aph *AdminAdminHandler) GetAdmins(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	admin, errs := aph.adminService.Admins()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(admin, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

