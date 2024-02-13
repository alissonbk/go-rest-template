// Package util_response is a util package to return pre-built responses
package util_response

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"net/http"
)

func InvalidJson() (httpStatus int, obj dto.Response[any]) {
	return http.StatusBadRequest, dto.BuildResponse[any](constant.InvalidRequest, "Failed to parse json into object", nil)
}

func InternalError(msg string) (httpStatus int, obj dto.Response[any]) {
	if msg == "" {
		msg = "Internal Error"
	}
	return http.StatusInternalServerError, dto.BuildResponse[any](constant.UnknownError, msg, nil)
}
