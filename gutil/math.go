package gutil

import (
	"github.com/spf13/cast"
	"math/rand"
)
import "time"
import "fmt"

func RandNumLen6() int64 {
	rand.Seed(time.Now().UnixNano())
	captchaLength := 6
	captcha := ""
	for i := 0; i < captchaLength; i++ {
		captcha += fmt.Sprintf("%d", rand.Intn(10))
	}
	return cast.ToInt64(captcha)
}
