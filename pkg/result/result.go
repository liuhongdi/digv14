package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

//返回的结果：
type ResultCont struct {
	Code	int  `json:"code"`     //提示代码
	Msg string  `json:"msg"`       //提示信息
	Data interface{} `json:"data"` //出错
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

//成功
func (r *Result) Success(data interface{}) {
	if (data == nil) {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK,res)
}

//出错
func (r *Result)Error(code int,msg string) {
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusBadRequest,res)
	r.Ctx.Abort()
}
