package jwt

import (
	"sso/internal/domain/models"
	"sso/internal/libs/jwt/helpers_jwt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_jwt(t *testing.T) {
	user1 := models.User{
		ID:    123,
		Email: "test@email.com",
	}
	app1 := models.App{
		ID:     123,
		Secret: "test_secret",
	}
	duration1 := time.Hour

	t.Run("valid token creation", func(t *testing.T) {
		tokenString, err := NewToken(user1, app1, duration1)
		if err != nil {
			t.Fatalf("failed to create token %v", err)
		}
		claims, err := helpers_jwt.ParseToken(tokenString, app1.Secret)
		if err != nil {
			t.Fatalf("failed to parse token %v", err)
		}
		assert.Equal(t, claims["uid"], float64(user1.ID))
		assert.Equal(t, claims["email"], user1.Email)
		assert.Equal(t, claims["app_id"], float64(app1.ID))

	})
}
