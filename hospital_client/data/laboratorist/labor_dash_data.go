package service


var baseURLDash = "http://localhost:8182/v1/pharm/profile/"
var baseURLDash2 = "http://localhost:8182/v1/pharm/multi/"

func GetDiagsNumber(id uint) int {
	laboratorist, _ := GetLaboratorist(id)
	diagnosis := laboratorist.Diagnosis
	length := len(diagnosis)
	return length
}

func GetPrescsNumber(pharmacistId uint) int {
	// pharmacist, _ := profileSRV.GetPharmacist(pharmacistId)
	// prescs := pharmacist.Prescription

	length := 15
	return length
}
