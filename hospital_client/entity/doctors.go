package entity

type AddPrescribtion struct {
	Prescription Prescription
	Pharmacist   []Pharmacist
}
type AddDiagonosis struct {
	Diagnosis    Diagnosis
	Laboratorist []Laboratorist
}
