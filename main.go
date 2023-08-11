package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
	"github.com/nickgatej/Photo-Journey/controllers"
	"github.com/nickgatej/Photo-Journey/migrations"
	"github.com/nickgatej/Photo-Journey/models"
	"github.com/nickgatej/Photo-Journey/templates"
	"github.com/nickgatej/Photo-Journey/views"
	"net/http"
	"os"
	"strconv"
)

type config struct {
	PSQL models.PostgresConfig
	SMTP models.SMTPConfig
	CSRF struct {
		Key    string
		Secure bool
	}
	Server struct {
		Address string
	}
}

func loadEnvConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	// TODO: Read the PSQL values from an ENV variable
	cfg.PSQL = models.DefaultPostgresConfig()

	// TODO: SMTP
	cfg.SMTP.Host = os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	cfg.SMTP.Port, err = strconv.Atoi(portStr)
	if err != nil {
		return cfg, err
	}
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")

	// TODO: Read the CSRF values from an ENV variable
	cfg.CSRF.Key = "9v2VMyJQtlF6xvYXUHIrnwbIMS1zPpu4"
	cfg.CSRF.Secure = false

	// TODO: Read the server values from an ENV variable
	cfg.Server.Address = ":3000"

	return cfg, nil
}

func main() {
	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	// Setup a database connection
	db, err := models.Open(cfg.PSQL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup our model services
	userService := &models.UserService{
		DB: db,
	}
	sessionService := &models.SessionService{
		DB: db,
	}
	pwResetService := &models.PasswordResetService{
		DB: db,
	}
	emailService := models.NewEmailService(cfg.SMTP)

	// Setup middlewares
	userMiddleware := controllers.UserMiddleware{
		SessionService: sessionService,
	}

	csrfMw := csrf.Protect(
		[]byte(cfg.CSRF.Key),
		// TODO: Fix this before deploying
		csrf.Secure(cfg.CSRF.Secure), // when set to true, this requires HTTPS connection
	)

	// Setup our controllers
	usersC := controllers.Users{
		UserService:          userService,
		SessionService:       sessionService,
		PasswordResetService: pwResetService,
		EmailService:         emailService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))
	usersC.Templates.ForgotPassword = views.Must(views.ParseFS(
		templates.FS,
		"forgot-pw.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.CheckYourEmail = views.Must(views.ParseFS(
		templates.FS,
		"check-your-email.gohtml", "tailwind.gohtml",
	))

	// Setup our router and routes
	router := chi.NewRouter()
	router.Use(csrfMw)
	router.Use(userMiddleware.SetUser)
	router.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	router.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	router.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	router.Get("/signup", usersC.New)
	router.Post("/signup", usersC.Create)
	router.Get("/signin", usersC.SignIn)
	router.Post("/signin", usersC.ProcessSignIn)
	router.Post("/signout", usersC.ProcessSignOut)
	router.Get("/forgot-pw", usersC.ForgotPassword)
	router.Post("/forgot-pw", usersC.ProcessForgotPassword)

	router.Route("/users/me", func(r chi.Router) {
		r.Use(userMiddleware.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Printf("Starting the server on %s...\n", cfg.Server.Address)
	err = http.ListenAndServe(cfg.Server.Address, router)
	if err != nil {
		panic(err)
	}
}
