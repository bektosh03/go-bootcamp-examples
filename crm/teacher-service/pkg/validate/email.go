package validate

import "net/mail"

// Email checks whether given email string is in valid format
func Email(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	return nil
}
