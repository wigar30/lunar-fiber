package user

import "lunar-commerce-fiber/internal/model"

type UserUseCase struct {
	userRepo model.UserRepositoryInterface
}

func NewUserUseCase(userRepo model.UserRepositoryInterface) model.UserUseCaseInterface {
	return &UserUseCase{
		userRepo: userRepo,
	}
}