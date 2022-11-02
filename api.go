package am

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/karfield/am-go-sdk/internal"
	"log"
	"os"
)

const (
	LogDebugLevel uint32 = iota
	LogInfoLevel
	LogWarningLevel
	LogErrorLevel
	LogFatalLevel
)

func BotID() string         { return os.Getenv("AM_BOT_ID") }
func InstanceID() string    { return os.Getenv("AM_INSTANCE_ID") }
func TaskCode() string      { return os.Getenv("AM_TASK_CODE") }
func TaskVersion() string   { return os.Getenv("AM_TASK_VERSION") }
func ComponentCode() string { return os.Getenv("AM_COMPONENT_CODE") }
func ComponentType() string { return os.Getenv("AM_COMPONENT_TYPE") }

func TraceID(ctx context.Context) string {
	if value := ctx.Value(traceIdKey{}); value != nil {
		if id, ok := value.(string); ok {
			return id
		}
	}
	return ""
}

func InputData(ctx context.Context) []byte {
	if value := ctx.Value(inputKey{}); value != nil {
		if data, ok := value.([]byte); ok {
			return data
		}
	}
	return nil
}

func Log(ctx context.Context, level uint32, message string, extra map[string]interface{}) {
	if clt := ctx.Value(baseClientKey{}); clt != nil {
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
				TraceId:   TraceID(ctx),
				Level:     internal.LogLevel(level),
				Message:   message,
				ExtraJson: extraJson,
			})
			if err != nil {
				log.Printf("fails to save logs: %s", err)
			}
		}
	}
}

func QuerySql(ctx context.Context, defaultDbCode *string, sql string, args interface{}) ([][]byte, error) {
	if clt := ctx.Value(sqlClientKey{}); clt != nil {
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
				&internal.ExecuteSqlRequest{
					TraceId: TraceID(ctx),
					Sql:     sql, JsonArguments: content, DefaultDbCode: defaultDbCode})
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
	if clt := ctx.Value(sqlClientKey{}); clt != nil {
		if clt, ok := clt.(internal.SqlIpcClient); ok {
			content, err := json.Marshal(args)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to marshal arguments error: %s", err)
			}
			result, err := clt.ExecuteSingleSql(context.Background(), &internal.ExecuteSqlRequest{
				TraceId: TraceID(ctx),
				Sql:     sql, JsonArguments: content, DefaultDbCode: defaultDbCode})
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
	if ocrClt := ctx.Value(ocrClientKey{}); ocrClt != nil {
		if ocrClt, ok := ocrClt.(internal.OcrIpcClient); ok {
			response, err := ocrClt.ResolveCaptchaImage(context.Background(), &internal.CaptchaImage{
				TraceId:  TraceID(ctx),
				Width:    width,
				Height:   height,
				Format:   format,
				Image:    image,
				Provider: provider,
			})
			if err != nil {
				return "", fmt.Errorf("failed to decode captcha image: %s", err)
			}
			return response.Answer, nil
		}
	}
	return "", nil
}

func NewCdpBrowser(ctx context.Context, opts ...chromedp.BrowserOption) (*chromedp.Browser, error) {
	if value := ctx.Value(cdpClientKey{}); value != nil {
		if cdpClt, ok := value.(internal.CdpIpcClient); ok {
			response, err := cdpClt.GetBrowserwsUrl(context.Background(), &internal.GetBrowserWsUrlRequest{})
			if err != nil {
				return nil, fmt.Errorf("fails to get browser ws-url: %E", err)
			}
			return chromedp.NewBrowser(ctx, response.GetWsUrl(), opts...)
		}
	}
	return nil, nil
}

func NewCdpAllocator(ctx context.Context) (context.Context, context.CancelFunc, error) {
	if value := ctx.Value(cdpClientKey{}); value != nil {
		if cdpClt, ok := value.(internal.CdpIpcClient); ok {
			response, err := cdpClt.GetBrowserDebuggerPort(context.Background(), &internal.GetBrowserDebuggerPortRequest{})
			if err != nil {
				return nil, nil, fmt.Errorf("fails to get browser ws-url: %s", err)
			}
			ctx, allocatorCancel := chromedp.NewRemoteAllocator(ctx, fmt.Sprintf("http://127.0.0.1:%d/", response.Port))
			ctx, ctxCancel := chromedp.NewContext(ctx)
			return ctx, func() {
				ctxCancel()
				allocatorCancel()
			}, nil
		}
	}
	return nil, nil, fmt.Errorf("cannot get cdp client key")
}
