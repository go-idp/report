package tidp

import (
	"os"
	"time"
)

const reportURL = "https://api.zcorky.com/logs/zmicro-tidps"
const reportURLLocal = "http://127.0.0.1:9000/api/open/v1/logs/zmicro-tidps"
const defaultReportTimeout = 10 * time.Minute

func getReportURL() string {
	if os.Getenv("MODE") == "local" {
		return reportURLLocal
	}

	return reportURL
}
