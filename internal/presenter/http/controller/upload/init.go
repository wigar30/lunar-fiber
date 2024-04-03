package upload

import "lunar-commerce-fiber/internal/model"

type UploadController struct {
	uploadUseCase model.UploadUseCaseInterface
	config        *model.EnvConfigs
}

func NewUploadController(uploadUseCase model.UploadUseCaseInterface, config *model.EnvConfigs) *UploadController {
	return &UploadController{
		uploadUseCase: uploadUseCase,
		config:        config,
	}
}
