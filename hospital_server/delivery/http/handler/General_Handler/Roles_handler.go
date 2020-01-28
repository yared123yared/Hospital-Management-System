package General_Handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	//_ "github.com/yaredsolomon/webProgram1/hospital/entity"

	"github.com/web1_group_project/hospital_server/General"
	"github.com/web1_group_project/hospital_server/entity"
)

type RolesHandler struct {
	roleService General.RoleService
}

// NewGeneralHandler returns new GeneralHandler object
func NewRolesHandler(gnService General.RoleService) *RolesHandler {
	return &RolesHandler{roleService: gnService}
}

// GetAppointments handles GET /v1/doctor/appointments request
func (rh *RolesHandler) Roles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {
	fmt.Println("thise is the role method")

	roles, errs := rh.roleService.Roles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rh *RolesHandler) Role(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am about to get single value")

	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	role, errs := rh.roleService.Role(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(role, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (rh *RolesHandler) RoleByName(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am about to get single value")

	name := ps.ByName("name")

	fmt.Println(name)

	role, errs := rh.roleService.RoleByName(name)
	//role, errs := aph.roleService.Role(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(role, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (rh *RolesHandler) UpdateRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	role, errs := rh.roleService.Role(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &role)

	roles, errs := rh.roleService.UpdateRole(role)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

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
func (rh *RolesHandler) DeleteRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am bout to delete ")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rh.roleService.DeleteRole(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (rh *RolesHandler) StoreRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am at the post method")

	l := r.ContentLength
	body := make([]byte, l)
	fmt.Println(" ia have changed the data to byte")
	fmt.Println(string(body))
	r.Body.Read(body)
	role := &entity.Role{}

	err := json.Unmarshal(body, role)
	fmt.Println("thise is the unmarchal jeson")
	fmt.Println(role)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	role, errs := rh.roleService.StoreRole(role)

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
