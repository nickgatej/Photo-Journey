package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/nickgatej/Photo-Journey/controllers"
	"github.com/nickgatej/Photo-Journey/views"
	"net/http"
	"path/filepath"
)

func main() {
	router := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	router.Get("/", controllers.StaticHandler(tpl))

	router.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	router.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
