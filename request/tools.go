package request

import (
	"fmt"
	"net/url"
)

func MapToQueryStr(params map[string]any) (string, error) {
	q := url.Values{}

	for k, v := range params {
		switch val := v.(type) {
		case string:
			q.Set(k, val)
		case int, int64, uint, uint64, float64, bool:
			q.Set(k, fmt.Sprint(val))
		case []string:
			for _, s := range val {
				q.Add(k, s)
			}
		default:
			return "", fmt.Errorf("unsupported type for key %q: %T", k, v)
		}
	}

	return q.Encode(), nil
}
