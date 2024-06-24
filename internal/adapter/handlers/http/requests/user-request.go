package requests

import "github.com/iffakhry/go-commerce-hexagonal/internal/core/domain"

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Role     string `json:"role" form:"role"`
}

func UserRequestToEntity(dataRequest UserRequest) domain.User {
	return domain.User{
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Name:     dataRequest.Name,
		Role:     dataRequest.Role,
	}
}
