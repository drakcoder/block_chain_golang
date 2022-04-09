package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	bc := new_block_chain()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/getData", func(c *fiber.Ctx) error {
		return c.JSON(bc)
	})
	app.Get("/mine", func(c *fiber.Ctx) error {
		id := c.Query("id")
		idi, error := strconv.Atoi(id)
		if error != nil {
			return c.SendStatus(502)
		}
		status := bc.mine(idi)
		fmt.Println(idi)
		if !status {
			return c.SendStatus(502)
		}
		return c.SendStatus(200)
	})
	app.Get("/add_block", func(c *fiber.Ctx) error {
		return c.Render("add_block", fiber.Map{})
	})
	app.Post("/add_block", func(c *fiber.Ctx) error {
		payload := struct {
			Data string `json:"Data"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			fmt.Println(err)
			return c.SendStatus(502)
		}
		bc.add_block(payload.Data)
		return c.SendStatus(201)
	})

	log.Fatal(app.Listen(":8080"))
}
