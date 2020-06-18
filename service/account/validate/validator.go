package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhongwen "github.com/go-playground/validator/v10/translations/zh"
	"micro/pkg"

	"micro/service/account/proto"
)

var validate *validator.Validate

func init() {
	zh2 := zh.New()
	uni := ut.New(zh2, zh2)
	trans, _ := uni.GetTranslator("zh")
	validate = validator.New()

	err := zhongwen.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		pkg.Logger.Error("register validator field error message translation Zh-CN failed")
	}
}

// 注册数据校验
func ValidatorRegister(req *proto.SignUpRequest) error {
	register := &Register{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		TraceId:  req.TranceId,
	}


	return validate.Struct(register)
}
