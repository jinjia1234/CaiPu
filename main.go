package main

import (
	"CaiPu/config"
	"CaiPu/service"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var (
	cfg *config.Configuration
	//svc *service.Service
)

type UpdataBody struct {
	Bianma    string `json:"bianma"`
	Caiming   string `json:"caiming"`
	Shiyong   string `json:"shiyong"`
	Changjing string `json:"changjing"`
	Leixing   string `json:"leixing"`
	Fangshi   string `json:"fangshi"`
	Shicai    string `json:"shicai"`
	Tese      string `json:"tese"`
	Dengji    string `json:"dengji"`
	Zhishu    string `json:"zhishu"`
	Beizhu    string `json:"beizhu"`
}

func main() {
	configPath := flag.String("config", "./config/config.toml", "config path")
	flag.Parse()
	cfg = config.Init(*configPath)
	service.New(cfg)

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

		return c.Render("database", fiber.Map{
			"ChangJing": data[0],
			"LeiXing":   data[1],
			"FangShi":   data[2],
			"TeSe":      data[3],
			"DengJi":    data[4],
			"ZhiShu":    data[5],
			"Quer":      service.QueryMultiRowDemo(),
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

	app.Post("/updata", func(c *fiber.Ctx) error {
		//c.Accepts("application/json")
		body := new(UpdataBody)

		err := c.BodyParser(body)
		if err != nil {
			fmt.Println(err.Error())
			return err

		}

		fmt.Println(body.Bianma)

		return c.SendString("成功")

	})

	err := app.Listen(":" + cfg.Port)
	if err != nil {
		return
	}

}
