package middlewares

import (
	"davinci/config"
	"davinci/module/user"

	. "davinci/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"davinci/pkg/utils"
	"net/http"
	"strings"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		// 定义不需要JWT验证的路径
		exceptPaths := map[string]bool{
			"/api/v1/login":          true,
			"/api/v1/register":       true,
			"/api/v3/login":          true,
			"/api/v3/register":       true,
			"/api/v3/configurations": true,
		}
		// 如果请求路径在白名单中，则不进行JWT验证，直接继续处理请求
		if _, ok := exceptPaths[path]; ok {
			c.Next()
			return
		}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt未获取"})
			ret := ResultWrapper(c)(nil, "未获取到jwt")(Unauthorized)
			c.JSON(http.StatusUnauthorized, ret)
			c.Abort()
			return
		}

		token := tokenString[7:]
		// 解析JWT令牌
		getToken, err := utils.ParseToken(token, []byte(config.SysYamlconfig.Jwt.PublicKey))

		if getToken != nil && getToken.Valid {
			//fmt.Println(getToken.Claims.(*utils.UserClaim).UserId)
			userId := getToken.Claims.(*utils.UserClaim).UserId

			// 将userId 转为int， 并调用 dao.DaoGetter.FindUserById
			userObj := user.NewModel()
			_, err := user.DaoGetter.FindUserById(userId, userObj)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
				c.Abort()
			} else {
				//// 生成 token
				prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
				curTime := time.Now().Add(time.Second * 60 * 60 * 24)
				token, _ := utils.GenerateToken(userObj.Id, prikey, curTime)

				c.Set("token", token)
				c.Set("auth_user", userObj)
				c.Next()
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "错误的token"})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("token过期或未启用")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token过期或未启用"})
			} else {
				fmt.Println("Couldn't handle this token:", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析此token"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析此token"})
		}
		c.Abort()
	}
}

// GetAuthUser 获取已经验证的用户
