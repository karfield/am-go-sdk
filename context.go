package am

import (
	"context"
	"github.com/karfield/am-go-sdk/internal"
)

type TaskContext struct {
	context.Context
	baseClt internal.BaseIpcClient
	sqlClt  internal.SqlIpcClient
	ocrClt  internal.OcrIpcClient
	cdpClt  internal.CdpIpcClient
	traceId string
	input   []byte
}

type (
	traceIdKey    struct{}
	inputKey      struct{}
	baseClientKey struct{}
	sqlClientKey  struct{}
	ocrClientKey  struct{}
	cdpClientKey  struct{}
)

func (ctx *TaskContext) Value(key any) any {
	switch key.(type) {
	case traceIdKey:
		return ctx.traceId
	case inputKey:
		return ctx.input
	case baseClientKey:
		return ctx.baseClt
	case sqlClientKey:
		return ctx.sqlClt
	case ocrClientKey:
		return ctx.ocrClt
	case cdpClientKey:
		return ctx.cdpClt
	}
	return nil
}
