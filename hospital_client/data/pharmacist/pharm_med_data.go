package pharmacist_data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/web1_group_project/hospital_client/entity"
)

var baseURLMed = "http://localhost:8182/v1/pharm/medicines/"

//GetMedicine returns single Medicine
func GetMedicine(id uint) (*entity.Medicine, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%d", baseURLMed, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	medicine := &entity.Medicine{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, medicine)
	if err != nil {
		return nil, err
	}
	return medicine, nil
}
func GetMedicines() ([]entity.Medicine, error) {
	client := &http.Client{}
	sp := ""
	URL := fmt.Sprintf("%s%s", baseURLMed, sp)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var medicines = make([]entity.Medicine, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &medicines)

	return medicines, nil
}

func PutMedicine(medicine *entity.Medicine) {
	client := &http.Client{}
	id := medicine.ID
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(medicine)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", baseURLMed+t, bytes.NewBuffer(json))
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

func DeleteMedicine(ID uint) {
	client := &http.Client{}
	t := strconv.Itoa(int(ID))
	// Create request
	req, err := http.NewRequest("DELETE", baseURLMed+t, nil)
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

func PostMedicine(medicine *entity.Medicine) {
	client := &http.Client{}

	json, err := json.Marshal(medicine)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", baseURLMed, bytes.NewBuffer(json))
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
