package main

import (
	"github.com/getach1/web1/Project/entities/handler"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handler.Appointment)
	http.HandleFunc("/profile", handler.Profile)
	http.HandleFunc("/doctors", handler.Doctors)
	http.HandleFunc("/prescription", handler.Prescription)
	http.HandleFunc("/request", handler.Request)
	http.HandleFunc("/request/new", handler.SendRequest)
	http.HandleFunc("/profile/update", handler.Update)
	_ = http.ListenAndServe(":8000", nil)
}
