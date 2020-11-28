package caliber

import (
	"fmt"
	"go.uber.org/atomic"
	"sync"
)

// concurrent counter
type ConcurrentCounter interface {
	Inc(key string)
	GetCounterMap() map[string]int64
}

////////////////////////////////////////
// AtomicCounter (syn.map + uber.atomic)
////////////////////////////////////////
type AtomicCounter struct {
	data sync.Map
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (ac *AtomicCounter) Inc_NotSafe(key string) {
	val, ok := ac.data.Load(key)
	var counter *atomic.Int64

	if !ok {
		counter = atomic.NewInt64(1)
		fmt.Printf("Add counter %v\n", key) // not thread-safe, output duplicated key added!!
		ac.data.Store(key, counter)
		return
	}

	counter = val.(*atomic.Int64)
	counter.Add(1)
	ac.data.Store(key, counter)
}

func (ac *AtomicCounter) Inc(key string) {
	val, loaded := ac.data.LoadOrStore(key, atomic.NewInt64(0))
	_ = loaded
	var counter = val.(*atomic.Int64)
	counter.Add(1)
	ac.data.Store(key, counter)
}

func (ac *AtomicCounter) GetCounterMap() map[string]int64 {
	result := make(map[string]int64)
	ac.data.Range(func(key, value interface{}) bool {
		result[key.(string)] = (value.(*atomic.Int64)).Load()
		return true
	})
	return result
}

////////////////////////////////////////
// MutexCounter (map + mutex)
////////////////////////////////////////
type MutexCounter struct {
	data  map[string]int64
	mutex sync.Mutex
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{
		data: make(map[string]int64),
	}
}

func (mc *MutexCounter) Inc(key string) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	mc.data[key] = mc.data[key] + 1
}

func (mc *MutexCounter) GetCounterMap() map[string]int64 {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	result := make(map[string]int64)
	for k, v := range mc.data {
		result[k] = v
	}
	return result
}
