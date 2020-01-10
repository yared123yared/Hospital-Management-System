package main

import (
	"github.com/getach1/web1/web1_group_project/hospital_client/delivery/http/handler"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseGlob("C:/Users/Gech/go/src/github.com/getach1/web1/web1_group_project/hospital_client/delivery/ui/template/petient/*"))
	patientHandler := handler.NewPatientHandler(tmpl)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", patientHandler.Appointment)
	http.HandleFunc("/profile", patientHandler.Profile)
	http.HandleFunc("/doctors", patientHandler.Doctors)
	http.HandleFunc("/prescription", patientHandler.Prescription)
	http.HandleFunc("/request", patientHandler.Request)
	http.HandleFunc("/request/new", patientHandler.SendRequest)
	http.HandleFunc("/profile/update", patientHandler.Update)
	_ = http.ListenAndServe(":8000", nil)
}
