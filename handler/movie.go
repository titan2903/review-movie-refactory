package handler

import (
	"net/http"
	"review_movie/entities"
	"review_movie/formatter"
	"review_movie/helper"
	"review_movie/input"
	"review_movie/service"

	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	service service.ServiceMovie
}

func NewMovieHandler(service service.ServiceMovie) *movieHandler {
	return &movieHandler{service}
}

func(m *movieHandler) CreateMovie(c *gin.Context) {
	var input input.CreateMovieInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponseError("Failed Create Movie", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(entities.User)
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

	formatter := formatter.FormatCreateMovieResponse(newMovie)

	response := helper.ApiResponseGeneral("Sucessfully Created Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(m *movieHandler) GetMovies(c *gin.Context) {
	
}