package main

import (
	"CaiPu/config"
	"CaiPu/dbsql"
	"CaiPu/http"
	"CaiPu/model"
	"CaiPu/service"
	"flag"
	"fmt"
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
		var (
			err   error
			foods []model.Food
		)
		foods, err = svc.FindFoods()
		if err != nil {
			return c.JSON(err)
		}
		return c.Render("menu", fiber.Map{
			"cdlist": foods,
		})
	})

	app.Get("/menus", func(c *fiber.Ctx) error {
		id := c.Query("id")
		times, changjing, zaocan, wucan, wancan := dbsql.QueryRowCp(id)
		fmt.Println("早餐", zaocan)
		return c.Render("menus", fiber.Map{
			"times":     times,
			"changjing": changjing,
			"zaocan":    zaocan,
			"wucan":     wucan,
			"wancan":    wancan,
		})
	})

	app.Get("/data", func(c *fiber.Ctx) error {
		var (
			err   error
			foods []model.Food
		)
		foods, err = svc.FindFoods()
		if err != nil {
			return c.JSON(err)
		}
		//conun, countHunCai, countSuCai, err := svc.CountFood()
		conun, countHunCai, countSuCai := dbsql.Count()

		if err != nil {
			return err
		}
		if err != nil {
			return c.JSON(err)
		}
		return c.Render("database", fiber.Map{
			"ChangJing":   data[0],
			"LeiXing":     data[1],
			"FangShi":     data[2],
			"TeSe":        data[3],
			"DengJi":      data[4],
			"ZhiShu":      data[5],
			"Quer1":       dbsql.QueryMultiRow(),
			"Quer":        foods,
			"countFood":   conun,
			"countHunCai": countHunCai,
			"countSuCai":  countSuCai,
		})
	})

	app.Post("/iast", func(c *fiber.Ctx) error {
		//c.Accepts("application/json")
		body := new(model.UpdataBody)

		err := c.BodyParser(body)
		if err != nil {
			fmt.Println(err.Error())
			return err

		}
		res := dbsql.IastInserId(body)
		fmt.Println(body)
		if res {
			return c.SendString("添加菜品成功！")
		} else {
			return c.SendString("添加菜品失败！")
		}
	})
	app.Post("/upda", func(c *fiber.Ctx) error {
		//c.Accepts("application/json")
		body := new(model.UpdataBody)

		err := c.BodyParser(body)
		if err != nil {
			fmt.Println(err.Error())
			return err

		}
		res := dbsql.Updata(body)
		fmt.Println(body)
		if res {
			return c.SendString("更新菜品成功！")
		} else {
			return c.SendString("更新菜品失败！")
		}
	})

	app.Post("baocuncd", func(c *fiber.Ctx) error {
		body := new(model.BaoCuncd)
		err := c.BodyParser(body)
		if err != nil {
			fmt.Println(err.Error())
			return err

		}
		_ = dbsql.IastInserCP(body)
		fmt.Println(body.ZaoCan)

		return c.SendString("0")

	})

	app.Post("/chaxun", func(c *fiber.Ctx) error {
		body := new(model.UpdataBody)
		err := c.BodyParser(body)
		bol := dbsql.QueryRow(body.Caiming)
		if err != nil {
			return err
		}
		if bol {
			return c.SendString("0")

		} else {
			return c.SendString("1")
		}
	})

	app.Get("/lishi", func(c *fiber.Ctx) error {
		res := dbsql.QueryMultiRowCP()
		//sum := len(res)
		var le []string
		for _, v := range res {
			le = append(le, v.Time+"("+v.ChangJing+"厅)")

		}
		fmt.Println(le)

		return c.Render("lishi", fiber.Map{
			"le": le,
		})
	})

	app.Post("/delete", func(c *fiber.Ctx) error {
		body := new(model.UpdataBody)

		err := c.BodyParser(body)
		bol := dbsql.Delete(body.Bianma)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if bol {
			return c.SendString("菜品删除成功！")

		} else {
			return c.SendString("菜品删除失败！")
		}
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
			"id": c.Query("id"),
		})
	})
	app.Post("/food_save", http.FoodSave) // 菜保存
	app.Post("/food_find", http.FoodFind) // 菜查询
	err := app.Listen(":" + cfg.Port)
	if err != nil {
		return
	}

}
