package main

import (
	"fmt"

	"github.com/BentleyOph/go-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/BentleyOph/go-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}
func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "leads.db") // Openning a database, creating if not exist
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("COnnection est to db")
	database.DB.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(4545)
	defer database.DB.Close()

}
