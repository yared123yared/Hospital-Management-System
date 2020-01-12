package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"

// 	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
// 	"github.com/webProj/sessions"
// 	"github.com/webProject2019/utils"
// )

// type UserHandler struct {
// 	templ *template.Template
// }

// func NewUserHandler(T *template.Template) *UserHandler {
// 	return &UserHandler{templ: T}
// }
// func (uh *UserHandler) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Post Login page")
// 	client := http.Client{}
// 	user := entity.Profile{}
// 	id, err := strconv.Atoi(r.FormValue("ID"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	user.ID = uint(id)
// 	user.Password = r.FormValue("password")
// 	user.Phone = ""
// 	user.Sex = ""
// 	user.Role = ""
// 	user.Email = ""
// 	user.Phone = ""
// 	output, err := json.MarshalIndent(user, "", "\t\t")
// 	session, _ := sessions.Store.Get(r, "session")
// 	if err != nil {
// 		sessions.Message("Internal Server Error", "danger", r, w)
// 		http.Redirect(w, r, "/login", 302)
// 	}

// 	req, err := http.NewRequest("POST", "http://localhost:9091/user/check", bytes.NewBuffer(output))
// 	response, err := client.Do(req)
// 	if http.StatusNotFound == response.StatusCode {
// 		sessions.Message("Invalid Username or Password", "danger", r, w)
// 		http.Redirect(w, r, "/login", 302)
// 	}
// 	if http.StatusUnprocessableEntity == response.StatusCode {
// 		sessions.Message("Internal Server Error", "danger", r, w)
// 		http.Redirect(w, r, "/login", 302)

// 	}
// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		sessions.Message("Internal Server Error", "danger", r, w)
// 		http.Redirect(w, r, "/login", 302)
// 	}

// 	var responseObject entity.Profile
// 	_ = json.Unmarshal(responseData, &responseObject)
// 	fmt.Println("login handler--", responseObject)
// 	session.Values["USERID"] = responseObject.ID
// 	//fmt.Println(user.Id)
// 	fmt.Println(session.Values["USERID"])
// 	_ = session.Save(r, w)

// 	fmt.Println("Redirection and scession created")
// 	http.Redirect(w, r, "/", 302)
// }
// func (uh *UserHandler) LoginGetHandler(w http.ResponseWriter, r *http.Request) {
// 	message, alert := sessions.Flash(r, w)
// 	fmt.Println("Logininininin")
// 	if r.Method == http.MethodGet {

// 		fmt.Println("Get Login page")
// 		_ = uh.templ.ExecuteTemplate(w, "login.html", struct {
// 			Alert utils.Alert
// 		}{
// 			Alert: utils.NewAlert(message, alert),
// 		})
// 	} else {
// 		uh.LoginPostHandler(w, r)
// 	}

// }
