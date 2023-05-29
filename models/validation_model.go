package models

import "time"

type TruoraPostResponse struct {
	Check struct {
		CheckID      string `json:"check_id"`
		CreationDate time.Time `json:"creation_date"`
	} `json:"check"`
}

type TruoraGetResponse struct {
	Check struct {
		CheckID       string    `json:"check_id"`
		Country       string    `json:"country"`
		CreationDate  time.Time `json:"creation_date"`
		PreviousCheck string    `json:"previous_check"`
		Score         int       `json:"score"`
		Scores        []struct {
			DataSet  string `json:"data_set"`
			Severity string `json:"severity"`
			Score    int    `json:"score"`
			Result   string `json:"result"`
		} `json:"scores"`
	} `json:"check"`
}
