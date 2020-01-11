package Patient

import (
	"bytes"
	"encoding/json"
	"fmt"
	entity2 "github.com/getach1/web1/web1_group_project/hospital_client/delivery/entity"
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

func FetchPetient(id int) (entity2.Petient, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		CheckErr(err)
		return entity2.Petient{}, err
	}
	userdata := entity2.Petient{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		CheckErr(err)
		return entity2.Petient{}, err
	}
	err = json.Unmarshal(body, &userdata)
	if err != nil {
		CheckErr(err)
		return entity2.Petient{}, err
	}
	return userdata, nil
}

func UpdateProfile(petient entity2.Petient) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, petient.ID)
	output, err := json.MarshalIndent(petient, "", "\t\t")
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(output))
	_, err = client.Do(req)
	CheckErr(err)

}

func FetchDoctors() ([]entity2.Doctor, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", doctorURL, nil)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		CheckErr(err)
		return []entity2.Doctor{}, err
	}
	doctordata := []entity2.Doctor{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		CheckErr(err)
		return []entity2.Doctor{}, err
	}
	err = json.Unmarshal(body, &doctordata)
	if err != nil {
		CheckErr(err)
		return []entity2.Doctor{}, err
	}
	return doctordata, nil
}
