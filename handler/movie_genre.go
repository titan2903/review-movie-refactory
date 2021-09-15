package handler

import (
	"net/http"
	"review_movie/formatter"
	"review_movie/helper"
	"review_movie/input"
	"review_movie/model"
	"review_movie/service"

	"github.com/gin-gonic/gin"
)

type movieGenreHandler struct {
	service service.ServiceMovieGenre
}

func NewMovieGenreHandler(service service.ServiceMovieGenre) *movieGenreHandler {
	return &movieGenreHandler{service}
}

func(mg *movieGenreHandler) CreateMovieGenre(c *gin.Context) {
	var input input.MovieGenreInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponseError("Failed Create Movie", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(model.User)
	if currentUser.Role != "admin" && currentUser.Role != "user" {
		response := helper.ApiResponseError("Role is not Admin", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	newMovieGenre, err := mg.service.CreateMovieGenre(input)
	if err != nil {
		response := helper.ApiResponseError("Failed Create movie", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	formatter := formatter.FormatCreateMovieGenreResponse(newMovieGenre)

	response := helper.ApiResponseGeneral("Sucessfully Created Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}