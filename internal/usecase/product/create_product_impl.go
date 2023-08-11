package product

import (
	"context"
	"errors"
	"fmt"

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
	getBrandName brand.GetByNameActionWithContext

	// Get size by name
	getSizeAction size.GetByNameActionWithContext

	// Get category by name
	getCategoryAction category.GetByNameActionWithContext
}

// Execute implements product.CreateProductUsecase.
func (u *CreateProductUsecaseImpl) Execute(request requests.ProductRequest) (product model.Product, err error) {

	// Check if brand name exists
	// return error if false
	errorChan := make(chan error, 3)

	var brand model.Brand
	var sizes []model.Size
	var category model.Category

	var taskCount int = 0
	checkingContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	brandContext, _ := context.WithCancel(checkingContext)    // Cancel should be done on parent
	sizeContext, _ := context.WithCancel(checkingContext)     // Cancel should be done on parent
	categoryContext, _ := context.WithCancel(checkingContext) // Cancel should be done on parent

	// Check existing name on repo
	taskCount++
	go u.getBrandNameRepo(brandContext, request.Brand, &brand, errorChan, &taskCount)

	// Check existing size on repo
	for _, v := range request.Sizes {
		insideSizeContext, _ := context.WithCancel(sizeContext) // Cancel should be done on parent's parent
		taskCount++
		go u.getSizeRepo(insideSizeContext, v, &sizes, errorChan, &taskCount)
	}

	// Check existing category on repo
	taskCount++
	go u.getCategoryRepo(categoryContext, request.Category, &category, errorChan, &taskCount)

	for e := range errorChan {
		if e != nil {
			// EMERGENCY! CANCEL ALL ROUTINES
			cancel()
			fmt.Println("chan error: ", e.Error())
			err = e
			return
		}
		taskCount--
		if taskCount == 0 {
			break
		}
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
		Sizes:       sizes,
	})
	if err != nil {
		return
	}

	return
}

func (u *CreateProductUsecaseImpl) getBrandNameRepo(ctx context.Context, brandName string, brand *model.Brand, errChan chan<- error, taskCount *int) {
	defer func() { *taskCount-- }()

	var err error
	*brand, err = u.getBrandName.GetByNameWithContext(ctx, brandName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errChan <- errors.New("NotFound: Brand not found")
		return
	}
	errChan <- err
}

func (u *CreateProductUsecaseImpl) getSizeRepo(ctx context.Context, sizeName string, sizes *[]model.Size, errChan chan<- error, taskCount *int) {
	defer func() { *taskCount-- }()

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

func (u *CreateProductUsecaseImpl) getCategoryRepo(ctx context.Context, categoryName string, category *model.Category, errChan chan<- error, taskCount *int) {
	defer func() { *taskCount-- }()

	var err error
	*category, err = u.getCategoryAction.GetByNameWithContext(ctx, categoryName)
	// BUG: This sometimes not emitting gorm.ErrRecordNotFound when inputting
	//      nonexistent errors.
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errChan <- errors.New("NotFound: Category not found")
		return
	}
	errChan <- err
}

func NewCreateProductUsecaseImpl(
	insertProductAction repositories.InsertAction[model.Product],
	searchByBrandAction brand.GetByNameActionWithContext,
	getSizeAction size.GetByNameActionWithContext,
	getCategoryAction category.GetByNameActionWithContext) usecase.CreateProductUsecase {
	return &CreateProductUsecaseImpl{
		insertProductAction: insertProductAction,
		getBrandName:        searchByBrandAction,
		getSizeAction:       getSizeAction,
		getCategoryAction:   getCategoryAction,
	}
}
