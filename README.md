## HTTP CMD

```go

import (
    "github.com/DGHeroin/httpcmd"
    "log"
)

func main() {
    r := httpcmd.NewEngine()
    r.GET("/status", func(c httpcmd.Context) {
        c.Response(200, httpcmd.F{
            "code": 0,
            "method":"get",
        })
    })
    r.POST("/p", func(c httpcmd.Context, fields httpcmd.F) {
        log.Println("bind:", fields)
        c.Response(200, httpcmd.F{
            "code": 0,
            "method":"post",
        })
    })
    r.Run()
}

```

```bash
curl -X POST 127.0.0.1:8080/p --data '{"name":"http cmd", "address":"github.com/DGHeroin/httpcmd"}' -H "Content-Type:application/json"
```