package main

import (
	_ "github.com/jackc/pgx/v4/stdlib" // import with side effects
	"github.com/joho/godotenv"
	"github.com/nickgatej/Photo-Journey/models"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	es.ForgotPassword("nick.gatej1@gmail.com", "https://photojourney.com/reset-pw?token=abc123")
}
