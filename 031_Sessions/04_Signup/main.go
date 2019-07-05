package main

import (
	"html/template"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

// var dbSessions = map[string]string{}
var dbSessions = map[string]session{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	//this is like a seeder file
	dbUsers["victorsteven"] = user{"victorsteven", bs, "Victor", "Steven", "007"}
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(res, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//process the form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		//username taken:
		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}

		//create session:
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}

		//store user in dbUsers:
		//turn the password into a slice of byte
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, r}
		//storing this user in the database of users
		dbUsers[un] = u

		//redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		//is there a username:
		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")
	//delete the session
	delete(dbSessions, c.Value)
	//remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, //delete the cookie now
	}
	http.SetCookie(res, c)

	//clean up dbSessions
	//substract the time our program started from the current time, and checked if it is greated than 30 seconds, if yes, remove the session
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go CleanSessions()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
