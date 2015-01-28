package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

//将未知格式json字符串解析成key value形式
func DecodeJsonMsg(jsonMsg string) (map[string]interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(jsonMsg))
	decodedMsg := make(map[string]interface{})
	if err := decoder.Decode(&decodedMsg); err != nil {
		fmt.Errorf("decode json msg '%s' error: %v", jsonMsg, err)
		return nil, err
	} else {
		return decodedMsg, err
	}
}
