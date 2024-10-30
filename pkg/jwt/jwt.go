package jwt

import (
	"admin-api/api/model"
	"admin-api/common/constant"
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	model.JwtAdmin
	jwt.StandardClaims
}

// 过期时间设置
const TokenExpireDuration = time.Hour * 24

// token秘钥
var Secret = []byte("admin-api")

var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

// GenerateTokenByAdmin 根据用户信息生成token
func GenerateTokenByAdmin(admin model.SysAdmin) (string, error) {
	var jwtAdmin = model.JwtAdmin{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}

	c := userStdClaims{
		jwtAdmin, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "backstage",                                // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// todo set redis
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*model.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtAdmin, err
}

// 返回id
func GetAdminId(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("can't get user id")
	}
	admin, ok := u.(*model.JwtAdmin)
	if ok {
		return admin.ID, nil
	}
	return 0, errors.New("can't convert to id struct")
}

// 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return string(string(0)), errors.New("can't get user name")
	}
	admin, ok := u.(*model.JwtAdmin)
	if ok {
		return admin.Username, nil
	}
	return string(string(0)), errors.New("can't convert to api name")
}

// 返回admin信息
func GetAdmin(c *gin.Context) (*model.JwtAdmin, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get api")
	}
	admin, ok := u.(*model.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}