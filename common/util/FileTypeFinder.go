package util

import "regexp"

func FileTypeFinder(fileName string) string {
	re := regexp.MustCompile(`\.[^.]+$`)
	match := re.FindString(fileName)
	return match
}
