package caliber

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//////////////////////////////
// Async result
//////////////////////////////

type AsyncDbResult struct {
	Data interface{}
	Err  error
}

func GetDbUrl(user, pass, host, defaultDb string) string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC&timeout=10s&readTimeout=30s&writeTimeout=60s&multiStatements=true",
		user, pass, host, defaultDb)
}

func GetGormDbV1(user, pass, host, defaultDb string) *gorm.DB {
	return SetupDb(GetDbUrl(user, pass, host, defaultDb))
}

func SetupDb(url string) *gorm.DB {
	return SetupDbWithLog(url, true)
}

func SetupDbWithLog(url string, useDefaultLog bool) *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	var err error
	db, err := gorm.Open("mysql", url)
	log.Infof("db = %v, err=%+v", db, err)
	if err != nil {
		log.Panic(err)
	}

	db.SingularTable(true)

	// 默认的 db.Debug():
	// 打印 SELECT count(*) FROM t_order WHERE (user_id = '8088')

	// 加上 db.SetLogger(log.StandardLogger()) 后
	// 打印 SELECT count(*) FROM t_order WHERE (user_id = ?) [8088]

	if !useDefaultLog {
		db.SetLogger(log.StandardLogger())
	}

	return db
}

func AsyncFetchDb(querier *gorm.DB, outItems interface{}) chan *AsyncDbResult {

	outC := make(chan *AsyncDbResult, 1)
	go func() {
		// do we need to defer close channel and why
		//defer close(outC)

		defer ShowElapsedTime(fmt.Sprintf("AsyncFetchDb gofunc(): outItems type = %v", reflect.TypeOf(outItems)))()

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
