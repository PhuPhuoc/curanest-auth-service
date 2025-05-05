package common

import (
	"encoding/json"
)

func ExtractErrorFromResponse(response map[string]interface{}) error {
	var appErr AppError
	jsonData, err := json.Marshal(response["error"])
	if err != nil {
		resp := NewInternalServerError().WithReason("failed to marshal error response: " + err.Error())
		return resp
	}

	if err := json.Unmarshal(jsonData, &appErr); err != nil {
		resp := NewInternalServerError().WithReason("failed to unmarshal error response: " + err.Error())
		return resp
	}

	return &appErr
}

func ExtractDataFromResponse[T any](response map[string]interface{}, key string) (*T, error) {
	rawData, ok := response[key].(map[string]interface{})
	if !ok {
		resp := NewInternalServerError().WithReason("data response is not in expected format")
		return nil, resp
	}

	var resp T
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		respErr := NewInternalServerError().WithReason("failed to marshal error response: " + err.Error())
		return nil, respErr
	}

	if err := json.Unmarshal(jsonData, &resp); err != nil {
		respErr := NewInternalServerError().WithReason("failed to unmarshal error response: " + err.Error())
		return nil, respErr
	}

	return &resp, nil
}

func ExtractListDataFromResponse[T any](response map[string]interface{}, key string) ([]T, error) {
	rawData, ok := response[key].([]interface{})
	if !ok {
		respErr := NewInternalServerError().WithReason("data response is not in expected format")
		return nil, respErr
	}

	var respList []T
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		respErr := NewInternalServerError().WithReason("failed to marshal error response: " + err.Error())
		return nil, respErr
	}

	if err := json.Unmarshal(jsonData, &respList); err != nil {
		respErr := NewInternalServerError().WithReason("failed to unmarshal error response: " + err.Error())
		return nil, respErr
	}

	return respList, nil
}
