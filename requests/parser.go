package requests

import (
	"encoding/json"
	"fmt"
)

type RawResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}

func ParseCommonPayload(buf []byte, v any) error {
	var rawResp RawResponse

	if err := json.Unmarshal(buf, &rawResp); err != nil {
		return fmt.Errorf("ParseCommonPayload fail to unmarshal raw buffer %v, err:%+v", string(buf), err.Error())
	}

	if rawResp.Code != 0 {
		return fmt.Errorf("ParseCommonPayload: errCode=%v, errMsg=%v", rawResp.Code, rawResp.Message)
	}

	if rawResp.Data == nil {
		// `Data` could be null, which means no data, not indicating error
		// this leave input `v` unchanged
		return nil
	}

	if err := json.Unmarshal(*rawResp.Data, v); err != nil {
		return fmt.Errorf("ParseCommonPayload fail to unmarshal payload, resp = %+v", rawResp)
	}

	return nil
}
