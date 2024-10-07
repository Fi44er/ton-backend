package main

import (
	"log"
	"os"

	"github.com/Fi44er/ton-backend/controller"
	"github.com/Fi44er/ton-backend/database"
	"github.com/Fi44er/ton-backend/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type server struct {
	app  *fiber.App
	port string
}

func (s *server) allRoutes() {
	api := s.app.Group("/api")
	transaction := api.Group("/transaction")
	transaction.Post("/", controller.Create)
	transaction.Get("/", controller.GetAll)
	transaction.Get("/:wallet", controller.GetByWallet)
	transaction.Delete("/:id", controller.Delete)
}

func NewServer(port string) *server {
	s := &server{
		app:  fiber.New(),
		port: port,
	}

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return s
}

func (s *server) run() {
	s.allRoutes()
	log.Fatal(s.app.Listen(":" + s.port))
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.ConnectDb()
	if err != nil {
		log.Fatalf("Connection error to database: %v", err)
	}

	if err = db.Db.AutoMigrate(&model.TransactionData{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT is not set in the environment variables")
	}

	s := NewServer(port)
	s.run()
}
