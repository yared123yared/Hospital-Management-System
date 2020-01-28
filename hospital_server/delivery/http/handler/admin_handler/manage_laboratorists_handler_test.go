package admin_handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	adminRepo "github.com/web1_group_project/hospital_server/Admin/repository"
	adminSrv "github.com/web1_group_project/hospital_server/Admin/service"
	adminHandler "github.com/web1_group_project/hospital_server/delivery/http/handler/admin_handler"
	"github.com/web1_group_project/hospital_server/entity"
)

func TestGetLaboratorists(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/laboratorists", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageLaboratoristsRepository(nil)
	profSrv := adminSrv.NewManageLaboratoristsService(ProfReq)
	profHandler := adminHandler.NewManageLaboratoristHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/laboratorists", profHandler.GetLaboratorists)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.LaboratoristMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSingleLaboratorist(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/laboratorists/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageLaboratoristsRepository(nil)
	profSrv := adminSrv.NewManageLaboratoristsService(ProfReq)
	profHandler := adminHandler.NewManageLaboratoristHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/laboratorists/:id", profHandler.GetSingleLaboratorist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.LaboratoristMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
func TestAddLaboratorist(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
		"Department": "Mock Department 01"
}`)
	req, err := http.NewRequest("POST", "/admin/laboratorsts/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/laboratorsts/", profHandler.AddDoctor)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestUpdateLaboratorist(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
	  }`)
	req, err := http.NewRequest("PUT", "/admin/laboratorists/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageLaboratoristsRepository(nil)
	profSrv := adminSrv.NewManageLaboratoristsService(ProfReq)
	profHandler := adminHandler.NewManageLaboratoristHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/admin/laboratorists/:id", profHandler.UpdateLaboratorist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeleteLaboratorist(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/admin/laboratorists/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageLaboratoristsRepository(nil)
	profSrv := adminSrv.NewManageLaboratoristsService(ProfReq)
	profHandler := adminHandler.NewManageLaboratoristHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/admin/laboratorists/:id", profHandler.DeleteLaboratorist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
