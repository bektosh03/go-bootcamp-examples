package validate

import (
	"fmt"
	"strings"
)

// PhoneNumber checks whether given number string is in valid format for UZ numbers
func PhoneNumber(number string) error {
	if len(number) != 13 {
		return fmt.Errorf("number should be 13 characters long, but was %d", len(number))
	}
	if !strings.HasPrefix(number, "+998") {
		return fmt.Errorf("number should have UZB prefix +998")
	}

	return nil
}
