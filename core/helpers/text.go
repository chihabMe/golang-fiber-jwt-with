package helpers

import "strings"

func Slugify(txt string) string {
	slug := strings.Replace(txt, " ", "-", -1)
	return slug
}
