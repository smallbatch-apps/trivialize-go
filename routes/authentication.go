package routes

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smallbatch-apps/trivialize-go/models"
)

func GetSessions(c *fiber.Ctx) error {
	return c.SendString("Get Sessions")
}

func LoginUser(c *fiber.Ctx) (string, error) {
	email := c.Params("email")
	password := c.Params("password")

	var user models.User
	err := user.FindUserByEmail(email)

	if err != nil {
		return "", c.Status(401).SendString("User not authorised")
	}

	err = user.ComparePassword(password)

	if err != nil {
		return "", c.Status(401).SendString("User not authorised")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":      email,
			"company_id": user.CompanyId,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LogoutUser(c *fiber.Ctx) error {
	return c.SendString("Logout User")
}
