package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	// We use this struct to store all the templates we're going to need to render different pages on the user's controller
	Templates struct {
		New Template
	}
}

// This Method will be used to render Signup Page
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// Take email from query params
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<p>Email: %s</p>", r.FormValue("email")) // name from HTML
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "<p>Password: %s</p>", r.FormValue("password"))
	if err != nil {
		return
	}
}
