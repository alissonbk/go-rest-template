package dto

import "com.github.alissonbk/go-rest-template/app/constant"

type Response[T any] struct {
	ResponseStatus  string `json:"response_status"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
}

func BuildResponse[T any](status constant.ResponseStatus, message string, data T) Response[T] {
	if message == "" {
		message = status.GetResponseStatus()
	}
	return Response[T]{
		ResponseStatus:  status.GetResponseStatus(),
		ResponseMessage: message,
		Data:            data,
	}
}
