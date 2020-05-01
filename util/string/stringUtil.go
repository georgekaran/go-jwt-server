package string

func IsAtLeastOneStringEmpty(fields... string) bool {
	for _, field := range fields {
		if field == "" {
			return true
		}
	}
	return false
}
