package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/karfield/am-go-sdk"
)

func main() {
	am.Run(func(ctx context.Context) (result []byte, output *string, error error) {
		am.Log(ctx, am.LogInfoLevel, "开始测试", nil)
		cdpCtx, cancel, err := am.NewCdpAllocator(ctx)
		if err != nil {
			return nil, nil, err
		}

		defer cancel()

		am.Log(ctx, am.LogInfoLevel, "开始测试网站", nil)

		if err := chromedp.Run(cdpCtx, chromedp.Navigate("https://www.baidu.com")); err != nil {
			am.Log(cdpCtx, am.LogErrorLevel, "打开失败", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, nil, err
		}

		am.Log(ctx, am.LogInfoLevel, "测试结束", nil)

		return am.InputData(ctx), nil, nil
	})
}
