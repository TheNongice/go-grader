package main

import (
	"fmt"

	"github.com/TheNongice/go-grader/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func introProg() {
	msg := `
	
   ██████      ██████  ██████   █████  ██████  ███████ ██████  
  ██          ██       ██   ██ ██   ██ ██   ██ ██      ██   ██ 
  ██          ██   ███ ██████  ███████ ██   ██ █████   ██████  
  ██          ██    ██ ██   ██ ██   ██ ██   ██ ██      ██   ██ 
   ██████      ██████  ██   ██ ██   ██ ██████  ███████ ██   ██ 
                                                                                                      
    C++ Grader (Judge_Server)
    GoLang Version -- Made in TH/A.
    Code by... @_ngixx's (TheNongice Wasawat)
    Contacts: ngixx@ngixx.in.th
	`
	fmt.Println(msg)
}

func main() {
	introProg()
	env := godotenv.Load()
	if env != nil {
		fmt.Println("Error loading .env file")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("grader f(x) is normally!")
	})

	router.JudgeService(app.Group("/api/"))
	// Manager
	router.ProblemSetService(app.Group("/api/prob/"))
	app.Listen(":8000")
}
