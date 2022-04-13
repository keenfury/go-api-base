package shared

import (
	"strconv"

	ae "github.com/keenfury/go-api-base/internal/api_error"
)

type (
	Output struct {
		JsonApi `json:"jsonapi"`
		Payload interface{} `json:"data,omitempty"`
		*Error  `json:"error,omitempty"`
	}

	JsonApi struct {
		Version string `json:"version"`
	}

	Error struct {
		Id     string `json:"Id,omitempty"`
		Title  string `json:"Title,omitempty"`
		Detail string `json:"Detail,omitempty"`
		Status string `json:"Status,omitempty"`
	}
)

func NewOutput(payload interface{}, apiError *ae.ApiError) Output {
	var err *Error
	if apiError != nil {
		err = &Error{Id: apiError.ApiErrorCode, Title: apiError.Title, Detail: apiError.Detail, Status: strconv.Itoa(apiError.StatusCode)}
	}
	output := Output{
		JsonApi: JsonApi{Version: "1.0"},
		Payload: payload,
		Error:   err,
	}
	return output
}
