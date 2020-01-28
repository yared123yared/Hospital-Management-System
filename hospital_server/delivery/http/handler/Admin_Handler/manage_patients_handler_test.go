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

func TestGetPatients(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/patient", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePatientsRepository(nil)
	profSrv := adminSrv.NewManagePatientsService(ProfReq)
	profHandler := adminHandler.NewManagePatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/patient", profHandler.GetPatients)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PetientMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSinglePatient(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/patient/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePatientsRepository(nil)
	profSrv := adminSrv.NewManagePatientsService(ProfReq)
	profHandler := adminHandler.NewManagePatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/patient/:id", profHandler.GetSinglePatient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PetientMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
func TestAddPatient(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
}`)

	req, err := http.NewRequest("POST", "/admin/patient/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePatientsRepository(nil)
	profSrv := adminSrv.NewManagePatientsService(ProfReq)
	profHandler := adminHandler.NewManagePatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/patient/", profHandler.AddPatient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestUpdatePatient(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"Uuid": 1,
		"BloodGroup": "Mock Patient 01"
		"Age": "1"
	  }`)
	req, err := http.NewRequest("PUT", "/admin/patient/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManagePatientsRepository(nil)
	profSrv := adminSrv.NewManagePatientsService(ProfReq)
	profHandler := adminHandler.NewManagePatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/admin/patient/:id", profHandler.UpdatePatient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeletePatient(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/admin/patient/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageDoctorsRepository(nil)
	profSrv := adminSrv.NewManageDoctorsService(ProfReq)
	profHandler := adminHandler.NewManageDoctorsHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/admin/patient/:id", profHandler.DeleteDoctor)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
