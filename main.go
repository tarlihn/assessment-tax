package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tarlihn/assessment-tax/database"
	"github.com/tarlihn/assessment-tax/tax"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func createDB() error {
	// Establish connection to the database
	p, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Db.Close()
	// Database connection successful
	log.Println("Connected to the database successfully!")

	createTb := `CREATE TABLE IF NOT EXISTS allowance ( id SERIAL PRIMARY KEY, personalDeduction INT);`
	_, err = p.Db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("create table success")

	row := p.Db.QueryRow("INSERT INTO  allowance (personalDeduction) values ($1)  RETURNING id", 60000)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return err
	}
	fmt.Println("insert todo success id : ", id)
	return nil
}
func main() {

	createDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	// User Routes
	// tax := e.Group("/tax")
	e.POST("/tax/calculation", tax.TaxController)

	//Admin Routes
	admin := e.Group("/admin")
	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		adminUsername := os.Getenv("ADMIN_USERNAME")
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		if adminUsername == "" {
			adminUsername = "adminTax" //for test
		}
		if adminPassword == "" {
			adminPassword = "admin!" //for test
		}
		return username == adminUsername && password == adminPassword, nil
	}))
	admin.POST("/deductions/personal", database.UpdatePersonalDeduction)

	// Retrieve port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portStr := fmt.Sprintf(":%s", port)
	// Start Echo server in a goroutine
	go func() {
		if err := e.Start(portStr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Error starting server: %v", err)
		}
	}()

	// Listen for interrupt signals (e.g., Ctrl+C) to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Shutdown Echo server
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Error shutting down server: %v", err)
	}
	fmt.Println("Server stopped gracefully")
}
