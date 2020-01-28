package Patient_Handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	patientHadler "github.com/web1_group_project/hospital_server/delivery/http/handler/Patient_Handler"
	"github.com/web1_group_project/hospital_server/entity"
	patientRepo "github.com/web1_group_project/hospital_server/petient/repository"
	patientSrv "github.com/web1_group_project/hospital_server/petient/service"
)

func TestGetPetients(t *testing.T) {
	req, err := http.NewRequest("GET", "/patient/patients", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPetientGormRepo(nil)
	profSrv := patientSrv.NewPetientService(ProfReq)
	profHandler := patientHadler.NewAdminPetientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/patient/patients", profHandler.GetPetients)
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
func TestGetSinglePetient(t *testing.T) {
	req, err := http.NewRequest("GET", "/patient/patients/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPetientGormRepo(nil)
	profSrv := patientSrv.NewPetientService(ProfReq)
	profHandler := patientHadler.NewAdminPetientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/patient/patients/:id", profHandler.GetSinglePetient)
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

func TestPostPetient(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
	  }`)
	req, err := http.NewRequest("POST", "/patient/patients/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPetientGormRepo(nil)
	profSrv := patientSrv.NewPetientService(ProfReq)
	profHandler := patientHadler.NewAdminPetientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/patient/patients/", profHandler.PostPetient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestPutPetient(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
	  }`)
	req, err := http.NewRequest("PUT", "/patient/patients/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPetientGormRepo(nil)
	profSrv := patientSrv.NewPetientService(ProfReq)
	profHandler := patientHadler.NewAdminPetientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/patient/patients/:id", profHandler.PutPetient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}

func TestDeletePetient(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/patient/patients/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPetientGormRepo(nil)
	profSrv := patientSrv.NewPetientService(ProfReq)
	profHandler := patientHadler.NewAdminPetientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/patient/patients/:id", profHandler.DeletePetient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
