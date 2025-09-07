package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonToValue(t *testing.T) {
	t.Run("struct value", func(t *testing.T) {
		input := `{"name":"Alice","age":30}`
		got, err := JsonToValue[MyStruct](input)

		require.NoError(t, err)
		require.Equal(t, MyStruct{Name: "Alice", Age: 30}, got)
	})

	t.Run("struct pointer", func(t *testing.T) {
		input := `{"name":"Bob","age":25}`
		got, err := JsonToValue[*MyStruct](input)

		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, &MyStruct{Name: "Bob", Age: 25}, got)
	})

	t.Run("map", func(t *testing.T) {
		input := `{"foo":"bar","num":42}`
		got, err := JsonToValue[map[string]any](input)

		require.NoError(t, err)
		require.Equal(t, map[string]any{
			"foo": "bar",
			"num": float64(42), // JSON numbers decode as float64
		}, got)
	})

	t.Run("slice", func(t *testing.T) {
		input := `[1,2,3]`
		got, err := JsonToValue[[]int](input)

		require.NoError(t, err)
		require.Equal(t, []int{1, 2, 3}, got)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		input := `{"name": "Charlie", "age": }`
		_, err := JsonToValue[MyStruct](input)

		require.Error(t, err)
	})
}

func Test2(t *testing.T) {

	// struct can't be converted to string
	t.Run("struct to string(expect failure)", func(t *testing.T) {
		input := `{"name":"Alice","age":30}`
		got, err := JsonToValue[string](input)
		t.Logf("got: %+v, err:%v\n", got, err)
		assert.Error(t, err)
	})

	// json string can be converted to go-string
	t.Run("json to value(string)", func(t *testing.T) {
		input := `"hello"`
		got, err := JsonToValue[string](input)
		t.Logf("got: %+v, err:%v\n", got, err)
		assert.NoError(t, err)
	})

}
