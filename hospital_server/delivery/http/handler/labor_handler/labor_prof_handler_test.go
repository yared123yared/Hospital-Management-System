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

func TestGetProfiles(t *testing.T) {
	req, err := http.NewRequest("GET", "/labor/profile", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockLabratoristProfileGormRepo(nil)
	profSrv := laborSrv.NewLabratoristProfileService(ProfReq)
	profHandler := laborHandler.NewLaborProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/labor/profile", profHandler.GetProfiles)
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

func TestGetSingleProfile(t *testing.T) {
	req, err := http.NewRequest("GET", "/labor/profile/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := laborRepo.NewMockLabratoristProfileGormRepo(nil)
	profSrv := laborSrv.NewLabratoristProfileService(ProfReq)
	profHandler := laborHandler.NewLaborProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/labor/profile/:id", profHandler.GetSingleProfile)
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

func TestPutProfile(t *testing.T) {
	jsonStr := []byte(`{
		"ID": 6,
		"Uuid": 1110,
		"Profile": {
			"ID": 1110,
			"FullName": "laborist6",
			"Password": "la@isttt",
			"Email": "labr1@gmail.com",
			"Phone": "09125",
			"Address": "Bahir Dar,Ethiopia",
			"Image": "lab1.png",
			"Sex": "Male",
			"RoleId": 5,
			"BirthDate": "1981-03-09T00:00:00+03:00",
			"Description": "laboratorist at the hospital"
		},
		"Diagnosis": []`)
	req, err := http.NewRequest("PUT", "/labor/profile/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	ProfReq := laborRepo.NewMockLabratoristProfileGormRepo(nil)
	profSrv := laborSrv.NewLabratoristProfileService(ProfReq)
	profHandler := laborHandler.NewLaborProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/labor/profile/:id", profHandler.PutProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PharmacistMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("1")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}

}

func TestDeleteProfile(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/labor/profile/1", nil)

	if err != nil {
		t.Fatal(err)
	}
	ProfReq := laborRepo.NewMockLabratoristProfileGormRepo(nil)
	profSrv := laborSrv.NewLabratoristProfileService(ProfReq)
	profHandler := laborHandler.NewLaborProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/labor/profile/:id", profHandler.DeleteProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.PharmacistMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Contains(body, []byte("1")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}

}
