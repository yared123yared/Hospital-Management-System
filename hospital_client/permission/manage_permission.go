package permission

import (
	"strings"
)

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/contact": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/login": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	"/about": permission{
		roles:   []string{"USER"},
		methods: []string{"GET"},
	},
	"/logout": permission{
		roles:   []string{"USER"},
		methods: []string{"POST"},
	},
	"/signup": permission{
		roles:   []string{"USER"},
		methods: []string{"GET", "POST"},
	},
	// <<<<<<<<<<<<<<<<<<<<<<<<<PHARMACIST>>>>>>>>>>>>>>
	"/pharmacist": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET"},
	},
	//Pharmacist Dashboard
	"/pharmacist/dashboard": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET"},
	},
	//Pharmacist Medicine
	"/pharmacist/medicine": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET"},
	},
	"/pharmacist/medicine/new": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET","POST"},
	},
	"/pharmacist/medicine/update": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET","PUT"},
	},
	"/pharmacist/medicine/delete": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET","DELETE"},
	},
	//Pharmacist profile
	"/pharmacist/profile": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET"},
	},
	"/pharmacist/profile/update": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET","PUT"},
	},
	//Pharmacist prescription
	"/pharmacist/prescription": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET"},
	},
	"/pharmacist/prescription/update": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET","PUT"},
	},
	"/pharmacist/prescription/delete": permission{
		roles:   []string{"PHARMACIST"},
		methods: []string{"GET","PUT"},
	},


	// <<<<<<<<<<<<<<<<<<<<<<<<<DOCTOR>>>>>>>>>>>>>>>>>>
	// doctor patient
	"/doctor": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","PUT","DELETE"},
	},
	"/doctor/patients": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","PUT","DELETE"},
	},
	"/doctor/patient/update": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","PUT"},
	},
	"/doctor/patient/delete": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"DELETE"},
	},
	"/doctor/patientNew": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","POST"},
	},

	// doctor appointment
	"/doctor/appointment": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","DELETE","PUT"},
	},
	"/doctor/appointmentNew": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","POST","PUT"},
	},

	//doctor prescribtions
	"/doctor/prescribtion": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","DELETE","PUT"},
	},
	"/doctor/prescribtionNew": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","POST","PUT"},
	},

	// doctor diagonosis
	"/doctor/diagonosis": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","DELETE","PUT"},
	},
	"/doctor/diagonosis/new": permission{
		roles:   []string{"DOCTOR"},
		methods: []string{"GET","POST"},
	},
	
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<PATIENT>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// patient roles
	"/patient": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET","PUT","DELETE","POST"},
	},
	"/patient/prescription": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET"},
	},
	"/patient/request": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET","PUT","DELETE","POST"},
	},
	"/patient/request/new": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET"},
	},
	
	"/patient/profile": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET","PUT","DELETE"},
	},
	"/patient/profile/update": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET","PUT","POST"},
	},
	"/patient/doctors": permission{
		roles:   []string{"PATIENT"},
		methods: []string{"GET"},
	},
	// <<<<<<<<<<<<<<<<<<<<<<ADMIN>>>>>>>>>>>>>>>>>>
	"/admin": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
}

// HasPermission checks if a given role has permission to access a given route for a given method
func HasPermission(path string, role string, method string) bool {
	if strings.HasPrefix(path, "/admin") {
		path = "/admin"
	}
	perm := authorities[path]
	checkedRole := checkRole(role, perm.roles)
	checkedMethod := checkMethod(method, perm.methods)
	if !checkedRole || !checkedMethod {
		return false
	}
	return true
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func checkMethod(method string, methods []string) bool {
	for _, m := range methods {
		if strings.ToUpper(m) == strings.ToUpper(method) {
			return true
		}
	}
	return false
}
