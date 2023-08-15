package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	Username string `json:"username"`
	Password string `json:"password"`
	Uid      string `json:"uid"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("The token has expired. Please log in again")
	TokenNotValidYet = errors.New("The token is invalid. Please log in again")
	TokenMalformed   = errors.New("The token is incorrect. Please log in again")
	TokenInvalid     = errors.New("This is not a token, please log in again")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("token")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// 解析token
		claims, err := j.ParserToken(tokenHeader)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code":    errmsg.ERROR,
					"message": TokenExpired.Error(),
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"code":    errmsg.ERROR,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("username", claims)
		c.Next()
	}
}
