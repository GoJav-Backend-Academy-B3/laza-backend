package review

import "github.com/phincon-backend/laza/domain/model"

func (r *ReviewRepo) GetReviewStatsByProduct(productID uint64) (averageRating float64, totalReviews int, err error) {
	var result struct {
		AverageRating float64
		TotalReviews  int
	}

	tx := r.db.
		Model(&model.Review{}).
		Select("AVG(rating) as average_rating, COUNT(*) as total_reviews").
		Where("product_id = ?", productID).
		Scan(&result)

	err = tx.Error
	if err != nil {
		return 0, 0, err
	}

	return result.AverageRating, result.TotalReviews, nil
}
