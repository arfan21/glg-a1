package middleware

import (
	"net/http"
	"strconv"

	"github.com/arfan21/hacktiv8-golang-jwt/database"
	"github.com/arfan21/hacktiv8-golang-jwt/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := model.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
