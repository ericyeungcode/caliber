package requests

import (
	"encoding/json"
	"fmt"
)

type BizResp struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}

func ParseBizPayload(buf []byte, v any) error {
	var bizResp BizResp

	if err := json.Unmarshal(buf, &bizResp); err != nil {
		return fmt.Errorf("ParseCommonPayload fail to unmarshal raw buffer %v, err:%+v", string(buf), err.Error())
	}

	if bizResp.Code != 0 {
		return fmt.Errorf("ParseCommonPayload: errCode=%v, errMsg=%v", bizResp.Code, bizResp.Message)
	}

	if bizResp.Data == nil {
		// `Data` could be null, which means no data, not indicating error
		// this leave input `v` unchanged
		return nil
	}

	if err := json.Unmarshal(*bizResp.Data, v); err != nil {
		return fmt.Errorf("ParseCommonPayload fail to unmarshal payload, resp = %+v", bizResp)
	}

	return nil
}
