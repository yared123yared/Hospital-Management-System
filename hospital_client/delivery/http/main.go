package main

import "net/http"

/*
import (
	"github.com/getach1/web1/web1_group_project/hospital_client/delivery/http/handler"
	"html/template"
	"net/http"
)
*/
/*
import (

	"github.com/yaredsolomon/webProgram1/hospital2/delivery/http/handler"

	"net/http"

	"html/template"

	//"github.com/betsegawlemma/restclient/data"

	//"github.com/yaredsolomon/webProgram1/go-rest-client-master/data"
)
*/
 */

func main() {
	/*
	tmpl := template.Must(template.ParseGlob("../ui/template/petient/*"))
	patientHandler:=handler.NewPatientHandler(tmpl)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", patientHandler.Appointment)
	http.HandleFunc("/profile", patientHandler.Profile)
	http.HandleFunc("/doctors", patientHandler.Doctors)
	http.HandleFunc("/prescription", patientHandler.Prescription)
	http.HandleFunc("/request", patientHandler.Request)
	http.HandleFunc("/request/new", patientHandler.SendRequest)
	http.HandleFunc("/profile/update", patientHandler.Update)
	//
	doctorPatientHandler := handler.NewpatientHandler(tmpl)
	doctorAppointmentHandler := handler.NewappointmentHandler(tmpl)
	doctorPrescribtionHandler := handler.NewprescribtionHandler(tmpl)
	doctorDiagonosisHandler := handler.NewdiagonosisHandler(tmpl)

	 */

/*
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/doctor", doctorPatientHandler.Index)
	mux.HandleFunc("/doctor/patients", doctorPatientHandler.Patients)
	mux.HandleFunc("/doctor/patientUpdate", doctorPatientHandler.UpdatePatient)
	mux.HandleFunc("/doctor/patientDelete", doctorPatientHandler.DeletePatient)
	mux.HandleFunc("/doctor/patientNew", doctorPatientHandler.AddNewPatient)

	mux.HandleFunc("/doctor/appointment", doctorAppointmentHandler.Appointment)
	mux.HandleFunc("/doctor/appointmentNew", doctorAppointmentHandler.AddNewAppointment)

	mux.HandleFunc("/doctor/prescribtion", doctorPrescribtionHandler.Prescribtions)
	mux.HandleFunc("/doctor/prescribtionNew", doctorPrescribtionHandler.AddNewPrescribtions)

	mux.HandleFunc("/doctor/diagonosis", doctorDiagonosisHandler.Diagonosises)
	mux.HandleFunc("/doctor/diagonosisNew", doctorDiagonosisHandler.AddNewDiagonosis)
	//	mux.HandleFunc("/doctor/appointmentNew", doctorAppointmentHandler.AddNewAppointment)
	//mux.HandleFunc("/doctor/patients", doctorPatientHandler.Patients)
	//mux.HandleFunc("/doctor/patientUpdate", doctorPatientHandler.UpdatePatient)
	//mux.HandleFunc("/doctor/patientDelete", doctorPatientHandler.DeletePatient)
	//mux.HandleFunc("/doctor/patientNew", doctorPatientHandler.AddNewPatient)


	http.ListenAndServe(":2241", mux)
	*/
	_ = http.ListenAndServe(":8000", nil)
}
