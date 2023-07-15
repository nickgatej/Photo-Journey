package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/nickgatej/Photo-Journey/controllers"
	"github.com/nickgatej/Photo-Journey/templates"
	"github.com/nickgatej/Photo-Journey/views"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
	router.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
	router.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
