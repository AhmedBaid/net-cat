package validators

import "unicode"

func Valid(name string) bool {
	for _, i := range name {
		if !unicode.IsLetter(i) {
			return false
		}
	}
	return true
}
