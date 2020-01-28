package Doctor_Handler

import (
	"fmt"
	"html/template"
	"io"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"github.com/web1_group_project/hospital_client/form"
	"github.com/web1_group_project/hospital_client/rtoken"
	//"github.com/betsegawlemma/restaurant/menu"
	
	_ "github.com/web1_group_project/hospital_client/data"
	Doctor_data "github.com/web1_group_project/hospital_client/data/Doctor"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/delivery/http/handler"
	
)

// MenuHandler handles menu related requests
type patientHandler struct {
	tmpl *template.Template
	UserHandler   *handler.UserHandler
  	LogedInDoctor *entity.Doctor
  	csrfSignKey    []byte
	//doctorSrv doctor.CategoryService
}

//NewMenuHandler initializes and returns new MenuHandler
func NewpatientHandler(T *template.Template,userHandler *handler.UserHandler,csKey []byte) *patientHandler {
	return &patientHandler{
		tmpl:T,
		UserHandler:userHandler,
		csrfSignKey:    csKey,
	}
	
}
func (mh *patientHandler)Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am about to fech data")
	
	doctor,err:= Doctor_data.Doctor(int(mh.UserHandler.LoggedInUser.ID))
	//fmt.Println(doctor)
	if err!=nil{
		panic(err)
	}
	
	mh.LogedInDoctor=doctor
	
    //petient:=*mh.LogedInDoctor
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
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		patientRegisterationForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		// patientRegisterationForm.Required("fullname", "email", "password", "confirmpassword")
		// patientRegisterationForm.MatchesPattern("email", form.EmailRX)
		// patientRegisterationForm.MatchesPattern("phone", form.PhoneRX)
		// patientRegisterationForm.MinLength("password", 8)
		// patientRegisterationForm.PasswordMatches("user_password", "confirm_password")
		 patientRegisterationForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !patientRegisterationForm.Valid() {
			mh.tmpl.ExecuteTemplate(w, "Doctor.patient.layout", patientRegisterationForm)
			return
		}

		// pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		// if pExists {
		// 	singnUpForm.VErrors.Add("phone", "Phone Already Exists")
		// 	mh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }
		// eExists := uh.userService.EmailExists(r.FormValue("email"))
		// if eExists {
		// 	singnUpForm.VErrors.Add("email", "Email Already Exists")
		// 	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("user_password")), 12)
		if err != nil {
			fmt.Println("hash password")
			patientRegisterationForm.VErrors.Add("password", "Password Could not be stored")
			mh.tmpl.ExecuteTemplate(w,"Doctor.patient.layout", patientRegisterationForm)
			return
		}

		// role, errs := uh.userRole.RoleByName("USER")

		// if len(errs) > 0 {
		// 	singnUpForm.VErrors.Add("role", "could not assign role to the user")
		// 	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }

	


		patient := entity.Petient{}
		patient.User.FullName = r.FormValue("full_name")
		patient.User.Password = string(hashedPassword)
		patient.User.Email = r.FormValue("email")
		patient.User.Phone = r.FormValue("contact_no")
		patient.User.Address = r.FormValue("address")
		// image store

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		writeFile(&mf, fh.Filename)
		patient.User.Image = fh.Filename
		patient.User.Sex = r.FormValue("sex")
		patient.User.RoleId = 2
		//patient.Profile.BirthDate = r.FormValue("bDate")
		patient.User.Description = r.FormValue("description")
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
	
	if r.Method == http.MethodGet {
		fmt.Println("thise is the get methd")
		patientAddForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}

		

		mh.tmpl.ExecuteTemplate(w, "Doctor.patient.layout", patientAddForm)

	}

}
}

// Menu handle request on route /menu
func (mh *patientHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	
		fmt.Println("thise is the delete methodddddddddddddddddddddddddd")

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = Doctor_data.DeleteUser(id)

		if err != nil {
			panic(err)
		}

	

	http.Redirect(w, r, "/doctor/patients", http.StatusSeeOther)
}


// Contact handle request on route /Contact
func (mh *patientHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("thise is the Updte method")

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		fmt.Println("Thise is the idddddddddddddd")
		fmt.Println(id)
		if err != nil {
			panic(err)
		}

		singleUser, err := Doctor_data.FetchUser(id)
		fmt.Println("fineeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
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
