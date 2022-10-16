package am

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/karfield/am-go-sdk/internal"
	"log"
)

const (
	LogDebugLevel uint32 = iota
	LogInfoLevel
	LogWarningLevel
	LogErrorLevel
	LogFatalLevel
)

func Log(ctx context.Context, level uint32, message string, extra map[string]interface{}) {
	if clt := ctx.Value(baseClient{}); clt != nil {
		if clt, ok := clt.(internal.BaseIpcClient); ok {
			var (
				extraJson []byte
				err       error
			)
			if extra != nil {
				extraJson, err = json.Marshal(extra)
				if err != nil {
					log.Printf("fails to marshal extra args %s", err)
					return
				}
			}
			_, err = clt.SaveLog(context.Background(), &internal.SaveLogRequest{
				Level:     internal.LogLevel(level),
				Message:   message,
				ExtraJson: extraJson,
			}, metaHeader(ctx))
			if err != nil {
				log.Printf("fails to save logs: %s", err)
			}
		}
	}
}

func QuerySql(ctx context.Context, defaultDbCode *string, sql string, args interface{}) ([][]byte, error) {
	if clt := ctx.Value(sqlClient{}); clt != nil {
		if clt, ok := clt.(internal.SqlIpcClient); ok {
			var content []byte
			var err error
			if args != nil {
				content, err = json.Marshal(args)
				if err != nil {
					return nil, fmt.Errorf("failed to marshal arguments, error: %s", err)
				}
			}
			result, err := clt.QuerySingleSql(context.Background(),
				&internal.ExecuteSqlRequest{Sql: sql, JsonArguments: content, DefaultDbCode: defaultDbCode}, metaHeader(ctx))
			if err != nil {
				return nil, fmt.Errorf("fails to query sql: %s, error: %s", sql, err)
			}
			if records := result.GetQueryResult(); records != nil {
				return records.GetJsonRecords(), nil
			} else if failure := result.GetFailure(); failure != nil {
				return nil, fmt.Errorf("%s", failure.GetMessage())
			}
		}
	}
	return nil, nil
}

func ExecuteSql(ctx context.Context, defaultDbCode *string, sql string, args interface{}) (lastInsertId uint64, rowsEffected uint64, err error) {
	if clt := ctx.Value(sqlClient{}); clt != nil {
		if clt, ok := clt.(internal.SqlIpcClient); ok {
			content, err := json.Marshal(args)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to marshal arguments error: %s", err)
			}
			result, err := clt.ExecuteSingleSql(context.Background(), &internal.ExecuteSqlRequest{Sql: sql, JsonArguments: content, DefaultDbCode: defaultDbCode}, metaHeader(ctx))
			if err != nil {
				return 0, 0, fmt.Errorf("failed to execute sql: %s, error: %s", sql, err)
			}
			if effects := result.GetExecResult(); effects != nil {
				return effects.GetLastInsertId(), effects.GetRowsEffected(), err
			} else if failure := result.GetFailure(); failure != nil {
				return 0, 0, fmt.Errorf("%s", failure.GetMessage())
			}
		}
	}
	return 0, 0, nil
}

func ResolveCaptchaImage(ctx context.Context, width, height uint32, format string, image []byte, provider *string) (string, error) {
	if ocrClt := ctx.Value(ocrClient{}); ocrClt != nil {
		if ocrClt, ok := ocrClt.(internal.OcrIpcClient); ok {
			response, err := ocrClt.ResolveCaptchaImage(context.Background(), &internal.CaptchaImage{
				Width:    width,
				Height:   height,
				Format:   format,
				Image:    image,
				Provider: provider,
			}, metaHeader(ctx))
			if err != nil {
				return "", fmt.Errorf("failed to decode captcha image: %s", err)
			}
			return response.Answer, nil
		}
	}
	return "", nil
}
