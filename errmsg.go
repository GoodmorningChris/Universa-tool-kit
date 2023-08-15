package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// ERROR_USERNAME_USED code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	ERROR_CAPTCHA_WRONG    = 1009
	ERROR_PASSWORD_INVALID = 1010
	// code= 2000... 文章模块的错误
)

var codeMsg = map[int]string{
	SUCCESS:              "Success!",
	ERROR:                "Fail!",
	ERROR_USERNAME_USED:  "The user name already exists!",
	ERROR_PASSWORD_WRONG: "Wrong password!",

	ERROR_USER_NOT_EXIST:   "The user does not exist!",
	ERROR_TOKEN_EXIST:      "Token does not exist. Please log in again",
	ERROR_TOKEN_RUNTIME:    "The token has expired. Please log in again",
	ERROR_TOKEN_WRONG:      "The token is incorrect. Please log in again",
	ERROR_TOKEN_TYPE_WRONG: "Token format is incorrect. Please log in again",
	ERROR_USER_NO_RIGHT:    "The user has no permission",
	ERROR_CAPTCHA_WRONG:    "Verification code error",
	ERROR_PASSWORD_INVALID: "The password does not comply with the rules",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
