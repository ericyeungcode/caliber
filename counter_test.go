package caliber

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"sync"
	"testing"
)

func MakeKeys() []string {
	var keyList []string
	for idx := 0; idx < 10; idx++ {
		keyList = append(keyList, RandStringN(10))
	}
	return keyList
}

func runCounterTest(cntr ConcurrentCounter, t *testing.T) {
	totalCall := 100000

	defer ShowElapsedTime("totalCall:%v, counter:%v", totalCall, reflect.TypeOf(cntr))()

	keyList := MakeKeys()
	numKey := len(keyList)

	var wg sync.WaitGroup
	wg.Add(totalCall)
	ready := make(chan struct{})
	for i := 0; i < totalCall; i++ {

		key := keyList[rand.Intn(numKey)]
		go func() {
			<-ready
			cntr.Inc(key)
			wg.Done()
		}()
	}

	close(ready) // trigger multiple go routines to start working
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
	runCounterTest(NewAtomicCounter(), t)
}

func TestMutexCounter_Inc(t *testing.T) {
	runCounterTest(NewMutexCounter(), t)
}

func runCounterBench(cntr ConcurrentCounter, b *testing.B) {
	keyList := MakeKeys()
	numKey := len(keyList)

	for i := 0; i < b.N; i++ {
		key := keyList[rand.Intn(numKey)]
		cntr.Inc(key)

	}
}

// BenchmarkAtomicCounter_Inc-12    	10284536	       120 ns/op
// after removal of "ac.data.Store(key, counter)", op time from 237 ns/op => 120 ns/op
func BenchmarkAtomicCounter_Inc(b *testing.B) {
	runCounterBench(NewAtomicCounter(), b)
}

// on MacOs, CUP i7, RAM 16G, mutex counter works better..
//BenchmarkMutexCounter_Inc-12    	18508624	        67.5 ns/op
func BenchmarkMutexCounter_Inc(b *testing.B) {
	runCounterBench(NewMutexCounter(), b)
}
