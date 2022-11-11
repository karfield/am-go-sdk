package main

import (
	"context"
	"encoding/json"
	"github.com/karfield/am-go-sdk"
)

func main() {
	am.Run(func(ctx context.Context) (result []byte, output *string, error error) {
		input := map[string]interface{}{}
		if err := json.Unmarshal(am.InputData(ctx), &input); err != nil {
			am.Log(ctx, am.LogFatalLevel, err.Error(), nil)
			return nil, nil, err
		}
		am.Log(ctx, am.LogInfoLevel, "收到API调用", input)
		return am.InputData(ctx), nil, nil
	})
}
