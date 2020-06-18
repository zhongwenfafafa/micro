package validate

type Register struct {
	Username string `validate:"required" comment:"用户名"`
	Password string `validate:"required,gte=6" comment:"密码"`
	Mobile string `validate:"required" comment:"手机号"`
	TraceId string `validate:"required"`
}
