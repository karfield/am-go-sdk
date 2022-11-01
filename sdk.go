package am

import (
	"context"
	"fmt"
	"github.com/karfield/am-go-sdk/internal"
	"google.golang.org/grpc"
	"log"
	"os"
	"runtime/debug"
	"strconv"
)

type RunOnce func(ctx context.Context) (result []byte, output *string, error error)

func Run(run RunOnce) {
	if run == nil {
		panic("missing run func")
	}
	port := os.Getenv("AM_PORT")
	if port == "" {
		panic("missing environment AM_PORT")
	} else {
		_, err := strconv.ParseInt(port, 10, 16)
		if err != nil {
			panic("illegal AM_PORT environment value")
		}
	}

	call := func(ctx *TaskContext) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("handle task failed: %v", r)
				_, err := ctx.baseClt.FinishWithFailure(context.Background(), &internal.ExecuteFailure{
					Error:      fmt.Sprintf("%v", r),
					PanicStack: debug.Stack(),
				})
				if err != nil {
					log.Printf("fails to report fatal message: %s", err)
				}
			}
		}()

		if result, output, err := run(ctx); err != nil {
			_, err := ctx.baseClt.FinishWithFailure(
				context.Background(),
				&internal.ExecuteFailure{
					Error: err.Error(),
				},
			)
			if err != nil {
				log.Printf("fails to report failure: %s", err)
			}
		} else {
			_, err := ctx.baseClt.FinishWithResult(context.Background(), &internal.ExecuteResult{
				PortIndicator: output,
				Output:        result,
			})
			if err != nil {
				log.Printf("fails to feedback result: %s", err)
			}
		}
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithPerRPCCredentials(ipcCredentials{}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	baseClient := internal.NewBaseIpcClient(conn)
	response, err := baseClient.Capabilities(context.Background(), &internal.CapabilitiesRequest{}, nil)
	if err != nil {
		log.Fatalf("fails to obtain bot capabilities: %s", err)
	}

	var sqlClient internal.SqlIpcClient
	if response.GetSql() {
		sqlClient = internal.NewSqlIpcClient(conn)
	}
	var ocrClient internal.OcrIpcClient
	if response.GetOcr() {
		ocrClient = internal.NewOcrIpcClient(conn)
	}

	if consumer, err := baseClient.ConsumeTask(context.Background(), &internal.ConsumeTaskRequest{}, nil); err != nil {
		log.Fatalf("unable to consume task: %s", err)
	} else {
		for {
			msg, err := consumer.Recv()
			if err != nil {
				log.Fatalf("unable to receive task from host: %s", err)
			}

			ctx := TaskContext{
				Context: context.Background(),
				traceId: msg.GetTraceId(),
				input:   msg.GetPayload(),
				baseClt: baseClient,
				sqlClt:  sqlClient,
				ocrClt:  ocrClient,
			}

			call(&ctx)
		}
	}
}

type ipcCredentials struct {
}

func (creds ipcCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"Process-Id":  fmt.Sprintf("%d", os.Getpid()),
		"Trace-Id":    os.Getenv("AM_TRACE_ID"),
		"Instance-Id": os.Getenv("AM_INSTANCE_ID"),
	}, nil
}

func (creds ipcCredentials) RequireTransportSecurity() bool {
	return false
}
