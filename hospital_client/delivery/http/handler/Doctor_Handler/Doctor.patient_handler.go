package Doctor_Handler

import (
	"fmt"
	_ "github.com/web1_group_project/hospital_client/data"
	Doctor_data "github.com/web1_group_project/hospital_client/data/Doctor"
	"github.com/web1_group_project/hospital_client/entity"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	//"github.com/betsegawlemma/restaurant/menu"
)

// MenuHandler handles menu related requests
type patientHandler struct {
	tmpl *template.Template
	//doctorSrv doctor.CategoryService
}

//NewMenuHandler initializes and returns new MenuHandler
func NewpatientHandler(T *template.Template) *patientHandler {
	return &patientHandler{tmpl: T}
}
func (mh *patientHandler) Index(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "Doctor.index.html", nil)
}

// Index handles request on route /
func (mh *patientHandler) Patients(w http.ResponseWriter, r *http.Request) {
	//pageraw := r.FormValue("page")
	//page, err := strconv.Atoi(pageraw)

	// if err != nil {
	// 	w.WriteHeader(http.StatusNoContent)
	// 	tmpl.ExecuteTemplate(w, "error.layout", nil)
	// }
	fmt.Println("i am about to fech data")
	//users, err := data.FetchUsers()
	//petient:=[]entity.Petient
	petient, err := Doctor_data.FetchUsers()
	fmt.Println(petient)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		//tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	mh.tmpl.ExecuteTemplate(w, "Doctor.patient.html", petient)

}

// About handles requests on route /about
func (mh *patientHandler) AddNewPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		patient := entity.Petient{}

		patient.Profile.FullName = r.FormValue("full_name")
		if r.FormValue("user_password") == r.FormValue("confirm_password") {
			patient.Profile.Password = r.FormValue("user_password")

		} else {
			mh.tmpl.ExecuteTemplate(w, "Doctor.patient.layout", nil)
		}

		patient.Profile.Email = r.FormValue("email")
		patient.Profile.Phone = r.FormValue("contact_no")
		patient.Profile.Address = r.FormValue("address")
		// image store

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		writeFile(&mf, fh.Filename)
		patient.Profile.Image = fh.Filename
		patient.Profile.Sex = r.FormValue("sex")
		patient.Profile.RoleId = 5
		//patient.Profile.BirthDate = r.FormValue("bDate")
		patient.Profile.Description = r.FormValue("description")
		patient.BloodGroup = r.FormValue("blood_group")
		//ctg.ID, _ = strconv.Atoi(r.FormValue("id"))
		patient.Age, _ = strconv.Atoi(r.FormValue("age"))
		//if err!=nil{
		//	panic(err)
		//}

		//ctg.Description = r.FormValue("description")
		err = Doctor_data.StorePatients(&patient)

		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("This is the errorrrrrrrrrrrrrr")
			panic(err)
		}

		http.Redirect(w, r, "/doctor/patients", http.StatusSeeOther)

	} else {

		mh.tmpl.ExecuteTemplate(w, "Doctor.patient.layout", nil)

	}

}

// Menu handle request on route /menu
func (mh *patientHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
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

	http.Redirect(w, r, "/doctor/patients", http.StatusSeeOther)
}

// Contact handle request on route /Contact
func (mh *patientHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(singleUser)
		if err != nil {
			fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRRRRRRR")
			panic(err)
		}

		mh.tmpl.ExecuteTemplate(w, "Doctor.patientUpdate.layout", singleUser)

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
		//	panic(err)
		//}

		http.Redirect(w, r, "/doctor/patients", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/doctor/patients", http.StatusSeeOther)
	}
}

// Admin handle request on route /admin
//func (mh *MenuHandler) Admin(w http.ResponseWriter, r *http.Request) {
//	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
//}
func writeFile(mf *multipart.File, fname string) {
	fmt.Println("i am at the image store")

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("This is first erorrrrrrrrrrrrrrrrrrrrrr")
		panic(err)
	}

	path := filepath.Join(wd, "../ui", "assets", "img", fname)
	image, err := os.Create(path)
	//	fmt.Println(*image)

	if err != nil {
		fmt.Println("This is second erorrrrrrrrrrrrrrrrrrrrrr")

		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
