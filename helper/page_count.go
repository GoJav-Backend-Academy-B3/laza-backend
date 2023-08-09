package helper

func PageCount(total int64, limit int64) int64 {
	pages := total / limit
	if total%limit > 0 {
		pages++
	}

	return pages
}
