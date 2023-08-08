package product

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/repositories/product"
	"github.com/phincon-backend/laza/domain/repositories/size"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"

	usecase "github.com/phincon-backend/laza/domain/usecases/product"
	icategory "github.com/phincon-backend/laza/internal/repo/category"
	iproduct "github.com/phincon-backend/laza/internal/repo/product"
	isize "github.com/phincon-backend/laza/internal/repo/size"
)

type CreateProductUsecaseImpl struct {

	// Insert to product table
	insertProductAction repositories.InsertAction[model.Product]

	// Search by brand
	searchByBrandAction product.SearchByBrandAction

	// Get product name in table
	searchByNameAction product.SearchByNameAction

	// Get size by name
	getSizeAction size.GetByNameAction

	// TODO: Add Category repo to check by name
	getCategoryByNameAction category.GetByNameAction
}

// Execute implements product.CreateProductUsecase.
func (u *CreateProductUsecaseImpl) Execute(request request.ProductRequest) (product model.Product, err error) {

	// Check if brand name exists
	// return error if false
	ps, err := u.searchByBrandAction.SearchByBrand(request.Brand, 0, 1)
	if err != nil {
		// TODO: Handle errors
		return
	}
	if len(ps) == 0 {
		return product, errors.New("NotFound: Brand name not found")
	}

	sizeModels := make([]model.Size, 0)
	for _, v := range request.Sizes {
		sz, err := u.getSizeAction.GetByName(v)
		// FIXME: Should use devs made repos instead of
		// gorms, but ok
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

	product, err = u.insertProductAction.Insert(model.Product{
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

func NewCreateProductUsecaseImpl(
	productRepo *iproduct.ProductRepo,
	sizeRepo *isize.SizeRepo,
	categoryRepo *icategory.CategoryRepo) usecase.CreateProductUsecase {
	return &CreateProductUsecaseImpl{
		insertProductAction:     productRepo,
		searchByBrandAction:     productRepo,
		searchByNameAction:      productRepo,
		getSizeAction:           sizeRepo,
		getCategoryByNameAction: categoryRepo,
	}
}
