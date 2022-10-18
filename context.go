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

func (t *TaskContext) Value(key any) any {
	switch key.(type) {
	case traceIdKey:
		return t.traceId
	case inputKey:
		return t.input
	case baseClientKey:
		return t.baseClt
	case sqlClientKey:
		return t.sqlClt
	case ocrClientKey:
		return t.ocrClt
	case cdpClientKey:
		return t.cdpClt
	}
	return nil
}
