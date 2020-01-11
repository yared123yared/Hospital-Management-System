package data

import (
	"encoding/json"
	"fmt"
	"github.com/getach1/web1/web1_group_project_old_new/hospital_client/entity"
	"io/ioutil"
	"net/http"
)

var baseURL_general = "http://localhost:4444/v1/doctor/pharmacists/"
var baseURL_general1 = "http://localhost:4444/v1/doctor/laboratorists/"

func Pharmacists() (*[]entity.Pharmacist, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL_general)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	usdata := &[]entity.Pharmacist{}
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, usdata)
	if err != nil {
		return nil, err
	}
	fmt.Println(usdata)
	return usdata, nil
}
func Laboratorists() (*[]entity.Laboratorist, error) {

	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL_general1)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	usdata := &[]entity.Laboratorist{}
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, usdata)
	if err != nil {
		return nil, err
	}
	fmt.Println(usdata)
	return usdata, nil
}
