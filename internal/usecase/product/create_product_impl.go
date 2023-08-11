package product

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/repositories/brand"
	"github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/repositories/size"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"

	usecase "github.com/phincon-backend/laza/domain/usecases/product"
)

type CreateProductUsecaseImpl struct {

	// Insert to product table
	insertProductAction repositories.InsertAction[model.Product]

	// get brand name
	getBrandName brand.GetByNameAction

	// Get size by name
	getSizeAction size.GetByNameAction

	// Get category by name
	getCategoryAction category.GetByNameAction
}

// Execute implements product.CreateProductUsecase.
func (u *CreateProductUsecaseImpl) Execute(request requests.ProductRequest) (product model.Product, err error) {

	// Check if brand name exists
	// return error if false
	brand, err := u.getBrandName.GetByName(request.Brand)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, errors.New("NotFound: Size not found")
	}

	sizeModels := make([]model.Size, 0)
	for _, v := range request.Sizes {
		sz, err := u.getSizeAction.GetByName(v)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("NotFound: Size not found")
		}
		sizeModels = append(sizeModels, sz)
	}

	category, err := u.getCategoryAction.GetByName(request.Category)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, errors.New("NotFound: Category not found")
	}

	file, err := request.Image.Open()
	defer file.Close()
	if err != nil {
		return
	}
	url, err := helper.UploadImageFile("product", file)
	if err != nil {
		return
	}

	product, err = u.insertProductAction.Insert(model.Product{
		Name:        request.Name,
		Description: request.Description,
		ImageUrl:    url,
		Price:       request.Price,
		CategoryId:  category.Id,
		BrandId:     brand.Id, // didapat dari search by brand, ambil salah satu brand id
		Sizes:       sizeModels,
	})
	if err != nil {
		return
	}

	return
}

func NewCreateProductUsecaseImpl(
	insertProductAction repositories.InsertAction[model.Product],
	searchByBrandAction brand.GetByNameAction,
	getSizeAction size.GetByNameAction,
	getCategoryByNameAction category.GetByNameAction) usecase.CreateProductUsecase {
	return &CreateProductUsecaseImpl{
		insertProductAction: insertProductAction,
		getBrandName:        searchByBrandAction,
		getSizeAction:       getSizeAction,
		getCategoryAction:   getCategoryByNameAction,
	}
}
