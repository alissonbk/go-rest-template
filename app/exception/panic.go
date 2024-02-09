package exception

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		statusKey := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch statusKey {
		case constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[interface{}](constant.DataNotFound, msg, nil))
			c.Abort()
		case constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, dto.BuildResponse[interface{}](constant.Unauthorized, msg, nil))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[interface{}](constant.UnknownError, msg, nil))
			c.Abort()
		}
	}
}

func PanicException(statusKey constant.ResponseStatus, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", statusKey, err)
	if err != nil {
		panic(err)
	}
}
