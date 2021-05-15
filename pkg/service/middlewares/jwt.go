package middlewares

import (
	"pishikot/pkg/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetJWTMiddleware(config models.Config) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Pishikot auth",
		Key:         []byte(config.Jwt.Key),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(models.User); ok {
				return jwt.MapClaims{
					"id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims["id"].(float64)
			return &models.User{
				ID:    uint(id),
				Login: claims["login"].(string),
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(
				code, gin.H{
					"message": message,
				},
			)
		},
		TokenLookup:   "header: Authorization, query:token",
		TimeFunc:      time.Now,
		TokenHeadName: "Bearer",
	})
	if err != nil {
		return nil, err
	}
	if err := authMiddleware.MiddlewareInit(); err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
