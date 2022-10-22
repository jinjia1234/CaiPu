package http

import (
	"CaiPu/config"
	"CaiPu/model"
	"CaiPu/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var (
	cfg *config.Configuration
	svc *service.Service
	db  *gorm.DB
)

func FoodSave(c *fiber.Ctx) error {
	//c.Accepts("application/json")
	body := new(struct {
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
	})
	err := c.BodyParser(body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = svc.CreateFood(model.Food{
		Name:      body.Caiming,
		BianMa:    body.Bianma,
		ChangJing: body.Changjing,
		LeiXing:   body.Leixing,
		FangShi:   body.Fangshi,
		ShiCai:    body.Shicai,
		TeSe:      body.Tese,
		DengJi:    body.Dengji,
		TuiJian:   body.Zhishu,
		BeiZhu:    body.Beizhu,
	})
	if err != nil {
		fmt.Println(err.Error())
		return c.SendString("失败")
	}
	return c.SendString("成功")
}

func FoodFind(c *fiber.Ctx) error {
	var (
		foods []model.Food
		err   error
	)
	foods, err = svc.FindFoods()
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(foods)
}
