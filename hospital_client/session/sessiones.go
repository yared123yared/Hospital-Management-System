package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

const (
	USERID  = "USERID"
	MESSAGE = "MESSAGE"
	ALERT   = "ALERT"
)

var Store = sessions.NewCookieStore([]byte("S3CR3TK3Y"))

func Message(message, alert string, r *http.Request, w http.ResponseWriter) {
	session, _ := Store.Get(r, "session")
	session.Values[MESSAGE] = message
	session.Values[ALERT] = alert
	session.Save(r, w)
}

func Flash(r *http.Request, w http.ResponseWriter) (string, string) {
	session, _ := Store.Get(r, "session")
	untypedMessage := session.Values["MESSAGE"]
	message, ok := untypedMessage.(string)
	if !ok {
		return "", ""
	}
	untypedAlert := session.Values["ALERT"]
	alert, ok := untypedAlert.(string)
	if !ok {
		return "", ""
	}
	delete(session.Values, "MESSAGE")
	delete(session.Values, "ALERT")
	session.Save(r, w)
	return message, alert
}

func SessionOptions(domain, path string, maxAge int, httpOnly bool) {
	Store.Options = &sessions.Options{
		Domain:   domain,
		Path:     path,
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
	}
}

func IsLogged(r *http.Request) (uint32, bool) {
	session, _ := Store.Get(r, "session")
	//fmt.Println("thise is the struct data")
	//fmt.Println(session.Values["USERID"])
	untypedUserId := session.Values["USERID"]
	userId, ok := untypedUserId.(uint32)
	if !ok {
		return 0, false
	}
	return userId, true
}
