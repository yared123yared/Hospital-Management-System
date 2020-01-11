package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yaredsolomon/webProgram1/hospital2/entity"
	"io/ioutil"
	"net/http"
)

var baseURL1 = "http://localhost:4444/v1/doctor/appointments/"

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
func Doctor(id int) (*entity.Doctor, error) {
	fmt.Println(id)
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL1, id)
	fmt.Println(URL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	userdata := &entity.Doctor{}
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
func DeleteAppointment(id int) error {
	fmt.Println("welcome to the deletemethod")
	fmt.Println("thise is the recived id")
	fmt.Println(id)
	//client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL, id)
	fmt.Println(URL)
	_, err := http.NewRequest("DELETE", URL, nil)
	fmt.Println("deleted")
	if err != nil {
		panic(err)
	}
	fmt.Println("no error")
	//res, err := client.Do(req)
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
func Doctors() (*[]entity.Doctor, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	usdata := &[]entity.Doctor{}
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

func UpdateDoctor(doctor *entity.Doctor, id int) error {
	fmt.Println("i am at the Update method")
	client := &http.Client{}
	fmt.Println("i am at the Update method2")
	URL := fmt.Sprintf("%s%d", baseURL1, id)
	fmt.Println("i am at the Update method3")
	output, err := json.MarshalIndent(doctor, "", "\t\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(output))
	if err != nil {
		panic(err)
	}
	fmt.Println("i am at the Update method5")
	fmt.Println(URL)
	_, err = client.Do(req)
	fmt.Println("i am at the Update method6")
	//res, err := client.Get(URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("i have finished with out error")
	fmt.Println(" i am about to left the update method")
	return nil
}
