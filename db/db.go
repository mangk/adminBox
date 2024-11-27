package db

import (
	"sync"
	"time"

	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _dbList map[string]*gorm.DB
var _dbInitOnce sync.Once

func DB(name ...string) *gorm.DB {
	_dbInitOnce.Do(func() {
		_dbList = make(map[string]*gorm.DB)
		for name, dbCfg := range config.DBCfg() {
			db, err := gorm.Open(dialectorBuild(dbCfg), &gorm.Config{
				Logger: logger.New(log.GormAdapter(), logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					IgnoreRecordNotFoundError: true,
					LogLevel:                  logger.LogLevel(dbCfg.LogMode),
					Colorful:                  false,
				}),
				NamingStrategy: &prefixNamingStrategy{
					NamingStrategy: schema.NamingStrategy{
						TablePrefix:   dbCfg.Prefix,
						SingularTable: dbCfg.Singular,
					},
					TablePrefix: dbCfg.Prefix,
				},
				DisableAutomaticPing:                     true,
				DisableForeignKeyConstraintWhenMigrating: true,
				// IgnoreRelationshipsWhenMigrating: true,
			})

			if err != nil {
				log.Panic("db init error", "name", name, "err", err)
			}

			conn, _ := db.DB()
			// 设置空闲连接池中连接的最大数量
			conn.SetMaxIdleConns(dbCfg.MaxIdleConn)
			// 设置打开数据库连接的最大数量
			conn.SetMaxOpenConns(dbCfg.MaxOpenConn)
			// 设置了连接可复用的最大时间
			conn.SetConnMaxLifetime(time.Hour)
			_dbList[name] = db
		}
	})

	dbName := "default"
	if len(name) == 1 {
		dbName = name[0]
	}

	if db, ok := _dbList[dbName]; ok {
		return db
	}

	log.Panic("db driver undefind", "name", dbName)
	return nil
}

func dialectorBuild(g config.DB) gorm.Dialector {
	switch g.Driver {
	case "mysql":
		return mysql.Open(g.Username + ":" + g.Password + "@tcp(" + g.Path + ":" + g.Port + ")/" + g.Dbname + "?" + g.Config)
	case "mssql":
		return sqlserver.Open("sqlserver://" + g.Username + ":" + g.Password + "@" + g.Path + ":" + g.Port + "?database=" + g.Dbname + "&encrypt=disable")
	case "pgsql":
		return postgres.Open("host=" + g.Path + " user=" + g.Username + " password=" + g.Password + " dbname=" + g.Dbname + " port=" + g.Port + " " + g.Config)
		//case "oracle":
		//	return "oracle://" + g.Username + ":" + g.Password + "@" + g.Path + ":" + g.Port + "/" + g.Dbname + "?" + g.Config
	}
	log.Panic("undefined db driver", "driver", g.Driver)
	return nil
}

// 自定义 NamingStrategy，结合 TablePrefix 和 TableName 方法
type prefixNamingStrategy struct {
	schema.NamingStrategy
	TablePrefix string
}

// func (f *prefixNamingStrategy) SchemaName(table string) string {
// 	return f.SchemaName(table)
// }

// func (f *prefixNamingStrategy) ColumnName(table string, column string) string {
// 	return f.ColumnName(table, column)
// }

// func (f *prefixNamingStrategy) JoinTableName(joinTable string) string {
// 	panic("not implemented") // TODO: Implement
// }

// func (f *prefixNamingStrategy) RelationshipFKName(_ schema.Relationship) string {
// 	panic("not implemented") // TODO: Implement
// }

// func (f *prefixNamingStrategy) CheckerName(table string, column string) string {
// 	return f.CheckerName(table, column)
// }

// func (f *prefixNamingStrategy) IndexName(table string, column string) string {
// 	panic("not implemented") // TODO: Implement
// }

// func (f *prefixNamingStrategy) UniqueName(table string, column string) string {
// 	panic("not implemented") // TODO: Implement
// }

// 实现 TableName 方法，使其在指定的表名前添加前缀
func (ns *prefixNamingStrategy) TableName(table string) string {
	return ns.TablePrefix + table
}

// 自定义 BeforeQuery Hook 添加表前缀
func (ns *prefixNamingStrategy) BeforeQuery(tx *gorm.DB) (err error) {
	tx.Statement.Table = ns.TablePrefix + tx.Statement.Table
	return nil
}
