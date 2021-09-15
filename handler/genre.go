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

type genreHandler struct {
	service service.ServiceGenre
}

func NewGenreHandler(service service.ServiceGenre) *genreHandler {
	return &genreHandler{service}
}

func(g *genreHandler) CreateGenre(c *gin.Context) {
	var input input.CreateGenreInput

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

	newGenre, err := g.service.CreateGenre(input)
	if err != nil {
		response := helper.ApiResponseError("Failed Create movie", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	formatter := formatter.FormatGenre(newGenre)

	response := helper.ApiResponseGeneral("Sucessfully Created Data!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(g *genreHandler) GetGenres(c *gin.Context) {

	genres, err := g.service.GetGenres()
	if err != nil {
		response := helper.ApiResponseError("Error to get genres", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponseGeneral("Successfully Get Genre List", "success", formatter.FormatGenres(genres))
	
	c.JSON(http.StatusOK, response)
}