package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

var baseURLDiag = "http://localhost:8182/v1/labor/diag/"
var baseURLDiag2 = "http://localhost:8182/v1/labor/multiDiag/"

//GetMedicine returns single Medicine
func GetDiagnosis(id uint) (*entity.Diagnosis, error) {
	fmt.Println("Enteritn")
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURLDiag, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	presc := &entity.Diagnosis{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, presc)
	if err != nil {
		return nil, err
	}
	fmt.Println("pres", presc)
	return presc, nil
}
func Diagnosis(id uint) ([]entity.Diagnosis, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURLDiag, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var prescs = make([]entity.Diagnosis, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &prescs)

	return prescs, nil
}
func GetDiagnosiss() ([]entity.Diagnosis, error) {
	client := &http.Client{}
	sp := ""
	URL := fmt.Sprintf("%s%s", baseURLDiag2, sp)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var prescs = make([]entity.Diagnosis, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &prescs)
	return prescs, nil
}

func PutDiagnosis(presc *entity.Diagnosis) {
	client := &http.Client{}
	id := presc.ID
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", baseURLDiag+t, bytes.NewBuffer(json))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Println("put done")
}

func DeleteDiagnosis(ID uint) {
	client := &http.Client{}
	t := strconv.Itoa(int(ID))
	// Create request
	req, err := http.NewRequest("DELETE", baseURLDiag+t, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

}

func PostDiagnosis(presc *entity.Diagnosis) {
	client := &http.Client{}

	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", baseURLDiag, bytes.NewBuffer(json))
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Println("post done")
}
