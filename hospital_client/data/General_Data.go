package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/web1_group_project/hospital_client/entity"
)

var baseURL_general = "http://localhost:8180/v1/doctor/pharmacists/"
var baseURL_general1 = "http://localhost:8180/v1/doctor/laboratorists/"
var baseURL_roles = "http://localhost:8180/v1/general/roles/"

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
	fmt.Println("i have done with pharmacist retrive method")
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
func Roles() (*[]entity.Role, error){
	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL_roles)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	
	if err != nil {
	
		return nil, err
	}
	roles := &[]entity.Role{}
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, roles)
	if err != nil {
		return nil, err
	}
	fmt.Println(roles)
	return roles, nil
}
