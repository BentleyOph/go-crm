package lead

import (
	"github.com/BentleyOph/go-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string	`json:"company"`
	Email   string	`json:"email"`
	Phone   int		`json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DB
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads) // send the leads as json
}
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}
func NewLead(c *fiber.Ctx) {
	db := database.DB
	lead := new(Lead)
	if err := c.BodyParser(lead);err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with given ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}
