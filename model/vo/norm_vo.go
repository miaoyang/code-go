package vo

// CaptchaVo 验证码返回vo
type CaptchaVo struct {
	CaptchaId         string `yaml:"captchaId"`
	CaptchaImgUrl     string `yaml:"captchaImgUrl"`
	CaptchaRefreshUrl string `yaml:"captchaRefreshUrl"`
}
