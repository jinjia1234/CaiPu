package dbsql

import (
	"CaiPu/model"
	"fmt"
)

// 查询多条数据示例
func QueryMultiRowDemo() []model.Foods {
	sqlStr := "select  * from food where id > ?"
	//var amp model.Foods
	var amp []model.Foods
	err := DbInit.Conn.Select(&amp, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return amp

	}
	return amp

	//var Q [...][11]string
	//var H = [11]string{}
	//for k, v := range amp {
	//	fmt.Println(v.LeiXing)
	//	H[0] = v.BianMa
	//	H[1] = v.Name
	//	H[2] = v.State
	//	H[3] = v.ChangJing
	//	H[4] = v.LeiXing
	//	H[5] = v.FangShi
	//	H[6] = v.ShiCai
	//	H[7] = v.TeSe
	//	H[8] = v.DengJi
	//	H[9] = v.TuiJian
	//	H[10] = v.BeiZhu
	//	Q[k] = H
	//}
	//return Q, nil

}
