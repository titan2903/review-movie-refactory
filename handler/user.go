package handler

import (
	"net/http"
	"review_movie/auth"
	"review_movie/formatter"
	"review_movie/helper"
	"review_movie/input"
	"review_movie/model"
	"review_movie/service"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.ServiceUser
	authService auth.Service
}

func NewUserHandler(userService service.ServiceUser, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func(h *userHandler) RegisterUser(c *gin.Context) {
	var input input.RegisterUserInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponseError("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Role != "admin" && input.Role != "user" {
		response := helper.ApiResponseError("Role value must admin or user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, _ := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponseError("Register account failed, Field Email, Password, FullName, Role is required!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUserRegisterResponse(newUser)

	response := helper.ApiResponseGeneral("Sucessfully Register!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(h *userHandler) Login(c *gin.Context) {
	var input input.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponseError("incorrect Username or Password", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponseError("incorrect Username or Password", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ApiResponseError("incorrect Username or Password", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	expireToken := time.Now().Add(time.Minute * 15)

	formatter := formatter.FormatUserLoginResponse(loggedinUser, token, expireToken,http.StatusOK)

	response := helper.ApiResponseLogin(formatter)
	c.JSON(http.StatusOK, response)
}

func(h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)
	formatter := formatter.FormatUserResponse(currentUser, "")

	response := helper.ApiResponseGeneral("Sucessfully Get Data!", "success", formatter)
	c.JSON(http.StatusOK, response)
}

func(h *userHandler) GetUserByEmail(c *gin.Context) {
	var input input.FindByEmailInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponseError("User not found", http.StatusNotFound, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	if currentUser.Role != "user" && currentUser.Role != "admin" {
		response := helper.ApiResponseError("Role is not user or admin", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	if currentUser.Email != input.Email {
		response := helper.ApiResponseError("Email not match with token", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	foundUser, _ := h.userService.FindUserByEmail(input)
	if err != nil {
		response := helper.ApiResponseError("User not found!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUserRegisterResponse(foundUser)

	response := helper.ApiResponseGeneral("Sucessfully Get Data!Sucessfully Get Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(h *userHandler) UpdateUser(c *gin.Context) {
	var inputUpdate input.UpdateUserInput

	err := c.ShouldBind(&inputUpdate)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponseError("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputEmail input.FindByEmailInput
	err = c.ShouldBindQuery(&inputEmail)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponseError("Email not found", http.StatusNotFound, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateInput := input.UpdateUserInput{}
	updateInput.FullName = inputUpdate.FullName
	updateInput.Email = inputUpdate.Email

	newUser, _ := h.userService.UpdateUser(updateInput, inputEmail)
	if err != nil {
		response := helper.ApiResponseError("Update Invalid", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUserRegisterResponse(newUser)

	response := helper.ApiResponseGeneral("Sucessfully Updated Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}