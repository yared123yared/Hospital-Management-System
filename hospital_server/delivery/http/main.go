package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"

	AdminRepository "github.com/web1_group_project/hospital_server/Admin/repository"
	AdminService "github.com/web1_group_project/hospital_server/Admin/service"
	"github.com/web1_group_project/hospital_server/Doctor/repository"
	"github.com/web1_group_project/hospital_server/Doctor/service"
	GeneralRepository "github.com/web1_group_project/hospital_server/General/repository"
	GeneralService "github.com/web1_group_project/hospital_server/General/service"
	laborrepo "github.com/web1_group_project/hospital_server/Laboratorist/repository"
	laborSrv "github.com/web1_group_project/hospital_server/Laboratorist/service"
	PharmacistRepo "github.com/web1_group_project/hospital_server/Pharmacist/repository"
	PharmacistService "github.com/web1_group_project/hospital_server/Pharmacist/service"
	Admin_Handler "github.com/web1_group_project/hospital_server/delivery/http/handler/Admin_Handler"
	"github.com/web1_group_project/hospital_server/delivery/http/handler/Doctor_Handler"
	General_Handler "github.com/web1_group_project/hospital_server/delivery/http/handler/General_Handler"
	"github.com/web1_group_project/hospital_server/delivery/http/handler/Patient_Handler"
	laborHandler "github.com/web1_group_project/hospital_server/delivery/http/handler/labor_handler"
	handler "github.com/web1_group_project/hospital_server/delivery/http/handler/pharm_handler"
	PetientRepository "github.com/web1_group_project/hospital_server/petient/repository"
	PetientService "github.com/web1_group_project/hospital_server/petient/service"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:user1@localhost/hospital14?sslmode=disable")

	if err != nil {
		panic(err)
	}
	router := httprouter.New()
	fmt.Println("hi")

	// thise is admin user path registeration
	adminUserRepo := AdminRepository.NewUserRepository(dbconn)
	adminUserServ := AdminService.NewUserService(adminUserRepo)
	adminUserHandler := Admin_Handler.NewUserHandler(adminUserServ)

	router.GET("/v1/admin/users/", adminUserHandler.Users)
	router.GET("/v1/admin/users/:id", adminUserHandler.User)
	router.PUT("/v1/admin/users/:id", adminUserHandler.UpdateUser)
	router.POST("/v1/admin/users/", adminUserHandler.StoreUser)
	router.DELETE("/v1/admin/users/:id", adminUserHandler.DeleteUser)
	router.GET("/v1/admin/usersByEmail/:email", adminUserHandler.UserByEmail)
	router.GET("/v1/admin/userRoles/", adminUserHandler.UserRoles)

	//

	// thise is general role path registeration
	roleRepo := GeneralRepository.NewRoleGormRepo(dbconn)
	roleServ := GeneralService.NewRoleService(roleRepo)
	roleHandler := General_Handler.NewRolesHandler(roleServ)

	router.GET("/v1/general/roles/", roleHandler.Roles)
	router.GET("/v1/general/roles/:id", roleHandler.Role)
	router.PUT("/v1/general/roles/:id", roleHandler.UpdateRole)
	router.POST("/v1/general/roles/", roleHandler.StoreRole)
	router.DELETE("/v1/general/roles/:id", roleHandler.DeleteRole)
	router.GET("/v1/general/rolesByName/:name", roleHandler.RoleByName)

	//
	// thise is general session path registeration
	sessionRepo := GeneralRepository.NewSessionGormRepo(dbconn)
	sessionServ := GeneralService.NewSessionService(sessionRepo)
	sessionHandler := General_Handler.NewSessionHandler(sessionServ)

	router.GET("/v1/general/sessions/:id", sessionHandler.Session)
	router.POST("/v1/general/sessions/", sessionHandler.StoreSession)
	router.DELETE("/v1/general/sessions/:id", sessionHandler.DeleteSession)
	//
	//

	//
	petientRepo := PetientRepository.NewPetientGormRepo(dbconn)
	petientServ := PetientService.NewPetientService(petientRepo)
	PetientHandler := Patient_Handler.NewAdminPetientHandler(petientServ)

	router.GET("/v1/admin/petients", PetientHandler.GetPetients)
	router.GET("/v1/admin/petients/:id", PetientHandler.GetSinglePetient)
	router.GET("/v1/admin/petients2/:id", PetientHandler.GetSinglePetient_uuid)
	router.PUT("/v1/admin/petients/:id", PetientHandler.PutPetient)
	router.POST("/v1/admin/petients", PetientHandler.PostPetient)
	router.DELETE("/v1/admin/petients/:id", PetientHandler.DeletePetient)

	//ADMIN DOCTORS
	adminDoctorRepo := AdminRepository.NewManageDoctorsRepository(dbconn)
	adminDoctorServ := AdminService.NewManageDoctorsService(adminDoctorRepo)
	adminDoctorHandler := Admin_Handler.NewManageDoctorsHandler(adminDoctorServ)

	router.GET("/v1/admin/doctors/:id", adminDoctorHandler.GetSingleDoctor)
	router.GET("/v1/admin/doctors", adminDoctorHandler.GetDoctors)
	router.PUT("/v1/admin/doctors/:id", adminDoctorHandler.UpdateDoctor)
	router.POST("/v1/admin/doctors", adminDoctorHandler.AddDoctor)
	router.DELETE("/v1/admin/doctors/:id", adminDoctorHandler.DeleteDoctor)

	//ADMIN PHARMACISTS
	adminPharmRepo := AdminRepository.NewManagePharmasistsRepository(dbconn)
	adminPharmSrv := AdminService.NewManagePharmasistsService(adminPharmRepo)
	adminPharmHandler := Admin_Handler.NewManagePharmasistHandler(adminPharmSrv)

	router.GET("/v1/admin/pharmacists", adminPharmHandler.GetPharmasists)
	router.GET("/v1/admin/pharmacists/:id", adminPharmHandler.GetSinglePharmasist)
	router.PUT("/v1/admin/pharmacists/:id", adminPharmHandler.UpdatePharmasist)
	router.POST("/v1/admin/pharmacists", adminPharmHandler.AddPharmasist)
	router.DELETE("/v1/admin/pharmacists/:id", adminPharmHandler.DeletePharmasist)

	//ADMIN LABORATORISTS
	adminLaborRepo := AdminRepository.NewManageLaboratoristsRepository(dbconn)
	adminLaborSrv := AdminService.NewManageLaboratoristsService(adminLaborRepo)
	adminLaborHnadler := Admin_Handler.NewManageLaboratoristHandler(adminLaborSrv)

	router.GET("/v1/admin/laboratorists", adminLaborHnadler.GetLaboratorists)
	router.GET("/v1/admin/laboratorists/:id", adminLaborHnadler.GetSingleLaboratorist)
	router.PUT("/v1/admin/laboratorists/:id", adminLaborHnadler.UpdateLaboratorist)
	router.POST("/v1/admin/laboratorists", adminLaborHnadler.AddLaboratorist)
	router.DELETE("/v1/admin/laboratorists/:id", adminLaborHnadler.DeleteLaboratorist)

	// patient request path registeration
	petientRequestRepo := PetientRepository.NewRequestGormRepo(dbconn)
	petientRequestServ := PetientService.NewRequestService(petientRequestRepo)
	PetientRequestHandler := Patient_Handler.NewPetientRequestHandler(petientRequestServ)
	router.GET("/v1/patient/requests", PetientRequestHandler.GetRequests)
	router.GET("/v1/patient/requests/:id", PetientRequestHandler.GetSingleRequest)
	router.POST("/v1/patient/requests", PetientRequestHandler.PostRequest)
	//
	// patient appointment path registeration
	petientAppointmentRepo := PetientRepository.NewAppointmentGormRepo(dbconn)
	petientAppointmentServ := PetientService.NewAppointmentService(petientAppointmentRepo)
	PetientAppointmentHandler := Patient_Handler.NewPetientAppointmentHandler(petientAppointmentServ)
	router.GET("/v1/patient/appointments", PetientAppointmentHandler.GetAppointments)
	router.GET("/v1/patient/appointments/:id", PetientAppointmentHandler.GetSingleAppointment)
	//
	// patient prescription path registeration
	petientPrescriptionRepo := PetientRepository.NewPrescriptionGormRepo(dbconn)
	petientPrescriptionServ := PetientService.NewPrescriptionService(petientPrescriptionRepo)
	PetientPrescriptionHandler := Patient_Handler.NewPetientPrescriptionHandler(petientPrescriptionServ)
	router.GET("/v1/patient/prescriptions", PetientPrescriptionHandler.GetSinglePrescription)
	router.GET("/v1/patient/prescriptions/:id", PetientPrescriptionHandler.GetPrescriptions)
	//
	// ** patient doctor path registeration
	petientDoctorRepo := PetientRepository.NewDoctorGormRepo(dbconn)
	petientDoctorServ := PetientService.NewDoctorService(petientDoctorRepo)
	PetientDoctorHandler := Patient_Handler.NewPetientDoctorHandler(petientDoctorServ)
	router.GET("/v1/patient/doctors", PetientDoctorHandler.GetDoctors)
	router.GET("/v1/patient/doctors/:id", PetientDoctorHandler.GetSingleDoctor)
	//
	// patient admin path registeration
	petientAdminRepo := PetientRepository.NewAdminGormRepo(dbconn)
	petientAdminServ := PetientService.NewAdminService(petientAdminRepo)
	PetientAdminHandler := Patient_Handler.NewPatientAdminHandler(petientAdminServ)
	router.GET("/v1/patient/admins/:id", PetientAdminHandler.GetAdmins)

	//
	// doctor patient path registeration
	patientRepo := repository.NewPatientGormRepo(dbconn)
	patientSrv := service.NewPatientService(patientRepo)
	doctorPatientHandler := Doctor_Handler.NewDoctorPatientHandler(patientSrv)

	router.GET("/v1/admin/patients/:id", doctorPatientHandler.GetSinglePatient)
	router.GET("/v1/admin/patients/", doctorPatientHandler.GetPatients)
	router.PUT("/v1/admin/patients/:id", doctorPatientHandler.PutPatient)
	router.POST("/v1/admin/patients/", doctorPatientHandler.PostPatient)
	router.DELETE("/v1/admin/patients/:id", doctorPatientHandler.DeletePatient)
	//
	// doctor general information path registeration
	appointmentRepo := repository.NewAppointmentGormRepo(dbconn)
	appointmentSrv := service.NewAppointmentService(appointmentRepo)
	doctorAppointmentHandler := Doctor_Handler.NewDoctorAppointmentHandler(appointmentSrv)

	router.GET("/v1/doctor/appointments/:id", doctorAppointmentHandler.GetSingleAppointment)
	router.GET("/v1/doctor/appointments/", doctorAppointmentHandler.GetAppointments)
	router.PUT("/v1/doctor/appointments/:id", doctorAppointmentHandler.PutAppointment)
	router.DELETE("/v1/doctor/appointments/:id", doctorAppointmentHandler.DeleteAppointment)
	//
	// general information path regiseration
	generalRepo := GeneralRepository.NewGeneralGormRepo(dbconn)
	generalSrv := GeneralService.NewGeneralService(generalRepo)
	generalHandler := General_Handler.NewGeneralHandler(generalSrv)

	router.GET("/v1/doctor/pharmacists/", generalHandler.GetPharmacists)
	router.GET("/v1/doctor/laboratorists/", generalHandler.GetLaboratorists)
	//
	// pharmacist medcine path registeration
	medRepo := PharmacistRepo.NewMedicineGormRepo(dbconn)
	medSrv := PharmacistService.NewMedicineService(medRepo)
	medHandler := handler.NewPharmMedicineHandler(medSrv)

	router.GET("/v1/pharm/medicines/:id", medHandler.GetSingleMedicine)
	router.GET("/v1/pharm/medMulti/:addedby", medHandler.GetMultipleMedicines)
	router.GET("/v1/pharm/medicines", medHandler.GetMedicines)
	router.PUT("/v1/pharm/medicines/:id", medHandler.PutMedicine)
	router.POST("/v1/pharm/medicines", medHandler.PostMedicine)
	router.DELETE("/v1/pharm/medicines/:id", medHandler.DeleteMedicine)
	//
	// pharmacist prescribtion path registeration
	prescRepo := PharmacistRepo.NewPrescriptionGormRepo(dbconn)
	prescSRV := PharmacistService.NewPrescriptionService(prescRepo)
	prescHandler := handler.NewPharmPrescriptionHandler(prescSRV)

	router.GET("/v1/pharm/presc/:id", prescHandler.GetSinglePrescription)
	router.GET("/v1/pharm/presc", prescHandler.GetMultiPrescriptions)
	router.PUT("/v1/pharm/presc/:id", prescHandler.PutPrescription)
	router.POST("/v1/pharm/presc", prescHandler.PostPrescription)
	router.GET("/v1/pharm/multiPre", prescHandler.GetMultiPrescriptions)
	router.DELETE("/v1/pharm/presc/:id", prescHandler.DeletePrescription)
	//
	// pharmacist profile path registeration
	profRepo := PharmacistRepo.NewPharmacistProfileGormRepo(dbconn)
	profSRV := PharmacistService.NewPharmacistProfileService(profRepo)
	profHandler := handler.NewPharmProfileHandler(profSRV)

	router.GET("/v1/pharm/profile/:id", profHandler.GetSingleProfile)
	router.GET("/v1/pharm/profile", profHandler.GetProfiles)
	router.PUT("/v1/pharm/profile/:id", profHandler.PutProfile)
	router.POST("/v1/pharm/profile", profHandler.PostProfile)
	router.DELETE("/v1/pharm/profile/:id", profHandler.DeleteProfile)
	//
	// laboratorist profile path registeration
	labProfRepo := laborrepo.NewLabratoristProfileGormRepo(dbconn)
	labProfSRV := laborSrv.NewLabratoristProfileService(labProfRepo)
	laborProdHandler := laborHandler.NewLaborProfileHandler(labProfSRV)

	router.GET("/v1/labor/profile/:id", laborProdHandler.GetSingleProfile)
	router.GET("/v1/labor/profile", laborProdHandler.GetProfiles)
	router.PUT("/v1/labor/profile/:id", laborProdHandler.PutProfile)
	router.POST("/v1/labor/profile", laborProdHandler.PostProfile)
	router.DELETE("/v1/labor/profile/:id", laborProdHandler.DeleteProfile)
	//
	// laboratorsit diagonosis path registeration
	labDiagRepo := laborrepo.NewDiagnosisGormRepo(dbconn)
	labdiagsrv := laborSrv.NewDiagnosisService(labDiagRepo)
	laborDiagHandler := laborHandler.NewLaborDiagnosisHandler(labdiagsrv)

	router.GET("/v1/labor/diag/:id", laborDiagHandler.GetSingleDiagnosis)
	router.GET("/v1/labor/diag", laborDiagHandler.GetDiagnosiss)
	router.GET("/v1/labor/multiDiag", laborDiagHandler.GetMultiDiagnosis)
	router.PUT("/v1/labor/diag/:id", laborDiagHandler.PutDiagnosis)
	router.POST("/v1/labor/diag", laborDiagHandler.PostDiagnosis)
	router.DELETE("/v1/labor/diag/:id", laborDiagHandler.DeleteDiagnosis)

	//
	http.ListenAndServe(":8180", router)

}
