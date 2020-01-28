package main

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/web1_group_project/hospital_client/delivery/http/handler"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/rtoken"

	//	"github.com/web1_group_project/hospital_client/delivery/http/handler/laboratorist_handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Admin_hanlder"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Doctor_Handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Patient_Handler"
	"github.com/web1_group_project/hospital_client/delivery/http/handler/Pharmacist_handler"
)

var tmpl_doctor = template.Must(template.ParseGlob("../ui/template/Doctor/*.html"))

//var tmpl_patient = template.Must(template.ParseGlob("../ui/template/petient/*.html"))

//var tmpl_admin = template.Must(template.ParseGlob("../ui/template/Admin/*.html"))
var tmpl_pharmacist = template.Must(template.ParseGlob("../ui/template/pharmacist/*.html"))

//var tmpl_laboratorsit = template.Must(template.ParseGlob("../ui/template/laboratorist/*.html"))
var tmpl = template.Must(template.ParseGlob("../ui/template/*.html"))
var templ_admin = template.Must(template.ParseGlob("../ui/template/Admin/*.html"))

var temple = template.Must(template.ParseGlob("../ui/template/pharmacist/*.html"))
var temple2 = template.Must(template.ParseGlob("../ui/template/laboratorist/*.html"))

var tmpl_patient = template.Must(template.ParseGlob("../ui/template/petient/*"))

func main() {
	fs := http.FileServer(http.Dir("../ui/assets/"))
	http.Handle("/assets/", http.StripPrefix(strings.TrimRight("/assets/", "/"), fs))

	// <<<<<<<<<<<<<<<<<<<<<<<LOGIN>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	sess := configSess()
	//mux := http.NewServeMux()
	uh := handler.NewUserHandler(tmpl, sess, csrfSignKey)
	http.HandleFunc("/login", uh.Login)
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	//
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<DOCTORS>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//Doctor.patient path registeration
	doctorHandler := Doctor_Handler.NewpatientHandler(tmpl_doctor, uh, csrfSignKey)
	http.Handle("/doctor", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.Index))))
	http.Handle("/doctor/patients", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.Patients))))
	http.Handle("/doctor/patientUpdate", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.UpdatePatient))))
	http.HandleFunc("/doctor/patient/delete", doctorHandler.DeletePatient)
	http.Handle("/doctor/patientNew", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.AddNewPatient))))
	//
	//Doctor appointment path regiseration
	//doctorAppointmentHandler := Doctor_Handler.NewappointmentHandler(tmpl_doctor, uh, csrfSignKey)
	http.Handle("/doctor/appointment", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.Appointment))))
	http.Handle("pharmacy/doctor/appointmentNew", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.AddNewAppointment))))
	//
	// Doctor prescribtion path registeration
	http.Handle("/doctor/prescribtion", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.Prescribtions))))
	http.Handle("/doctor/prescribtionNew", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorHandler.AddNewPrescribtions))))
	//
	// doctor diagonosis path registeration
	doctorDiagonosisHandler := Doctor_Handler.NewdiagonosisHandler(tmpl_doctor)
	http.Handle("/doctor/diagonosis", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorDiagonosisHandler.Diagonosises))))
	http.Handle("/doctor/diagonosis/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(doctorDiagonosisHandler.AddNewDiagonosis))))
	//
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<PATIENT>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	patientHandler := Patient_Handler.NewPatientHandler(tmpl_patient, uh, csrfSignKey)
	http.Handle("/patient", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.Appointment))))
	http.Handle("/patient/profile", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.Profile))))
	http.Handle("/patient/doctors", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.Doctors))))
	http.Handle("/patient/prescription", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.Prescription))))
	http.Handle("/patient/request", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.Request))))
	http.Handle("/patient/request/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.NewRequest))))
	//mux.Handle("/patient/request/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.SendRequest))))
	http.Handle("/patient/profile/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(patientHandler.Update))))
	//
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<PHARMACIST>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//patientHandler := Patient_Handler.NewPatientHandler(tmpl_patient,uh,csrfSignKey)

	pharmacisthandler := Pharmacist_handler.NewPharmTempHandler(temple, uh, csrfSignKey)
	//Pharmacist Dashboard
	http.Handle("/pharmacist", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.Index))))
	http.Handle("/pharmacist/dashboard", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.Dashboard))))
	//
	//Pharmacist Medicine
	http.Handle("/pharmacist/medicine", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.CatHandler))))
	http.Handle("/pharmacist/medicine/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.AddNewCat))))
	http.Handle("/pharmacist/medicine/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.UpdateCat))))
	http.Handle("/pharmacist/medicine/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.DleteMedicine))))

	//Pharmacist profile
	http.Handle("/pharmacist/profile", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.ProHandler))))
	http.Handle("/pharmacist/profile/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.PharmProfileUpdate))))
	//Pharmacist prescription
	http.Handle("/pharmacist/prescription", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.Prescription))))
	http.Handle("/pharmacist/prescription/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.PrescriptionUpdate))))
	http.Handle("/pharmacist/prescription/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(pharmacisthandler.DeletePrescription))))
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<LABORATORIST>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//laboratoristHandler := laboratorist_handler.NewLaborTempHandler(temple2)
	//Laboratorist dashboard
	//http.HandleFunc("/laboratorist/dashboard", laboratoristHandler.LaborDashHandler)
	//
	//Laboratorist Diagnosis
	// http.HandleFunc("/laboratorist/diagnosis", laboratoristHandler.LaborDiagnosisHandler)
	// http.HandleFunc("/laboratorist/updateDiagnosis", laboratoristHandler.LaborDiagnosisUpdateHandler)
	//
	//Laboratorist Profile
	// http.HandleFunc("/laboratorist/profile", laboratoristHandler.LaborProfileHandler)
	// http.HandleFunc("/laboratorist/profile/update", laboratoristHandler.LaborProfileUpdateHandler)
	//
	//
	adminHandler := Admin_hanlder.NewAdminTempHandler(templ_admin, uh, csrfSignKey)
	http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.AdminIndex))))
	http.Handle("/admin/doctors", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.DoctorTempHandler))))
	http.Handle("/admin/doctors/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.AddDoctorTempHandler))))

	http.Handle("/admin/pharmacists", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.PharmacistTempHandler))))
	http.Handle("/admin/pharmacists/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.PharmacistNewTempHandler))))
	// http.Handle("admin/pharmacists", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.PharmacistTempHandler))))

	http.Handle("/admin/laboratorists", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.LaboratoristTempHandler))))
	http.Handle("/admin/laboratorists/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(adminHandler.LaboratoristNewTempHandler))))

	http.ListenAndServe(":8287", nil)

}
func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}

	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
