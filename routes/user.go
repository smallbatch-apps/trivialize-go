package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/smallbatch-apps/trivialize-go/database"
	"github.com/smallbatch-apps/trivialize-go/models"
)

type User struct {
	// this is not the user model, see this as the serialiser
	ID       uuid.UUID `json:"id" `
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func CreateResponseUser(user models.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(201).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id string, user *models.User) error {
	if len(id) == 0 {
		return errors.New("id is required")
	}
	database.Database.Db.Find(&user, "id = ?", id)

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if len(id) == 0 {
		return c.Status(400).JSON("id is required")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	database.Database.Db.First(&user, id)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
