package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


func main(){

	app := fiber.New()

	// Method1 define simple api route >> ToDO ::: to  be moved to routes 
	app.Get("/abdel", func(c *fiber.Ctx) error {
		return c.SendString("Hi Abdel From Fiber")
	})

	app.Get("/", func(c *fiber.Ctx) error{
		return c.SendString("This is the main page, Abdel!")
	})


	// Method2: Define routes 
	router1 := app.Group("/router1")
	router1.Get("/asd",func(c *fiber.Ctx) error{
		return c.SendString("This is asd from router1")
	})
	router1.Get("/asd2",func(c *fiber.Ctx) error{
		return c.SendString("This is asd2 from router1")
	})

	// Method3: 
	app.Route("/router2", func(api fiber.Router){
		api.Get("/asd", func(c *fiber.Ctx) error{
			return c.SendString("This is asd from router2")
		})
		api.Get("/asd2", func(c *fiber.Ctx) error{
			return c.SendString("This is asd2 from router2")
		})
	})

	// Starting the server 
	// TODO

	/**
		- interfaces, structs, models
		- schemas, validation  
		- gorm
		- auth middleware , oauth2, jwt 
		- status, form 
		- file handling 
		- stream, coroutines 
	**/

	log.Fatal(app.Listen(":5000"))
}
