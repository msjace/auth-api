package adapter

import (
	"github.com/msjace/auth-api/adapter/connection"
	"github.com/msjace/auth-api/adapter/handler"
	"github.com/msjace/auth-api/adapter/middleware"
	"github.com/msjace/auth-api/application/usecase"
	"github.com/msjace/auth-api/config"
	"github.com/msjace/auth-api/infrastructure/persistence"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func Server() {

	config.LoadConfig()

	dbClient := connection.ConnectDB()
	defer dbClient.Close()

	redisClient := connection.ConnectRedis()
	defer redisClient.Close()

	authPersistence := persistence.NewAuthRedisPersistence(dbClient, redisClient)
	userUseCase := usecase.NewUserUseCase(authPersistence)
	ah := handler.NewAuthHandler(userUseCase)

	http.Handle("/login", middleware.ChainMiddleware(http.HandlerFunc(ah.LoginHandler), []middleware.Middleware{middleware.CorsMiddleware}))
	http.Handle("/verify", middleware.ChainMiddleware(http.HandlerFunc(ah.VerifyHandler), []middleware.Middleware{middleware.CorsMiddleware}))
	http.Handle("/refresh", middleware.ChainMiddleware(http.HandlerFunc(ah.RefreshHandler), []middleware.Middleware{middleware.CorsMiddleware}))
	http.Handle("/logout", middleware.ChainMiddleware(http.HandlerFunc(ah.LogoutHandler), []middleware.Middleware{middleware.CorsMiddleware}))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), nil))
}
