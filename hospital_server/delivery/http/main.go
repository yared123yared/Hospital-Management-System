/*

package main
import(

	//

	"github.com/yaredsolomon/webProgram1/hospital/entity"


	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main(){


	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital4?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	errs:=dbconn.CreateTable(&entity.Profile{}).GetErrors()
	errs=dbconn.CreateTable(&entity.Pharmacist{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Petient{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Admin{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Doctor{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Laboratorist{}).AddForeignKey("uuid","profiles(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Prescription{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("phrmacist_Id","pharmacists(Id)","cascade","cascade").GetErrors()

	errs=dbconn.CreateTable(&entity.Appointment{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("patient_uname","profiles(full_name)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Diagnosis{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("laboratorist_Id","laboratorists(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Medicine{}).AddForeignKey("added_By","pharmacists(Id)","cascade","cascade").GetErrors()
	errs=dbconn.CreateTable(&entity.Request{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").AddForeignKey("approved_By","admins(Id)","cascade","cascade").GetErrors()









	if errs!=nil {
		panic(errs)

	}



}




*/

package main

import (
	"github.com/monday271/hospital_server/Doctor/repository"
	"github.com/monday271/hospital_server/Doctor/service"
	"github.com/monday271/hospital_server/delivery/http/handler"
	"net/http"
	/*
	 */

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	//"github.com/yaredsolomon/webProgram1/hospital/Registeration/repository"
	//"github.com/yaredsolomon/webProgram1/hospital/Registeration/service"
	//"github.com/yaredsolomon/webProgram1/hospital/request/repository"
	//"github.com/yaredsolomon/webProgram1/hospital/request/service"
	//"github.com/yaredsolomon/webProgram1/hospital/delivery/http/handler"
	// _"github.com/yaredsolomon/webProgram1/sathurday18/comment/repository"
	// _"github.com/yaredsolomon/webProgram1/sathurday18/comment/service"
	//_"github.com/yaredsolomon/webProgram1/hospital/entity"
)

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital4?sslmode=disable")

	if err != nil {
		panic(err)
	}

	// defer dbconn.Close()
	// errs:=dbconn.CreateTable(&entity.Profile{}).GetErrors()
	// if errs!=nil {
	// 	panic(errs)

	// }

	//petients
	// errs:=dbconn.CreateTable(&entity.Petient{}).AddForeignKey("Uuid","profiles(Id)","cascade","cascade").GetErrors()
	// if errs!=nil {
	// 	panic(errs)
	// }

	// errs1:=dbconn.CreateTable(&entity.PetientHistory{}).AddForeignKey("patient_Id","petients(uuid)","cascade","cascade").GetErrors()
	// if errs1!=nil {
	// 	panic(errs1)
	// }
	//
	// errs1:=dbconn.CreateTable(&entity.Prescription{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").GetErrors()
	// if errs1!=nil {
	// 	panic(errs1)uint  `gorm:"not null"`
	// }

	//
	// errs1:=dbconn.CreateTable(&entity.Diagnosis{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").AddForeignKey("doctor_Id","doctors(Id)","cascade","cascade").GetErrors()
	// if errs1!=nil {
	// 	panic(errs1)
	// }
	//
	// errs1:=dbconn.CreateTable(&entity.Appointment{}).s.GetErrors()
	// if errs1!=nil {
	// 	panic(errs1)
	// }

	//errs:=dbconn.CreateTable(&entity.Cate{}).GetErrors()

	// commentRepo := comment.repository.NewCommentGormRepo(dbconn)
	// commentSrv := service.NewCommentService(commentRepo)

	// adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)

	// router := httprouter.New()

	// router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	// router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	// router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	// router.POST("/v1/admin/comments", adminCommentHandler.PostComment)
	// router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)
	//
	// patient registeration
	/*
		patientRepo := repository.NewPatientGormRepo(dbconn)
		patientSrv := service.NewPatientService(patientRepo)

		doctorPatientHandler := handler.NewDoctorPatientHandler(patientSrv)

		router := httprouter.New()

		router.GET("/v1/admin/users/", doctorPatientHandler.GetSinglePatient)
		router.GET("/v1/admin/users", doctorPatientHandler.GetPatients)
		router.PUT("/v3/admin/users/:id", doctorPatientHandler.PutPatient)
		router.POST("/v1/admin/users", doctorPatientHandler.PostPatient)
		router.DELETE("/v2/admin/users/:id", doctorPatientHandler.DeletePatient)

		http.ListenAndServe(":8980", router)
	*/

	appointmentRepo := repository.NewAppointmentGormRepo(dbconn)
	appointmentSrv := service.NewAppointmentService(appointmentRepo)

	doctorAppointmentHandler := handler.NewDoctorAppointmentHandler(appointmentSrv)

	router := httprouter.New()

	router.GET("/v1/doctor/appointments/:id ", doctorAppointmentHandler.GetSingleAppointment)
	router.GET("/v1/doctor/appointments", doctorAppointmentHandler.GetAppointments)
	router.PUT("/v1/doctor/appointments/:id", doctorAppointmentHandler.PutAppointment)

	router.DELETE("/v1/doctor/appointments/:id", doctorAppointmentHandler.DeleteAppointment)

	http.ListenAndServe(":8980", router)
}
