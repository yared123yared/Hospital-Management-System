package data
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/web1_group_project/hospital_client/entity"
)
var baseURL_session = "http://localhost:8180/v1/general/sessions/"
func StoreSession(session *entity.Session) (*entity.Session, error){
	fmt.Println("i have entered to the session method")
	client := &http.Client{}
	URL := fmt.Sprintf("%s", baseURL_session)
	output, err := json.MarshalIndent(session, "", "\t\t")

	fmt.Println("i have entered to the session method2")
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(output))
	fmt.Println("i have entered to the session method3")
	res, err := client.Do(req)
	//res, err := client.Get(URL)
	if err != nil {
		return nil,err
	}
	fmt.Println("i have entered to the session method4")
	usession := &entity.Session{}
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("i have entered to the session method5")
	fmt.Println(body)
	err = json.Unmarshal(body, usession)
	if err != nil {
		panic(err)
	}
	fmt.Println("i have entered to the session method6")
	fmt.Println(usession)
	return usession, nil

	
}
func DeleteSession(sessionID string) (error){
	fmt.Println("welcome to the DeleteSession")
	fmt.Println("thise is the recived id")

	id, err := strconv.Atoi(sessionID)

	fmt.Println(id)
	client := &http.Client{}
	URL := fmt.Sprintf("%s%d", baseURL_session, id)
	fmt.Println(URL)
	req, err := http.NewRequest("DELETE", URL, nil)
	fmt.Println("deleted")
	if err != nil {
		panic(err)
	}
	fmt.Println("no error")
	_, err = client.Do(req)
	
	return nil
}