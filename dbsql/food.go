package dbsql

import (
	"CaiPu/model"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

func QueryFoodMetaMultiRow() []model.FoodMeta {
	sqlStr := "select  * from food_meta"
	var s []model.FoodMeta
	err := DbInit.Conn.Select(&s, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return s
	}
	return s
}

func QueryFoods(where map[string]string) []model.Foods {
	whereStr := ""
	state := where["state"]
	if state != "" {
		whereStr += " state='" + state + "' and "
	}
	ChangJing := where["ChangJing"]
	if ChangJing != "" {
		whereStr += " ChangJing='" + ChangJing + "' and "
	}
	LeiXing := where["LeiXing"]
	if LeiXing != "" {
		whereStr += " LeiXing='" + LeiXing + "' and "
	}
	FangShi := where["FangShi"]
	if FangShi != "" {
		whereStr += " FangShi='" + FangShi + "' and "
	}
	TeSe := where["TeSe"]
	if TeSe != "" {
		whereStr += " TeSe='" + TeSe + "' and "
	}
	DengJi := where["DengJi"]
	if DengJi != "" {
		whereStr += " DengJi='" + DengJi + "' and "
	}
	TuiJian := where["TuiJian"]
	if TuiJian != "" {
		whereStr += " TuiJian='" + TuiJian + "' and "
	}
	sqlStr := "select  * from food where " + whereStr + " 1=1 "
	var amp []model.Foods
	err := DbInit.Conn.Select(&amp, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return amp
	}
	return amp
}

// 查询多条数据
func QueryMultiRow() []model.Foods {
	sqlStr := "select  * from food where id > ?"
	//var amp model.Foods
	var amp []model.Foods
	err := DbInit.Conn.Select(&amp, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return amp
	}
	return amp
}

func QueryMultiRowCP() []model.BaoCuncdS {
	sqlStr := "select  * from book where id > ?"
	//var amp model.Foods
	var amp []model.BaoCuncdS
	err := DbInit.Conn.Select(&amp, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return amp
	}
	return amp

}

func getsum(n string) int {
	sqlStr := "select  * from food where LeiXing = ?"
	var amp []model.Foods
	err := DbInit.Conn.Select(&amp, sqlStr, n)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return 0
	}
	zol := len(amp)
	return zol
}

func Count() (int, int, int, int, int) {
	sqlStr := "select  * from food where id > ?"
	//var amp model.Foods
	var amp []model.Foods
	err := DbInit.Conn.Select(&amp, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return 0, 0, 0, 0, 0

	}
	s := len(amp)
	return s, getsum("荤菜"), getsum("素菜"), getsum("汤"), getsum("早餐")
}

//添加数据
func IastInserId(c *model.UpdataBody) bool {
	sqlstr := "insert into food(name,state,BianMa,ChangJing,LeiXing,FangShi,ShiCai,TeSe,DengJi,TuiJian,BeiZhu,update_at)values (?,?,?,?,?,?,?,?,?,?,?,?)"
	r, err := DbInit.Conn.Exec(sqlstr, c.Caiming, c.Shiyong, c.Bianma, c.Changjing, c.Leixing, c.Fangshi, c.Shicai, c.Tese, c.Dengji, c.Zhishu, c.Beizhu, time.Now())
	if err != nil {
		fmt.Println("SQL insert err", err)
		return false
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("SQL LastInsertId ERR", err)
		return false
	}
	fmt.Println("insert succ:", id)

	return true

}

//保存菜谱
func IastInserCP(c *model.BaoCuncd) bool {

	var (
		changjing = c.ChangJing
		times     = c.Time
		zaocan    = c.ZaoCan
		wucan     = c.WuCan
		wancan    = c.WanCan
	)
	sqlstr := "insert into book(ChangJing,time,zaocan,wucan,wancan,update_at)values (?,?,?,?,?,?)"
	r, err := DbInit.Conn.Exec(sqlstr, changjing, times, zaocan, wucan, wancan, time.Now())

	if err != nil {
		fmt.Println("SQL insert err", err)
		return false
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("SQL LastInsertId ERR", err)
		return false
	}
	fmt.Println("insert succ:", id)
	return true
}

// 查询单条数据
func QueryRow(name string) bool {
	sqlStr := "select name from food where name=?"
	var u model.Names
	err := DbInit.Conn.Get(&u, sqlStr, name)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return false
	}
	fmt.Printf("name:%s ", u.Name)
	return true
}

func strtojson(v string) [9][5]string {
	//var s [9]string
	var x [9][5]string
	for i := 0; i < 9; i++ {
		var s [5]string
		for ii := 0; ii < 5; ii++ {

			value := gjson.Get(v, strconv.Itoa(ii)+"."+strconv.Itoa(i))
			//println(value.String())
			s[ii] = value.String()
		}
		x[i] = s
	}
	return x
}

