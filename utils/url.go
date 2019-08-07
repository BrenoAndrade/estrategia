package utils

// MakeURL a
func MakeURL(base string, adds ...string) string {
	var url string
	for _, add := range adds {
		url += "/" + add
	}
	return base + url
}
