package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует заголовок авторизации"})
			return
		}

		// Проверяем формат токена
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Некорректный формат токена"})
			return
		}

		tokenString := bearerToken[1]

		// Проверяем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Валидация алгоритма
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Некорректный алгоритм подписи токена")
			}
			return []byte("secret"), nil
		})

		// Извлечение информации из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка при извлечении paylodad"})
		}

		// Извлечение user_id из claims
		UserIDfloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Не удалось извлечь user_id из токена"})
			return
		}
		userID := int(UserIDfloat)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Не удалось извлечь user_id из токена"})
			return
		}

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Некорректный или просроченный токен"})
			return
		}

		// Добавляем user_id в контекст
		c.Set("user_id", userID)

		// Если токен корректный, передаем обработку запроса дальше
		c.Next()
	}
}
