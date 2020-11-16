package main

import (
        "fmt"
        "github.com/gofiber/fiber/v2"
        "github.com/gofiber/template/html"
        "github.com/Woof/database"
        "github.com/Woof/dog"
        "github.com/Woof/home"
        "github.com/Woof/adoption"
        "github.com/jinzhu/gorm"
        _"github.com/jinzhu/gorm/dialects/postgres"

       )

func initDatabase() {
  var err error
  database.DBConn, err = gorm.Open("#Database Details")
  if err != nil {
    panic("Failed to Connect to Database")
  }
  database.DBConn.AutoMigrate(&adoption.Adopter{})
  database.DBConn.AutoMigrate(&adoption.PlaceAdoption{})
  database.DBConn.AutoMigrate(&home.Doubt{})
  fmt.Println("Connection to database established.")
}

func setupRoutes(app *fiber.App){
  app.Get("/", home.GetContact)
  app.Post("/", home.PostContact)
  app.Get("/dog/:id", dog.GetDog)
  app.Get("/dogs", dog.GetDogs)
  app.Get("/adopter/:id", adoption.GetAdopter)
  app.Post("/adopter/:id", adoption.PostAdopter)
  app.Get("/placeAdoption", adoption.GetPlaceAdoption)
  app.Post("/placeAdoption", adoption.PostPlaceAdoption)
  app.Get("/thankyou", adoption.GetThankyou)
  app.Get("/error", adoption.GetError)
  app.Get("*", adoption.GetNotFound)
}

func main(){
  engine := html.New("./templates", ".html")
  app := fiber.New(fiber.Config{
    Views: engine,
  })
  app.Static("/", "./static")
  initDatabase()
  defer database.DBConn.Close()
  setupRoutes(app)

  app.Listen(":8080")

}
