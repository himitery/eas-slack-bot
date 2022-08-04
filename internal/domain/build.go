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
		BuildUrl        *string `json:"buildUrl,omitempty"`
		LogsS3KeyPrefix string  `json:"logsS3KeyPrefix"`
	} `json:"artifacts"`
	Metadata struct {
		AppName         string  `json:"appName"`
		Username        string  `json:"username"`
		Workflow        string  `json:"workflow"`
		AppVersion      string  `json:"appVersion"`
		CliVersion      string  `json:"cliVersion"`
		SdkVersion      string  `json:"sdkVersion"`
		BuildProfile    string  `json:"buildProfile"`
		Distribution    string  `json:"distribution"`
		AppIdentifier   string  `json:"appIdentifier"`
		GitCommitHash   string  `json:"gitCommitHash"`
		ReleaseChannel  *string `json:"releaseChannel,omitempty"`
		AppBuildVersion string  `json:"appBuildVersion"`
		TrackingContext struct {
			Platform         string `json:"platform"`
			AccountId        string `json:"account_id"`
			DevClient        bool   `json:"dev_client"`
			ProjectId        string `json:"project_id"`
			TrackingId       string `json:"tracking_id"`
			ProjectType      string `json:"project_type"`
			DevClientVersion string `json:"dev_client_version"`
		} `json:"trackingContext"`
		GitCommitMessage      *string `json:"gitCommitMessage,omitempty"`
		CredentialsSource     string  `json:"credentialsSource"`
		ReactNativeVersion    *string `json:"reactNativeVersion,omitempty"`
		IsGitWorkingTreeDirty bool    `json:"isGitWorkingTreeDirty"`
	} `json:"metadata"`
	Metrics struct {
		Memory                   *int     `json:"memory,omitempty"`
		BuildEndTimestamp        *int64   `json:"buildEndTimestamp,omitempty"`
		TotalDiskReadBytes       *int     `json:"totalDiskReadBytes,omitempty"`
		BuildStartTimestamp      *int64   `json:"buildStartTimestamp,omitempty"`
		TotalDiskWriteBytes      *int     `json:"totalDiskWriteBytes,omitempty"`
		CpuActiveMilliseconds    *float64 `json:"cpuActiveMilliseconds,omitempty"`
		BuildEnqueuedTimestamp   *int64   `json:"buildEnqueuedTimestamp,omitempty"`
		TotalNetworkEgressBytes  *int     `json:"totalNetworkEgressBytes,omitempty"`
		TotalNetworkIngressBytes *int     `json:"totalNetworkIngressBytes,omitempty"`
	} `json:"metrics"`
	Error *struct {
		Message   string `json:"message"`
		ErrorCode string `json:"errorCode"`
	} `json:"error"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ExpirationDate time.Time `json:"expirationDate"`
	Priority       *string   `json:"priority,omitempty"`
}
