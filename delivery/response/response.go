package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-basic/errors"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": errors.NORMAL.Code,
		"msg":  errors.NORMAL.Msg,
		"data": data,
	})
}

func Fail(c *gin.Context, err any) {
	if v, ok := err.(validator.ValidationErrorsTranslations); ok {
		e := errors.VALIDATION_ERROR
		e.Data = v
		_ = c.AbortWithError(http.StatusUnprocessableEntity, e)
	} else {
		_ = c.Error(err.(error))
	}
}
