package models

import "time"

type TruoraPostResponse struct {
	Check struct {
		CheckID      string    `json:"check_id"`
		CreationDate time.Time `json:"creation_date"`
	} `json:"check"`
}

type TruoraGetResponse struct {
	Check struct {
		CheckID       string    `json:"check_id"`
		Country       string    `json:"country"`
		CreationDate  time.Time `json:"creation_date"`
		PreviousCheck string    `json:"previous_check"`
		Score         float64   `json:"score"`
		Scores        []struct {
			DataSet  string  `json:"data_set"`
			Severity string  `json:"severity"`
			Score    float64 `json:"score"`
			Result   string  `json:"result"`
		} `json:"scores"`
	} `json:"check"`
}

type TruoraErrorResponse struct {
	Code     int    `json:"code"`
	HttpCode int    `json:"http_code"`
	Message  string `json:"message"`
}
