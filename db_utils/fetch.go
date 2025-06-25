package db_utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ericyeungcode/caliber/common"
	"gorm.io/gorm"
)

//////////////////////////////
// Async result
//////////////////////////////

type AsyncDbResult struct {
	Data any
	Err  error
}

func AsyncFetchDb(querier *gorm.DB, outItems any) chan *AsyncDbResult {

	outC := make(chan *AsyncDbResult, 1)
	go func() {
		// do we need to defer close channel and why?
		// to tell `for ... := range outC` to finish loop
		//defer close(outC)

		defer common.ShowElapsedTime("AsyncFetchDb gofunc(): outItems type = %v", reflect.TypeOf(outItems))()

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
