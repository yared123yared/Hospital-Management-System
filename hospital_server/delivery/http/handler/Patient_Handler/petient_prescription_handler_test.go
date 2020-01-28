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

func TestGetPrescriptions(t *testing.T) {
	req, err := http.NewRequest("GET", "/patient/prescription", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPrescriptionGormRepo(nil)
	profSrv := patientSrv.NewPrescriptionService(ProfReq)
	profHandler := patientHadler.NewPetientPrescriptionHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/patient/prescription", profHandler.GetPrescriptions)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PrescriptionMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSinglePrescription(t *testing.T) {
	req, err := http.NewRequest("GET", "/patient/prescription/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := patientRepo.NewMockPrescriptionGormRepo(nil)
	profSrv := patientSrv.NewPrescriptionService(ProfReq)
	profHandler := patientHadler.NewPetientPrescriptionHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/patient/prescription/:id", profHandler.GetSinglePrescription)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PrescriptionMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
