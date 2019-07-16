/**
  create by yy on 2019-07-16
*/

package gin

import (
	"net/http"
	"testing"
)

type ResponseSuccess struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessBack(ctx *Context, data interface{}) {
	rsp := &ResponseSuccess{
		Code: 0,
		Msg:  "",
		Data: data,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type TestController struct{}

func (test *TestController) Index(ctx *Context) {
	SuccessBack(ctx, "ok")
}

func TestAutoRouteExecute(t *testing.T) {
	ginServer := Default()
	ginServer.AutoRouter(&TestController{})
	ginServer.Run("127.0.0.1:8081")
}
