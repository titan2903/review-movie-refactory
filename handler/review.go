package handler

import (
	"fmt"
	"net/http"
	"review_movie/formatter"
	"review_movie/helper"
	"review_movie/input"
	"review_movie/model"
	"review_movie/service"

	"github.com/gin-gonic/gin"
)

type reviewHandler struct {
	service service.ServiceReview
}

func NewReviewHandler(service service.ServiceReview) *reviewHandler {
	return &reviewHandler{service}
}

func (r *reviewHandler) CreateReview(c *gin.Context) {
	var input input.CreateReviewInput
	fmt.Printf("data %v", input)
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponseError("Failed Create Review", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(model.User)
	input.UserID = currentUser.ID

	if currentUser.Role != "user" {
		response := helper.ApiResponseError("Role is not user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	newReview, err := r.service.CreateReview(input)
	if err != nil {
		response := helper.ApiResponseError("Failed Create Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	formatter := formatter.FormatCreateReviewResponse(newReview)

	response := helper.ApiResponseGeneral("Sucessfully Add Reviews For this Movie!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(r *reviewHandler) GetReviewByMovieID(c *gin.Context) {
	
}