package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"net/http"
	"server/model"
	"time"
)

var store = base64Captcha.NewMemoryStore(10240, 3*time.Minute)

// mathConfig 生成图形化算术验证码配置
func mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          100,
		Width:           300,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return mathType
}

// 生成验证码
func CaptchaImage(c *gin.Context) {
	var driver base64Captcha.Driver = mathConfig()
	// 创建验证码并传入创建的类型的配置，以及存储的对象
	b := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := b.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "error": err.Error(), "msg": "验证码生成失败"})
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"captchaEnabled": true, "msg": "操作成功", "code": 200, "img": b64s, "uuid": id}) //captchaEnabled和前端一起作用，用于开关验证码功能
}

// jwt登录鉴权
func Login(c *gin.Context) {
	var req model.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "error": err.Error(), "msg": "参数错误"})
		return
	}
	var captcha bool = store.Verify(req.Uuid, req.Code, true) // 验证验证码是否正确
	if req.Username != "admin" || req.PassWord != "admin123" || !captcha {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "用户名或密码或验证码错误"})
		return
	}
}
