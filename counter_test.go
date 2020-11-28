package caliber

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"sync"
	"testing"
)

func runCounterTest(cntr ConcurrentCounter, t *testing.T) {
	totalCall := 100000

	defer ShowElapsedTime("totalCall:%v, counter:%v", totalCall, reflect.TypeOf(cntr))()

	var keyList []string
	for idx := 0; idx < 10; idx++ {
		keyList = append(keyList, RandStringN(10))
	}
	numKey := len(keyList)

	var wg sync.WaitGroup
	wg.Add(totalCall)
	for i := 0; i < totalCall; i++ {

		key := keyList[rand.Intn(numKey)]
		go func() {
			cntr.Inc(key)
			wg.Done()
		}()
	}

	wg.Wait()

	chkTotal := 0
	dataMap := cntr.GetCounterMap()
	for k, v := range dataMap {
		fmt.Printf("%v: %v\n", k, v)
		chkTotal += int(v)
	}
	assert.Equal(t, chkTotal, totalCall)
}

func TestAtomicCounter_Inc(t *testing.T) {
	runCounterTest(&AtomicCounter{}, t)
}

func TestMutexCounter_Inc(t *testing.T) {
	runCounterTest(NewMutexCounter(), t)
}
