package validator

import "regexp"

func ValidateEmail(email string) (bool, error) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return regex.MatchString(email), nil
}
