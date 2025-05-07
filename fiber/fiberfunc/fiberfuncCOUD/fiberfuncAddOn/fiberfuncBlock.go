package fiberfuncCOUDaddOn

import (
	"bre-api/addOn/colortext"
	"bre-api/fiber/fiberfunc/fiberfuncConfig"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func BlockFunc(c *fiber.Ctx) error {
	Ckey, err := getEnv()
	if err != nil {
		colortext.IsLogErr(err)
		c.Status(fiberfuncConfig.ErrorToEnv.Startus)
		return c.JSON(fiberfuncConfig.ErrorToEnv.Err)
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (any, error) {
		return []byte(Ckey), nil
	})
	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
