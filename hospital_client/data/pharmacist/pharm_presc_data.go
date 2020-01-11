package pharmacist_data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/web1_group_project/hospital_client/entity"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var baseURLPres = "http://localhost:8182/v1/pharm/presc/"
var baseURLPres2 = "http://localhost:8182/v1/pharm/multiPre/"

//GetMedicine returns single Medicine
func GetPrescription(id uint) (*entity.Prescription, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%d", baseURLPres, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	presc := &entity.Prescription{}
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
func GetPrescriptions(id uint) ([]entity.Prescription, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURLPres2, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var prescs = make([]entity.Prescription, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &prescs)

	return prescs, nil
}
func Prescriptions() ([]entity.Prescription, error) {
	client := &http.Client{}
	sp := ""
	URL := fmt.Sprintf("%s%s", baseURLPres2, sp)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var prescs = make([]entity.Prescription, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &prescs)
	return prescs, nil
}

func PutPrescription(presc *entity.Prescription) {
	client := &http.Client{}
	id := presc.ID
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", baseURLPres+t, bytes.NewBuffer(json))
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

func DeletePrescription(ID uint) {
	client := &http.Client{}
	t := strconv.Itoa(int(ID))
	// Create request
	req, err := http.NewRequest("DELETE", baseURLPres+t, nil)
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

func PostPrescription(presc *entity.Prescription) {
	client := &http.Client{}

	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", baseURLPres, bytes.NewBuffer(json))
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
