package domain

import "time"

type SubmitResult struct {
	Id               string `json:"id"`
	AppId            string `json:"appId"`
	InitiatingUserId string `json:"initiatingUserId"`
	TurtleBuildId    string `json:"turtleBuildId"`
	Platform         string `json:"platform"`
	Status           string `json:"status"`
	SubmissionInfo   struct {
		Error struct {
			Message   string `json:"message"`
			ErrorCode string `json:"errorCode"`
		} `json:"error"`
		LogsUrl string `json:"logsUrl"`
	} `json:"submissionInfo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
