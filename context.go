package am

import (
	"github.com/karfield/am-go-sdk/internal"
	"time"
)

type TaskContext struct {
	baseClt internal.BaseIpcClient
	sqlClt  internal.SqlIpcClient
	ocrClt  internal.OcrIpcClient
	cdpClt  internal.CdpIpcClient
	traceId string
	input   []byte
}

func (t *TaskContext) Deadline() (deadline time.Time, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskContext) Done() <-chan struct{} {
	//TODO implement me
	panic("implement me")
}

func (t *TaskContext) Err() error {
	return nil
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
