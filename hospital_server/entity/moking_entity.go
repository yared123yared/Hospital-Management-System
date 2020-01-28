package entity

import "time"

// ProfileMock mokes Profile
var ProfileMock = User{
	ID:          1,
	FullName:    "Mock Name 01",
	Password:    "password 01",
	Email:       "Mock Email 01",
	Phone:       "Mock Phone 01",
	Address:     "Mock Address 01",
	Image:       "Mock Image 01",
	Sex:         "Mock sex 01",
	RoleId:      3,
	BirthDate:   time.Time{},
	Description: "Mock Description 01",
}

// PharmacistMock mocks Pharmacist
var PharmacistMock = Pharmacist{
	ID:           1,
	Uuid:         10,
	User:         User{},
	Medicine:     []Medicine{},
	Prescription: []Prescription{},
}

//PrescriptionMock mocks Description
var PrescriptionMock = Prescription{
	ID:             1,
	PatientId:      1,
	PatientName:    "Mock patient 01",
	DoctorId:       1,
	PhrmacistId:    1,
	PrescribedDate: time.Time{},
	MedicineName:   "Mock Medicine 01",
	Description:    "Mock prescription 01 Description",
	GivenStatus:    "Mock prescritpin 01 givenstatus",
	GivenDate:      time.Time{},
}

//MedicineMock mocks Medicine
var MedicineMock = Medicine{
	ID:           1,
	CategoryName: "Mock CategoryName 01",
	MedicineName: "Mock MedicineName 01",
	ExpiredAt:    time.Time{},
	Amount:       1,
	AddedBy:      1,
}

//RoleMock mocks Role
var RoleMock = Role{
	ID:    1,
	Users: []User{},
}

//DoctorMock mocks Doctor
var DoctorMock = Doctor{
	ID:           1,
	User:         User{},
	Uuid:         1,
	Department:   "Mock Department 01",
	Prescription: []Prescription{},
	Diagnosis:    []Diagnosis{},
	Appointment:  []Appointment{},
	Pharmacist:   []Pharmacist{},
}

//LaboratoristMock mocks Laboratorist
var LaboratoristMock = Laboratorist{
	ID:        1,
	Uuid:      0,
	User:      User{},
	Diagnosis: []Diagnosis{},
}

//DiagnosisMock mocks Diagnosis
var DiagnosisMock = Diagnosis{
	ID:             1,
	PatientId:      1,
	PatientName:    "Mock ",
	DoctorId:       0,
	LaboratoristId: 0,
	Description:    "",
	DiagonosesDate: time.Time{},
	Reponse:        "",
	ResponseDate:   time.Time{},
}
var PetientMock = Petient{
	ID:           1,
	Uuid:         10,
	User:         User{},
	BloodGroup:   "Mock Patient 01",
	Age:          1,
	Prescription: []Prescription{},
	Diagnosis:    []Diagnosis{},
	Appointment:  []Appointment{},
	Request:      []Request{},
}
var AppointmentMock = Appointment{
	ID:          1,
	PatientId:   1,
	PatientName: "Mock patient 01",
	DoctorId:    1,
	Date:        time.Time{},
}
var RequestMock = Request{
	ID:            1,
	PatientId:     1,
	PatientName:   "Mock patient 01",
	DoctorId:      1,
	ApproveStatus: "Mock patient 01 approve status",
	ApprovedBy:    1,
}

var AdminMock = Admin{
	ID:      1,
	Uuid:    1,
	User:    User{},
	Request: []Request{},
}
