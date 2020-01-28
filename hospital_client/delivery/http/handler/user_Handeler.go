package handler

import (
	"context"
	"html/template"
	"net/http"
	"net/url"

	// "strconv"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	//"github.com/betsegawlemma/restaurant/user"

	"github.com/web1_group_project/hospital_client/data"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/form"
	"github.com/web1_group_project/hospital_client/permission"
	"github.com/web1_group_project/hospital_client/rtoken"
	"github.com/web1_group_project/hospital_client/session"
)

// UserHandler handler handles user related requests
type UserHandler struct {
	tmpl *template.Template
	//userService    user.UserService
	//sessionService user.SessionService
	userSess     *entity.Session
	LoggedInUser *entity.User
	// userRole       user.RoleService
	csrfSignKey []byte
}

type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

// NewUserHandler returns new UserHandler object
func NewUserHandler(t *template.Template, usrSess *entity.Session, csKey []byte) *UserHandler {
	return &UserHandler{tmpl: t, userSess: usrSess, csrfSignKey: csKey}
}

// Authenticated checks if a user is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Authorized checks if a user has proper authority to access a give route
func (uh *UserHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.LoggedInUser == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		//roles, errs := uh.userService.UserRoles(uh.loggedInUser)
		roles, errs := data.UserRoles(uh.LoggedInUser)
		if errs != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			fmt.Println("ERRRRRRRRRRRRRRRRR")
			fmt.Println(r.FormValue("_csrf"))
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// Login hanldes the GET/POST /login requests
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("thise is the login method")
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		fmt.Println("thise is the get methd")
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}

		uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
		fmt.Println("i am about to excute the template")
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		//usr, errs := uh.userService.UserByEmail(r.FormValue("email"))
		usr, errs := data.UserByEmail(r.FormValue("email"))
		fmt.Println("thise is the user data")
		fmt.Println(usr)
		if errs != nil {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("password")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		// if(usr.Password!=r.FormValue("password")){
		// 	loginForm.VErrors.Add("generic", "Your email address or password is wrong")
		//  	uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
		//  	return
		// }

		uh.LoggedInUser = usr
		claims := rtoken.Claims(usr.Email, uh.userSess.Expires)
		session.Create(claims, uh.userSess.UUID, uh.userSess.SigningKey, w)
		// newSess, errs := uh.sessionService.StoreSession(uh.userSess)
		newSess, errs := data.StoreSession(uh.userSess)
		if errs != nil {
			loginForm.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		uh.userSess = newSess
		// roles, _ := uh.userService.UserRoles(usr)
		roles, errs := data.UserRoles(usr)
		fmt.Println("thise is the role data")
		fmt.Println(roles)

		if uh.checkAdmin(roles) {

			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		if uh.checkDoctor(roles) {
			http.Redirect(w, r, "/doctor", http.StatusSeeOther)
			return
		} else if uh.checkPharmacist(roles) {
			http.Redirect(w, r, "/pharmacist", http.StatusSeeOther)
			return
		} else if uh.checkLaboratorist(roles) {
			http.Redirect(w, r, "/laboratorist", http.StatusSeeOther)
			return
		} else if uh.checkPatient(roles) {

			http.Redirect(w, r, "/patient", http.StatusSeeOther)
			return
		}
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

// Logout hanldes the POST /logout requests
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	session.Remove(userSess.UUID, w)
	//uh.sessionService.DeleteSession(userSess.UUID)
	data.DeleteSession(userSess.UUID)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

/*
// Signup hanldes the GET/POST /signup requests
func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		singnUpForm.Required("fullname", "email", "password", "confirmpassword")
		singnUpForm.MatchesPattern("email", form.EmailRX)
		singnUpForm.MatchesPattern("phone", form.PhoneRX)
		singnUpForm.MinLength("password", 8)
		singnUpForm.PasswordMatches("password", "confirmpassword")
		singnUpForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !singnUpForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		if pExists {
			singnUpForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}
		eExists := uh.userService.EmailExists(r.FormValue("email"))
		if eExists {
			singnUpForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			singnUpForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		role, errs := uh.userRole.RoleByName("USER")

		if len(errs) > 0 {
			singnUpForm.VErrors.Add("role", "could not assign role to the user")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		user := &entity.User{
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   role.ID,
		}
		_, errs = uh.userService.StoreUser(user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
*/
func (uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		return false
	}
	userSess := uh.userSess
	c, err := r.Cookie(userSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

/*
// AdminUsers handles Get /admin/users request
func (uh *UserHandler) AdminPatients(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	// users, errs := uh.userService.Users()
	users, errs :=data.Patients()

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Users   []entity.Petient
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		Users:   users,
		CSRF:    token,
	}
	uh.tmpl.ExecuteTemplate(w, "admin.manage.patients.layout", tmplData)
}

// AdminUsersNew handles GET/POST /admin/users/new request
func (uh *UserHandler) AdminPatientsNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		//roles, errs := uh.userRole.Roles()
		roles, errs := data.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		accountForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			//Roles   []entity.Role
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			//Roles:   roles,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate the form contents
		accountForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		accountForm.Required("fullname", "email", "password", "confirmpassword")
		accountForm.MatchesPattern("email", form.EmailRX)
		accountForm.MatchesPattern("phone", form.PhoneRX)
		accountForm.MinLength("password", 8)
		accountForm.PasswordMatches("password", "confirmpassword")
		accountForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !accountForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		if pExists {
			accountForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		eExists := uh.userService.EmailExists(r.FormValue("email"))
		if eExists {
			accountForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			accountForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			accountForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		user := &entity.User{
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   uint(roleID),
		}
		_, errs := uh.userService.StoreUser(user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

// AdminUsersUpdate handles GET/POST /admin/users/update?id={id} request
func (uh *UserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		user, errs := uh.userService.User(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		roles, errs := uh.userRole.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		role, errs := uh.userRole.Role(user.RoleID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("userid", idRaw)
		values.Add("fullname", user.FullName)
		values.Add("email", user.Email)
		values.Add("role", string(user.RoleID))
		values.Add("phone", user.Phone)
		values.Add("rolename", role.Name)

		upAccForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entity.Role
			User    *entity.User
			CSRF    string
		}{
			Values:  values,
			VErrors: form.ValidationErrors{},
			Roles:   roles,
			User:    user,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		upAccForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		upAccForm.Required("fullname", "email", "phone")
		upAccForm.MatchesPattern("email", form.EmailRX)
		upAccForm.MatchesPattern("phone", form.PhoneRX)
		upAccForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !upAccForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		userID := r.FormValue("userid")
		uid, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		user, errs := uh.userService.User(uint(uid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		eExists := uh.userService.EmailExists(r.FormValue("email"))
		if (user.Email != r.FormValue("email")) && eExists {
			upAccForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		pExists := uh.userService.PhoneExists(r.FormValue("phone"))

		if (user.Phone != r.FormValue("phone")) && pExists {
			upAccForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			upAccForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		usr := &entity.User{
			ID:       user.ID,
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: user.Password,
			RoleID:   uint(roleID),
		}
		_, errs = uh.userService.UpdateUser(usr)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

// AdminUsersDelete handles Delete /admin/users/delete?id={id} request
func (uh *UserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := uh.userService.DeleteUser(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

*/
func (uh *UserHandler) checkAdmin(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("Admin") {
			return true
		}
	}
	return false
}
func (uh *UserHandler) checkDoctor(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("doctor") {
			return true
		}
	}
	return false
}
func (uh *UserHandler) checkPharmacist(rs []entity.Role) bool {
	fmt.Println("thise is the pharmacist method")
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("PHARMACIST") {
			fmt.Println("the user is pharmacist")
			return true
		}
	}
	return false
}
func (uh *UserHandler) checkLaboratorist(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("laboratorist") {
			return true
		}
	}
	return false
}
func (uh *UserHandler) checkPatient(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("patient") {
			return true
		}
	}
	return false
}
