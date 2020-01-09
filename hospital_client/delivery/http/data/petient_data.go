package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	entity "github.com/getach1/web1/Project/entities"
	"io/ioutil"
	"log"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var baseURL = "http://localhost:8100/v1/admin/petients/"
var doctorURL = "http://localhost:8100/v1/admin/doctors/"

//var requestURL="http://localhost:8100/v1/admin/requests/"

func FetchPetient(id int) (entity.Petient, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		CheckErr(err)
		return entity.Petient{}, err
	}
	userdata := entity.Petient{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		CheckErr(err)
		return entity.Petient{}, err
	}
	err = json.Unmarshal(body, &userdata)
	if err != nil {
		CheckErr(err)
		return entity.Petient{}, err
	}
	return userdata, nil
}

func UpdateProfile(petient entity.Petient) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, petient.ID)
	output, err := json.MarshalIndent(petient, "", "\t\t")
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(output))
	_, err = client.Do(req)
	CheckErr(err)

}

func FetchDoctors() ([]entity.Doctor, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", doctorURL, nil)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		CheckErr(err)
		return []entity.Doctor{}, err
	}
	doctordata := []entity.Doctor{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		CheckErr(err)
		return []entity.Doctor{}, err
	}
	err = json.Unmarshal(body, &doctordata)
	if err != nil {
		CheckErr(err)
		return []entity.Doctor{}, err
	}
	return doctordata, nil
}
