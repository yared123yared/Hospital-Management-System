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

func TestGetAppointments(t *testing.T) {

	req, err := http.NewRequest("GET", "/doctor/appointent", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockAppointmentGormRepo(nil)
	profSrv := doctorSrv.NewAppointmentService(ProfReq)
	profHandler := doctorHadler.NewDoctorAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/doctor/appointent", profHandler.GetAppointments)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.AppointmentMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSingleAppointment(t *testing.T) {
	req, err := http.NewRequest("GET", "/doctor/appointent/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockAppointmentGormRepo(nil)
	profSrv := doctorSrv.NewAppointmentService(ProfReq)
	profHandler := doctorHadler.NewDoctorAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/doctor/appointent/:id", profHandler.GetSingleAppointment)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.AppointmentMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestPutAppointment(t *testing.T) {
	var jsonStr = []byte(`{
    "ID": 1,
    "PatientId": 1,
    "PatientName": "patient1",
    "DoctorId": 1,
    "Date": "2020-09-09T00:00:00+03:00"
    }`)
	req, err := http.NewRequest("PUT", "/doctor/appointent/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockAppointmentGormRepo(nil)
	profSrv := doctorSrv.NewAppointmentService(ProfReq)
	profHandler := doctorHadler.NewDoctorAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/doctor/appointent/:id", profHandler.PutAppointment)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}

func TestDeleteAppointment(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/doctor/appointent/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := doctorRepo.NewMockAppointmentGormRepo(nil)
	profSrv := doctorSrv.NewAppointmentService(ProfReq)
	profHandler := doctorHadler.NewDoctorAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/doctor/appointent/:id", profHandler.DeleteAppointment)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
