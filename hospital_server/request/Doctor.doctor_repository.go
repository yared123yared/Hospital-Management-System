package request





import  "github.com/yaredsolomon/webProgram1/hospital/entity"


// UserRepository specifies application user related database operations
type AppointmentRepository interface {
	Appointments() ([]entity.Doctor, []error)
	Appointment(id uint) (*entity.Doctor, []error)
	UpdateAppointment(user *entity.Doctor) (*entity.Doctor, []error)
	DeleteAppointment(id uint) (*entity.Doctor, []error)
	
}

