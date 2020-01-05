package main

import (
	"html/template"
	"net/http"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Petient struct {
	ID        int
	FirstName string
	LastName  string
	Image     string
	Address   string
	Email     string
	Phone     string
	BirthDate time.Time
}

func doctors(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	err := tmpl.ExecuteTemplate(w, "patient.view.doctor", nil)
	checkErr(err)
}
func appointment(w http.ResponseWriter, _ *http.Request) {

	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	err := tmpl.ExecuteTemplate(w, "patient.view.appointment", nil)
	checkErr(err)
}
func prescription(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	err := tmpl.ExecuteTemplate(w, "patient.view.prescription", nil)
	checkErr(err)
}
func profile(w http.ResponseWriter, _ *http.Request) {
	petient := Petient{1, "Getachew", "Tebikew", "prescription.png", "Addis Ababa", "xy@z.com", "+1113444", time.Now()}
	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	err := tmpl.ExecuteTemplate(w, "patient.view.profile.update", petient)
	checkErr(err)
}
func request(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	err := tmpl.ExecuteTemplate(w, "patient.view.request", nil)
	checkErr(err)
}

func update(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	err := tmpl.ExecuteTemplate(w, "patient.view.profile.update", nil)
	checkErr(err)
}

func main() {
	fs := http.FileServer(http.Dir("../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", appointment)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/doctors", doctors)
	http.HandleFunc("/prescription", prescription)
	http.HandleFunc("/request", request)
	//http.HandleFunc("/update",update)
	_ = http.ListenAndServe(":8000", nil)
}
