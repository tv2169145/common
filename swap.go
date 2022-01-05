package common

import "encoding/json"

func SwapTo(request, model interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, model)
}
