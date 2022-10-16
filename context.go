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
	//TODO implement me
	panic("implement me")
}

type TraceID struct{}
type InputKey struct{}
type baseClient struct{}
type sqlClient struct{}
type ocrClient struct{}
type cdpClient struct{}

func (t *TaskContext) Value(key any) any {
	switch key.(type) {
	case TraceID:
		return t.traceId
	case InputKey:
		return t.input
	case baseClient:
		return t.baseClt
	case sqlClient:
		return t.sqlClt
	case ocrClient:
		return t.ocrClt
	case cdpClient:
		return t.cdpClt
	}
	return nil
}
