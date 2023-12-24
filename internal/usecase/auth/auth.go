package auth

import (
	"lunar-commerce-fiber/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (au AuthUseCase) Login(c *fiber.Ctx, req *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := au.userRepo.GetUserByEmail(req.Email)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// println(err.Error())
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "Invalid Login Credentials. Please Try Again",
		}
	}

	roleID, err := strconv.Atoi(user.RoleID)
	if err != nil {
		return nil, err
	}

	token, err := GenerateAccessToken(&*au.config, user.ID, roleID)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken: token,
	}, nil
}