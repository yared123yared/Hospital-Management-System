package Doctor_Handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	doctorRepo "github.com/web1_group_project/hospital_server/Doctor/repository"
	doctorSrv "github.com/web1_group_project/hospital_server/Doctor/service"
	doctorHadler "github.com/web1_group_project/hospital_server/delivery/http/handler/Doctor_Handler"
	"github.com/web1_group_project/hospital_server/entity"
)

func TestGetPatients(t *testing.T) {
	req, err := http.NewRequest("GET", "/doctor/patient", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockPatientGormRepo(nil)
	profSrv := doctorSrv.NewPatientService(ProfReq)
	profHandler := doctorHadler.NewDoctorPatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/doctor/patient", profHandler.GetPatients)
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
	req, err := http.NewRequest("GET", "/doctor/patient/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockPatientGormRepo(nil)
	profSrv := doctorSrv.NewPatientService(ProfReq)
	profHandler := doctorHadler.NewDoctorPatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/doctor/patient/:id", profHandler.GetSinglePatient)
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

func TestPostPatient(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
	  }`)
	req, err := http.NewRequest("POST", "/doctor/patient/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockPatientGormRepo(nil)
	profSrv := doctorSrv.NewPatientService(ProfReq)
	profHandler := doctorHadler.NewDoctorPatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/doctor/patient/", profHandler.PostPatient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestPutPatient(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
		}`)
	req, err := http.NewRequest("PUT", "/doctor/patient/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockPatientGormRepo(nil)
	profSrv := doctorSrv.NewPatientService(ProfReq)
	profHandler := doctorHadler.NewDoctorPatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/doctor/patient/:id", profHandler.PutPatient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestDeletePatient(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/doctor/patient/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockPatientGormRepo(nil)
	profSrv := doctorSrv.NewPatientService(ProfReq)
	profHandler := doctorHadler.NewDoctorPatientHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/doctor/patient/:id", profHandler.DeletePatient)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
