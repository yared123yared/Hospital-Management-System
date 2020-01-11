package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/web1_group_project/hospital_client/entity"

	"io/ioutil"
	"net/http"
)

var baseURL_login = "http://localhost:4479/v1/general/users/"

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

func GetUser(user *entity.Profile) (*entity.Profile, error) {
	fmt.Println("i am at the Update method")
	client := &http.Client{}
	fmt.Println("i am at the Update method2")
	URL := fmt.Sprintf("%s", baseURL_login)
	fmt.Println("i am at the Update method3")

	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	req, err := http.NewRequest("GET", URL, bytes.NewBuffer(output))
	if err != nil {
		panic(err)
	}
	fmt.Println("i am at the Update method5")
	fmt.Println(URL)
	res, err := client.Do(req)
	//res, err := client.Get(URL)

	if err != nil {
		return nil, err
	}

	usdata := &entity.Profile{}
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
