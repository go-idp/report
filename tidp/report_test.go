package tidp

import (
	"encoding/json"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
)

func TestReport(t *testing.T) {
	raw := `{
    "CI": "true",
    "EUNOMIA_ACTOR": "USER",
    "EUNOMIA_ACTOR_ID": "1000",
    "EUNOMIA_BUILD_AUTHOR": "USER",
    "EUNOMIA_BUILD_AUTHOR_EMAIL": "example@example.cn",
    "EUNOMIA_BUILD_AUTHOR_ID": "106",
    "EUNOMIA_BUILD_FROM": "eunomia",
    "EUNOMIA_BUILD_ID": "1000_1234",
    "EUNOMIA_BUILD_TIMESTAMP": "1730085383",
    "EUNOMIA_BUILD_TRIGGER": "idp",
    "EUNOMIA_CI": "true",
    "EUNOMIA_DEPLOYMENT_CPU_LIMIT": "2",
    "EUNOMIA_DEPLOYMENT_ID": "1234",
    "EUNOMIA_DEPLOYMENT_MEMORY_LIMIT": "4096",
    "EUNOMIA_DEPLOY_COMPONENT_CLUSTER_VERSION": "dev",
    "EUNOMIA_DEPLOY_COMPONENT_MODULE_NAME": "example",
    "EUNOMIA_EXPORT_DIST": "1000_1234.tar.gz",
    "EUNOMIA_EXPORT_DIST_URL": "https://storage.idp.example.cn/1000_1234.tar.gz",
    "EUNOMIA_FAILED_HOOK": "https://idp.example.cn/api/open/v1/webhook/flows/deployments/1234/fail",
    "EUNOMIA_GIT_BRANCH": "master",
    "EUNOMIA_GIT_CHECKOUT_TYPE": "branch",
    "EUNOMIA_GIT_COMMIT": "b960eb66075dcc",
    "EUNOMIA_GIT_REPO": "http://10.27.249.150:8888/example/example",
    "EUNOMIA_GIT_REPOSITORY": "http://10.27.249.150:8888/example/example",
    "EUNOMIA_GIT_TAG": "",
    "EUNOMIA_INFO_SUPPLY_HOOK": "https://idp.example.cn/api/open/v1/webhook/flows/deployments/1234/info_supply",
    "EUNOMIA_JOB_ID": "50412",
    "EUNOMIA_LOG_HOOK": "https://idp.example.cn/api/open/v1/webhook/flows/deployments/1234/log",
    "EUNOMIA_NAME": "example 构建",
    "EUNOMIA_PROJECT_ENABLE_AUTO_TEST": "false",
    "EUNOMIA_PROJECT_ID": "452",
    "EUNOMIA_RUNNER_ARCH": "amd64",
    "EUNOMIA_RUNNER_NAME": "Eunomia CI",
    "EUNOMIA_RUNNER_OS": "Linux",
    "EUNOMIA_RUNNER_PLATFORM": "linux",
    "EUNOMIA_RUNNER_USER": "runner",
    "EUNOMIA_SKIP_BUILD_STAGE": "true",
    "EUNOMIA_START_HOOK": "https://idp.example.cn/api/open/v1/webhook/flows/deployments/1234/start",
    "EUNOMIA_SUCCEED_HOOK": "https://idp.example.cn/api/open/v1/webhook/flows/deployments/1234/succeed",
    "EUNOMIA_TASK_ID": "1000",
    "EUNOMIA_TASK_NAME": "example 构建",
    "EUNOMIA_TASK_TEST_TASKIDS": "",
    "EUNOMIA_TASK_TEST_URL": "",
    "EUNOMIA_TEST_REPORT_HOOK": "https://idp.example.cn/api/open/v1/webhook/flows/deployments/1234/test_report",
    "EUNOMIA_USER_EMAIL": "example@example.cn",
    "EUNOMIA_USER_ID": "106",
    "EUNOMIA_USER_NICKNAME": "USER",
    "EUNOMIA_VERSION": "v1",
    "branch": "master",
    "eunomia_task_name": "example 构建",
    "eunomia_user_email": "example@example.cn",
    "eunomia_user_id": "106",
    "eunomia_user_nickname": "USER"
}`

	environment := map[string]string{}
	if err := json.Unmarshal([]byte(raw), &environment); err != nil {
		t.Fatal(err)
	}

	response, err := Report(&ReportRequest{
		Environment: environment,
	})
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("response is nil")
	}

	fmt.PrintJSON(response)
}
