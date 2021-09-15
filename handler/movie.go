package handler

import (
	"net/http"
	"review_movie/helper"
	"review_movie/movie"
	"review_movie/user"

	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	service movie.Service
}

func NewMovieHandler(service movie.Service) *movieHandler {
	return &movieHandler{service}
}

func(m *movieHandler) CreateMovie(c *gin.Context) {
	var input movie.CreateMovieInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponseError("Failed Create Movie", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(user.User)
	if currentUser.Role != "admin" && currentUser.Role != "user" {
		response := helper.ApiResponseError("Role is not Admin", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	newMovie, err := m.service.CreateMovie(input)
	if err != nil {
		response := helper.ApiResponseError("Failed Create movie", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	formatter := movie.FormatCreateMovieResponse(newMovie)

	response := helper.ApiResponseGeneral("Sucessfully Created Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(m *movieHandler) GetMovies(c *gin.Context) {
	
}