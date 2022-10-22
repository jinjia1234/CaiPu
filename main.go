package main

import (
	"CaiPu/config"
	"CaiPu/dbsql"
	"CaiPu/http"
	"CaiPu/model"
	"CaiPu/service"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var (
	cfg *config.Configuration
	svc *service.Service
)

func main() {
	configPath := flag.String("config", "./config/config.toml", "config path")
	flag.Parse()
	cfg = config.Init(*configPath)
	svc = service.New(cfg)
	dbsql.New(cfg)

	var tid, data = config.JichuanInit()

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
		/select
	*/
	app.Get("/", func(c *fiber.Ctx) error {
		//_, m, err := svc.FindFoodById(40)
		//if err != nil {
		//	log.Errorf("svc.FindFoodById(%d) error(%v)", 40, err)
		//}
		//println(m.Name)
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
		var (
			foods []model.Food
			err   error
		)
		foods, err = svc.FindFoods()
		if err != nil {
			return c.JSON(err)
		}

		return c.Render("database", fiber.Map{
			"ChangJing": data[0],
			"LeiXing":   data[1],
			"FangShi":   data[2],
			"TeSe":      data[3],
			"DengJi":    data[4],
			"ZhiShu":    data[5],
			"Quer":      dbsql.QueryMultiRowDemo(),
			"Quer1":     foods,
		})
	})

	app.Get("/chives", func(c *fiber.Ctx) error {
		return c.Render("archives", fiber.Map{
			"Tdd": tid,
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
	app.Post("/food_save", http.FoodSave) // 菜保存
	app.Post("/food_find", http.FoodFind) // 菜查询
	err := app.Listen(":" + cfg.Port)
	if err != nil {
		return
	}

}
