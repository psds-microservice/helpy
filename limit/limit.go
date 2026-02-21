package limit

// ClampLimit returns limit clamped to [defaultIfZero, max].
// If limit <= 0, returns defaultIfZero. If limit > max, returns max.
// Shared by search-service and operator-directory-service for pagination.
func ClampLimit(limit, defaultIfZero, max int) int {
	if limit <= 0 {
		return defaultIfZero
	}
	if limit > max {
		return max
	}
	return limit
}
