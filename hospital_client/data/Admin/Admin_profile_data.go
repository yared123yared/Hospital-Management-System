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

	"github.com/web1_group_project/hospital_client/entity"
)

var baseURLProfile = "http://localhost:8182/admin/profiles"

func GetProfile(id uint) (*entity.Profile, error) {
	fmt.Println("Enteritn")
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURLProfile, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	presc := &entity.Profile{}
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
func PostProfile(presc *entity.Profile) {
	client := &http.Client{}

	json, err := json.Marshal(presc)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", baseURLProfile, bytes.NewBuffer(json))
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

func writeFile(mf *multipart.File, fname string) {
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
