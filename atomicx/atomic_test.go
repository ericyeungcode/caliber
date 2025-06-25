package atomicx

import (
	"testing"
)

func TestAtomicValue_Int(t *testing.T) {
	var a AtomicValue[int]

	// Test initial Load (should be zero)
	if val := a.Load(); val != 0 {
		t.Errorf("expected zero value, got %d", val)
	}

	// Store a value
	a.Store(42)
	if val := a.Load(); val != 42 {
		t.Errorf("expected 42, got %d", val)
	}

	// Swap value
	old := a.Swap(100)
	if old != 42 {
		t.Errorf("expected old value 42, got %d", old)
	}
	if val := a.Load(); val != 100 {
		t.Errorf("expected 100 after swap, got %d", val)
	}

	// CompareAndSwap success
	swapped := a.CompareAndSwap(100, 200)
	if !swapped {
		t.Errorf("expected CompareAndSwap to succeed")
	}
	if val := a.Load(); val != 200 {
		t.Errorf("expected 200 after CAS, got %d", val)
	}

	// CompareAndSwap fail
	swapped = a.CompareAndSwap(999, 300)
	if swapped {
		t.Errorf("expected CompareAndSwap to fail")
	}
	if val := a.Load(); val != 200 {
		t.Errorf("expected 200 to remain unchanged, got %d", val)
	}
}

func TestAtomicValue_Struct(t *testing.T) {
	type User struct {
		ID   int64
		Name string
	}

	var a AtomicValue[*User]

	// Initial load
	if val := a.Load(); val != nil {
		t.Errorf("expected nil, got %v", val)
	}

	user1 := &User{ID: 1, Name: "Alice"}
	user2 := &User{ID: 2, Name: "Bob"}

	a.Store(user1)
	if val := a.Load(); val != user1 {
		t.Errorf("expected %v, got %v", user1, val)
	}

	// Swap
	old := a.Swap(user2)
	if old != user1 {
		t.Errorf("expected old value %v, got %v", user1, old)
	}

	// CAS
	ok := a.CompareAndSwap(user2, user1)
	if !ok || a.Load() != user1 {
		t.Errorf("expected CAS to succeed and value to be %v", user1)
	}
}
