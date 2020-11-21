package adoption

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
	//"github.com/jinzhu/gorm"
	//_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/Woof/database"
)

type Dog struct {
	ID       string `gorm:"primary_key"`
	Name     string
	Breed    string
	Gender   string
	DOB      time.Time
	Age      int
	Info     string
	Image    string
	Location string
	Adopters Adopter
}

type Adopter struct {
	RequestID string `gorm:"primary_key"`
	Name      string
	Gender    string
	Address   string
	PhoneNo   string
	Email     string
	DogID     string
}

type PlaceAdoption struct {
	PlaceID   string `gorm:"primary_key"`
	Name      string
	Address   string
	PhoneNo   string
	Email     string
	DogName   string
	DogAge    int
	DogGender string
	Breed     string
	Image     string
	Adopters []Adopter
}

func GetAdopter(c *fiber.Ctx) error {
	c.ClearCookie()
	db := database.DBConn

	var dog Dog
	dogid := c.Params("id")

	if err := db.Find(&dog, dogid).Error; err != nil {
		log.Println(err)
		return c.Redirect("/error")
	}

	return c.Render("adopter", dog)

}

func PostAdopter(c *fiber.Ctx) error {
	c.ClearCookie()
	db := database.DBConn

	adopter := new(Adopter)
	if err := c.BodyParser(adopter); err != nil {
		log.Println(err)
		return c.Redirect("/error")
	}

	var count int
	db.Model(&Adopter{}).Count(&count)
	requestid := strconv.Itoa(count + 1)

	adopter.RequestID = requestid
	if err := db.Create(&adopter).Error; err != nil {
		log.Println(err)
		return c.Redirect("/error")
	}

	fmt.Println(adopter)
	return c.Redirect("/thankyou")
}

func GetPlaceAdoption(c *fiber.Ctx) error {
	c.ClearCookie()

	return c.Render("placeAdoption", fiber.Map{})
}

func PostPlaceAdoption(c *fiber.Ctx) error {
	db := database.DBConn

	placeAdoption := new(PlaceAdoption)
	if err := c.BodyParser(placeAdoption); err != nil {
		log.Println(err)
		return c.Redirect("/error")
	}

	var count int
	db.Model(&PlaceAdoption{}).Count(&count)
	placeid := strconv.Itoa(count + 1)

	placeAdoption.PlaceID = placeid
	if form, err := c.MultipartForm(); err == nil {
		if token := form.Value["token"]; len(token) > 0 {
			fmt.Println(token[0])
		}

		files := form.File["image"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			if err := c.SaveFile(file, fmt.Sprintf("adoption/ToPlace/"+placeAdoption.PlaceID+"_"+placeAdoption.DogName+".jpg")); err != nil {
				log.Println(err)
				return c.Redirect("/error")
			}
		}
	}
	placeAdoption.Image = fmt.Sprintf("dogImages/Allen" + placeAdoption.PlaceID + "_" + placeAdoption.DogName + ".jpg")
	fmt.Println(placeAdoption)


	if err := db.Create(&placeAdoption).Error; err != nil {
		log.Println(err)

		return c.Redirect("/error")
	}
	return c.Redirect("/thankyou")

}

func GetThankyou(c *fiber.Ctx) error {
	return c.Render("thankyou", fiber.Map{})
}

func GetError(c *fiber.Ctx) error {
	return c.Render("error", fiber.Map{})
}

func GetNotFound(c *fiber.Ctx) error {
	return c.Render("notFound", fiber.Map{})
}
