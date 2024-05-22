#### API-RESPONSE

#### Setup
1. install
```shell
go get github.com/jeffcail/go-api-response
```

2. usage An example for echo frame. other is same.
response.go
```go
import (
    github.com/jeffcail/go-api-response
)

func ToJson(c echo.Context) {
    return c.JSON(http.StatusOK, res)
}

```

router.go
```go

e.GET("/web/ping", func(c echo.Context) error {

    linkId := c.Get(_cons.SystemLinkId)
    logger.AdminLog.Info(fmt.Sprintf("%s%s", linkId, "/app/ping"))

    return _response.ToJson(c, _response.Res.ResponseSuccess("ping成功"))
})

```
