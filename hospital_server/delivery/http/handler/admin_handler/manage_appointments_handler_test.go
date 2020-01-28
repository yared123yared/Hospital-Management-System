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

func TestGetAppointments(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/appointments", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageAppointmetRepository(nil)
	profSrv := adminSrv.NewManageAppointmetService(ProfReq)
	profHandler := adminHandler.NewManageAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/appointments", profHandler.GetAppointments)
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

func TestGetSingleAppointment(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/appointments/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageAppointmetRepository(nil)
	profSrv := adminSrv.NewManageAppointmetService(ProfReq)
	profHandler := adminHandler.NewManageAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/appointments/:id", profHandler.GetSingleAppointment)
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
func TestAddAppointment(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
}`)
	req, err := http.NewRequest("POST", "/admin/appointments/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageAppointmetRepository(nil)
	profSrv := adminSrv.NewManageAppointmetService(ProfReq)
	profHandler := adminHandler.NewManageAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/appointments/", profHandler.AddAppointment)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestUpdateAppointment(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"Date": "2020-09-09T00:00:00+03:00"
	  }`)
	req, err := http.NewRequest("PUT", "/admin/appointments/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageAppointmetRepository(nil)
	profSrv := adminSrv.NewManageAppointmetService(ProfReq)
	profHandler := adminHandler.NewManageAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/admin/appointments/:id", profHandler.UpdateAppointment)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeleteAppointment(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/admin/appointments/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageAppointmetRepository(nil)
	profSrv := adminSrv.NewManageAppointmetService(ProfReq)
	profHandler := adminHandler.NewManageAppointmentHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/admin/appointments/:id", profHandler.DeleteAppointment)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
