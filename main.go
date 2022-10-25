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
	"github.com/tidwall/gjson"
)

var (
	cfg *config.Configuration
	svc *service.Service
)

func main() {
	//cmdLine := "dir"
	//cmd := exec.Command("cmd.exe", "/c", "start "+cmdLine)
	//if stdout, err := cmd.StdoutPipe(); err != nil { //获取输出对象，可以从该对象中读取输出结果
	//	log.Fatal(err)
	//}
	//
	//defer stdout.Close()                // 保证关闭输出流
	//if err := cmd.Start(); err != nil { // 运行命令
	//
	//	log.Fatal(err)
	//
	//}
	//err1 := cmd.Run()
	//fmt.Printf("%s, error:%v \n", cmdLine, err1)
	//return
	configPath := flag.String("config", "./config/config.toml", "config path")
	flag.Parse()
	cfg = config.Init(*configPath)
	svc = service.New(cfg)
	dbsql.New(cfg)

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

		res := dbsql.QueryMultiRowCP()
		//sum := len(res)
		var le []string
		for _, v := range res {
			le = append(le, v.Time+"("+v.ChangJing+"厅)")

		}
		fmt.Println(le)
		return c.Render("menu", fiber.Map{
			"cdlist": dbsql.QueryMultiRow(),
			"le":     le,
		})
	})
	app.Get("/deletels", func(c *fiber.Ctx) error {
		times := c.Query("id")
		dbsql.Deletels(times)
		return c.SendString("2")
	})

	app.Get("/menus", func(c *fiber.Ctx) error {
		id := c.Query("id")
		res := dbsql.QueryMultiRowCP()
		//sum := len(res)
		var le []string
		for _, v := range res {
			le = append(le, v.Time+"("+v.ChangJing+"厅)")

		}

		_, times, changjing, zaocan, wucan, wancan := dbsql.QueryRowCp(id)
		return c.Render("menus", fiber.Map{
			"times":     times,
			"changjing": changjing,
			"zaocan":    zaocan,
			"wucan":     wucan,
			"wancan":    wancan,
			"cdlist":    dbsql.QueryMultiRow(),
			"le":        le,
		})
	})

	app.Get("/data/*", func(c *fiber.Ctx) error {
		var (
			err   error
			foods []model.Food
		)
		where := make(map[string]string)
		state := c.Query("state")

		if state != "" {
			where["state"] = state
		}
		ChangJing := c.Query("ChangJing")

		if ChangJing != "" {
			where["ChangJing"] = ChangJing
		}
		LeiXing := c.Query("LeiXing")
		if LeiXing != "" {
			where["LeiXing"] = LeiXing
		}
		FangShi := c.Query("FangShi")
		if FangShi != "" {
			where["FangShi"] = FangShi
		}
		TeSe := c.Query("TeSe")
		if TeSe != "" {
			where["ChangJing"] = TeSe
		}
		DengJi := c.Query("DengJi")
		if DengJi != "" {
			where["DengJi"] = DengJi
		}
		TuiJian := c.Query("TuiJian")
		if TuiJian != "" {
			where["TuiJian"] = TuiJian
		}
		foods, err = svc.FindFoods()
		if err != nil {
			return c.JSON(err)
		}
		conun, countHunCai, countSuCai, counT, counZao := dbsql.Count()
		foodMeta := dbsql.QueryFoodMetaMultiRow()
		if err != nil {
			return err
		}
		if err != nil {
			return c.JSON(err)
		}
		return c.Render("database", fiber.Map{
			"ChangJing":   gjson.Parse(foodMeta[1].MetaValue).Array(),
			"LeiXing":     gjson.Parse(foodMeta[2].MetaValue).Array(),
			"FangShi":     gjson.Parse(foodMeta[3].MetaValue).Array(),
			"TeSe":        gjson.Parse(foodMeta[4].MetaValue).Array(),
			"DengJi":      gjson.Parse(foodMeta[5].MetaValue).Array(),
			"ZhiShu":      gjson.Parse(foodMeta[6].MetaValue).Array(),
			"Quer":        dbsql.QueryFoods(where),
			"Quer1":       foods,
			"countFood":   conun,
			"countHunCai": countHunCai,
			"countSuCai":  countSuCai,
			"counT":       counT,
			"counZao":     counZao,
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
		var res = true

		bol, _, _, _, _, _ := dbsql.QueryRowCps(body.Time, body.ChangJing)
		if bol {
			dbsql.Updatals(body)
		} else {
			res = dbsql.IastInserCP(body)
		}

		if res {
			return c.SendString("保存菜谱成功！")
		} else {
			return c.SendString("保存菜谱失败！")
		}

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
		//var tid, _ = config.JichuanInit()

		foodMeta := dbsql.QueryFoodMetaMultiRow()
		tid := config.Jichuantoarr(foodMeta)
		return c.Render("archives", fiber.Map{
			"Tdd": tid,
		})
	})

	app.Post("/chives", func(c *fiber.Ctx) error {
		body := new(model.FoodMetaSave)
		err := c.BodyParser(body)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		var res = true
		fmt.Println(body)
		dbsql.UpdataJc(body)
		if res {
			return c.SendString("保存菜谱成功！")
		} else {
			return c.SendString("保存菜谱失败！")
		}
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
