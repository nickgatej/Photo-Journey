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

	// GET / via the StaticHandler closure
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	router.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	router.Get("/faq", controllers.StaticHandler(tpl))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
