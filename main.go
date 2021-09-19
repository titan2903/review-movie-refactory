package main

import (
	"fmt"
	"log"
	"review_movie/auth"
	"review_movie/entities"
	"review_movie/handler"
	"review_movie/middleware"
	"review_movie/repository"
	"review_movie/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPassword := myEnv["DB_PASSWORD"]
	dbHost := myEnv["DB_HOST"]
	dbName := myEnv["DB_NAME"]
	dbUsername := myEnv["DB_USERNAME"]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&entities.User{}, &entities.Review{}, &entities.MovieGenre{}, &entities.Movie{}, &entities.Genre{})

	//! Auth
	authService := auth.NewService()

	//! Users
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewServiceUser(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	genreRepository := repository.NewRepositoryGenre(db)
	genreService := service.NewServiceGenre(genreRepository)
	genreHandler := handler.NewGenreHandler(genreService)

	reviewRepository := repository.NewRepositoryReview(db)
	reviewService := service.NewServiceReview(reviewRepository)
	reviewHandler := handler.NewReviewHandler(reviewService)

	movieRepository := repository.NewRepositoryMovie(db)
	movieService := service.NewServiceMovie(movieRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	moviegenreRepository := repository.NewRepositoryMovieGenre(db)
	moviegenreService := service.NewServiceMovieGenre(moviegenreRepository)
	moviegenreHandler := handler.NewMovieGenreHandler(moviegenreService)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware()) // ! Allow cors

	apiUser := router.Group("/api/v1/users")
	apiUser.POST("/register", userHandler.RegisterUser)
	apiUser.POST("/login", userHandler.Login)
	apiUser.POST("/fetch", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser)

	apiMovieReview := router.Group("/api/v1/movie_reviews")
	apiMovieReview.PUT("/user", middleware.AuthMiddleware(authService, userService), userHandler.UpdateUser)
	apiMovieReview.GET("/user", middleware.AuthMiddleware(authService, userService), userHandler.GetUserByEmail)
	apiMovieReview.POST("/genre", middleware.AuthMiddleware(authService, userService), genreHandler.CreateGenre)
	apiMovieReview.GET("/genre", middleware.AuthMiddleware(authService, userService), genreHandler.GetGenres)
	apiMovieReview.GET("/movies", middleware.AuthMiddleware(authService, userService), movieHandler.GetMovies)
	apiMovieReview.POST("/movies", middleware.AuthMiddleware(authService, userService), movieHandler.CreateMovie)
	apiMovieReview.POST("/movies/genre", middleware.AuthMiddleware(authService, userService), moviegenreHandler.CreateMovieGenre)
	apiMovieReview.POST("/review", middleware.AuthMiddleware(authService, userService), reviewHandler.CreateReview)
	apiMovieReview.GET("/review", middleware.AuthMiddleware(authService, userService), reviewHandler.GetReviewByMovieID)

	router.Run(":5000")
}