package controllers

import (
	"github.com/nickgatej/Photo-Journey/views"
	"net/http"
)

type Users struct {
	// We use this struct to store all the templates we're going to need to render different pages on the user's controller
	Templates struct {
		New views.Template
	}
}

// This Method will be used to render Signup Page
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// We need a view to render
	u.Templates.New.Execute(w, nil) //
}
