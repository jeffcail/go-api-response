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

var result = goApiResponse.Response
var resultList = goApiResponse.ResponsePage

func ToJsonSuccess(c echo.Context, msg string, data ...interface{}) error {
    return c.JSON(http.StatusOK, result.ResponseSuccess(msg))
}

func ToJsonSuccessList(c echo.Context, msg string, count int64, data ...interface{}) error {
    var d interface{}
    if data != nil {
        d = data[0]
    } else {
        d = data
    }
    return c.JSON(http.StatusOK, result.ResponseSuccess(msg, resultList.Pagination(count, d)))
}

func ToJsonFailed(c echo.Context, code int, msg string) error {
    return c.JSON(http.StatusOK, result.ResponseError(code, msg))
}


```

router.go
```go

type A struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

e.GET("/web/ping", func(c echo.Context) error {

    linkId := c.Get(_cons.SystemLinkId)
    logger.AdminLog.Info(fmt.Sprintf("%s%s", linkId, "/app/ping"))

    //return _response.ToJsonSuccess(c, "")
    //return _response.ToJsonFailed(c, 5000, "failed")
    
    arr := make([]*A, 0)
    a1 := &A{
    Name:  "qq",
    Email: "qq.com",
    }
    a2 := &A{
    	Name:  "163",
    	Email: "163.com",
    }
    arr = append(arr, a1, a2)
    
    return _response.ToJsonSuccessList(c, "success", 2, arr)

    return _response.ToJson(c, _response.Res.ResponseSuccess("ping成功"))
})

```
