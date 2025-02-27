package validators

func ValidateLength(str string) bool {
	if len(str) < 3 || len(str) > 15 {
		return false
	}
	return true
}
