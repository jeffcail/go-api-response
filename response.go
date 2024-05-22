package go_api_response

import (
	"time"
)

var Response *result

type result struct {
	Lasting string      `json:"lasting"`
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r *result) ResponseSuccess(message string, data ...interface{}) *result {
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = data
	}

	if message == "" {
		message = "success"
	}
	return &result{
		Lasting: time.Now().Format(time.DateTime),
		Status:  true,
		Code:    1000,
		Msg:     message,
		Data:    d,
	}
}

func (r *result) ResponseError(code int, message string, data ...interface{}) *result {
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = data
	}

	return &result{
		Lasting: time.Now().Format(time.DateTime),
		Status:  false,
		Code:    code,
		Msg:     message,
		Data:    d,
	}
}

// echo frame usage
//func ToJson(c echo.Context, res *result) error {
//	return c.JSON(http.StatusOK, res)
//}

//
//func ToXml(c echo.Context, res *result) error {
//	return c.XML(http.StatusOK, res)
//}

var ResponsePage *pageList

type pageList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func (p *pageList) Pagination(count int64, list interface{}) *pageList {
	return &pageList{
		Total: count,
		List:  list,
	}
}
