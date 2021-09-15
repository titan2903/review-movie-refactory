package handler

import (
	"net/http"
	"review_movie/helper"
	"review_movie/moviegenre"
	"review_movie/user"

	"github.com/gin-gonic/gin"
)

type movieGenreHandler struct {
	service moviegenre.Service
}

func NewMovieGenreHandler(service moviegenre.Service) *movieGenreHandler {
	return &movieGenreHandler{service}
}

func(mg *movieGenreHandler) CreateMovieGenre(c *gin.Context) {
	var input moviegenre.MovieGenreInput

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

	newMovieGenre, err := mg.service.CreateMovieGenre(input)
	if err != nil {
		response := helper.ApiResponseError("Failed Create movie", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	formatter := moviegenre.FormatCreateMovieGenreResponse(newMovieGenre)

	response := helper.ApiResponseGeneral("Sucessfully Created Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}