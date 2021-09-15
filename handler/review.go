package handler

import (
	"fmt"
	"net/http"
	"review_movie/helper"
	"review_movie/review"
	"review_movie/user"

	"github.com/gin-gonic/gin"
)

type reviewHandler struct {
	service review.Service
}

func NewReviewHandler(service review.Service) *reviewHandler {
	return &reviewHandler{service}
}

func (r *reviewHandler) CreateReview(c *gin.Context) {
	var input review.CreateReviewInput
	fmt.Printf("data %v", input)
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponseError("Failed Create Review", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(user.User)
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

	formatter := review.FormatCreateReviewResponse(newReview)

	response := helper.ApiResponseGeneral("Sucessfully Add Reviews For this Movie!", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(r *reviewHandler) GetReviewByMovieID(c *gin.Context) {
	
}