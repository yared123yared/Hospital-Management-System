package Doctor_Handler

import (
	"fmt"
	//"html/template"
	"net/http"
	"net/url"
	"strconv"
	_ "time"

	//"github.com/betsegawlemma/restaurant/menu"

	"github.com/web1_group_project/hospital_client/data"
	_ "github.com/web1_group_project/hospital_client/data"
	Doctor_data "github.com/web1_group_project/hospital_client/data/Doctor"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/form"
	"github.com/web1_group_project/hospital_client/rtoken"
)

// MenuHandler handles menu related requests

// Index handles request on route /
func (mh *patientHandler) Prescribtions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("thise is the prescribtions method")
	prescriptions := mh.LogedInDoctor
	fmt.Println(prescriptions)

	mh.tmpl.ExecuteTemplate(w, "Doctor.prescribtion.html", prescriptions)
	//fmt.Println(appointment)
}
func (mh *patientHandler) SinglePrescribtion(prescribtion []entity.Prescription, id uint) entity.Prescription {
	singlePrescribtion := entity.Prescription{}
	for _, app := range prescribtion {
		if app.ID == id {
			singlePrescribtion = app

		}

	}
	return singlePrescribtion
}
func (mh *patientHandler) AddNewPrescribtions(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	//fmt.Println(doctor)
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)
	fmt.Println("Thise is the idddddddddddddd")
	fmt.Println(id)
	doctor := mh.LogedInDoctor
	//doctor.Pharmacist=pharm
	pharmacist, err := data.Pharmacists()
	if err != nil {
		fmt.Println("ERORRRRRRRRRRRRRRRRR")
		panic(err)
	}
	fmt.Println("prescribtionesssssssssss1")
	fmt.Println(doctor.Prescription)
	fmt.Println(uint(id))
	prescribtion1 := mh.SinglePrescribtion(doctor.Prescription, uint(id))
	fmt.Println("prescribtionesssssssssss2")
	dataToAdd := entity.AddPrescribtion{}
	fmt.Println("prescribtionesssssssssss3")
	dataToAdd.Pharmacist = *pharmacist
	dataToAdd.Prescription = prescribtion1
	fmt.Println(dataToAdd)
	if r.Method == http.MethodPost {
		fmt.Println(" i  am at the post method")
		// here will go the post method
		medName := r.FormValue("mdName")
		description := r.FormValue("description")

		pharId, err := strconv.Atoi(r.FormValue("pharmacist_id"))
		if err != nil {
			panic(err)
		}
		prescribtion1.MedicineName = medName
		prescribtion1.Description = description
		prescribtion1.PhrmacistId = uint(pharId)
		//prescribtion2 := GetModifiedPrescribtion(doctor.Prescription, prescribtion1, uint(id))

		//doctor.Prescription = prescribtion2
		err = Doctor_data.UpdatePrescribtion(&prescribtion1, int(prescribtion1.ID))
		if err != nil {
			panic(err)
		}
		fmt.Println(doctor)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			//tmpl.ExecuteTemplate(w, "error.layout", nil)
		}
		//doctor, err := data.Doctor(3)

		fmt.Println("i have done with the update")

		http.Redirect(w, r, "/doctor/prescribtion", http.StatusSeeOther)

	} else {

		prescribtionAddForm := struct {
			Values       url.Values
			VErrors      form.ValidationErrors
			CSRF         string
			Prescription entity.Prescription
			Pharmacist   []entity.Pharmacist
		}{
			Values:       nil,
			VErrors:      nil,
			CSRF:         token,
			Prescription: prescribtion1,
			Pharmacist:   *pharmacist,
		}
		mh.tmpl.ExecuteTemplate(w, "Doctor.add_prescribtion.html", prescribtionAddForm)

	}
}
func GetModifiedPrescribtion(prescribtion []entity.Prescription, appointment1 entity.Prescription, u uint) []entity.Prescription {
	fmt.Println("i am at the modifie method")
	a := []entity.Prescription{}
	for _, app := range prescribtion {
		if app.ID == u {
			a = append(a, appointment1)
		} else {
			a = append(a, app)
		}
	}
	fmt.Println("i am about to left the modified method")
	fmt.Println(a)
	return a
}

