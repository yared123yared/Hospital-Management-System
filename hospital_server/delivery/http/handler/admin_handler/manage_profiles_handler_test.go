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

func TestGetProfiles(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/profile", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageProfileRepository(nil)
	profSrv := adminSrv.NewManageProfileService(ProfReq)
	profHandler := adminHandler.NewManageProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/profile", profHandler.GetProfiles)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.ProfileMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSingleProfile(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/profile/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageProfileRepository(nil)
	profSrv := adminSrv.NewManageProfileService(ProfReq)
	profHandler := adminHandler.NewManageProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/admin/profile/:id", profHandler.GetSingleProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	expected := entity.ProfileMock
	body, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("01")) {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
func TestAddProfile(t *testing.T) {
	var jsonStr = []byte(`{
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
	req, err := http.NewRequest("POST", "/admin/profile/", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageProfileRepository(nil)
	profSrv := adminSrv.NewManageProfileService(ProfReq)
	profHandler := adminHandler.NewManageProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/profile/", profHandler.AddProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
func TestUpdateProfile(t *testing.T) {
	var jsonStr = []byte(`{
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
	req, err := http.NewRequest("PUT", "/admin/profile/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageProfileRepository(nil)
	profSrv := adminSrv.NewManageProfileService(ProfReq)
	profHandler := adminHandler.NewManageProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/admin/profile/:id", profHandler.UpdateProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}

}

func TestDeleteProfile(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/admin/profile/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := adminRepo.NewMockManageProfileRepository(nil)
	profSrv := adminSrv.NewManageProfileService(ProfReq)
	profHandler := adminHandler.NewManageProfileHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/admin/profile/:id", profHandler.DeleteProfile)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
