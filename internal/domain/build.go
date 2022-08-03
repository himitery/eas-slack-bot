package domain

import "time"

type BuildResult struct {
	Id               string  `json:"id"`
	AppId            string  `json:"appId"`
	InitiatingUserId string  `json:"initiatingUserId"`
	CancelingUserId  *string `json:"cancelingUserId"`
	Platform         string  `json:"platform"`
	Status           string  `json:"status"`
	Artifacts        struct {
		BuildUrl        string `json:"buildUrl"`
		LogsS3KeyPrefix string `json:"logsS3KeyPrefix"`
	} `json:"artifacts"`
	Metadata struct {
		AppName         string `json:"appName"`
		Username        string `json:"username"`
		Workflow        string `json:"workflow"`
		AppVersion      string `json:"appVersion"`
		CliVersion      string `json:"cliVersion"`
		SdkVersion      string `json:"sdkVersion"`
		BuildProfile    string `json:"buildProfile"`
		Distribution    string `json:"distribution"`
		AppIdentifier   string `json:"appIdentifier"`
		GitCommitHash   string `json:"gitCommitHash"`
		ReleaseChannel  string `json:"releaseChannel"`
		AppBuildVersion string `json:"appBuildVersion"`
		TrackingContext struct {
			Platform         string `json:"platform"`
			AccountId        string `json:"account_Id"`
			DevClient        bool   `json:"dev_Client"`
			ProjectId        string `json:"project_id"`
			TrackingId       string `json:"tracking_id"`
			ProjectType      string `json:"project_type"`
			DevClientVersion string `json:"dev_client_version"`
		} `json:"trackingContext"`
		CredentialsSource     string `json:"credentialsSource"`
		IsGitWorkingTreeDirty bool   `json:"isGitWorkingTreeDirty"`
	} `json:"metadata"`
	Metrics struct {
		Memory                   int     `json:"memory"`
		BuildEndTimestamp        int64   `json:"buildEndTimestamp"`
		TotalDiskReadBytes       int     `json:"totalDiskReadBytes"`
		BuildStartTimestamp      int64   `json:"buildStartTimestamp"`
		TotalDiskWriteBytes      int     `json:"totalDiskWriteBytes"`
		CpuActiveMilliseconds    float64 `json:"cpuActiveMilliseconds"`
		BuildEnqueuedTimestamp   int64   `json:"buildEnqueuedTimestamp"`
		TotalNetworkEgressBytes  int     `json:"totalNetworkEgressBytes"`
		TotalNetworkIngressBytes int     `json:"totalNetworkIngressBytes"`
	} `json:"metrics"`
	Error struct {
		Message   string `json:"message"`
		ErrorCode string `json:"errorCode"`
	} `json:"error"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ExpirationDate time.Time `json:"expirationDate"`
}
