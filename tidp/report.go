package tidp

import (
	"fmt"
	"time"

	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/debug"
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/safe"
)

type ReportRequest struct {
	Environment map[string]string `json:"environment"`
}

type ReportResponse struct {
	// Approval status, can be "approved", "rejected" or "pending"
	ApprovalStatus string `json:"approval_status"`

	// Delay in miniseconds
	ApprovalDelay int `json:"approval_delay"`

	// Reason for approval or rejection
	ApprovalReason string `json:"approval_reason"`
}

func (r *ReportResponse) Approved() bool {
	return r.ApprovalStatus == "approved"
}

func (r *ReportResponse) Rejected() bool {
	return r.ApprovalStatus == "rejected"
}

func (r *ReportResponse) Pending() bool {
	return r.ApprovalStatus == "pending"
}

func (r *ReportResponse) Reason() string {
	if r.ApprovalReason == "" {
		return "网络异常，请稍后重试"
	}

	return r.ApprovalReason
}

func (r *ReportResponse) Delay() time.Duration {
	return time.Duration(r.ApprovalDelay) * time.Millisecond
}

// Report reports the data to the feishu group
func Report(req *ReportRequest) (res *ReportResponse, err error) {
	res = &ReportResponse{}

	err = safe.Do(func() error {
		data := BuildData(req.Environment)

		response, err := fetch.Post(reportURL, &fetch.Config{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: data,
			//
			Timeout: 3 * 60 * time.Second,
		})
		if err != nil {
			debug.Debug("failed to fetch post in report: %v", err)
			return err
		}
		if !response.Ok() {
			debug.Debug("failed to fetch post in report: %s", response.String())
			return fmt.Errorf("failed to fetch post in report: %s", response.String())
		}

		if err = response.UnmarshalJSON(res); err != nil {
			debug.Debug("failed to report: %s", response.String())
			return fmt.Errorf("failed to report: %s", response.String())
		}

		debug.Debug("report success: [id: %s] %s", data.BuildID, data.BuildTaskName)
		return nil
	})
	if err != nil {
		debug.Debug("failed to report: %v", err)

		// if timeout, auto approve
		if strings.Contains(err.Error(), "context deadline exceeded") {
			res.ApprovalStatus = "approved"
			return res, nil
		}

		// if not found, auto approve
		if strings.Contains(err.Error(), "Not Found") {
			res.ApprovalStatus = "approved"
			return res, nil
		}

		return nil, err
	}

	return res, err
}
