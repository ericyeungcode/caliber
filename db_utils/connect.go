package db_utils

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDbUrl(user, pass, host, defaultDb string) string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC&timeout=10s&readTimeout=30s&writeTimeout=60s&multiStatements=true",
		user, pass, host, defaultDb)
}

func ConnectMysqlDsn(dsn string, maxConn int) *gorm.DB {
	connectConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		//Logger: logger.New(),
	}

	var err error
	db, err := gorm.Open(mysql.Open(dsn), connectConfig)

	log.Infof("db = %v, err=%+v", db, err)
	if err != nil {
		log.Panic(err)
	}

	rawDb, err := db.DB()
	if err != nil {
		log.Panic(err)
	}

	rawDb.SetMaxIdleConns(maxConn)
	rawDb.SetMaxOpenConns(maxConn)
	rawDb.SetConnMaxLifetime(6 * time.Hour)
	rawDb.SetConnMaxIdleTime(1 * time.Hour)

	// 默认的 db.Debug():
	// 打印 SELECT count(*) FROM t_order WHERE (user_id = '8088')

	// 加上 db.SetLogger(log.StandardLogger()) 后
	// 打印 SELECT count(*) FROM t_order WHERE (user_id = ?) [8088]

	//if !useDefaultLog {
	//	db.set(log.StandardLogger())
	//}

	return db
}

func ConnectMysql(user, pass, host, defaultDb string, maxConn int) *gorm.DB {
	return ConnectMysqlDsn(GetDbUrl(user, pass, host, defaultDb), maxConn)
}
