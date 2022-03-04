package app

import (
	"github.com/gin-gonic/gin"
	"github.com/khalil-farashiani/microservice_oauth-api/src/clients/cassandra"
	"github.com/khalil-farashiani/microservice_oauth-api/src/domain/access_token"
	"github.com/khalil-farashiani/microservice_oauth-api/src/http"
	"github.com/khalil-farashiani/microservice_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if err != nil {
		panic(dbErr)
	}
	session.Close()
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
