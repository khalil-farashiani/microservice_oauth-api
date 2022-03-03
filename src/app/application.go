package app

import (
	"github.com/khalil-farashiani/microservice_oauth-api/src/domain/access_token"
	"github.com/khalil-farashiani/microservice_oauth-api/src/repository/db"
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
}
