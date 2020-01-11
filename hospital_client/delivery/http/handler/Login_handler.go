package handler

import (
	"fmt"
	"github.com/yaredsolomon/webProgram1/hospital2/delivery/utils"

	"github.com/yaredsolomon/webProgram1/hospital2/delivery/http/data"
	"github.com/yaredsolomon/webProgram1/hospital2/delivery/session"
	"github.com/yaredsolomon/webProgram1/hospital2/entity"
	"html/template"

	"net/http"

	"strconv"
	//"github.com/betsegawlemma/restaurant/menu"
)

// MenuHandler handles menu related requests
type loginHandler struct {
	tmpl *template.Template
	//doctorSrv doctor.CategoryService
}

//NewMenuHandler initializes and returns new MenuHandler
func NewLoginHandler(T *template.Template) *loginHandler {
	return &loginHandler{tmpl: T}
}
func (lh *loginHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	lh.tmpl.ExecuteTemplate(w, "Login.html", nil)
}

// Index handles request on route /
func (lh *loginHandler) Users(w http.ResponseWriter, r *http.Request) {
	users := entity.Profile{}
	id, err := strconv.Atoi(r.FormValue("id"))
	password := r.FormValue("password")

	if err != nil {
		panic(err)
	}
	users.ID = uint(id)
	users.Password = password

	//if err != nil {
	//	w.WriteHeader(http.StatusNoContent)
	//	tmpl.ExecuteTemplate(w, "error.layout", nil)
	//}
	fmt.Println("i am about to fech data")
	//users, err := data.FetchUsers()
	//petient:=[]entity.Petient

	users1, err := data.GetUser(&users)
	session1, _ := session.Store.Get(r, "session")

	fmt.Println(users1)
	if users1 == nil {
		fmt.Println("incorrect user name and password")
		session.Message("Invalid Username or Password", "danger", r, w)
		http.Redirect(w, r, "/", 302)
	} else {
		session1.Values["USERID"] = users.ID
		fmt.Println("thise is the login user id")
		fmt.Println(session1.Values["USERID"])
		//id:=session1.Values["USERID"]

		role := users1.RoleId
		_ = session1.Save(r, w)
		//role=4
		if role == 4 {
			fmt.Print("Admin")
		} else if role == 2 {
			fmt.Print("doctor")
			http.Redirect(w, r, "/doctor", 302)
		} else if role == 3 {
			fmt.Print("Pharmacist")
			http.Redirect(w, r, "/patients", 302)
		} else if role == 1 {
			fmt.Print("Laboratorist")
			http.Redirect(w, r, "/doctor/appointment", 302)
		} else if role == 5 {
			fmt.Print("Petient")
			http.Redirect(w, r, "/laboratorsit", 302)
		}
	}

	//mh.tmpl.ExecuteTemplate(w, "Doctor.patient.html", petient)

}
func (uh *loginHandler) LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := session.Flash(r, w)
	if r.Method == http.MethodGet {

		fmt.Println("Get Login page")
		_ = uh.tmpl.ExecuteTemplate(w, "Login.html", struct {
			Alert utils.Alert
		}{
			Alert: utils.NewAlert(message, alert),
		})
	} else {
		uh.Users(w, r)
	}

}
