package Doctor_Handler

import (
	"fmt"
	data "github.com/getach1/web1/web1_group_project-master/hospital_client/data/Doctor"
	"github.com/getach1/web1/web1_group_project-master/hospital_client/entity"
	"html/template"
	"net/http"
	"strconv"
	_ "time"
)

// MenuHandler handles menu related requests
type diagonosisHandler struct {
	tmpl *template.Template
	//doctorSrv doctor.CategoryService
}

//NewMenuHandler initializes and returns new MenuHandler
func NewdiagonosisHandler(T *template.Template) *diagonosisHandler {
	return &diagonosisHandler{tmpl: T}
}

// Index handles request on route /
func (mh *diagonosisHandler) Diagonosises(w http.ResponseWriter, r *http.Request) {
	fmt.Println("thise is the Diagonosises method")
	diagonosises, err := data.Doctor(1)
	fmt.Println(diagonosises)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		//tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	mh.tmpl.ExecuteTemplate(w, "Doctor.checkUp.html", diagonosises)
	//fmt.Println(appointment)
}
func (mh *diagonosisHandler) SingleDiagonosis(diagonosis []entity.Diagnosis, id uint) entity.Diagnosis {
	singleDiagonosis := entity.Diagnosis{}
	for _, app := range diagonosis {
		if app.ID == id {
			singleDiagonosis = app

		}

	}
	return singleDiagonosis
}
func (mh *diagonosisHandler) AddNewDiagonosis(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(doctor)
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)
	fmt.Println("Thise is the idddddddddddddd")
	fmt.Println(id)
	if err != nil {
		panic(err)
	}
	doctor, err := data.Doctor(1)
	//doctor.Pharmacist=pharm
	laboratorist, err := data.Laboratorists()
	if err != nil {
		panic(err)
	}
	diagonosis1 := mh.SingleDiagonosis(doctor.Diagnosis, uint(id))
	dataToAdd := entity.AddDiagonosis{}
	dataToAdd.Laboratorist = *laboratorist
	dataToAdd.Diagnosis = diagonosis1
	fmt.Println(dataToAdd)
	if r.Method == http.MethodPost {
		fmt.Println(" i  am at the post method")
		// here will go the post method
		//medName:=r.FormValue("mdName")
		description := r.FormValue("description")

		labId, err := strconv.Atoi(r.FormValue("laboratorist_id"))
		if err != nil {
			panic(err)
		}
		//diagonosis1.MedicineName=medName
		diagonosis1.Description = description
		diagonosis1.LaboratoristId = uint(labId)
		diagonosis2 := GetModifiedDiagonosis(doctor.Diagnosis, diagonosis1, uint(id))

		doctor.Diagnosis = diagonosis2
		err = data.UpdateDoctor(doctor, 1)
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

		http.Redirect(w, r, "/doctor/diagonosis", http.StatusSeeOther)

	} else {

		//doctor, err := data.Doctor(3)

		fmt.Println(" i am at the get method")
		fmt.Println(dataToAdd.Diagnosis)

		mh.tmpl.ExecuteTemplate(w, "Doctor.add_checkUp.html", dataToAdd)

	}
}
func GetModifiedDiagonosis(diagonosis []entity.Diagnosis, diagonosis1 entity.Diagnosis, u uint) []entity.Diagnosis {
	fmt.Println("i am at the modifie method")
	a := []entity.Diagnosis{}
	for _, app := range diagonosis {
		if app.ID == u {
			a = append(a, diagonosis1)
		} else {
			a = append(a, app)
		}
	}
	fmt.Println("i am about to left the modified method")
	fmt.Println(a)
	return a
}
