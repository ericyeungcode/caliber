package caliber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	v, err := Div(10, 2)
	assert.NoError(t, err)
	assert.True(t, v == 5)

	v, err = Div(22.5, 2.5)
	assert.NoError(t, err)
	assert.True(t, v == 9)

	v, err = Div(3.14, 2)
	assert.NoError(t, err)
	assert.True(t, v == 1.57)

	_, err = Div(10, 0)
	assert.Error(t, err)

	v, err = Div(int64(15), int64(3))
	assert.NoError(t, err)
	assert.True(t, v == 5)

	v, err = Div(float32(7.5), float32(2.5))
	assert.NoError(t, err)
	assert.True(t, v == 3)
}
