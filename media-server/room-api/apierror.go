package main

import "net/http"

const (
	APIBadParameter          = iota
	APIResourceUnavailable   = iota
	APIUnknownServiceFailure = iota
)

// ErrorDetail details on one API error in a list
type ErrorDetail struct {
	Code    int    `json:"code"`
	Field   string `json:"field"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// APIErrorMessage normalized error message for all API calls
type APIErrorMessage struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Details []*ErrorDetail `json:"details"`
}

func Error(errorCode int, field string, value string, message string) *APIErrorMessage {
	e := new(APIErrorMessage)
	e.Code = errorCode
	e.Message = message
	var errorDetails []*ErrorDetail
	errorDetail := new(ErrorDetail)
	errorDetail.Code = errorCode
	errorDetail.Field = field
	errorDetail.Value = value
	errorDetail.Message = message
	errorDetails = append(errorDetails, errorDetail)
	e.Details = errorDetails
	return e
}

func resolveServiceError(serviceError *ServiceError) func(field string, value string) (int, *APIErrorMessage) {
	return func(field string, value string) (int, *APIErrorMessage) {
		switch serviceError.StatusCode {
		case ServerNotAvailable:
			return http.StatusServiceUnavailable,
				Error(APIResourceUnavailable, field, value, serviceError.Error())
		case RoomNotFound:
			return http.StatusNotFound,
				Error(APIBadParameter, field, value, serviceError.Error())
		case ServerFull:
			return http.StatusServiceUnavailable,
				Error(APIResourceUnavailable, field, value, serviceError.Error())
		case RoomFull:
			return http.StatusServiceUnavailable,
				Error(APIResourceUnavailable, field, value, serviceError.Error())
		case RemoteCallFailed:
			return http.StatusBadGateway,
				Error(APIUnknownServiceFailure, field, value, serviceError.Error())
		}
		return http.StatusInternalServerError,
			Error(APIUnknownServiceFailure, field, value, serviceError.Error())
	}
}
