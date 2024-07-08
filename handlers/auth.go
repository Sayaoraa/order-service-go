package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"order-tracking/config"
	"order-tracking/models"
	"time"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("secret")

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	config.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return echo.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return echo.ErrUnauthorized
	}

	claims := &JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя с указанными username и password
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Имя пользователя"
// @Param password formData string true "Пароль пользователя"
// @Success 200 {object} models.User
// @Router /register [post]
func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	config.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}
