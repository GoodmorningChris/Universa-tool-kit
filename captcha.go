package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	captcha "github.com/mojocn/base64Captcha"
	"net/http"
)

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverString {
	driver := new(captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 5
	driver.ShowLineOptions = captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine
	driver.Length = 4
	driver.Source = ""
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

func GenerateCaptchaHandler(c *gin.Context) {
	var driver = NewDriver().ConvertFonts()
	d := captcha.NewCaptcha(driver, store)
	_, content, answer := d.Driver.GenerateIdQuestionAnswer()
	toke := uuid.New()
	item, _ := d.Driver.DrawCaptcha(content)
	d.Store.Set(toke.String(), answer)
	//item.WriteTo(w)
	c.JSON(
		http.StatusOK, gin.H{
			"code":    errmsg.SUCCESS,
			"captcha": item.EncodeB64string(),
			"token":   toke,
		},
	)
}

// 验证
func CaptchaVerifyHandle(token string, code string) int {
	if store.Verify(token, code, true) && token != "" && code != "" {
		return errmsg.SUCCESS
	} else {
		return errmsg.ERROR_CAPTCHA_WRONG
	}
}
