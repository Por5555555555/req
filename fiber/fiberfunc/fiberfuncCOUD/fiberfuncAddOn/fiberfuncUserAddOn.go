package fiberfuncCOUDaddOn

import (
	"bre-api/addOn/colortext"
	"bre-api/config"
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/grom/databases"
	"bre-api/grom/handler"
	"bre-api/grom/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var key string
var CookieLifeTime = config.CookieLifeTime

func CreateUser(c *fiber.Ctx) error {
	//New User and get from json
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}

	//Password to hashFunc
	password, err := hashPassword(*user)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToHashPassword.Startus)
		return c.JSON(fiberfuncConfig.ErrorToHashPassword.Err)
	}
	user.PasswordUser = string(password)

	_, err = handler.Create(user)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

func LoginUser(c *fiber.Ctx) error {
	//get User with email and get data from json
	var input models.User
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}

	db := databases.GetConn()
	if err := db.Where("Email = ?", input.Email).First(&user).Error; err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordUser), []byte(input.PasswordUser)); err != nil {
		fmt.Println(err)
		c.Status(fiberfuncConfig.ErrorToHashPassword.Startus)
		return c.JSON("Sus")
	}

	colortext.IsOk("Ok")
	//createToken token
	tokenValue, err := createToken(user)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToEnv.Startus)
		colortext.IsWanging(err)
		colortext.IsLogErr("Bad Config Env")
		return c.JSON("Error to Bad Config Server")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokenValue,
		Expires:  time.Now().Add(CookieLifeTime),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "None", // or "None" if using cross-site cookies
		// origin
		// Lax
		// None
	})

	return c.JSON(fiber.Map{"message": "login success"})
}

func hashPassword(user models.User) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.PasswordUser), bcrypt.DefaultCost)
	return string(password), err
}

func getEnv() (string, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		return "", err
	}
	ok := os.Getenv("what")
	if ok == "" {
		return "", errors.New("Error Get Env")
	}
	return ok, nil
}

func createToken(user models.User) (string, error) {
	key, err := getEnv()
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(CookieLifeTime).Unix()

	t, err := token.SignedString([]byte(key))
	return t, err
}
func CheckJwt(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token"})
	}

	secret, err := getEnv()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot load secret"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	return c.JSON(fiber.Map{"message": "Authenticated"})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "None", // or "None" if using cross-site cookies
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "logout success"})
}
