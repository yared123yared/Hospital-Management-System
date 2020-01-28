package pharmacist_data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/web1_group_project/hospital_client/entity"
)

var baseURLProf = "http://localhost:8180/v1/pharm/profile/"

//GetPharmacist returns pharmacist
func GetPharmacist(id uint) (*entity.Pharmacist, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%d", baseURLProf, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	pharmacist := &entity.Pharmacist{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, pharmacist)
	if err != nil {
		return nil, err
	}
	return pharmacist, nil

}

//PutPharmacist updates Pharmacist
func PutPharmacist(pharmacist *entity.Pharmacist) {
	client := &http.Client{}
	id := pharmacist.ID
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(pharmacist)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", baseURLProf+t, bytes.NewBuffer(json))
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

func WriteFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
