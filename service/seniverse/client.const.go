package seniverse

const apiUrlV3 = "https://api.seniverse.com/v3/"

const apiUrlV4 = "https://api.seniverse.com/v4"

const LogTable = "seniverse"

type ApiError struct {
	Status     string `json:"status"`
	StatusCode string `json:"status_code"` // 心知状态码
}
