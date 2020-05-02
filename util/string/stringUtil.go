package string

func HasStringEmpty(fields... string) bool {
	for _, field := range fields {
		if field == "" {
			return true
		}
	}
	return false
}
