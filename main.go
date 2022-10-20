package main

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/template/html"

func main() {

	app := fiber.New(fiber.Config{
		Views: html.New("./template", ".html"),
	})

	app.Static("/public", "/public")

	/*
		/主目录
		/menu 菜单编排
		/data 菜品资料库
		/chives	基础信息
		/basic 菜单存档
	*/

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, Worl!",
		})
	})

	app.Get("/menu", func(c *fiber.Ctx) error {
		return c.Render("menu", fiber.Map{
			"Title": "Hello, Worl!",
		})
	})

	app.Get("/data", func(c *fiber.Ctx) error {
		return c.Render("database", fiber.Map{
			"Title": "Hello, Worl!",
		})
	})
	app.Get("/chives", func(c *fiber.Ctx) error {
		return c.Render("archives", fiber.Map{
			"Title": "Hello, Worl!",
		})
	})
	app.Get("/basic", func(c *fiber.Ctx) error {
		return c.Render("basic", fiber.Map{
			"Title": "Hello, Worl!",
		})
	})
	app.Get("/select", func(c *fiber.Ctx) error {
		return c.Render("select", fiber.Map{
			"Title": "Hello, Worl!",
		})
	})

	err := app.Listen(":9668")
	if err != nil {
		return
	}

}
