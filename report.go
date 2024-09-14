package report

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/datetime"
	"github.com/go-zoox/debug"
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/ip"
	"github.com/go-zoox/once"
)

var currentPublicIP string
var currentPrivateIP string
var latestReportTime time.Time

// Report reports the data to the feishu group
func Report(namespace, title string, data any) error {
	once.Do("get ip", func() {
		currentPublicIP, _ = ip.GetPublicIP()
		currentPrivateIP, _ = ip.GetInternalIP()
	})

	// signature := md5.Md5([]byte(fmt.Sprintf("%s%s", namespace, title)), crypto.MD5)

	if time.Since(latestReportTime) < minReportInterval {
		debug.Debug("report too frequently")
		return nil
	}

	latestReportTime = time.Now()

	title = fmt.Sprintf(reportTitleFormat, namespace, title)
	content, err := json.MarshalIndent(map[string]any{
		"metadata": map[string]any{
			"public_ip":  currentPublicIP,
			"private_ip": currentPrivateIP,
		},
		"data": data,
	}, "", "  ")
	if err != nil {
		debug.Debug("failed to marshal report data: %v", err)
		return err
	}

	response, err := fetch.Post(reportURL, &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: map[string]any{
			"msg_type": "post",
			"content": map[string]any{
				"post": map[string]any{
					"zh_cn": map[string]any{
						"title": title,
						"content": [][]map[string]any{
							{
								{
									"tag": "text",
									"text": fmt.Sprintf(
										strings.Join([]string{
											"详细数据：%s",
											"发生时间：%s",
										}, "\n"),
										//
										string(content),
										datetime.Now().Format(),
									),
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		debug.Debug("failed to fetch post in report: %v", err)
		return err
	}
	if !response.Ok() {
		debug.Debug("failed to fetch post in report: %s", response.String())
		return fmt.Errorf("failed to fetch post in report: %s", response.String())
	}

	if response.Get("code").Int() != 0 {
		debug.Debug("failed to report: %s", response.String())
		return fmt.Errorf("failed to report: %s", response.String())
	}

	debug.Debug("report success: [namespace: %s] %s", namespace, title)
	return nil
}
