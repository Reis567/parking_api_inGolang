package repository

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

func (us *userRepository) FindUserByEmail(email string) (model.UserDomainInterface *rest_err.RestErr) {
	return nil,nil
}