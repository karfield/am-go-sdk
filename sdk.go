package am

import (
	"context"
	"fmt"
	"github.com/karfield/am-go-sdk/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

type RunOnce func(ctx context.Context) (output string, result []byte, error error)

func Run(port int, run RunOnce) {
	if run == nil {
		panic("missing run func")
	}

	call := func(ctx *TaskContext) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("handle task failed: %v", r)
				_, err := ctx.baseClt.FinishWithFailure(context.Background(), &internal.ExecuteFailure{
					TraceId: ctx.traceId,
					Stage:   "",
					Error:   fmt.Sprintf("%v", r),
				})
				if err != nil {
					log.Printf("fails to report fatal message: %s", err)
				}
			}
		}()

		if port, result, err := run(ctx); err != nil {
			_, err := ctx.baseClt.FinishWithFailure(
				context.Background(),
				&internal.ExecuteFailure{
					TraceId: ctx.traceId,
					Stage:   "",
					Error:   err.Error(),
				},
			)
			if err != nil {
				log.Printf("fails to report failure: %s", err)
			}
		} else {
			if port == "" {
				port = "output"
			}

			_, err := ctx.baseClt.FinishWithResult(context.Background(), &internal.ExecuteResult{
				TraceId:       ctx.traceId,
				PortIndicator: &port,
				Output:        result,
			})
			if err != nil {
				log.Printf("fails to feedback result: %s", err)
			}
		}
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	baseClient := internal.NewBaseIpcClient(conn)
	response, err := baseClient.Capabilities(context.Background(), &internal.CapabilitiesRequest{}, metaHeader())
	if err != nil {

	}

	var sqlClient internal.SqlIpcClient
	if response.GetSql() {
		sqlClient = internal.NewSqlIpcClient(conn)
	}
	var ocrClient internal.OcrIpcClient
	if response.GetOcr() {
		ocrClient = internal.NewOcrIpcClient(conn)
	}

	if consumer, err := baseClient.ConsumeTask(context.Background(), &internal.ConsumeTaskRequest{}, metaHeader()); err != nil {
		log.Fatalf("unable to consume task: %s", err)
	} else {
		for {
			msg, err := consumer.Recv()
			if err != nil {
				log.Fatalf("unable to receive task from host: %s", err)
			}

			ctx := TaskContext{
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

func metaHeader() grpc.CallOption {
	md := metadata.New(map[string]string{})
	md.Append("pid", fmt.Sprintf("%d", os.Getpid()))
	return grpc.Header(&md)
}
