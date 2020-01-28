package pharm_handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	PharmacistRepo "github.com/web1_group_project/hospital_server/Pharmacist/repository"
	PharmacistService "github.com/web1_group_project/hospital_server/Pharmacist/service"
	handler2 "github.com/web1_group_project/hospital_server/delivery/http/handler/pharm_handler"
	"github.com/web1_group_project/hospital_server/entity"
)

func TestGetMedicines(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharm/medicine", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmMedRepo(nil)
	profSrv := PharmacistService.NewMedicineService(ProfReq)
	profHandler := handler2.NewPharmMedicineHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/medicine", profHandler.GetMedicines)
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

func TestGetSingleMedicine(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharm/medicine/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmMedRepo(nil)
	profSrv := PharmacistService.NewMedicineService(ProfReq)
	profHandler := handler2.NewPharmMedicineHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/medicine/:id", profHandler.GetSingleMedicine)
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

func TestGetMultipleMedicines(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharm/medicine/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmMedRepo(nil)
	profSrv := PharmacistService.NewMedicineService(ProfReq)
	profHandler := handler2.NewPharmMedicineHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/pharm/medicine/:addedby", profHandler.GetMultipleMedicines)
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

func TestPostMedicine(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"CategoryName": "China Product",
		"MedicineName": "pencilin",
		"ExpiredAt": "2022-09-12T00:00:00+03:00",
		"Amount": 12,
		"AddedBy": 1
}`)
	req, err := http.NewRequest("POST", "/pharm/medicine", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmMedRepo(nil)
	profSrv := PharmacistService.NewMedicineService(ProfReq)
	profHandler := handler2.NewPharmMedicineHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/pharm/medicine", profHandler.PostMedicine)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusCreated {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
	// expected := entity.PharmacistMock
	// body, err := ioutil.ReadAll(rr.Body)

	// if err != nil {
	// 	t.Fatal(err)
	// }
	// expected := entity.PharmacistMock
	// body, err := ioutil.ReadAll(rr.Body)

	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if !bytes.Contains(body, []byte("01")) {
	// 	t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	// }
}

func TestUpdateMedicine(t *testing.T) {
	var jsonStr = []byte(`{
		"ID": 1,
		"CategoryName": "China Product",
		"MedicineName": "pencilin",
		"ExpiredAt": "2022-09-12T00:00:00+03:00",
		"Amount": 12,
		"AddedBy": 1
}`)
	req, err := http.NewRequest("PUT", "/pharm/medicine/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmMedRepo(nil)
	profSrv := PharmacistService.NewMedicineService(ProfReq)
	profHandler := handler2.NewPharmMedicineHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.PUT("/pharm/medicine/:id", profHandler.PutMedicine)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}

func TestDeletMedicine(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/pharm/medicine/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	ProfReq := PharmacistRepo.NewMockPharmMedRepo(nil)
	profSrv := PharmacistService.NewMedicineService(ProfReq)
	profHandler := handler2.NewPharmMedicineHandler(profSrv)

	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.DELETE("/pharm/medicine/:id", profHandler.DeleteMedicine)
	router.ServeHTTP(rr, req)

	if statu := rr.Code; statu != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code : got %v want %v", statu, http.StatusOK)
	}
}
