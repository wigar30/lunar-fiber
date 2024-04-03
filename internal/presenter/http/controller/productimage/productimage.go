package productimage

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (pc *ProductImageController) CreateProductImage(c *fiber.Ctx) error {
	var payloads []model.CreateProductImage

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	for _, files := range form.File {
		// Loop through each file in the files slice
		for a, file := range files {
			// Open the file
			f, err := file.Open()
			if err != nil {
				return err
			}
			defer f.Close()

			// Read the text from the form field
			tenantId := form.Value["tenantId"]
			productId := form.Value["productId"]
			titleID := form.Value["title"]

			// Create a new FormData instance and populate it with the text and file data
			payload := model.CreateProductImage{
				TenantID:  tenantId[a],
				ProductID: productId[a],
				StatusID:  "1",
				Title:     titleID[a],
				ImageFile: file,
			}

			// Append the FormData instance to the payloads slice
			payloads = append(payloads, payload)
		}
	}

	resp, err := pc.productImageUseCase.CreateBulkProductImage(c, payloads)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
