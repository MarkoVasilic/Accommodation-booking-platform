package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/repositories"
	generate "github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/tokens"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type PublicService struct {
	PublicRepository *repositories.PublicRepository
}

var Validate = validator.New()

func VerifyPassword(userpassword string, givenpassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenpassword), []byte(userpassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Password is incorrect"
		valid = false
	}
	return valid, msg
}

func (service *PublicService) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := Validate.Struct(user)
	if validationErr != nil {
		fmt.Println(validationErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	count, err := service.PublicRepository.CountByEmail(*user.Email)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	} else {
		inserterr := service.PublicRepository.CreateUser(&user)

		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			return
		}
		c.JSON(http.StatusCreated, "Successfully Signed Up!!")
	}
}

func (service *PublicService) VerifyUser(c *gin.Context) {
	var user models.User
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	founduser, err := service.PublicRepository.GetUserByEmail(*user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is incorrect"})
		return
	}
	PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
	defer cancel()
	if !PasswordIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		fmt.Println(msg)
		return
	}
	token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID, models.Role(*founduser.Role))
	defer cancel()
	fmt.Println(token)
	generate.UpdateAllTokens(service.PublicRepository.UserCollection, token, refreshToken, founduser.User_ID)
	c.JSON(http.StatusOK, founduser)
}

func (service *PublicService) GetUserById(c *gin.Context, id string) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	founduser, err := service.PublicRepository.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that id doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, founduser)
}
