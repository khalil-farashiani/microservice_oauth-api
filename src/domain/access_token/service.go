package access_token

import "github.com/khalil-farashiani/microservice_oauth-api/src/utils/errors"

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func (s *service) GetByID(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}

func NewService(repository Repository) Service {
	return &service{}
}
