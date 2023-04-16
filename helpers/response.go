package helpers

import (
	"net/http"

	"github.com/Digisata/dts-hactiv8-golang-chap3/models"
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, models.SuccessResult{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	})
}

func FailResponse(c *gin.Context, code int, err error) {
	switch code {
	case http.StatusUnauthorized:
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.UnauthorizedResult{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Message: err.Error(),
		})
	case http.StatusBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, models.BadRequestResult{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
	case http.StatusNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, models.NotFoundResult{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Message: err.Error(),
		})
	case http.StatusInternalServerError:
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.InternalServerErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
	}
}