// 查询单条数据(菜谱)
func QueryRowCp(name string) (bool, string, string, [9][5]string, [9][5]string, [9][5]string) {
	sqlStr := "select ChangJing ,time ,zaocan, wucan ,wancan from book where time=?"
	var u model.BaoCuncdCP
	err := DbInit.Conn.Get(&u, sqlStr, name)
	var (
		times     = ""
		changjing = ""
	)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return false, times, changjing, [9][5]string{}, [9][5]string{}, [9][5]string{}
	}
	times = u.Time
	changjing = u.ChangJing
	zaocan := strtojson(u.ZaoCan)
	wucan := strtojson(u.WuCan)
	wancan := strtojson(u.WanCan)

	fmt.Printf("name:%s ", u)

	return true, times, changjing, zaocan, wucan, wancan
}

// 查询单条数据(菜谱两个条件)
func QueryRowCps(name string, c string) (bool, string, string, [9][5]string, [9][5]string, [9][5]string) {
	sqlStr := "select ChangJing ,time ,zaocan, wucan ,wancan from book where time=? and ChangJing =?"
	var u model.BaoCuncdCP
	err := DbInit.Conn.Get(&u, sqlStr, name, c)
	var (
		times     = ""
		changjing = ""
	)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return false, times, changjing, [9][5]string{}, [9][5]string{}, [9][5]string{}
	}
	times = u.Time
	changjing = u.ChangJing
	zaocan := strtojson(u.ZaoCan)
	wucan := strtojson(u.WuCan)
	wancan := strtojson(u.WanCan)

	fmt.Printf("name:%s ", u)

	return true, times, changjing, zaocan, wucan, wancan
}

//更新数据
func Updata(c *model.UpdataBody) bool {
	sqlstr := "update food set name=?,state=?,BianMa=?,ChangJing=?,LeiXing=?,FangShi=?,ShiCai=?,TeSe=?,DengJi=?,TuiJian=?,BeiZhu=?,update_at=? where name = ?"
	res, err := DbInit.Conn.Exec(sqlstr, c.Caiming, c.Shiyong, c.Bianma, c.Changjing, c.Leixing, c.Fangshi, c.Shicai, c.Tese, c.Dengji, c.Zhishu, c.Beizhu, time.Now(), c.Caiming)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return false

	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return false
	}
	fmt.Println("update succ:", row)
	return true
}

//更新数据（菜谱）
func Updatals(c *model.BaoCuncd) bool {
	sqlstr := "update book set ChangJing=?,time=?,ZaoCan=?,wucan=?,wancan=? ,update_at=? where time = ? and ChangJing = ?"
	res, err := DbInit.Conn.Exec(sqlstr, c.ChangJing, c.Time, c.ZaoCan, c.WuCan, c.WanCan, time.Now(), c.Time, c.ChangJing)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return false

	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return false
	}
	fmt.Println("update succ:", row)
	return true

}

//更新数据（菜谱）
func UpdataJc(c *model.FoodMetaSave) bool {

	var w [1]string
	w[0] = ""

	//["ChangJing","LeiXing","FangShi","TeSe","DengJi","ZhiShu","Ting"]
	nr := c.DengJi
	sqlstr := "update food_meta set meta_value=? ,update_at=? where  id = ?"
	for i := 2; i < 9; i++ {
		if i == 2 {
			nr = c.ChangJing
		} else if i == 3 {
			nr = c.LeiXing
		} else if i == 4 {
			nr = c.FangShi
		} else if i == 5 {
			nr = c.TeSe
		} else if i == 6 {
			nr = c.DengJi
		} else if i == 7 {
			nr = c.ZhiShu
		} else if i == 8 {
			nr = c.Ting
		}
		bytes, _ := json.Marshal(nr)

		res, err := DbInit.Conn.Exec(sqlstr, string(bytes), time.Now(), i)
		if err != nil {
			fmt.Println("exec failed, ", err)

		}
		row, err := res.RowsAffected()
		fmt.Println("update succ:", row)

	}

	//if err != nil {
	//	fmt.Println("rows failed, ", err)
	//	return false
	//}
	//fmt.Println("update succ:", row)
	return true

}

//删除数据
func Delete(name string) bool {
	sqlstr := "delete from food where BianMa=?"
	res, err := DbInit.Conn.Exec(sqlstr, name)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return false
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return false
	}

	fmt.Println("delete succ: ", row)
	if row == 0 {
		return false
	} else {
		return true
	}

}

//删除数据(历史菜谱)
func Deletels(name string) bool {
	sqlstr := "delete from book where time=?"
	res, err := DbInit.Conn.Exec(sqlstr, name)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return false
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return false
	}

	fmt.Println("delete succ: ", row)
	if row == 0 {
		return false
	} else {
		return true
	}

}
