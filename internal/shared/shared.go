package shared

import (
	"github.com/keenfury/go-api-base/internal/util"
)

type (
	Output struct {
		Success bool        `json:"success"`
		Payload interface{} `json:"payload"`
		Counts  Count       `json:"counts"`
	}

	Count struct {
		Count       int `json:"count"`
		PageSize    int `json:"page_size"`
		CurrentPage int `json:"current_page"`
		PageCount   int `json:"pageCount"`
	}
)

func NewOutput(payload interface{}, success bool, pageSize, currentPage, pageCount int) Output {
	count := util.GetTypeCount(payload)
	output := Output{
		Success: success,
		Payload: payload,
		Counts: Count{
			Count:       count,
			PageSize:    pageSize,
			CurrentPage: currentPage,
			PageCount:   pageCount,
		},
	}
	return output
}
