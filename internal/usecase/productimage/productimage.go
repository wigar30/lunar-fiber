package productimage

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (pu *ProductImageUseCase) CreateBulkProductImage(c *fiber.Ctx, productImages []model.CreateProductImage) (bool, error) {
	tx := pu.productImageRepo.BeginTransaction()

	for _, p := range productImages {
		if p.ImageFile != nil {
			resp, err := pu.uploadUseCase.UploadImage(c, p.ImageFile, "product_image")
			if errC, ok := err.(*model.ErrorResponse); ok {
				pu.productImageRepo.RollbackTransaction(tx)
				return false, &model.ErrorResponse{
					Code:    errC.Code,
					Message: errC.Error(),
				}
			}

			p.Image = resp.FilePath
		}
	}

	err := pu.productImageRepo.CreateBulkProductImage(tx, productImages)
	if errC, ok := err.(*model.ErrorResponse); ok {
		pu.productImageRepo.RollbackTransaction(tx)
		return false, &model.ErrorResponse{
			Code:    errC.Code,
			Message: errC.Error(),
		}
	}

	return true, pu.productImageRepo.CommitTransaction(tx)
}
