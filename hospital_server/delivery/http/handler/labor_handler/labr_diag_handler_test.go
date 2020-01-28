package labor_handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	laborRepo "github.com/web1_group_project/hospital_server/Laboratorist/repository"
	laborSrv "github.com/web1_group_project/hospital_server/Laboratorist/service"
	laborHandler "github.com/web1_group_project/hospital_server/delivery/http/handler/labor_handler"
	"github.com/web1_group_project/hospital_server/entity"
)

func TestGetDiagnosiss(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharm/diagnosis", nil)
	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockDiagnosisGormRepo(nil)
	profSrv := laborSrv.NewDiagnosisService(ProfReq)
	profHandler := laborHandler.NewLaborDiagnosisHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/diagnosis", profHandler.GetDiagnosiss)
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

func TestGetSingleDiagnosis(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharm/diagnosis/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockDiagnosisGormRepo(nil)
	profSrv := laborSrv.NewDiagnosisService(ProfReq)
	profHandler := laborHandler.NewLaborDiagnosisHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/diagnosis/:id", profHandler.GetSingleDiagnosis)
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

func TestPostDiagnosis(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"LaboratoristId": 1,
		"Description": "description",
		"DiagonosesDate": "2020-01-01T00:00:00+03:00",
		"Reponse": "response",
		"ResponseDate": "2020-01-01T00:00:00+03:00"
	}`)
	req, err := http.NewRequest("POST", "/pharm/medicine", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockDiagnosisGormRepo(nil)
	profSrv := laborSrv.NewDiagnosisService(ProfReq)
	profHandler := laborHandler.NewLaborDiagnosisHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/pharm/medicine", profHandler.PostDiagnosis)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestPutDiagnosis(t *testing.T) {

	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"LaboratoristId": 1,
		"Description": "description",
		"DiagonosesDate": "2020-01-01T00:00:00+03:00",
		"Reponse": "response",
		"ResponseDate": "2020-01-01T00:00:00+03:00"
	}`)
	req, err := http.NewRequest("PUT", "/pharm/medicine/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockDiagnosisGormRepo(nil)
	profSrv := laborSrv.NewDiagnosisService(ProfReq)
	profHandler := laborHandler.NewLaborDiagnosisHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/pharm/medicine/:id", profHandler.PutDiagnosis)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeleteDiagnosis(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"PatientId": 1,
		"PatientName": "patient1",
		"DoctorId": 1,
		"LaboratoristId": 1,
		"Description": "description",
		"DiagonosesDate": "2020-01-01T00:00:00+03:00",
		"Reponse": "response",
		"ResponseDate": "2020-01-01T00:00:00+03:00"
	}`)
	req, err := http.NewRequest("DELETE", "/pharm/medicine/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockDiagnosisGormRepo(nil)
	profSrv := laborSrv.NewDiagnosisService(ProfReq)
	profHandler := laborHandler.NewLaborDiagnosisHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/pharm/medicine/:id", profHandler.PutDiagnosis)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
