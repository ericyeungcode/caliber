package db_utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ericyeungcode/caliber"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//////////////////////////////
// Async result
//////////////////////////////

type AsyncDbResult struct {
	Data any
	Err  error
}

func GetDbUrl(user, pass, host, defaultDb string) string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC&timeout=10s&readTimeout=30s&writeTimeout=60s&multiStatements=true",
		user, pass, host, defaultDb)
}

func ConnectMysql(dsn string, maxConn int) *gorm.DB {
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

func AsyncFetchDb(querier *gorm.DB, outItems any) chan *AsyncDbResult {

	outC := make(chan *AsyncDbResult, 1)
	go func() {
		// do we need to defer close channel and why?
		// to tell `for ... := range outC` to finish loop
		//defer close(outC)

		defer caliber.ShowElapsedTime(fmt.Sprintf("AsyncFetchDb gofunc(): outItems type = %v", reflect.TypeOf(outItems)))()

		err := querier.Debug().Find(outItems).Error
		outC <- &AsyncDbResult{
			Data: outItems,
			Err:  err,
		}
	}()

	return outC
}

func RecvAsyncResult(resultC chan *AsyncDbResult, waitSeconds int) *AsyncDbResult {
	var result *AsyncDbResult
	select {
	case r := <-resultC:
		result = r
	case <-time.After(time.Second * time.Duration(waitSeconds)):
		result = &AsyncDbResult{
			Data: nil,
			Err:  fmt.Errorf("AsyncFetchDb timeout"),
		}
	}

	return result
}
