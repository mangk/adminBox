package model

import (
	"fmt"

	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/log"
	"gorm.io/gorm"
)


type ModelWithoutID struct {
	ID int `json:"-" gorm:"type:int(11);primaryKey;comment:主键" uri:"id"`

	Cb int `json:"cb,omitempty" gorm:"type:int(11);comment:创建者"`
	Ub int `json:"ub,omitempty" gorm:"type:int(11);comment:更新者"`
	Db int `json:"-" gorm:"type:int(11);comment:删除者"`

	Ct *LocalTime     `json:"ct,omitempty" gorm:"type:datetime;default:CURRENT_TIMESTAMP;autoCreateTime;comment:创建时间"`
	Ut *LocalTime     `json:"ut,omitempty" gorm:"type:datetime;default:NULL ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime;comment:更新时间"`
	Dt gorm.DeletedAt `json:"-" gorm:"type:datetime;index;comment:删除时间"`

	CbName string `json:"cb_name,omitempty" gorm:"-:all"` // 创建者
	UbName string `json:"ub_name,omitempty" gorm:"-:all"` // 更新者
	DbName string `json:"db_name,omitempty" gorm:"-:all"` // 删除者
}

func (m ModelWithoutID) GetId() int {
	return m.ID
}

// 指定 driver 直接执行 sql 并将结果作为 map 返回
func (ModelWithoutID) SQLQuery(driver, sql string) (data []map[string]interface{}, err error) {
	log.Infof(fmt.Sprintf("SQLQuery: [%s] [%s]", driver, sql))
	mdb, _ := db.DB(driver).DB()
	dbObj, err := mdb.Prepare(sql)
	if err != nil {
		return
	}
	defer dbObj.Close()

	rows, err := dbObj.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	// 列的个数
	count := len(columns)

	// 返回值 Map切片
	data = make([]map[string]interface{}, 0)
	// 一条数据的各列的值（需要指定长度为列的个数，以便获取地址）
	values := make([]interface{}, count)
	// 一条数据的各列的值的地址
	valPointers := make([]interface{}, count)
	for rows.Next() {

		// 获取各列的值的地址
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}

		// 获取各列的值，放到对应的地址中
		rows.Scan(valPointers...)

		// 一条数据的Map (列名和值的键值对)
		entry := make(map[string]interface{})

		// Map 赋值
		for i, col := range columns {
			var v interface{}

			// 值复制给val(所以Scan时指定的地址可重复使用)
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				// 字符切片转为字符串
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		data = append(data, entry)
	}

	return
}
