package Admin

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

var baseURLDctrs = "http://localhost:8180/v1/admin/doctors/"

func GetDoctor(id uint) (*entity.Doctor, error) {
	fmt.Println("Enteritn")
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURLDctrs, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	presc := &entity.Doctor{}
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
func Doctors() ([]entity.Doctor, error) {
	fmt.Println("Geting all The doctors")
	client := &http.Client{}
	sp := ""
	URL := fmt.Sprintf("%s%s", baseURLDctrs, sp)
	req, _ := http.NewRequest("GET", URL, nil)
	res, _ := client.Do(req)
	var prescs = make([]entity.Doctor, 0)
	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &prescs)

	return prescs, nil
}

func PutDoctor(presc *entity.Doctor) {
	client := &http.Client{}
	id := presc.Uuid
	t := strconv.Itoa(int(id))
	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", baseURLDctrs+t, bytes.NewBuffer(json))
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
	fmt.Println("id put", id)
	fmt.Println("put done")
}

func DeleteDoctor(ID uint) {
	client := &http.Client{}
	t := strconv.Itoa(int(ID))
	// Create request
	req, err := http.NewRequest("DELETE", baseURLDctrs+t, nil)
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

func PostDoctor(presc *entity.Doctor) {
	client := &http.Client{}

	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", baseURLDctrs, bytes.NewBuffer(json))
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
	fmt.Println("post done with", presc)
}
func WriteFile(mf *multipart.File, fname string) {
	fmt.Println("i am at the image store")

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("This is first erorrrrrrrrrrrrrrrrrrrrrr")
		panic(err)
	}

	path := filepath.Join(wd, "../ui", "assets", "img", fname)
	image, err := os.Create(path)
	//	fmt.Println(*image)

	if err != nil {
		fmt.Println("This is second erorrrrrrrrrrrrrrrrrrrrrr")

		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
