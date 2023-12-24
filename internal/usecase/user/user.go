package user

import "lunar-commerce-fiber/internal/model"

func (uu UserUseCase) GetUserByID(ID int64) (*model.UserResponse, error) {
	user, err := uu.userRepo.GetUserByID(ID, true)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return &model.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		RoleID: user.RoleID,
		StatusID: user.StatusID,
		Role: user.Role,
		Status: user.Status,
	}, nil
}