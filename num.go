package caliber

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

// /////////////////////////////////////////
// format API output numbers
// /////////////////////////////////////////

const (
	ApiOutputDecimalPlace = 8
)

var FloatFormat = fmt.Sprintf("%%.%vf", ApiOutputDecimalPlace)

func DecimalToVoStr(value decimal.Decimal) string {
	return value.StringFixed(ApiOutputDecimalPlace)
}

func DecimalToVoStrDefaultBlank(value decimal.Decimal) string {
	if value.Equal(decimal.Zero) {
		return ""
	}
	return value.StringFixed(ApiOutputDecimalPlace)
}

func FloatToVoStr(value float64) string {
	return fmt.Sprintf(FloatFormat, value)
}

///////////////////////////////////////////
// string to int/int64/float64/decimal
///////////////////////////////////////////

func AToInt(s string) (int, error) {
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("Failed to convert integer: %v", s)
	}

	return val, nil
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
		log.Panicf("Failed to parse [%v] to int", s)
	}
	return val
}

func AToInt64(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("Failed to convert long: %v", s)
	}

	return val, nil
}

func AToInt64Empty(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0
	}

	return val
}

func MustAToI64(s string) int64 {
	val, err := AToInt64(s)
	if err != nil {
		log.Panicf("Failed to parse [%v] to int64", s)
	}
	return val
}

func AToF(s string) (float64, error) {
	val, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, fmt.Errorf("Failed to convert float64: %v", s)
	}

	return val, nil
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
		log.Panicf("Failed to parse [%v] to float64", s)
	}
	return val
}

func AToD(s string) (decimal.Decimal, error) {
	return decimal.NewFromString(s)
}

func AToDEmpty(s string) decimal.Decimal {
	v, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero
	}
	return v
}

func MustAToD(s string) decimal.Decimal {
	val, err := decimal.NewFromString(s)
	if err != nil {
		log.Panicf("Failed to parse [%v] to decimal", s)
	}
	return val
}

///////////////////////////////////////////
// int/int64/float64/decimal to string
///////////////////////////////////////////

func IToA(i int) string {
	return strconv.Itoa(i)
}

func I64ToA(i int64) string {
	return strconv.FormatInt(i, 10)
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

// /////////////////////////////////////////
// generic div
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Div[A, B Number](a A, b B) (float64, error) {
	if float64(b) == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}
