package pharm_handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"

	PharmacistRepo "github.com/web1_group_project/hospital_server/Pharmacist/repository"
	PharmacistService "github.com/web1_group_project/hospital_server/Pharmacist/service"
	handler2 "github.com/web1_group_project/hospital_server/delivery/http/handler/pharm_handler"
	"github.com/web1_group_project/hospital_server/entity"
)

func TestGetProfiles(t *testing.T) {
	ReqUrl := url.PathEscape("/pharm/profile")
	req, err := http.NewRequest("GET", ReqUrl, nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmacistProfileGormRepo(nil)
	profSrv := PharmacistService.NewPharmacistProfileService(ProfReq)
	profHandler := handler2.NewPharmProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/profile", profHandler.GetProfiles)
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

func TestGetSingleProfiles(t *testing.T) {

	req, err := http.NewRequest("GET", "/pharm/profile/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmacistProfileGormRepo(nil)
	profSrv := PharmacistService.NewPharmacistProfileService(ProfReq)
	profHandler := handler2.NewPharmProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/profile/:id", profHandler.GetSingleProfile)
	router.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
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
func TestPostProfile(t *testing.T) {

	var jsonStr = []byte(`{"ID": 6,
	"Uuid": 1028,
	"Profile": {
			"ID": 1028,
			"FullName": "Abe123",
			"Password": "a@alex123",
			"Email": "alex@gmail.com",
			"Phone": "0911223344",
			"Address": "Addis Ababa,Ethiopia",
			"Image": "alexProf.png",
			"Sex": "Male",
			"RoleId": 3,
			"BirthDate": "1980-01-21T00:00:00+03:00",
			"Description": "Pharmacist at black lion Hospital"
	},
	"Medicine": [],
	"Prescription": []
}`)

	req, err := http.NewRequest("POST", "/pharm/profile", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	ProfReq := PharmacistRepo.NewMockPharmacistProfileGormRepo(nil)
	profSrv := PharmacistService.NewPharmacistProfileService(ProfReq)
	profHandler := handler2.NewPharmProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/pharm/profile", profHandler.PostProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusCreated)
	}
	expected := entity.PharmacistMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Contains(body, []byte("10")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
