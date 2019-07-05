package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Username string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{} //user ID, user
// var dbSessions = make(map[string]string) //session ID, user ID
var dbSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":7000", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
	}

	//if the user exist already, get rhe user
	var u user

	//get the username from the session
	if un, ok := dbSessions[c.Value]; ok {
		//get the user from the  username
		u = dbUsers[un]
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		//create the user with the above values:
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]

	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
