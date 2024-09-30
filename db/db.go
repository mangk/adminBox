package db

import (
	"time"

	"github.com/mangk/adminBox/moduleRegister"

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

func init() {
	moduleRegister.ModuleAdd(db{})
}

type db struct{}

func (db) InitModule() {
	_dbList = make(map[string]*gorm.DB)
	for name, dbCfg := range config.DBCfg() {
		db, err := gorm.Open(dialectorBuild(dbCfg), &gorm.Config{
			Logger: logger.New(log.GormAdapter(), logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.LogLevel(dbCfg.LogMode),
				Colorful:                  false,
			}),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   dbCfg.Prefix,
				SingularTable: dbCfg.Singular,
			},
			DisableAutomaticPing:                     true,
			DisableForeignKeyConstraintWhenMigrating: true,
			IgnoreRelationshipsWhenMigrating:         true,
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
}

func DB(name ...string) *gorm.DB {
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
