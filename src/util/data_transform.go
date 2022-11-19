package util

func StringListToString(stringList []string) string {
	result := ""
	for _, str := range stringList {
		result += str
	}
	return result
}
