package handler
import (
	"encoding/json"
	"fmt"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/entity"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/petient"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// PetientRequestHandler handles request related http requests
type PetientRequestHandler struct {
	requestService petient.RequestService
}

// NewPetientRequestHandler returns new PetientRequestHandler object
func NewPetientRequestHandler(cmntService petient.RequestService) *PetientRequestHandler {
	return &PetientRequestHandler{requestService: cmntService}
}

// GetRequests handles GET /v1/admin/requests request
func (aph *PetientRequestHandler) GetRequests(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	requests, errs := aph.requestService.Requests()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(requests, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetSingleRequest handles GET /v1/admin/requests/:id request
func (aph *PetientRequestHandler) GetSingleRequest(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	request, errs := aph.requestService.Request(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(request, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostRequest handles POST /v1/admin/requests request
func (aph *PetientRequestHandler) PostRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	request := &entity.Request{}

	err := json.Unmarshal(body, request)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	request, errs := aph.requestService.StoreRequest(request)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/requests/%d", request.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