//
//// Menu handle request on route /menu
//func (mh *patientHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
//  if r.Method == http.MethodGet {
//
//    idRaw := r.URL.Query().Get("id")
//
//    id, err := strconv.Atoi(idRaw)
//
//    if err != nil {
//      panic(err)
//    }
//
//    err = data.DeleteUser(id)
//
//    if err != nil {
//      panic(err)
//    }
//
//  }
//
//  http.Redirect(w, r, "/doctor/appointments", http.StatusSeeOther)
//}
//
//// Contact handle request on route /Contact
//func (mh *patientHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
//  if r.Method == http.MethodGet {
//
//    idRaw := r.URL.Query().Get("id")
//
//    id, err := strconv.Atoi(idRaw)
//    fmt.Println("Thise is the idddddddddddddd")
//    fmt.Println(id)
//    if err != nil {
//      panic(err)
//    }
//
//    singleUser, err := data.FetchUser(id)
//    fmt.Println("fineeeeeeeeeeeeeeeeeeeee");
//
//    if err != nil {
//      fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRRRRRRR");
//      panic(err)
//    }
//
//    mh.tmpl.ExecuteTemplate(w, "Doctor.patient.html", singleUser)
//
//  } else if r.Method == http.MethodPost {
//
//    //ctg := entity.Role{}
//    //ctg.ID, _ = strconv.Atoi(r.FormValue("id"))
//    //ctg.Name = r.FormValue("name")
//    //
//    //
//    //
//    //err := rh.roleService.UpdateRole(ctg)
//    //
//    //if err != nil {
//    //  panic(err)
//    //}
//
//    http.Redirect(w, r, "/doctor/appointments", http.StatusSeeOther)
//
//  } else {
//    http.Redirect(w, r, "/doctor/appointments", http.StatusSeeOther)
//  }
//}
//
//// Admin handle request on route /admin
////func (mh *MenuHandler) Admin(w http.ResponseWriter, r *http.Request) {
////  mh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
////}
//func (mh *patientHandler) SingleAppointment(appointment []entity.Appointment,id uint)entity.Appointment {
//  singleAppointment:=entity.Appointment{}
//  for _,app:=range appointment{
//    if app.ID==id{
//      singleAppointment=app
//
//    }
//
//
//  }
//  return singleAppointment
//}
//func (mh *patientHandler) AddNewAppointment(w http.ResponseWriter, r *http.Request) {
//  fmt.Println("i am at hte add new appointment method")
//  //pageraw := r.FormValue("page")
//  //page, err := strconv.Atoi(pageraw)
//
//  // if err != nil {
//  //   w.WriteHeader(http.StatusNoContent)
//  //   tmpl.ExecuteTemplate(w, "error.layout", nil)
//  // }
//  idRaw := r.URL.Query().Get("id")
//
//  id, err := strconv.Atoi(idRaw)
//  doctor, err := data.Doctor(9)
//  if err!=nil{
//    panic(err)
//  }
//  appointment1:=mh.SingleAppointment(doctor.Appointment,uint(id))
//  fmt.Println("thise is the single value")
//
//  date:=r.FormValue("appDate")
//  fmt.Println(date)
//  //str:="2014-11-12"
//  appointment1.PatientUname="yared"
//  if err!=nil{
//    panic(err)
//  }
//  fmt.Println(appointment1)
//  appointment2:=GetModifiedAppointment(doctor.Appointment,appointment1,uint(id))
//
//  doctor.Appointment=appointment2
//  err=data.UpdateDoctor(doctor,9)
//  if err!=nil{
//    panic(err)
//  }
//  fmt.Println(doctor)
//  if err != nil {
//    w.WriteHeader(http.StatusNoContent)
//    //tmpl.ExecuteTemplate(w, "error.layout", nil)
//  }
//
//  fmt.Println("i have done with the update")
//
//  mh.tmpl.ExecuteTemplate(w, "Doctor.appointment.html", doctor)
//
//
//  //fmt.Println(appointment)
//
//
//
//
//
//
//
//
//
//}
//
//func GetModifiedAppointment(appointment []entity.Appointment, appointment1 entity.Appointment, u uint) []entity.Appointment{
//  fmt.Println("i am at the modifie method")
//  a:=[]entity.Appointment{}
//  for _,app:=range appointment{
//    if app.ID==u{
//      a=append(a,appointment1)
//    }else{
//      a=append(a,app)
//    }
//  }
//  fmt.Println("i am about to left the modified method")
//  fmt.Println(a)
//  return  a
//}
//
