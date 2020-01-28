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

func TestGetPharmasists(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/pharmasists", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePharmasistsRepository(nil)
	profSrv := adminSrv.NewManagePharmasistsService(ProfReq)
	profHandler := adminHandler.NewManagePharmasistHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/pharmasists", profHandler.GetPharmasists)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PharmacistMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSinglePharmasist(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/pharmasists/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePharmasistsRepository(nil)
	profSrv := adminSrv.NewManagePharmasistsService(ProfReq)
	profHandler := adminHandler.NewManagePharmasistHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/pharmasists/:id", profHandler.GetSinglePharmasist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PharmacistMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
func TestAddPharmasist(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
		"Department": "Mock Department 01"
}`)

	req, err := http.NewRequest("POST", "/admin/pharmasists/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePharmasistsRepository(nil)
	profSrv := adminSrv.NewManagePharmasistsService(ProfReq)
	profHandler := adminHandler.NewManagePharmasistHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/pharmasists/", profHandler.AddPharmasist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestUpdatePharmasist(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
		"Department": "Mock Department 01"
	  }`)
	req, err := http.NewRequest("PUT", "/admin/pharmasists/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePharmasistsRepository(nil)
	profSrv := adminSrv.NewManagePharmasistsService(ProfReq)
	profHandler := adminHandler.NewManagePharmasistHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/admin/pharmasists/:id", profHandler.UpdatePharmasist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeletePharmasist(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/admin/pharmasists/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePharmasistsRepository(nil)
	profSrv := adminSrv.NewManagePharmasistsService(ProfReq)
	profHandler := adminHandler.NewManagePharmasistHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/admin/pharmasists/:id", profHandler.DeletePharmasist)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
