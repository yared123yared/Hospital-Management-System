package main

import (
	"github.com/web1_group_project/hospital_client/delivery/http/handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Doctor_Handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Patient_Handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/pharmacist_handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Admin_hanlder"
	"html/template"
	"net/http"
)

var tmpl_doctor = template.Must(template.ParseGlob("../ui/template/Doctor/*.html"))
var tmpl_patient = template.Must(template.ParseGlob("../ui/template/petient/*.html"))

//var tmpl_admin = template.Must(template.ParseGlob("../ui/template/Admin/*.html"))
var tmpl_pharmacist = template.Must(template.ParseGlob("../ui/template/pharmacist/*.html"))

//var tmpl_laboratorsit = template.Must(template.ParseGlob("../ui/template/laboratorist/*.html"))
var tmpl = template.Must(template.ParseGlob("../ui/template/*.html"))
var templ_admin = template.Must(template.ParseGlob("../ui/template/Admin/*.html"))

func main() {
	//doctor handlers
	doctorPatientHandler := Doctor_Handler.NewpatientHandler(tmpl_doctor)
	doctorAppointmentHandler := Doctor_Handler.NewappointmentHandler(tmpl_doctor)
	doctorPrescribtionHandler := Doctor_Handler.NewprescribtionHandler(tmpl_doctor)
	doctorDiagonosisHandler := Doctor_Handler.NewdiagonosisHandler(tmpl_doctor)
	//loginhandler
	loginHandler := handler.NewLoginHandler(tmpl)
	//patient handler
	patientHandler := Patient_Handler.NewPatientHandler(tmpl_patient)
	//pharmacistHandler
	pharmacisstHandler := pharmacist_handler.NewPharmTempHandler(tmpl_pharmacist)
	//aborHandler := lbrhdlr.NewLaborTempHandler(temple2)
	//adminHandler := adminhdlr.NewAdminTempHandler(temple3)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// login page path registeration

	mux.HandleFunc("/", loginHandler.LoginGetHandler)

	// patient handler path registeration
	mux.HandleFunc("/appointment", patientHandler.Appointment)
	mux.HandleFunc("/profile", patientHandler.Profile)
	mux.HandleFunc("/doctors", patientHandler.Doctors)
	mux.HandleFunc("/prescription", patientHandler.Prescription)
	mux.HandleFunc("/request", patientHandler.Request)
	mux.HandleFunc("/request/new", patientHandler.SendRequest)
	mux.HandleFunc("/profile/update", patientHandler.Update)

	//doctor handler path registeration

	//Doctor.patient registeration pathes
	mux.HandleFunc("/doctor", doctorPatientHandler.Index)
	mux.HandleFunc("/doctor/patients", doctorPatientHandler.Patients)
	mux.HandleFunc("/doctor/patientUpdate", doctorPatientHandler.UpdatePatient)
	mux.HandleFunc("/doctor/patientDelete", doctorPatientHandler.DeletePatient)
	mux.HandleFunc("/doctor/patientNew", doctorPatientHandler.AddNewPatient)

	//Doctor appointment path regiseration
	mux.HandleFunc("/doctor/appointment", doctorAppointmentHandler.Appointment)
	mux.HandleFunc("/doctor/appointmentNew", doctorAppointmentHandler.AddNewAppointment)

	// Doctor prescribtion path registeration
	mux.HandleFunc("/doctor/prescribtion", doctorPrescribtionHandler.Prescribtions)
	mux.HandleFunc("/doctor/prescribtionNew", doctorPrescribtionHandler.AddNewPrescribtions)

	// doctor diagonosis path registeration
	mux.HandleFunc("/doctor/diagonosis", doctorDiagonosisHandler.Diagonosises)
	mux.HandleFunc("/doctor/diagonosisNew", doctorDiagonosisHandler.AddNewDiagonosis)

	//pharmacist path registeration
	mux.HandleFunc("/pharmacist", pharmacisstHandler.Index)
	mux.HandleFunc("/cat", pharmacisstHandler.CatHandler)
	mux.HandleFunc("/prof", pharmacisstHandler.ProHandler)
	mux.HandleFunc("/addcat", pharmacisstHandler.AddNewCat)
	mux.HandleFunc("/updateCat", pharmacisstHandler.UpdateCat)
	mux.HandleFunc("/deleteCat", pharmacisstHandler.DleteMedicine)
	mux.HandleFunc("/updateProv", pharmacisstHandler.UpdateProv)
	mux.HandleFunc("/pharmProf/update", pharmacisstHandler.PharmProfileUpdate)
	mux.HandleFunc("/prescription1", pharmacisstHandler.Prescription)
	mux.HandleFunc("/updatePres", pharmacisstHandler.PrescriptionUpdate)
	mux.HandleFunc("/deletePres", pharmacisstHandler.DeletePrescription)
	mux.HandleFunc("/dashboard", pharmacisstHandler.Dashboard)

	//mux.HandleFunc("/dashboardLabor", laborHandler.LaborDashHandler)
	//mux.HandleFunc("/diagnosisLabor", laborHandler.LaborDiagnosisHandler)
	//mux.HandleFunc("/profileLabor", laborHandler.LaborProfileHandler)
	//mux.HandleFunc("/laborProf/update", laborHandler.LaborProfileUpdateHandler)
	//mux.HandleFunc("/labor/updateDiag", laborHandler.LaborDiagnosisUpdateHandler)
	//Admin
	//http.HandleFunc("/adminDoctors", adminHandler.DoctorTempHandler)
	//http.HandleFunc("/admin/addNewDoctor", adminHandler.AddDoctorTempHandler)

/*888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888*/
/*88888888888888888888888888888888 PETIENTT HANDLER  8888888888888888888888888888888888888888888888888888888888888888888888888888888888888888*/
/*88888888888888888888888888888888 PATIENT HANDLER   888888888888888888888888888888888888888888888888888888888888888888888888888888*/
/*888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888*/
	adminHandler := Admin_hanlder.NewAdminTempHandler(templ_admin)

	//Admin doctors
	http.HandleFunc("/adminDoctors", adminHandler.DoctorTempHandler)
	http.HandleFunc("/admin/addNewDoctor", adminHandler.AddDoctorTempHandler)
	http.HandleFunc("/adminDocUpdate", adminHandler.UpdateDoctorTempHandler)
	http.HandleFunc("/adminDocDelete", adminHandler.DeleteDoctorTempHandler)

	//Admin pharmacists
	http.HandleFunc("/adminPharmacists", adminHandler.PharmacistTempHandler)
	http.HandleFunc("/admin/addNewPharmacist", adminHandler.PharmacistNewTempHandler)
	http.HandleFunc("/adminPharmUpdate", adminHandler.UpdatePharmacistTempHandler)
	http.HandleFunc("/adminpharmDelete", adminHandler.DeletePharmacistTempHandler)
	//admin Laboratorist
	http.HandleFunc("/adminLaboratorists", adminHandler.LaboratoristTempHandler)
	http.HandleFunc("/admin/addNewLaboratorist", adminHandler.LaboratoristNewTempHandler)
	http.HandleFunc("/adminLabormUpdate", adminHandler.UpdateLLaboratoristTempHandler)
	http.HandleFunc("/adminLaborDelete", adminHandler.DeleteLaboratoristTempHandler)
	tmpl := template.Must(template.ParseGlob("C:/Users/Gech/go/src/github.com/getach1/web1/web1_group_project_old_new/hospital_client/delivery/ui/template/petient/*"))
	patientHandler = Patient_Handler.NewPatientHandler(tmpl)
	fsn := http.FileServer(http.Dir("C:/Users/Gech/go/src/github.com/getach1/web1/web1_group_project_old_new/hospital_client/delivery/ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fsn))
	mux.HandleFunc("/", patientHandler.Appointment)
	mux.HandleFunc("/profile", patientHandler.Profile)
	mux.HandleFunc("/doctors", patientHandler.Doctors)
	mux.HandleFunc("/prescription", patientHandler.Prescription)
	mux.HandleFunc("/request", patientHandler.Request)
	mux.HandleFunc("/request_new", patientHandler.NewRequest)
	mux.HandleFunc("/request/new", patientHandler.SendRequest)
	mux.HandleFunc("/profile/update", patientHandler.Update)


	http.ListenAndServe(":2241", mux)

	_ = http.ListenAndServe(":8000", nil)
}
