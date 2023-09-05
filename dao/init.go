package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  //禁止dateTime精度，mysql 5.6之前数据库不支持
		DontSupportRenameIndex:    true,  //禁止重命名索引，重命名索引要把索引先删除再重建，5.7之前不支持
		DontSupportRenameColumn:   true,  //用change重命名列，8之前数据库不支持
		SkipInitializeWithVersion: false, //
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)  //设置连接池
	sqlDb.SetMaxIdleConns(100) //打开连接数
	sqlDb.SetConnMaxIdleTime(time.Second * 30)
	_db = db

	//主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      //写操作
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, //读操作
		Policy:   dbresolver.RandomPolicy{},
	}))

	migration()
}

func NewDbClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
