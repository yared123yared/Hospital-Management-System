package main

import (
	aptrepo "github.com/getach1/web1/hospital/Appointment/repository"
	aptserv "github.com/getach1/web1/hospital/Appointment/service"
	"github.com/getach1/web1/hospital/delivery/http/handler"
	prescsrepo "github.com/getach1/web1/hospital/prescribtion/repository"
	prescserv "github.com/getach1/web1/hospital/prescribtion/service"
	reqRepo "github.com/getach1/web1/hospital/request/repository"
	reqServ "github.com/getach1/web1/hospital/request/service"
	peRepo "github.com/getach1/web1/web1_group_project/hospital_server/petient/repository"
	peServ "github.com/getach1/web1/web1_group_project/hospital_server/petient/service"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	//dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital4?sslmode=disable")
	dbconn, err := gorm.Open("postgres", "postgres://postgres:gechman@localhost/hospital_2?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	//CREATE TABLES
/*
	errs:=dbconn.CreateTable(&entity.Profile{}).GetErrors()
	errs=dbconn.CreateTable(&entity.Pharmacist{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Petient{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Admin{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Doctor{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Laboratorist{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Prescription{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("phrmacist_Id","pharmacists(Id)","cascade","cascade").AddForeignKey("patient_name","profiles(full_name)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Appointment{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("patient_name","profiles(full_name)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Diagnosis{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("laboratorist_Id","laboratorists(Id)","cascade","cascade").AddForeignKey("patient_name","profiles(full_name)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Medicine{}).AddForeignKey("added_By","pharmacists(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Request{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("approved_By","admins(Id)","cascade","cascade").AddForeignKey("patient_name","profiles(full_name)","cascade","cascade").GetErrors()
	if errs!=nil {
		panic(errs)

	}*/



//INSERT TO TABLES
/*
petient:=entity.Petient{
	ID:           0,
	Uuid:         1,
	Profile:      entity.Profile{},
	BloodGroup:   "AB",
	Age:          30,
	Prescription: nil,
	Diagnosis:    nil,
	Appointment:  nil,
	Request:      nil,
}
dbconn.Debug().Save(&petient)
*/
/*
profile:=entity.Profile{
	ID:          0,
	FullName:    "abebe hello",
	Password:    "12345FSDFDJ",
	Email:       "GETACHEW@G.COM",
	Phone:       "12343jkdskj439434",
	Address:     "adama ABABA",
	Image:       "doctor.png",
	Sex:         "MALE",

	Role:        "PETIENT",
	BirthDate:   time.Time{},
	Description: "INPETIENT IN THIS HOSPITAL",
}

dbconn.Debug().Save(&profile)
*/

/*doctor:=entity.Doctor{
	ID:           0,
	Profile:      entity.Profile{},
	Uuid:         3,
	Department:   "surgery",
	Prescription: nil,
	Diagnosis:    nil,
	Appointment:  nil,
}
parma:=entity.Pharmacist{
	ID:           0,
	Uuid:         1,
	Profile:      entity.Profile{},
	Medicine:     nil,
	Prescription: nil,
}
lab:=entity.Laboratorist{
	ID:        0,
	Uuid:      1,
	Profile:   entity.Profile{},
	Diagnosis: nil,
}
dbconn.Save(&doctor)
dbconn.Save(&parma)
dbconn.Save(&lab)

*/
	petientRepo := peRepo.NewPetientGormRepo(dbconn)
	petientServ := peServ.NewPetientService(petientRepo)

	adminPetientHandler := handler.NewAdminPetientHandler(petientServ)
	router := httprouter.New()
	router.GET("/v1/admin/petients", adminPetientHandler.GetPetients)
	router.GET("/v1/admin/petients/:id", adminPetientHandler.GetSinglePetient)

	router.PUT("/v1/admin/petients/:id", adminPetientHandler.PutPetient)
	router.POST("/v1/admin/petients", adminPetientHandler.PostPetient)
	router.DELETE("/v1/admin/petients/:id", adminPetientHandler.DeletePetient)

	requestRepo := reqRepo.NewRequestGormRepo(dbconn)
	requestServ := reqServ.NewRequestService(requestRepo)

	adminRequestHandler := handler.NewAdminRequestHandler(requestServ)
	router.GET("/v1/admin/requests", adminRequestHandler.GetRequests)
	router.GET("/v1/admin/requests/:id", adminRequestHandler.GetSingleRequest)
	router.PUT("/v1/admin/requests/:id", adminRequestHandler.PutRequest)
	router.POST("/v1/admin/requests", adminRequestHandler.PostRequest)
	router.DELETE("/v1/admin/requests/:id", adminRequestHandler.DeleteRequest)

	appointmentRepo := aptrepo.NewAppointmentGormRepo(dbconn)
	appointmentServ := aptserv.NewAppointmentService(appointmentRepo)

	adminAppointmentHandler := handler.NewAdminAppointmentHandler(appointmentServ)
	router.GET("/v1/admin/appointments", adminAppointmentHandler.GetAppointments)
	router.GET("/v1/admin/appointments/:id", adminAppointmentHandler.GetSingleAppointment)
	router.PUT("/v1/admin/appointments/:id", adminAppointmentHandler.PutAppointment)
	router.POST("/v1/admin/appointments", adminAppointmentHandler.PostAppointment)
	router.DELETE("/v1/admin/appointments/:id", adminAppointmentHandler.DeleteAppointment)

	prescriptionRepo := prescsrepo.NewPrescriptionGormRepo(dbconn)
	prescriptionServ := prescserv.NewPrescriptionService(prescriptionRepo)

	adminPrescriptiontHandler := handler.NewAdminPrescriptionHandler(prescriptionServ)
	router.GET("/v1/admin/prescriptions", adminPrescriptiontHandler.GetPrescriptions)
	router.GET("/v1/admin/prescriptions/:id", adminPrescriptiontHandler.GetSinglePrescription)
	router.PUT("/v1/admin/prescriptions/:id", adminPrescriptiontHandler.PutPrescription)
	router.POST("/v1/admin/prescriptions", adminPrescriptiontHandler.PostPrescription)
	router.DELETE("/v1/admin/prescriptions/:id", adminPrescriptiontHandler.DeletePrescription)

	http.ListenAndServe(":8100", router)
	/*
		patientRepo := repository.NewPatientGormRepo(dbconn)
		patientSrv := service.NewPatientService(patientRepo)
		doctorPatientHandler := handler.NewDoctorPatientHandler(patientSrv)
		// thise is the general doctor information
		appointmentRepo := repository.NewAppointmentGormRepo(dbconn)
		appointmentSrv := service.NewAppointmentService(appointmentRepo)
		doctorAppointmentHandler := handler.NewDoctorAppointmentHandler(appointmentSrv)

		router := httprouter.New()

		router.GET("/v1/admin/users/", doctorPatientHandler.GetSinglePatient)
		router.GET("/v1/admin/users", doctorPatientHandler.GetPatients)
		router.PUT("/v3/admin/users/:id", doctorPatientHandler.PutPatient)
		router.POST("/v1/admin/users", doctorPatientHandler.PostPatient)
		router.DELETE("/v2/admin/users/:id", doctorPatientHandler.DeletePatient)
		//
		router.GET("/v1/doctor/appointments/:id ", doctorAppointmentHandler.GetSingleAppointment)
		router.GET("/v1/doctor/appointments", doctorAppointmentHandler.GetAppointments)
		router.PUT("/v1/doctor/appointments/:id", doctorAppointmentHandler.PutAppointment)
		router.DELETE("/v1/doctor/appointments/:id", doctorAppointmentHandler.DeleteAppointment)

		http.ListenAndServe(":8480", router)*/
}
