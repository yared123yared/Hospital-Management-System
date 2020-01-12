package main

import (
	"github.com/getach1/web1/web1_group_project-master/hospital_server/delivery/http/handler"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	PetientRepository "github.com/web1_group_project/hospital_server/petient/repository"
	PetientService "github.com/web1_group_project/hospital_server/petient/service"
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


	admin:=entity.Admin{
		ID:      0,
		Uuid:    1,
		Profile: entity.Profile{},
		Request: nil,
	}
	dbconn.Save(&admin)*/

	petientRepo := PetientRepository.NewPetientGormRepo(dbconn)
	petientServ := PetientService.NewPetientService(petientRepo)
	PetientHandler := handler.NewAdminPetientHandler(petientServ)
	router := httprouter.New()
	router.GET("/v1/admin/petients", PetientHandler.GetPetients)
	router.GET("/v1/admin/petients/:id", PetientHandler.GetSinglePetient)
	router.PUT("/v1/admin/petients/:id", PetientHandler.PutPetient)
	router.POST("/v1/admin/petients", PetientHandler.PostPetient)
	router.DELETE("/v1/admin/petients/:id", PetientHandler.DeletePetient)

	petientRequestRepo := PetientRepository.NewRequestGormRepo(dbconn)
	petientRequestServ := PetientService.NewRequestService(petientRequestRepo)
	PetientRequestHandler := handler.NewPetientRequestHandler(petientRequestServ)
	router.GET("/v1/patient/requests", PetientRequestHandler.GetRequests)
	router.GET("/v1/patient/requests/:id", PetientRequestHandler.GetSingleRequest)
	router.POST("/v1/patient/requests", PetientRequestHandler.PostRequest)

	petientAppointmentRepo := PetientRepository.NewAppointmentGormRepo(dbconn)
	petientAppointmentServ := PetientService.NewAppointmentService(petientAppointmentRepo)
	PetientAppointmentHandler := handler.NewPetientAppointmentHandler(petientAppointmentServ)
	router.GET("/v1/patient/appointments", PetientAppointmentHandler.GetAppointments)
	router.GET("/v1/patient/appointments/:id", PetientAppointmentHandler.GetSingleAppointment)

	petientPrescriptionRepo := PetientRepository.NewPrescriptionGormRepo(dbconn)
	petientPrescriptionServ := PetientService.NewPrescriptionService(petientPrescriptionRepo)
	PetientPrescriptionHandler := handler.NewPetientPrescriptionHandler(petientPrescriptionServ)
	router.GET("/v1/patient/prescriptions", PetientPrescriptionHandler.GetSinglePrescription)
	router.GET("/v1/patient/prescriptions/:id", PetientPrescriptionHandler.GetPrescriptions)


	petientDoctorRepo := PetientRepository.NewDoctorGormRepo(dbconn)
	petientDoctorServ := PetientService.NewDoctorService(petientDoctorRepo)
	PetientDoctorHandler := handler.NewPetientDoctorHandler(petientDoctorServ)
	router.GET("/v1/patient/doctors", PetientDoctorHandler.GetDoctors)
	router.GET("/v1/patient/doctors/:id", PetientDoctorHandler.GetSingleDoctor)


	petientAdminRepo := PetientRepository.NewAdminGormRepo(dbconn)
	petientAdminServ := PetientService.NewAdminService(petientAdminRepo)
	PetientAdminHandler := handler.NewPetientAdminHandler(petientAdminServ)
	router.GET("/v1/patient/admins/:id", PetientAdminHandler.GetAdmins)
	http.ListenAndServe(":8100", router)
	/*
			patientRepo := repository.NewPatientGormRepo(dbconn)
			patientSrv := service.NewPatientService(patientRepo)
			doctorPatientHandler := handler.NewDoctorPatientHandler(patientSrv)
			// thise is the general doctor information
			appointmentRepo := repository.NewAppointmentGormRepo(dbconn)
			appointmentSrv := service.NewAppointmentService(appointmentRepo)
			doctorAppointmentHandler := handler.NewDoctorAppointmentHandler(appointmentSrv)
		//
			generalRepo := repository.NewGeneralGormRepo(dbconn)
			generalSrv := service.NewGeneralService(generalRepo)
			generalHandler := handler.NewGeneralHandler(generalSrv)

			router := httprouter.New()

			router.GET	(	"/v1/admin/users/:id", doctorPatientHandler.GetSinglePatient)
			router.GET	(	"/v1/admin/users/", doctorPatientHandler.GetPatients)
			router.PUT	(	"/v1/admin/users/:id", doctorPatientHandler.PutPatient)
			router.POST	(	"/v1/admin/users/", doctorPatientHandler.PostPatient)
			router.DELETE(	"/v1/admin/users/:id", doctorPatientHandler.DeletePatient)
			//
			router.GET("/v1/doctor/appointments/:id", doctorAppointmentHandler.GetSingleAppointment)
			router.GET("/v1/doctor/appointments/", doctorAppointmentHandler.GetAppointments)
			router.PUT("/v1/doctor/appointments/:id", doctorAppointmentHandler.PutAppointment)
			router.DELETE("/v1/doctor/appointments/:id", doctorAppointmentHandler.DeleteAppointment)
			//
			router.GET("/v1/doctor/pharmacists/", generalHandler.GetPharmacists)
			router.GET("/v1/doctor/laboratorists/", generalHandler.GetLaboratorists)

			http.ListenAndServe(":8480", router)*/
}
