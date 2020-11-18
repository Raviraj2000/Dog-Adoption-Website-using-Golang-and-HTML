package home

import (
	"fmt"
	"github.com/Woof/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type Doubt struct {
	DoubtID string `gorm:"primary_key"`
	Name    string
	Email   string
	Message string
}

func GetContact(c *fiber.Ctx) error {
	c.ClearCookie()

	return c.Render("home", fiber.Map{})
}

func PostContact(c *fiber.Ctx) error {
	c.ClearCookie()
	db := database.DBConn

	doubt := new(Doubt)
	if err := c.BodyParser(doubt); err != nil {
		log.Println(err)
		return c.Redirect("/error")
	}

	var count int
	if err := db.Model(&Doubt{}).Count(&count).Error; err != nil {
		log.Println(err)
		c.Redirect("/error")
	}
	doubtid := strconv.Itoa(count + 1)

	doubt.DoubtID = doubtid
	if err := db.Create(&doubt).Error; err != nil {
		log.Println(err)
		return c.Redirect("/error")
	}
	fmt.Println(doubt)

	return c.Redirect("/thankyou")
}
