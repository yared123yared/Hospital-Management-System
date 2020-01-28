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

func TestGetDoctors(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/doctors", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/doctors", profHandler.GetDoctors)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.DoctorMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSingleDoctor(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/doctors/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/doctors/:id", profHandler.GetSingleDoctor)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.DoctorMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
func TestAddDoctor(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
		"Department": "Mock Department 01"
}`)
	req, err := http.NewRequest("POST", "/admin/doctors/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/doctors/", profHandler.AddDoctor)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestUpdateDoctor(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
		"Department": "Mock Department 01"
	  }`)
	req, err := http.NewRequest("PUT", "/admin/doctors/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/admin/doctors/:id", profHandler.UpdateDoctor)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeleteDoctor(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/admin/doctors/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/admin/doctors/:id", profHandler.DeleteDoctor)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
