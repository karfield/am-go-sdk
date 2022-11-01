package main

import (
	"context"
	"github.com/karfield/am-go-sdk"
)

func main() {
	am.Run(func(ctx context.Context) (result []byte, output *string, error error) {
		am.Log(ctx, am.LogInfoLevel, "Test bot log functionality", nil)

		return am.InputData(ctx), nil, nil
	})
}
