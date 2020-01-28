package General

import "github.com/web1_group_project/hospital_server/entity"

type GeneralRepository interface {
	Pharmacists() ([]entity.Pharmacist, []error)
	Laboratorists() ([]entity.Laboratorist, []error)
	Users(id int, password string) (*entity.User, []error)
}

// RoleRepository speifies application user role related database operations
type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionRepository specifies logged in user session related database operations
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}

