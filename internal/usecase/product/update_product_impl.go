package product

import (
	"context"
	"errors"

	"sync/atomic"

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

type UpdateProductUsecaseImpl struct {

	// Update to product table
	updateProductAction repositories.UpdateAction[model.Product]

	// get brand name
	getBrandName brand.GetByNameActionWithContext

	// Get size by name
	getSizeAction size.GetByNameActionWithContext

	// Get category by name
	getCategoryAction category.GetByNameActionWithContext
}

// Execute implements product.UpdateProductUsecase.
func (u *UpdateProductUsecaseImpl) Execute(id uint64, request requests.ProductRequest) (product model.Product, err error) {

	// Check if brand name exists
	// return error if false
	errorChan := make(chan error, 3)

	var brand model.Brand
	var sizes []model.Size
	var category model.Category

	var taskCount atomic.Int32
	checkingContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Check existing name on repo
	taskCount.Add(1)
	go u.getBrandNameRepo(checkingContext, request.Brand, &brand, errorChan, &taskCount)

	// Check existing size on repo
	for _, v := range request.Sizes {
		taskCount.Add(1)
		go u.getSizeRepo(checkingContext, v, &sizes, errorChan, &taskCount)
	}

	// Check existing category on repo
	taskCount.Add(1)
	go u.getCategoryRepo(checkingContext, request.Category, &category, errorChan, &taskCount)

	for e := range errorChan {
		if e != nil {
			cancel()
			err = e
			return
		}
		if taskCount.Load() == 0 {
			close(errorChan)
			break
		}
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
		BrandId:     brand.Id, // didapat dari search by brand, ambil salah satu brand id
		Sizes:       sizes,
	})

	if err != nil {
		return
	}

	return
}

func (u *UpdateProductUsecaseImpl) getBrandNameRepo(ctx context.Context, brandName string, brand *model.Brand, errChan chan<- error, taskCount *atomic.Int32) {
	defer taskCount.Add(-1)

	var err error
	*brand, err = u.getBrandName.GetByNameWithContext(ctx, brandName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errChan <- errors.New("NotFound: Brand not found")
		return
	}
	errChan <- err
}

func (u *UpdateProductUsecaseImpl) getSizeRepo(ctx context.Context, sizeName string, sizes *[]model.Size, errChan chan<- error, taskCount *atomic.Int32) {
	defer taskCount.Add(-1)

	sz, err := u.getSizeAction.GetByNameWithContext(ctx, sizeName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errChan <- errors.New("NotFound: Size not found")
		return
	} else if err != nil {
		errChan <- err
		return
	}
	*sizes = append(*sizes, sz)
	errChan <- nil
}

func (u *UpdateProductUsecaseImpl) getCategoryRepo(ctx context.Context, categoryName string, category *model.Category, errChan chan<- error, taskCount *atomic.Int32) {
	defer taskCount.Add(-1)

	var err error
	*category, err = u.getCategoryAction.GetByNameWithContext(ctx, categoryName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errChan <- errors.New("NotFound: Category not found")
		return
	}
	errChan <- err
}

func NewUpdateProductUsecaseImpl(
	updateProductAction repositories.UpdateAction[model.Product],
	getBrandAction brand.GetByNameActionWithContext,
	getSizeAction size.GetByNameActionWithContext,
	getCategoryAction category.GetByNameActionWithContext) usecase.UpdateProductUsecase {
	return &UpdateProductUsecaseImpl{
		updateProductAction: updateProductAction,
		getBrandName:        getBrandAction,
		getSizeAction:       getSizeAction,
		getCategoryAction:   getCategoryAction,
	}
}
