package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type myError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TestPair(t *testing.T) {
	p := NewPair("Dave", 40)
	t.Logf("first:%v, second:%v", p.First, p.Second)
	assert.Equal(t, "Dave", p.First)
	assert.Equal(t, 40, p.Second)

	x := NewPair(168, &myError{Code: 70003, Message: "general error"})
	t.Logf("count:%v, err:%v", x.First, x.Second)
	assert.Equal(t, 168, x.First)
	assert.Equal(t, 70003, x.Second.Code)
	assert.Equal(t, "general error", x.Second.Message)
}
