package Admin_Handler

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 

// 	"github.com/julienschmidt/httprouter"
// 	"github.com/webProj/Login"
// )

// // AdminCommentHandler handles comment related http requests
// type AllLoginHandler struct {
// 	loginService Login.LoginService
// }

// // NewAdminCommentHandler returns new AdminCommentHandler object
// func NewAllLoginHandler(cmntService Login.LoginService) *AllLoginHandler {
// 	return &AllLoginHandler{loginService: cmntService}
// }

// // GetComments handles GET /v1/admin/comments request
// func (alh *AllLoginHandler) UserLogin(w http.ResponseWriter,
// 	r *http.Request, _ httprouter.Params) {

// 	id, err := strconv.Atoi(r.FormValue("id"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	password := r.FormValue("password")
// 	profile, errs := alh.loginService.UsersLogin(id, password)
// 	if errs != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}
// 	fmt.Println("Profile")
// 	fmt.Println(profile)
// 	//	output, err := json.MarshalIndent(profile, "", "\t\t")
// 	// fmt.Println("output")

// 	// if err != nil {
// 	// 	w.Header().Set("Content-Type", "application/json")
// 	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	// 	return
// 	// }

// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.Write(output)

// 	// fmt.Println(output)
// 	// return

// }
