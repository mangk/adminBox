package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/log"
	"gorm.io/gorm"
)

type M interface {
	GetId() int
	TableName() string
}

type Model struct {
	ID int `json:"id" gorm:"type:int(11);primaryKey;comment:主键" uri:"id"`

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

func (m Model) GetId() int {
	return m.ID
}

// 指定 driver 直接执行 sql 并将结果作为 map 返回
func (Model) SQLQuery(driver, sql string) (data []map[string]interface{}, err error) {
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

func Find(tx *gorm.DB) (data []map[string]interface{}, err error) {
	rows, err := tx.Rows()
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

const (
	TimeZone   = "Asia/Shanghai"
	TimeFormat = "2006-01-02 15:04:05"
)

type LocalTime time.Time

func (t LocalTime) Now() *LocalTime {
	tt := LocalTime(time.Now())
	return &tt
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	loc, _ := time.LoadLocation(TimeZone)
	tTime := time.Time(t).In(loc)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(TimeFormat))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("cannot convert %v to timestamp", v)
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return nil
	}

	loc, _ := time.LoadLocation(TimeZone)
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), loc)
	if err != nil {
		return err
	}
	*t = LocalTime(now)
	return nil
}

func (lt LocalTime) GobEncode() ([]byte, error) {
	return time.Time(lt).MarshalBinary()
}

func (lt *LocalTime) GobDecode(data []byte) error {
	var t time.Time
	if err := t.UnmarshalBinary(data); err != nil {
		return err
	}
	*lt = LocalTime(t)
	return nil
}

func (t LocalTime) Format(layout ...string) string {
	if len(layout) > 0 {
		return time.Time(t).Format(layout[0])
	}
	loc, _ := time.LoadLocation(TimeZone)
	tTime := time.Time(t).In(loc)
	return tTime.Format(TimeFormat)
}

func (t LocalTime) Year() int {
	return time.Time(t).Year()
}

func (t LocalTime) Month() time.Month {
	return time.Time(t).Month()
}

func (t LocalTime) Day() int {
	return time.Time(t).Day()
}
