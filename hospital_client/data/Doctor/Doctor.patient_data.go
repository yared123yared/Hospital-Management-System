package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getach1/web1/web1_group_project-master/hospital_client/entity"
	"io/ioutil"
	"net/http"
)

var baseURL = "http://localhost:1885/v1/admin/users/"

// User represents User data
//type User struct {
//	UserId    int    `json:"userId"`
//	ID        int    `json:"id"`
//	Title     string `json:"title"`
//	Body string `json:"body"`
//
//}jhg

//// SingleData represents a single User
//type SingleData struct {
//	user User
//}
//
//// CollectionData represents collection of Users
//type CollectionData struct {
//	users []User
//}

//FetchUser fetchs a single user by its id

func StorePatients(petient *entity.Petient) error {
	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL)
	output, err := json.MarshalIndent(petient, "", "\t\t")
	//l := 10
	//body := make([]byte, l)
	//r.Body.Read(body)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(output))
	_, err = client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return err
	}
	//userdata := &entity.Petient{}
	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.Unmarshal(body, userdata)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("thise is the struct data")
	//fmt.Println(userdata)
	return nil
}
func FetchUser(id int) (*entity.Petient, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, id)

	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	userdata := &entity.Petient{}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, userdata)
	if err != nil {
		return nil, err
	}
	fmt.Println("thise is the struct data")
	fmt.Println(userdata)
	return userdata, nil
}
func DeleteUser(id int) error {
	fmt.Println("welcome to the deletemethod")
	fmt.Println("thise is the recived id")
	fmt.Println(id)
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, id)
	fmt.Println(URL)
	req, err := http.NewRequest("DELETE", URL, nil)
	fmt.Println("deleted")
	if err != nil {
		panic(err)
	}
	fmt.Println("no error")
	_, err = client.Do(req)
	////res, err := client.Get(URL)

	//userdata := &entity.Petient{}
	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	//if err != nil {
	//	return  err
	//}
	//
	//err = json.Unmarshal(body, userdata)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("thise is the struct data")
	//fmt.Println(&userdata)
	return nil
}

// FetchUsers fetchs all users on a given page
func FetchUsers() (*[]entity.Petient, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	usdata := &[]entity.Petient{}
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
