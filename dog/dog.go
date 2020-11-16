package dog

import (
        "fmt"
        "log"
        "time"
        "github.com/gofiber/fiber/v2"
        //"github.com/jinzhu/gorm"
       _"github.com/jinzhu/gorm/dialects/postgres"
        "github.com/Woof/database"
        "github.com/Woof/adoption"
       )

type Dog struct{

ID string `gorm:"primary_key"`
Name string
Breed string
Gender string
DOB time.Time
Age int
Info string
Image string
Location string
Adopters []adoption.Adopter
}


func GetDogs(c *fiber.Ctx) error{
  db := database.DBConn
  var dogs []Dog
  if err := db.Find(&dogs).Error; err != nil{
    log.Println(err)
    return c.Redirect("/error")
  }
  fmt.Println(dogs)
  return c.Render("dogs", fiber.Map{
    "dogs":dogs,
  })
}

func GetDog(c *fiber.Ctx) error{
  c.ClearCookie()
  db := database.DBConn

  var dog Dog
  dogid := c.Params("id")

  if err := db.Find(&dog, dogid).Error; err != nil {
    log.Println(err)
    return c.Redirect("/error")
  }

  return c.Render("dog",dog)

}
