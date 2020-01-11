package pharmacist



import (

)


func GetMedsNumber(id uint) int {
	pharmacist, _ := GetPharmacist(id)
	medicines := pharmacist.Medicine
	length := len(medicines)
	return length
}

func GetPrescsNumber(pharmacistId uint) int {
	pharmacist, _ := GetPharmacist(pharmacistId)
	prescs := pharmacist.Prescription

	length := len(prescs)
	return length
}
