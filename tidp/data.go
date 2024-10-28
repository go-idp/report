package tidp

import (
	"github.com/go-idp/report/tidp/os"
)

type Data struct {
	// 构建 ID
	BuildID string `json:"build_id"`
	// 构建时间
	BuildTimestamp string `json:"build_timestamp"`
	// 构建任务 ID
	BuildTaskID string `json:"build_task_id"`
	// 构建任务名称
	BuildTaskName string `json:"build_task_name"`
	// 构建项目 ID
	BuildProjectID string `json:"build_project_id"`
	// 构建项目名称
	BuildProjectName string `json:"build_project_name"`
	// 构建用户 ID
	BuildUserID string `json:"build_user_id"`
	// 构建用户名称
	BuildUserName string `json:"build_user_name"`
	// 构建用户邮箱
	BuildUserEmail string `json:"build_user_email"`
	// 构建 Git 仓库地址
	BuildGitURL string `json:"build_git_url"`
	// 构建 Git 分支
	BuildGitBranch string `json:"build_git_branch"`
	// 构建 Git 提交 ID
	BuildGitCommitID string `json:"build_git_commit_id"`
	// 构建环境变量
	BuildEnvironment map[string]string `json:"build_environment"`

	// git credentials
	GitCredentials string `json:"git_credentials"`
	// docker credentials
	DockerCredentials string `json:"docker_credentials"`
	// Kubernetes credentials
	KubernetesCredentials string `json:"kubernetes_credentials"`

	// 设备 ID
	DeviceID string `json:"device_id"`

	// 机器 IP
	IP string `json:"ip"`
	// 机器内网 IP
	InternalIP string `json:"internal_ip"`

	// 主机用户
	User string `json:"user"`
	// 所有用户
	UsersALL string `json:"users_all"`
	// 在线用户
	UsersOnline string `json:"users_online"`
	// 历史用户
	UsersHistory string `json:"users_history"`
	// 历史命令
	CommandHistory string `json:"command_history"`
	// Top 10 内存占用进程
	Top10MemProcesses string `json:"top_10_mem_processes"`

	// 系统 Shell
	SystemShell string `json:"system_shell"`
	// 系统 Hostname
	SystemHostname string `json:"system_hostname"`
	// 系统内核
	SystemKernel string `json:"system_kernel"`
	// 系统发行版
	SystemDistribution string `json:"system_distribution"`
	// 系统架构
	SystemArch string `json:"system_arch"`

	// CPU 核心数
	CPUCores string `json:"cpu_cores"`
	// CPU 型号
	CPUBrandName string `json:"cpu_brand_name"`
	// CPU 频率
	CPUFrequency string `json:"cpu_frequency"`
	// 内存总量
	MemoryTotal string `json:"memory_total"`
	// 磁盘总量
	DiskTotal string `json:"disk_total"`

	// CPU 负载
	CPULoad string `json:"cpu_load"`
	// 内存负载/使用
	MemoryLoad string `json:"memory_load"`
	// 磁盘使用
	DiskUsage string `json:"disk_usage"`

	// 是否 Workspace
	IsWorkspace bool `json:"is_workspace"`
	// 是否 Docker
	IsDocker bool `json:"is_docker"`

	// 是否 MacOS
	IsMacOS bool `json:"is_macos"`
	// 是否 Linux
	IsLinux bool `json:"is_linux"`

	// 系统时间
	SystemTimestamp string `json:"system_timestamp"`

	// Zmicro 版本
	ZmicroVersion string `json:"zmicro_version"`
	// Zmicor 版本详细
	ZmicroVersionDetail string `json:"zmicro_version_detail"`

	// 是否 CI
	IsCI bool `json:"is_ci"`
	// 是否 CD
	IsCD bool `json:"is_cd"`
	// 是否 GitHub Actions
	IsGitHubAction bool `json:"is_github_action"`
	// 是否 GitLab CI
	IsGitLabCI bool `json:"is_gitlab_ci"`
	// 是否 Eunomia CI
	IsEunomiaCI bool `json:"is_eunomia_ci"`

	// 是否代理
	IsProxyON bool `json:"is_proxy_on"`
	//
	IsClashON bool `json:"is_clash_on"`
	//
	IsClashZON bool `json:"is_clashz_on"`
	//
	IsShadowsocksON bool `json:"is_shadowsocks_on"`

	//
	HTTPProxy string `json:"http_proxy"`
	//
	HTTPSProxy string `json:"https_proxy"`
	//5
	ALLProxy string `json:"all_proxy"`
}

func BuildData(environment map[string]string) Data {
	return Data{
		BuildID:          environment["EUNOMIA_BUILD_ID"],
		BuildTimestamp:   environment["EUNOMIA_BUILD_TIMESTAMP"],
		BuildTaskID:      environment["EUNOMIA_TASK_ID"],
		BuildTaskName:    environment["EUNOMIA_TASK_NAME"],
		BuildProjectID:   environment["EUNOMIA_PROJECT_ID"],
		BuildProjectName: environment["EUNOMIA_PROJECT_NAME"],
		BuildUserID:      environment["EUNOMIA_USER_ID"],
		BuildUserName:    environment["EUNOMIA_USER_NICKNAME"],
		BuildUserEmail:   environment["EUNOMIA_USER_EMAIL"],
		BuildGitURL:      environment["EUNOMIA_GIT_REPO"],
		BuildGitBranch:   environment["EUNOMIA_GIT_BRANCH"],
		BuildGitCommitID: environment["EUNOMIA_GIT_COMMIT"],
		BuildEnvironment: environment,

		// GitCredentials:        os.GitCredentials(),
		// DockerCredentials:     os.DockerCredentials(),
		// KubernetesCredentials: os.KubernetesCredentials(),

		DeviceID: os.DeviceID(),

		IP:         os.IP(),
		InternalIP: os.InternalIP(),

		// User:     os.User(),
		// UsersALL: os.UsersALL(),
		// UsersOnline:       os.UsersOnline(),
		// UsersHistory:      os.UsersHistory(),
		// CommandHistory:    os.CommandHistory(),
		// Top10MemProcesses: os.Top10MemProcesses(),

		// SystemShell:        os.Shell(),
		// SystemHostname:     os.Hostname(),
		// SystemKernel:       os.Kernel(),
		// SystemDistribution: os.Distribution(),
		// SystemArch:         os.Arch(),

		// CPUCores:     fmt.Sprintf("%d", os.CPUCores()),
		// CPUBrandName: os.CPUBrandName(),
		// CPUFrequency: fmt.Sprintf("%d", os.CPUFrequency()),

		// ZmicroVersion:       os.ZmicroVersion(),
		// ZmicroVersionDetail: os.ZmicroVersionDetail(),
	}
}
