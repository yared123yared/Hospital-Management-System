package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/web1_group_project/hospital_server/Doctor/repository"
	"github.com/web1_group_project/hospital_server/Doctor/service"
	laborrepo "github.com/web1_group_project/hospital_server/Laboratorist/repository"
	laborSrv "github.com/web1_group_project/hospital_server/Laboratorist/service"
	"github.com/web1_group_project/hospital_server/delivery/http/handler"
	"github.com/web1_group_project/hospital_server/delivery/http/handler/Doctor_Handler"
	"github.com/web1_group_project/hospital_server/delivery/http/handler/Patient_Handler"
	PetientRepository "github.com/web1_group_project/hospital_server/petient/repository"
	PetientService "github.com/web1_group_project/hospital_server/petient/service"

	PharmacistRepo "github.com/web1_group_project/hospital_server/Pharmacist/repository"
	PharmacistService "github.com/web1_group_project/hospital_server/Pharmacist/service"

	"net/http"

	laborHandler "github.com/web1_group_project/hospital_server/delivery/http/handler/labor_handler"
	handler "github.com/web1_group_project/hospital_server/delivery/http/handler/pharm_handler"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital12?sslmode=disable")

	if err != nil {
		panic(err)
	}
	fmt.Println("hi")
	petientRepo := PetientRepository.NewPetientGormRepo(dbconn)
	petientServ := PetientService.NewPetientService(petientRepo)
	PetientHandler := Patient_Handler.NewAdminPetientHandler(petientServ)
	router := httprouter.New()
	router.GET("/v1/admin/petients", PetientHandler.GetPetients)
	router.GET("/v1/admin/petients/:id", PetientHandler.GetSinglePetient)
	router.PUT("/v1/admin/petients/:id", PetientHandler.PutPetient)
	router.POST("/v1/admin/petients", PetientHandler.PostPetient)
	router.DELETE("/v1/admin/petients/:id", PetientHandler.DeletePetient)

	petientRequestRepo := PetientRepository.NewRequestGormRepo(dbconn)
	petientRequestServ := PetientService.NewRequestService(petientRequestRepo)
	PetientRequestHandler := Patient_Handler.NewPetientRequestHandler(petientRequestServ)
	router.GET("/v1/patient/requests", PetientRequestHandler.GetRequests)
	router.GET("/v1/patient/requests/:id", PetientRequestHandler.GetSingleRequest)
	router.POST("/v1/patient/requests", PetientRequestHandler.PostRequest)

	petientAppointmentRepo := PetientRepository.NewAppointmentGormRepo(dbconn)
	petientAppointmentServ := PetientService.NewAppointmentService(petientAppointmentRepo)
	PetientAppointmentHandler := Patient_Handler.NewPetientAppointmentHandler(petientAppointmentServ)
	router.GET("/v1/patient/appointments", PetientAppointmentHandler.GetAppointments)
	router.GET("/v1/patient/appointments/:id", PetientAppointmentHandler.GetSingleAppointment)

	petientPrescriptionRepo := PetientRepository.NewPrescriptionGormRepo(dbconn)
	petientPrescriptionServ := PetientService.NewPrescriptionService(petientPrescriptionRepo)
	PetientPrescriptionHandler := Patient_Handler.NewPetientPrescriptionHandler(petientPrescriptionServ)
	router.GET("/v1/patient/prescriptions", PetientPrescriptionHandler.GetSinglePrescription)
	router.GET("/v1/patient/prescriptions/:id", PetientPrescriptionHandler.GetPrescriptions)

	petientDoctorRepo := PetientRepository.NewDoctorGormRepo(dbconn)
	petientDoctorServ := PetientService.NewDoctorService(petientDoctorRepo)
	PetientDoctorHandler := Patient_Handler.NewPetientDoctorHandler(petientDoctorServ)
	router.GET("/v1/patient/doctors", PetientDoctorHandler.GetDoctors)
	router.GET("/v1/patient/doctors/:id", PetientDoctorHandler.GetSingleDoctor)

	petientAdminRepo := PetientRepository.NewAdminGormRepo(dbconn)
	petientAdminServ := PetientService.NewAdminService(petientAdminRepo)
	PetientAdminHandler := Patient_Handler.NewPatientAdminHandler(petientAdminServ)
	router.GET("/v1/patient/admins/:id", PetientAdminHandler.GetAdmins)
	http.ListenAndServe(":8100", router)

	patientRepo := repository.NewPatientGormRepo(dbconn)
	patientSrv := service.NewPatientService(patientRepo)
	doctorPatientHandler := Doctor_Handler.NewDoctorPatientHandler(patientSrv)
	// thise is the general doctor information
	appointmentRepo := repository.NewAppointmentGormRepo(dbconn)
	appointmentSrv := service.NewAppointmentService(appointmentRepo)
	doctorAppointmentHandler := Doctor_Handler.NewDoctorAppointmentHandler(appointmentSrv)
	//
	generalRepo := repository.NewGeneralGormRepo(dbconn)
	generalSrv := service.NewGeneralService(generalRepo)
	generalHandler := handler.NewGeneralHandler(generalSrv)

	router.GET("/v1/admin/users/:id", doctorPatientHandler.GetSinglePatient)
	router.GET("/v1/admin/users/", doctorPatientHandler.GetPatients)
	router.PUT("/v1/admin/users/:id", doctorPatientHandler.PutPatient)
	router.POST("/v1/admin/users/", doctorPatientHandler.PostPatient)
	router.DELETE("/v1/admin/users/:id", doctorPatientHandler.DeletePatient)
	//
	router.GET("/v1/doctor/appointments/:id", doctorAppointmentHandler.GetSingleAppointment)
	router.GET("/v1/doctor/appointments/", doctorAppointmentHandler.GetAppointments)
	router.PUT("/v1/doctor/appointments/:id", doctorAppointmentHandler.PutAppointment)
	router.DELETE("/v1/doctor/appointments/:id", doctorAppointmentHandler.DeleteAppointment)
	//
	router.GET("/v1/doctor/pharmacists/", generalHandler.GetPharmacists)
	router.GET("/v1/doctor/laboratorists/", generalHandler.GetLaboratorists)

	medRepo := PharmacistRepo.NewMedicineGormRepo(dbconn)
	medSrv := PharmacistService.NewMedicineService(medRepo)

	medHandler := handler.NewPharmMedicineHandler(medSrv)

	router := httprouter.New()
	router.GET("/v1/pharm/medicines/:id", medHandler.GetSingleMedicine)
	router.GET("/v1/pharm/medMulti/:addedby", medHandler.GetMultipleMedicines)
	router.GET("/v1/pharm/medicines", medHandler.GetMedicines)
	router.PUT("/v1/pharm/medicines/:id", medHandler.PutMedicine)
	router.POST("/v1/pharm/medicines", medHandler.PostMedicine)
	router.DELETE("/v1/pharm/medicines/:id", medHandler.DeleteMedicine)

	prescRepo := PharmacistRepo.NewPrescriptionGormRepo(dbconn)
	prescSRV := PharmacistService.NewPrescriptionService(prescRepo)

	prescHandler := handler.NewPharmPrescriptionHandler(prescSRV)

	router.GET("/v1/pharm/presc/:id", prescHandler.GetSinglePrescription)
	router.GET("/v1/pharm/presc", prescHandler.GetMultiPrescriptions)
	router.PUT("/v1/pharm/presc/:id", prescHandler.PutPrescription)
	router.POST("/v1/pharm/presc", prescHandler.PostPrescription)
	router.GET("/v1/pharm/multiPre", prescHandler.GetMultiPrescriptions)
	router.DELETE("/v1/pharm/presc/:id", prescHandler.DeletePrescription)

	profRepo := PharmacistRepo.NewPharmacistProfileGormRepo(dbconn)
	profSRV := PharmacistService.NewPharmacistProfileService(profRepo)

	profHandler := handler.NewPharmProfileHandler(profSRV)

	router.GET("/v1/pharm/profile/:id", profHandler.GetSingleProfile)
	router.GET("/v1/pharm/profile", profHandler.GetProfiles)
	router.PUT("/v1/pharm/profile/:id", profHandler.PutProfile)
	router.POST("/v1/pharm/profile", profHandler.PostProfile)
	router.DELETE("/v1/pharm/profile/:id", profHandler.DeleteProfile)

	labProfRepo := laborrepo.NewLabratoristProfileGormRepo(dbconn)
	labProfSRV := laborSrv.NewLabratoristProfileService(labProfRepo)

	laborProdHandler := laborHandler.NewLaborProfileHandler(labProfSRV)

	router.GET("/v1/labor/profile/:id", laborProdHandler.GetSingleProfile)
	router.GET("/v1/labor/profile", laborProdHandler.GetProfiles)
	router.PUT("/v1/labor/profile/:id", laborProdHandler.PutProfile)
	router.POST("/v1/labor/profile", laborProdHandler.PostProfile)
	router.DELETE("/v1/labor/profile/:id", laborProdHandler.DeleteProfile)

	labDiagRepo := laborrepo.NewDiagnosisGormRepo(dbconn)
	labdiagsrv := laborSrv.NewDiagnosisService(labDiagRepo)
	laborDiagHandler := laborHandler.NewLaborDiagnosisHandler(labdiagsrv)

	router.GET("/v1/labor/diag/:id", laborDiagHandler.GetSingleDiagnosis)
	router.GET("/v1/labor/diag", laborDiagHandler.GetDiagnosiss)
	router.GET("/v1/labor/multiDiag", laborDiagHandler.GetMultiDiagnosis)
	router.PUT("/v1/labor/diag/:id", laborDiagHandler.PutDiagnosis)
	router.POST("/v1/labor/diag", laborDiagHandler.PostDiagnosis)
	router.DELETE("/v1/labor/diag/:id", laborDiagHandler.DeleteDiagnosis)

	http.ListenAndServe(":8480", router)

}
