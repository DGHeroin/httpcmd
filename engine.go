package httpcmd

import (
    "github.com/gin-gonic/gin"
)

type (
    Engine interface {
        Run(addr ...string) error
        RunTLS(addr, certFile, keyFile string) error
        GET(cmd string, handler func(c Context))
        POST(cmd string, handler func(c Context, fields F))
    }
    engine struct {
        r *gin.Engine
    }
)

func (e *engine) Run(addr ...string) error {
    return e.r.Run(addr...)
}
func (e *engine) RunTLS(addr, certFile, keyFile string) error {
    return e.r.RunTLS(addr, certFile, keyFile)
}
func (e *engine) GET(cmd string, handler func(c Context)) {
    r := e.r
    r.GET(cmd, e.middleWare, func(c *gin.Context) {
        ctx := &context{
            gc: c,
        }
        handler(ctx)
    })
}
func (e *engine) POST(cmd string, handler func(c Context, fields F)) {
    r := e.r
    r.POST(cmd, e.middleWare, func(c *gin.Context) {
        ctx := &context{
            gc: c,
        }
        fields := F{}
        if err := c.Bind(&fields); err != nil {
            c.JSON(200, gin.H{
                "code": 1,
                "err":  err.Error(),
            })
            return
        }
        handler(ctx, fields)
    })
}
func (e *engine) middleWare(c *gin.Context) {
    // TODO
}
func NewEngine() Engine {
    gin.SetMode(gin.ReleaseMode)
    eng := &engine{
        r: gin.New(),
    }
    eng.r.NoRoute(func(c *gin.Context) {
        c.JSON(200, gin.H{
            "code": -1,
            "err":  "cmd not found",
        })
    })
    return eng
}
