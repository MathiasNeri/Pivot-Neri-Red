package red

func Capitalize(s *string) string {
	result := ""
	isFirstChar := true
	for _, char := range *s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if isFirstChar {
				if char >= 'a' && char <= 'z' {
					char -= 32
				}
				isFirstChar = false
			} else {
				if char >= 'A' && char <= 'Z' {
					char += 32
				}
			}
			result += string(char)
		} else {
			isFirstChar = true
			result += string(char)
		}
	}
	return result
}

func isAlpha(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
}
