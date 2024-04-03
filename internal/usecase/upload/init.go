package upload

import "lunar-commerce-fiber/internal/model"

type UploadUseCase struct {
}

func NewUploadUseCase() model.UploadUseCaseInterface {
	return &UploadUseCase{}
}
