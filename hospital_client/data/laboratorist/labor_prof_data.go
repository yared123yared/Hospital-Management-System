package service

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

	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

var baseURLProf = "http://localhost:8182/v1/labor/profile/"

//GetLaboratorist returns pharmacist
func GetLaboratorist(id uint) (*entity.Laboratorist, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%d", baseURLProf, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	laboratorist := &entity.Laboratorist{}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, laboratorist)
	if err != nil {
		return nil, err
	}
	return laboratorist, nil

}

//PutLaboratorist updates Pharmacist
func PutLaboratorist(laboratorist *entity.Laboratorist) {
	client := &http.Client{}
	id := laboratorist.ID
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(laboratorist)
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
