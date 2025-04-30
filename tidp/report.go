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
	Script      string
	Environment map[string]string `json:"environment"`
}

type ReportResponse struct {
	// Approval status, can be "approved", "rejected" or "pending"
	ApprovalStatus string `json:"approval_status"`

	// Delay in miniseconds
	ApprovalDelay int `json:"approval_delay"`

	// Reason for approval or rejection
	ApprovalReason string `json:"approval_reason"`

	// Inject Scripts
	ApprovalInjectScripts ApprovalInjectScripts `json:"approval_inject_scripts"`
}

type ApprovalInjectScripts struct {
	Before string `json:"before"`
	After  string `json:"after"`
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

func (r *ReportResponse) InjectScripts() []string {
	return []string{
		r.ApprovalInjectScripts.Before,
		r.ApprovalInjectScripts.After,
	}
}

func (r *ReportResponse) InjectScriptsBefore() string {
	return r.ApprovalInjectScripts.Before
}

func (r *ReportResponse) InjectScriptsAfter() string {
	return r.ApprovalInjectScripts.After
}

// Report reports the data to the feishu group
func Report(req *ReportRequest) (res *ReportResponse, err error) {
	res = &ReportResponse{}

	err = safe.Do(func() error {
		data := BuildData(req.Script, req.Environment)
		response, err := fetch.Post(getReportURL(), &fetch.Config{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: data,
			//
			Timeout: defaultReportTimeout,
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
