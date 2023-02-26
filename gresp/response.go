package gresp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/heqiaomu/gtools/gconsts"
	"github.com/heqiaomu/gtools/gerror"
	"github.com/heqiaomu/gtools/gvalidator"
	"net/http"
	"strings"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

func ReturnAny(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.JSON(httpCode, data)
}

//ReturnJsonFromString 将json字符窜以标准json格式返回（例如，从redis读取json格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

//Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, gconsts.CurdStatusOkCode, msg, data)
}
func SuccessAny(c *gin.Context, data interface{}) {
	ReturnAny(c, http.StatusOK, data)
}

//Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusAccepted, dataCode, msg, data)
	c.Abort()
}

// ErrorTokenBaseInfo token 基本的格式错误
func ErrorTokenBaseInfo(c *gin.Context) {
	ReturnJson(c, http.StatusAccepted, http.StatusBadRequest, gerror.ErrorsTokenBaseInfo, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//ErrorTokenAuthFail token 权限校验失败
func ErrorTokenAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusAccepted, http.StatusUnauthorized, gerror.ErrorsNoAuthorization, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//ErrorTokenRefreshFail token不符合刷新条件
func ErrorTokenRefreshFail(c *gin.Context) {
	ReturnJson(c, http.StatusAccepted, http.StatusUnauthorized, gerror.ErrorsRefreshTokenFail, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

//token 参数校验错误
func TokenErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusAccepted, gconsts.ValidatorParamsCheckFailCode, gconsts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// ErrorCasbinAuthFail 鉴权失败，返回 405 方法不允许访问
func ErrorCasbinAuthFail(c *gin.Context, msg interface{}) {
	ReturnJson(c, http.StatusAccepted, http.StatusMethodNotAllowed, gerror.ErrorsCasbinNoAuthorization, msg)
	c.Abort()
}

//ErrorParam 参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusAccepted, gconsts.ValidatorParamsCheckFailCode, gconsts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// ErrorSystem 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusAccepted, gconsts.ServerOccurredErrorCode, gconsts.ServerOccurredErrorMsg+msg, data)
	c.Abort()
}

// ValidatorError 翻译表单参数验证器出现的校验错误
func ValidatorError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		wrongParam := gvalidator.RemoveTopStruct(errs.Translate(gvalidator.Trans))
		ReturnJson(c, http.StatusAccepted, gconsts.ValidatorParamsCheckFailCode, gconsts.ValidatorParamsCheckFailMsg, wrongParam)
	} else {
		errStr := err.Error()
		// multipart:nextpart:eof 错误表示验证器需要一些参数，但是调用者没有提交任何参数
		if strings.ReplaceAll(strings.ToLower(errStr), " ", "") == "multipart:nextpart:eof" {
			ReturnJson(c, http.StatusAccepted, gconsts.ValidatorParamsCheckFailCode, gconsts.ValidatorParamsCheckFailMsg, gin.H{"tips": gerror.ErrorNotAllParamsIsBlank})
		} else {
			ReturnJson(c, http.StatusAccepted, gconsts.ValidatorParamsCheckFailCode, gconsts.ValidatorParamsCheckFailMsg, gin.H{"tips": errStr})
		}
	}
	c.Abort()
}
