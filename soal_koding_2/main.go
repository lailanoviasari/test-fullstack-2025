// soal koding 2

package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	_"github.com/redis/go-redis/v9"
)

func main() {
	app := fiber.New()

	// Login Endpoint
	app.Post("/login", login)

	// Public Endpoint
	app.Get("/", accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("login"),
	}))

	// Private Route 
	app.Get("/restricted", restricted)

	app.Listen(":3000")
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Lempar unsuthorized status
	if user != "AbertoDoniSianturi" || pass != "f7c3bc1d808e0" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// membuat claim
	claims := jwt.MapClaims {
		"realname" 	: "Aberto Doni Sianturi",
		"email" 	: "adss@gmail.com",
		"exp" 		: time.Now().Add(time.Hour*72).Unix(),
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("login"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token" : t})

} 

func accessible(c *fiber.Ctx) error {
	return c.SendString("Dapat diakses")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	realname := claims["realname"].(string)

	return c.SendString("Welcome " + realname)
}


