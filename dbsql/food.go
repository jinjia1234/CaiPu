package dbsql

import (
	"CaiPu/model"
	"fmt"
	"time"
)

// 查询多条数据示例
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

func IastInserId(c *model.UpdataBody) bool {
	r, err := DbInit.Conn.Exec("insert into food(name,state,BianMa,ChangJing,LeiXing,FangShi,ShiCai,TeSe,DengJi,TuiJian,BeiZhu,update_at)values (?,?,?,?,?,?,?,?,?,?,?,?)",
		c.Caiming, c.Shiyong, c.Bianma, c.Changjing, c.Leixing, c.Fangshi, c.Shicai, c.Tese, c.Dengji, c.Zhishu, c.Beizhu, time.Now())
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

// 查询单条数据示例

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

func Delete(name string) bool {

	res, err := DbInit.Conn.Exec("delete from food where BianMa=?", name)
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
	return true

}
