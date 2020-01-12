package Admin_hanlder

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/web1_group_project/hospital_client/data/Admin"
	"github.com/web1_group_project/hospital_client/entity"
)

type AdminTempHandler struct {
	tmpl *template.Template
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewAdminTempHandler(T *template.Template) *AdminTempHandler {
	return &AdminTempHandler{tmpl: T}
}

func (ach *AdminTempHandler) DoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	doctors, _ := Admin.Doctors()
	ach.tmpl.ExecuteTemplate(w, "admin.manage.doctors.layout", doctors)

}
func (ach *AdminTempHandler) AddDoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		med := entity.Doctor{}
		// i, _ := strconv.ParseUint(r.FormValue("id"), 10, 64)
		fname := r.FormValue("firstname")
		lname := r.FormValue("lastname")

		med.Profile.FullName = fname + " " + lname
		med.Profile.BirthDate = time.Now()
		if r.FormValue("password") == r.FormValue("conpassword") {
			med.Profile.Password = r.FormValue("password")

		} else {
			ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", nil)
		}
		med.Profile.Password = r.FormValue("password")
		med.Profile.Email = r.FormValue("email")
		med.Profile.Phone = r.FormValue("phone")
		med.Profile.Address = r.FormValue("address")
		med.Profile.Sex = r.FormValue("sex")
		med.Profile.RoleId = 2
		med.Profile.Description = r.FormValue("description")
		// med.Uuid = med.Profile.ID
		// med.ID = med.Profile.ID
		// prof := med.Profile
		// Admin.PostProfile(&prof)
		mf, fh, err := r.FormFile("image")
		if err != nil {
			panic(err)
		}
		defer mf.Close()
		med.Profile.Image = fh.Filename

		Admin.WriteFile(&mf, fh.Filename)
		med.Department = r.FormValue("department")
		Admin.PostDoctor(&med)
		// fmt.Println(prof.ID, "sdfjhdsjhfjdshfgdfskjhgjhfsdkjghkjdfshgkjh")
		http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)

	} else {
		fmt.Println("temlating")
		ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", nil)
		fmt.Println("temlating2")

	}
}
func (ach *AdminTempHandler) UpdateDoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := Admin.GetDoctor(uint(id))
		fmt.Println("updateing", parms)

		ach.tmpl.ExecuteTemplate(w, "admin.update.manage.doctors.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))
		fmt.Println("id", id)
		medicine, _ := Admin.GetDoctor(uint(id))
		fmt.Println("medinen", medicine)

		med := entity.Doctor{}

		med.Profile.ID = medicine.Profile.ID
		fname := r.FormValue("fullname")

		med.Profile.FullName = fname
		med.Profile.BirthDate = time.Now()
		if r.FormValue("password") == r.FormValue("conpassword") {
			med.Profile.Password = r.FormValue("password")

		} else {
			ach.tmpl.ExecuteTemplate(w, "admin.update.doctors.layout", nil)
		}
		med.Profile.Password = r.FormValue("password")
		med.Profile.Email = r.FormValue("email")
		med.Profile.Phone = r.FormValue("phone")
		med.Profile.Address = r.FormValue("address")
		med.Profile.Sex = r.FormValue("sex")
		med.Profile.RoleId = 2
		med.Profile.Description = r.FormValue("description")
		var image, _, _ = r.FormFile("image")
		if image == nil {
			med.Profile.Image = r.FormValue("hidimage")
		} else {
			mf, fh, err := r.FormFile("image")
			if err != nil {
				panic(err)
			}
			defer mf.Close()

			med.Profile.Image = fh.Filename

			Admin.WriteFile(&mf, fh.Filename)

		}
		med.Department = r.FormValue("department")
		fmt.Println("TO be done", med)
		Admin.PutDoctor(&med)

		http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)
	}
}

func (ach *AdminTempHandler) DeleteDoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		fmt.Println(id, "ifhdsjkhfjdh")

		Admin.DeleteDoctor(uint(id))

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)
}
