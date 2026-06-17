package numx

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAToInt(t *testing.T) {
	v, err := AToInt("42")
	assert.NoError(t, err)
	assert.Equal(t, 42, v)

	v, err = AToInt("-7")
	assert.NoError(t, err)
	assert.Equal(t, -7, v)

	_, err = AToInt("abc")
	assert.Error(t, err)

	_, err = AToInt("")
	assert.Error(t, err)
}

func TestAToIntEmpty(t *testing.T) {
	assert.Equal(t, 42, AToIntEmpty("42"))
	assert.Equal(t, 0, AToIntEmpty(""))
	assert.Equal(t, 0, AToIntEmpty("abc"))
}

func TestMustAToInt(t *testing.T) {
	assert.Equal(t, 42, MustAToInt("42"))
	assert.Panics(t, func() { MustAToInt("abc") })
}

func TestAToUInt(t *testing.T) {
	v, err := AToUInt("42")
	assert.NoError(t, err)
	assert.Equal(t, uint(42), v)

	_, err = AToUInt("-1")
	assert.Error(t, err)

	_, err = AToUInt("abc")
	assert.Error(t, err)
}

func TestAToUIntEmpty(t *testing.T) {
	assert.Equal(t, uint(42), AToUIntEmpty("42"))
	assert.Equal(t, uint(0), AToUIntEmpty("-1"))
	assert.Equal(t, uint(0), AToUIntEmpty("abc"))
}

func TestMustAToUInt(t *testing.T) {
	assert.Equal(t, uint(42), MustAToUInt("42"))
	assert.Panics(t, func() { MustAToUInt("-1") })
}

func TestAToInt64(t *testing.T) {
	v, err := AToInt64("9223372036854775807")
	assert.NoError(t, err)
	assert.Equal(t, int64(9223372036854775807), v)

	v, err = AToInt64("-42")
	assert.NoError(t, err)
	assert.Equal(t, int64(-42), v)

	_, err = AToInt64("abc")
	assert.Error(t, err)
}

func TestAToInt64Empty(t *testing.T) {
	assert.Equal(t, int64(42), AToInt64Empty("42"))
	assert.Equal(t, int64(0), AToInt64Empty(""))
	assert.Equal(t, int64(0), AToInt64Empty("abc"))
}

func TestMustAToInt64(t *testing.T) {
	assert.Equal(t, int64(42), MustAToInt64("42"))
	assert.Panics(t, func() { MustAToInt64("abc") })
}

func TestAToUInt64(t *testing.T) {
	v, err := AToUInt64("18446744073709551615")
	assert.NoError(t, err)
	assert.Equal(t, uint64(18446744073709551615), v)

	_, err = AToUInt64("-1")
	assert.Error(t, err)

	_, err = AToUInt64("abc")
	assert.Error(t, err)
}

func TestAToUInt64Empty(t *testing.T) {
	assert.Equal(t, uint64(42), AToUInt64Empty("42"))
	assert.Equal(t, uint64(0), AToUInt64Empty("-1"))
	assert.Equal(t, uint64(0), AToUInt64Empty("abc"))
}

func TestMustAToUInt64(t *testing.T) {
	assert.Equal(t, uint64(42), MustAToUInt64("42"))
	assert.Panics(t, func() { MustAToUInt64("abc") })
}

func TestAToF(t *testing.T) {
	v, err := AToF("3.14")
	assert.NoError(t, err)
	assert.InDelta(t, 3.14, v, 1e-9)

	v, err = AToF("-0.5")
	assert.NoError(t, err)
	assert.InDelta(t, -0.5, v, 1e-9)

	_, err = AToF("abc")
	assert.Error(t, err)
}

func TestAToFEmpty(t *testing.T) {
	assert.InDelta(t, 3.14, AToFEmpty("3.14"), 1e-9)
	assert.Equal(t, float64(0), AToFEmpty(""))
	assert.Equal(t, float64(0), AToFEmpty("abc"))
}

func TestMustAToF(t *testing.T) {
	assert.InDelta(t, 3.14, MustAToF("3.14"), 1e-9)
	assert.Panics(t, func() { MustAToF("abc") })
}

func TestAToD(t *testing.T) {
	v, err := AToD("3.14")
	assert.NoError(t, err)
	assert.True(t, v.Equal(decimal.RequireFromString("3.14")))

	_, err = AToD("abc")
	assert.Error(t, err)
}

func TestAToDEmpty(t *testing.T) {
	assert.True(t, AToDEmpty("3.14").Equal(decimal.RequireFromString("3.14")))
	assert.True(t, AToDEmpty("").Equal(decimal.Zero))
	assert.True(t, AToDEmpty("abc").Equal(decimal.Zero))
}

func TestMustAToD(t *testing.T) {
	assert.True(t, MustAToD("3.14").Equal(decimal.RequireFromString("3.14")))
	assert.Panics(t, func() { MustAToD("abc") })
}

func TestAToNullD(t *testing.T) {
	assert.False(t, AToNullD("").Valid)
	assert.False(t, AToNullD("abc").Valid)
	nd := AToNullD("3.14")
	assert.True(t, nd.Valid)
	assert.True(t, nd.Decimal.Equal(decimal.RequireFromString("3.14")))
}

func TestIntToA(t *testing.T) {
	assert.Equal(t, "42", IntToA(42))
	assert.Equal(t, "-7", IntToA(-7))
	assert.Equal(t, "0", IntToA(0))
}

func TestUIntToA(t *testing.T) {
	assert.Equal(t, "42", UIntToA(42))
	assert.Equal(t, "0", UIntToA(0))
}

func TestInt64ToA(t *testing.T) {
	assert.Equal(t, "9223372036854775807", Int64ToA(9223372036854775807))
	assert.Equal(t, "-42", Int64ToA(-42))
}

func TestUInt64ToA(t *testing.T) {
	assert.Equal(t, "18446744073709551615", UInt64ToA(18446744073709551615))
	assert.Equal(t, "0", UInt64ToA(0))
}

func TestFToA(t *testing.T) {
	assert.Equal(t, "3.14", FToA(3.14))
	assert.Equal(t, "-0.5", FToA(-0.5))
	assert.Equal(t, "0", FToA(0))
}

func TestDToA(t *testing.T) {
	assert.Equal(t, "3.14", DToA(decimal.RequireFromString("3.14")))
	assert.Equal(t, "0", DToA(decimal.Zero))
}

func TestFToD(t *testing.T) {
	assert.True(t, FToD(3.14).Equal(decimal.NewFromFloat(3.14)))
	assert.True(t, FToD(0).Equal(decimal.Zero))
}

func TestDToF(t *testing.T) {
	assert.InDelta(t, 3.14, DToF(decimal.RequireFromString("3.14")), 1e-9)
	assert.Equal(t, float64(0), DToF(decimal.Zero))
}
