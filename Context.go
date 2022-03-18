package httpcmd

import "github.com/gin-gonic/gin"

type (
    Context interface {
        Response(httpStatusCode int, fields F)
    }
    context struct {
        gc *gin.Context
    }
    F map[string]interface{}
)

func (c *context) Response(httpStatusCode int, fields F) {
    c.gc.JSON(httpStatusCode, fields)
}
