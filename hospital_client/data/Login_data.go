package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/web1_group_project/hospital_client/entity"
)

var baseURL_login = "http://localhost:8180/v1/general/users/"
var baseURL_userByEmail = "http://localhost:8180//v1/admin/usersByEmail/"
var baseURL_userRoles="http://localhost:8180/v1/admin/userRoles/"


func GetUser(user *entity.User) (*entity.User, error) {
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

	usdata := &entity.User{}
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
func UserByEmail(email string) (*entity.User, error){
	
	user := &entity.User{}
	user.Email=email
	fmt.Println("i am at userByEmail method")
	client := &http.Client{}
	fmt.Println("i am at the userByEmail method2")
	URL := fmt.Sprintf("%s%s", baseURL_userByEmail,email)
	fmt.Println("i am at the userByEmail method3")

	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	req, err := http.NewRequest("GET", URL, bytes.NewBuffer(output))
	if err != nil {
		panic(err)
	}
	fmt.Println("i am at the userByEmail method5")
	fmt.Println(URL)
	res, err := client.Do(req)
	//res, err := client.Get(URL)

	if err != nil {
		return nil, err
	}

	usdata := &entity.User{}
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
func UserRoles(user *entity.User) ([]entity.Role, error){
	
	fmt.Println("i am at UserRoles method")
	client := &http.Client{}
	fmt.Println("i am at the UserRoles method2")
	URL := fmt.Sprintf("%s", baseURL_userRoles)
	fmt.Println("i am at the UserRoles method3")

	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	req, err := http.NewRequest("GET", URL, bytes.NewBuffer(output))
	if err != nil {
		panic(err)
	}
	fmt.Println("i am at the UserRoles method5")
	fmt.Println(URL)
	res, err := client.Do(req)
	//res, err := client.Get(URL)

	if err != nil {
		return nil, err
	}
	fmt.Println("i am at the UserRoles method6")
	usrole := &[]entity.Role{}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	fmt.Println("i am at the UserRoles method7")
	err = json.Unmarshal(body, usrole)
	if err != nil {
		panic(err)
	}
	fmt.Println("i am at the UserRoles method8")
	fmt.Println(*usrole)
	return *usrole, nil


}

func Patients() ([]entity.User, []error){
	return nil,nil
}

