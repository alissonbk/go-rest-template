package exception

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// PanicHandler is responsible for auto returning a response with
// a predefined status code and message when a #PanicException occurs
func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		logrus.Errorf("Error occurred in the following context:\npath: %s\n%s\nparams: %s", c.FullPath(), fmtContextHandlers(c.HandlerNames()), c.Params)
		str := fmt.Sprint(err)

		var statusKey int
		var msg string
		if isPanicException(str) {
			strArr := strings.Split(str, ":")
			strArr[0] = strings.Split(strArr[0], "PanicException ")[1]
			statusKey, err = strconv.Atoi(strArr[0])
			if err != nil {
				statusKey = constant.UnknownError.GetNumber()
			}

			msg = strings.Trim(strings.Join(strArr[1:], " "), " ")
			logrus.Error(msg)
		} else {
			logrus.Warn("this panic was not handled by PanicException")
			panic(err)
		}

		switch statusKey {
		case constant.DBDuplicatedKey.GetNumber():
			c.JSON(http.StatusConflict, dto.BuildResponse[interface{}](
				constant.DBDuplicatedKey,
				msg,
			))
		case constant.DBNoRowsAffected.GetNumber():
			c.JSON(http.StatusNotModified, nil)
		case constant.ParsingFailed.GetNumber():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](
				constant.ParsingFailed,
				"Failed to parse you payload into an object"))
			c.Abort()
		case constant.DataNotFound.GetNumber():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](constant.DataNotFound, msg))
			c.Abort()
		case constant.Unauthorized.GetNumber():
			c.JSON(http.StatusUnauthorized, dto.BuildResponse[interface{}](constant.Unauthorized, msg))
			c.Abort()
		case constant.InvalidRequest.GetNumber():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](constant.InvalidRequest, msg))
		default:
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[interface{}](
				constant.UnknownError,
				"interal error"))
			c.Abort()
		}
	}
}

// TODO: define when the panic handler should expose the message or not (to avoid exposing server information to the client)
func PanicException(statusKey constant.ResponseStatus, message string) {
	err := errors.New(message)
	err = fmt.Errorf("PanicException %d: %w", statusKey.GetNumber(), err)
	if err != nil {
		panic(err)
	}
}

func isPanicException(errStr string) bool {
	split := strings.Split(errStr, " ")
	return len(split) > 0 && split[0] == "PanicException"
}

func fmtContextHandlers(handlers []string) string {
	str := "handlers stack:"
	for _, h := range handlers {
		str += "\n\t->" + h
	}

	return str
}

func GoroutinePanicHandler() {
	if err := recover(); err != nil {
		logrus.Warn("recovering from panic in a goroutine")
		logrus.Error(err)
	}
}
