package numx

/*

String to number Patterns: T is Int, UInt, Int64, UInt64, F, D (decimal)

- ATo<T>
- ATo<T>Empty
- MustATo<T>

Number to string Patterns: T is Int, UInt, Int64, UInt64, F, D (decimal)

- <T>ToA

*/

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

///////////////////////////////////////////
// convert strings to numbers
///////////////////////////////////////////

func AToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func AToIntEmpty(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

func MustAToInt(s string) int {
	val, err := AToInt(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse [%v] to int", s))
	}
	return val
}

func AToUInt(s string) (uint, error) {
	val, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func AToUIntEmpty(s string) uint {
	val, err := AToUInt(s)
	if err != nil {
		return 0
	}
	return val
}

func MustAToUInt(s string) uint {
	val, err := AToUInt(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse [%v] to uint", s))
	}
	return val
}

func AToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func AToInt64Empty(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func MustAToInt64(s string) int64 {
	val, err := AToInt64(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse [%v] to int64", s))
	}
	return val
}

func AToUInt64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func AToUInt64Empty(s string) uint64 {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func MustAToUInt64(s string) uint64 {
	val, err := AToUInt64(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse [%v] to uint64", s))
	}
	return val
}

func AToF(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func AToFEmpty(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return val
}

func MustAToF(s string) float64 {
	val, err := AToF(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse [%v] to float64", s))
	}
	return val
}

func AToD(s string) (decimal.Decimal, error) {
	return decimal.NewFromString(s)
}

func AToDEmpty(s string) decimal.Decimal {
	// Fast-path the common "" (not-populated) case: NewFromString("")
	// would otherwise run the full parse and allocate an error before we
	// discard it.
	if s == "" {
		return decimal.Zero
	}
	v, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero
	}
	return v
}

func MustAToD(s string) decimal.Decimal {
	val, err := decimal.NewFromString(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse [%v] to decimal", s))
	}
	return val
}

// AToNullD parses s into a decimal.NullDecimal. Empty or malformed input
// yields NullDecimal{Valid:false} (the nullable counterpart of AToDEmpty,
// which collapses the same inputs to decimal.Zero).
func AToNullD(s string) decimal.NullDecimal {
	if s == "" {
		return decimal.NullDecimal{}
	}
	v, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.NullDecimal{}
	}
	return decimal.NewNullDecimal(v)
}

///////////////////////////////////////////
// convert numbers to strings
///////////////////////////////////////////

func IntToA(i int) string {
	return strconv.Itoa(i)
}

func UIntToA(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func Int64ToA(i int64) string {
	return strconv.FormatInt(i, 10)
}

func UInt64ToA(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// to get formatted output, use FloatToVoStr()
func FToA(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// to get formatted output, use DecimalToVoStr()
func DToA(d decimal.Decimal) string {
	return d.String()
}

///////////////////////////////////////////
// float64 to decimal and vise vesa
///////////////////////////////////////////

func FToD(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}

func DToF(d decimal.Decimal) float64 {
	return d.InexactFloat64()
}
