package exception

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// PanicHandler is responsible for auto returning a response with
// a predefined status code and message when a #PanicException occurs
func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		statusKey, err := strconv.Atoi(strArr[0])
		if err != nil {
			statusKey = constant.UnknownError.GetNumber()
		}

		msg := strings.Trim(strArr[1], " ")

		switch statusKey {
		case constant.DBNoRowsAffected.GetNumber():
			c.JSON(http.StatusNotModified, nil)
		case constant.ParsingFailed.GetNumber():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](
				constant.ParsingFailed,
				"Failed to parse you payload into an object",
				nil))
			c.Abort()
		case constant.DataNotFound.GetNumber():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](constant.DataNotFound, msg, nil))
			c.Abort()
		case constant.Unauthorized.GetNumber():
			c.JSON(http.StatusUnauthorized, dto.BuildResponse[interface{}](constant.Unauthorized, msg, nil))
			c.Abort()
		case constant.InvalidRequest.GetNumber():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](constant.InvalidRequest, msg, nil))
		default:
			fmt.Println(statusKey)
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[interface{}](
				constant.UnknownError,
				msg,
				nil))
			c.Abort()
		}
	}
}

func PanicException(statusKey constant.ResponseStatus, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%d: %w", statusKey.GetNumber(), err)
	if err != nil {
		panic(err)
	}
}
