package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(res http.ResponseWriter, req *http.Request) user {
	//get Cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "sessions",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength //this is 30 seconds
	http.SetCookie(res, c)

	//if the user exists already, get the user:
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		// u = dbUsers[un]
		//when a new page is visited, refresh the lastActivity time to be now time
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	//refresh session
	c.MaxAge = sessionLength

	http.SetCookie(res, c)
	return ok
}

func CleanSessions() {
	fmt.Println("Before Clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("After Clean")
	showSessions()
}

func showSessions() {
	fmt.Println("*******")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
