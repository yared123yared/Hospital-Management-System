package entity

import "time"

// Category represents Food Menu Category
type Profile struct {
	ID          uint  
	FullName   string `gorm:"type:varchar(255);not null"`
	UserName      string `gorm:"type:varchar(255);not null"`
	Password      string `gorm:"type:varchar(255);"`
	Email      string `gorm:"type:varchar(255);"`
	Phone      string `gorm:"type:varchar(255);"`
	Address      string `gorm:"type:varchar(255);"`
	Image      string `gorm:"type:varchar(255);"`
	Sex      string `gorm:"type:varchar(255);"`
	Role      string `gorm:"type:varchar(255);"`
	BirthDate    time.Time
	Description string

}
type Doctor struct {
	ID       uint `gorm:"not null"`
	Profile  Profile `gorm:"ForeignKey:ID"`
	Uuid uint 
	Department string `gorm:"type:varchar(255);not null"`

	//DoctorHistory []DoctorHistory
	//PetientHistory []PetientHistory
	Prescription []Prescription  
	Diagnosis    []Diagnosis   
}
type Appointment struct {
	ID              uint
	PatientId  uint `gorm:"not null"`
	PatientUname string `gorm:"type:varchar(255);not null"`
	DoctorId  uint `gorm:"not null"`
	Date            time.Time
}
type Petient struct {
	ID       uint `gorm:"not null"`
	Uuid uint 
	Profile  Profile `gorm:"ForeignKey:Uuid"`

	BloodGroup string `gorm:"type:varchar(255);not null"`
	Age int
	Prescription    []Prescription  `gorm:"ForeignKey:PatientId"`
	Diagnosis       []Diagnosis   `gorm:"ForeignKey:PatientId"`
	Appointment []Appointment `gorm:"ForeignKey:PatientId"`
}
type Pharmacist struct {
	ID       uint
	Profile  Profile  
	PharmacistHistory []PetientHistory
}
type Laboratorist struct {
	ID       uint
	Profile  Profile 
	LaboratoristHistory []LaboratoristHistory
}

type Admin struct {
	ID         uint      
	Appointment []Appointment `gorm:"many2many:admin_appointment"`
	Request    []Request
	Profile    Profile
}
type PharmasistHistory struct {
	ID           uint          

	Medicine     []Medicine    `gorm:"many2many:ph_medcines"`
	Prescription []Prescription  `gorm:"many2many:ph_prescribtion"`
}
type LaboratoristHistory struct {
	ID           uint          

	Diagnosis    []Diagnosis  `gorm:"many2many:lbh_diagonasis"` 
}
type DoctorHistory struct {
	ID           uint      


}
type Request struct {
	ID      uint      
	Patient []Petient 
}
type AdminHistory struct{
	ID      	uint       

	Request     Request 
	Profile     []Profile  
}

type Prescription struct {
	ID                 uint
	PatientId        uint `gorm:"not null"`
	DoctorId   uint `gorm:"not null"`
	PhrmacistId    uint 
	Date               time.Time
	MedicineName       string  `gorm:"type:varchar(255);"`
	Description        string  `gorm:"type:varchar(255);"`
	GivenStatus 	 string  `gorm:"type:varchar(255);"`
}

type Medicine struct {
	ID           uint
	CategoryName string  `gorm:"type:varchar(255);not null"`
	MedicineName string  `gorm:"type:varchar(255);not null"`
	ExpiredAt    time.Time
	Amount       uint
	AddedBy      string  `gorm:"type:varchar(255);not null"`
}
type PetientHistory struct {
	ID              uint 
	PatientId  uint `gorm:"not null"`

}

type Diagnosis struct {
	ID                  uint
	PatientId     		uint `gorm:"not null"`
	DoctorUserName      string `gorm:"type:varchar(255);not null"`
	LabratoristUserName string `gorm:"type:varchar(255);"`
	Date                time.Time
	Reponse             string `gorm:"type:varchar(255);"`
	Description         string `gorm:"type:varchar(255);"`
}

type Error struct {
	Code    int
	Message string
}

