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
	icategory "github.com/phincon-backend/laza/internal/repo/category"
	iproduct "github.com/phincon-backend/laza/internal/repo/product"
	isize "github.com/phincon-backend/laza/internal/repo/size"
)

type UpdateProductUsecaseImpl struct {

	// Update to product table
	updateProductAction repositories.UpdateAction[model.Product]

	// get brand name
	getBrandName brand.GetByNameAction

	// Get size by name
	getSizeAction size.GetByNameAction

	// Get category by name
	getCategoryAction category.GetByNameAction
}

// Execute implements product.UpdateProductUsecase.
func (u *UpdateProductUsecaseImpl) Execute(id uint64, request requests.ProductRequest) (product model.Product, err error) {

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

	category, err := u.getCategoryByNameAction.GetByName(request.Category)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, errors.New("NotFound: Category not found")
	}

	file, err := request.Image.Open()
	defer file.Close()
	if err != nil {
		// TODO: Should return error here
		return
	}
	url, err := helper.UploadImageFile("product", file)
	if err != nil {
		// TODO: Should return error here
		return
	}

	product, err = u.updateProductAction.Update(id, model.Product{
		Name:        request.Name,
		Description: request.Description,
		ImageUrl:    url,
		Price:       request.Price,
		CategoryId:  category.Id,
		BrandId:     ps[0].BrandId, // didapat dari search by brand, ambil salah satu brand id
		Sizes:       sizeModels,
	})

	if err != nil {
		return
	}

	return
}

func NewUpdateProductUsecaseImpl(
	productRepo *iproduct.ProductRepo,
	sizeRepo *isize.SizeRepo,
	categoryRepo *icategory.CategoryRepo) usecase.UpdateProductUsecase {
	return &UpdateProductUsecaseImpl{
		updateProductAction:     productRepo,
		searchByBrandAction:     productRepo,
		searchByNameAction:      productRepo,
		getSizeAction:           sizeRepo,
		getCategoryByNameAction: categoryRepo,
	}
}
