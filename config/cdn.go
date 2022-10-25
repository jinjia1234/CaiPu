package config

import (
	"CaiPu/model"
	"fmt"
	"github.com/tidwall/gjson"
)

func JichuanInit() ([7][7]string, [7][]string) {
	//h1 := [7]string{"通用", "荤菜", "爆炒", "麻辣", "高", "★★★★★", "主厅"}

	/*
		使用场景  0
		菜品类型  1
		烹饪方式  2
		菜品特色  3
		成本等级  4
		推荐指数  5
	*/

	arrs := [7][7]string{}
	arr := [7][]string{
		{"通用", "主厅专用", "职工专用", "特制专属"},
		{"荤菜", "素菜", "汤", "早餐"},
		{"爆炒", "炝炒", "煲炖", "蒸熏", "煎炸", "凉拌", "鲜卤"},
		{"麻辣", "鲜嫩", "清淡", "原香"},
		{"高", "较高", "一般", "较低", "便宜"},
		{"★★★★★", "★★★★☆", "★★★☆☆", "★★☆☆☆", "★☆☆☆☆"},
		{"主厅", "职工厅"}}

	for i := 0; i < 7; i++ {
		for ii := 0; ii < len(arr[i]); ii++ {
			arrs[ii][i] = arr[i][ii]
		}
	}

	return arrs, arr

}

func Jichuantoarr(c []model.FoodMeta) [10][7]string {

	arr := [7][10]string{}
	for i := 1; i < 8; i++ {
		for k, v := range gjson.Parse(c[i].MetaValue).Array() {
			fmt.Println(v)
			arr[i-1][k] = v.String()
		}
	}
	arrs := [10][7]string{}
	for i := 0; i < 7; i++ {
		for ii := 0; ii < 10; ii++ {
			arrs[ii][i] = arr[i][ii]
		}
	}
	return arrs

}
