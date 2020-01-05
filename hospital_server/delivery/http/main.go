/*
package main
import(

	//

	"github.com/yaredsolomon/webProgram1/hospital/entity"
	

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main(){


	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital3?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	errs:=dbconn.CreateTable(&entity.Pharmacist{},
		&entity.Petient{},&entity.Profile{},
		&entity.Prescription{},&entity.Admin{},
		&entity.PetientHistory{},&entity.AdminHistory{},&entity.DoctorHistory{},&entity.Doctor{},
		&entity.PharmasistHistory{},&entity.LaboratoristHistory{},&entity.Laboratorist{},&entity.Appointment{},
		&entity.Diagnosis{},&entity.Medicine{},&entity.Request{}).GetErrors()
	 if errs!=nil {
 		panic(errs)
		
	 }
	 

}
*/


package main

import (
	"net/http"
	//

	// "github.com/betsegawlemma/restaurant-rest/comment/repository"
	// "github.com/betsegawlemma/restaurant-rest/comment/service"
	// "github.com/betsegawlemma/restaurant-rest/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/yaredsolomon/webProgram1/hospital/Registeration/repository"
	"github.com/yaredsolomon/webProgram1/hospital/Registeration/service"
	"github.com/yaredsolomon/webProgram1/hospital/delivery/http/handler"
	 _"github.com/yaredsolomon/webProgram1/sathurday18/comment/repository"
	// "github.com/yaredsolomon/webProgram1/sathurday18/user/repository"
	 _"github.com/yaredsolomon/webProgram1/sathurday18/comment/service"
	// "github.com/yaredsolomon/webProgram1/sathurday18/user/service"
	// "github.com/yaredsolomon/webProgram1/sathurday18/delivery/http/handler"
	_"github.com/yaredsolomon/webProgram1/hospital/entity"
)

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/hospital3?sslmode=disable")

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
	// errs1:=dbconn.CreateTable(&entity.Prescription{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").GetErrors()
	// if errs1!=nil {
	// 	panic(errs1)
	// }
	//
	// errs1:=dbconn.CreateTable(&entity.Diagnosis{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").GetErrors()
	// if errs1!=nil {
	// 	panic(errs1)
	// }
	//
	// errs1:=dbconn.CreateTable(&entity.Appointment{}).AddForeignKey("patient_Id","petients(Id)","cascade","cascade").GetErrors()
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
}
