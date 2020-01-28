package Doctor_Handler

import (
	"fmt"
	// "html/template"
	"net/http"
	"strconv"
	"time"

	//"github.com/betsegawlemma/restaurant/menu"

	Doctor_data "github.com/web1_group_project/hospital_client/data/Doctor"
	"github.com/web1_group_project/hospital_client/entity"
	// "github.com/web1_group_project/hospital_client/delivery/http/handler"
)

// MenuHandler handles menu related requests
// type appointmentHandler struct {

//   tmpl *template.Template
//   UserHandler   *handler.UserHandler
//     LogedInDoctor *entity.Doctor
//     csrfSignKey    []byte
//   //doctorSrv doctor.CategoryService
// }

// //NewMenuHandler initializes and returns new MenuHandler
// func NewappointmentHandler(T *template.Template,userHandler *handler.UserHandler,csKey []byte) *appointmentHandler {
//   return &appointmentHandler{
//     tmpl:T,
//     UserHandler:userHandler,
//     csrfSignKey:    csKey,

//   }
// }
func (mh *patientHandler) appointmentIndex(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "Doctor.index.html", nil)
}

// Index handles request on route /
func (mh *patientHandler) Appointment(w http.ResponseWriter, r *http.Request) {

	//pageraw := r.FormValue("page")
	//page, err := strconv.Atoi(pageraw)

	// if err != nil {
	//   w.WriteHeader(http.StatusNoContent)
	//   tmpl.ExecuteTemplate(w, "error.layout", nil)
	// }
	fmt.Println("i am at the appointment method")
	//users, err := data.FetchUsers()
	//petient:=[]entity.Petient
	//doctor1, err := Doctor_data.Doctor(doctor.U)
	fmt.Println(mh.LogedInDoctor)
	// if err != nil {
	//   w.WriteHeader(http.StatusNoContent)
	//   //tmpl.ExecuteTemplate(w, "error.layout", nil)
	// }

	//appointment:=doctor.Appointment
	//fmt.Println(appointment)

	mh.tmpl.ExecuteTemplate(w, "Doctor.appointment.html", mh.LogedInDoctor)

	//fmt.Println(appointment)
}

// About handles requests on route /about
//func (mh *patientHandler) AddNewAppointment(w http.ResponseWriter, r *http.Request) {
//  fmt.Println("thise is the new appointment  method")
//  //mh.tmpl.ExecuteTemplate(w, "Doctor.add_patient.html", nil)
//}

// Menu handle request on route /menu
func (mh *patientHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = Doctor_data.DeleteUser(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/doctor/appointments", http.StatusSeeOther)
}

// Contact handle request on route /Contact
func (mh *patientHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		fmt.Println("Thise is the idddddddddddddd")
		fmt.Println(id)
		if err != nil {
			panic(err)
		}

		singleUser, err := Doctor_data.FetchUser(id)
		fmt.Println("fineeeeeeeeeeeeeeeeeeeee")

		if err != nil {
			fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRRRRRRR")
			panic(err)
		}

		mh.tmpl.ExecuteTemplate(w, "Doctor.patient.html", singleUser)

	} else if r.Method == http.MethodPost {

		//ctg := entity.Role{}
		//ctg.ID, _ = strconv.Atoi(r.FormValue("id"))
		//ctg.Name = r.FormValue("name")
		//
		//
		//
		//err := rh.roleService.UpdateRole(ctg)
		//
		//if err != nil {
		//  panic(err)
		//}

		http.Redirect(w, r, "/doctor/appointments", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/doctor/appointments", http.StatusSeeOther)
	}
}

// Admin handle request on route /admin
//func (mh *MenuHandler) Admin(w http.ResponseWriter, r *http.Request) {
//  mh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
//}
func (mh *patientHandler) SingleAppointment(appointment []entity.Appointment, id uint) entity.Appointment {
	singleAppointment := entity.Appointment{}
	for _, app := range appointment {
		if app.ID == id {
			singleAppointment = app

		}

	}
	return singleAppointment
}
func (mh *patientHandler) AddNewAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am at hte add new appointment method")
	//pageraw := r.FormValue("page")
	//page, err := strconv.Atoi(pageraw)

	// if err != nil {
	//   w.WriteHeader(http.StatusNoContent)
	//   tmpl.ExecuteTemplate(w, "error.layout", nil)
	// }
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)
	doctor := mh.LogedInDoctor
	if err != nil {
		panic(err)
	}
	appointment1 := mh.SingleAppointment(doctor.Appointment, uint(id))
	fmt.Println("thise is the single value")

	date := r.FormValue("appDate")
	fmt.Println(date)
	//str:="2014-11-12"

	appointment1.Date = time.Now()
	if err != nil {
		panic(err)
	}
	fmt.Println(appointment1)

	err = Doctor_data.UpdateAppointment(&appointment1, int(appointment1.ID))
	if err != nil {
		panic(err)
	}
	fmt.Println(doctor)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		//tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	fmt.Println("i have done with the update")

	mh.tmpl.ExecuteTemplate(w, "Doctor.appointment.html", doctor)

	//fmt.Println(appointment)

}

/*
func GetModifiedAppointment(appointment []entity.Appointment, appointment1 entity.Appointment, u uint) []entity.Appointment {
  fmt.Println("i am at the modifie method")
  a := []entity.Appointment{}
  for _, app := range appointment {
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
*/
