package main

import (
	"fmt"
	"log"
	"review_movie/auth"
	"review_movie/genre"
	"review_movie/handler"
	"review_movie/middleware"
	"review_movie/movie"
	"review_movie/moviegenre"
	"review_movie/review"
	"review_movie/user"

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
	dsn := fmt.Sprintf("root:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{}, &review.Review{}, &moviegenre.MovieGenre{}, &movie.Movie{}, &genre.Genre{})

	//! Auth
	authService := auth.NewService()

	//! Users
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	genreRepository := genre.NewRepository(db)
	genreService := genre.NewService(genreRepository)
	genreHandler := handler.NewGenreHandler(genreService)

	reviewRepository := review.NewRepository(db)
	reviewService := review.NewService(reviewRepository)
	reviewHandler := handler.NewReviewHandler(reviewService)

	movieRepository := movie.NewRepository(db)
	movieService := movie.NewService(movieRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	moviegenreRepository := moviegenre.NewRepository(db)
	moviegenreService := moviegenre.NewService(moviegenreRepository)
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