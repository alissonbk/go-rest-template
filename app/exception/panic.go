package exception

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"com.github.alissonbk/go-rest-template/app/constant"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func buildResponse(status constant.ResponseStatus, message string) map[string]string {
	return map[string]string{
		"response_status":  status.GetResponseStatus(),
		"response_message": message,
	}
}

// PanicHandler is responsible for auto returning a response with
// a predefined status code and message when a #PanicException occurs
func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		logrus.WithFields(logrus.Fields{
			"path":     c.FullPath(),
			"handlers": fmtContextHandlers(c.HandlerNames()),
			"params":   c.Params,
			"error":    err,
		}).Error("Panic occurred during request processing")

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
			c.JSON(http.StatusConflict, buildResponse(
				constant.DBDuplicatedKey,
				msg,
			))
		case constant.DBNoRowsAffected.GetNumber():
			c.JSON(http.StatusNotModified, nil)
		case constant.ParsingFailed.GetNumber():
			c.JSON(http.StatusBadRequest, buildResponse(
				constant.ParsingFailed,
				"Failed to parse you payload into an object"))
		case constant.DataNotFound.GetNumber():
			c.JSON(http.StatusBadRequest, buildResponse(constant.DataNotFound, msg))
		case constant.Unauthorized.GetNumber():
			c.JSON(http.StatusUnauthorized, buildResponse(constant.Unauthorized, msg))
		case constant.InvalidRequest.GetNumber():
			c.JSON(http.StatusBadRequest, buildResponse(constant.InvalidRequest, msg))
		default:
			c.JSON(http.StatusInternalServerError, buildResponse(
				constant.UnknownError,
				"interal error"))
		}
		c.Abort()
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
	if len(handlers) == 0 {
		return "No handlers in stack"
	}
	var sb strings.Builder
	sb.WriteString("Handler stack:\n")
	for i, handler := range handlers {
		sb.WriteString(fmt.Sprintf("\t%d. %s\n", i+1, handler))
	}

	return sb.String()
}

func GoroutinePanicHandler() {
	if err := recover(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Warn("Recovered from panic in a goroutine")
	}
}
