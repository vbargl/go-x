package stringutil

// FirstNonEmpty returns first non empty string.
// It is useful for precomputed default values.
func FirstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
