package modules

import "golang.org/x/crypto/bcrypt"

func HasingPassword(s string) (string, error) {
	password := []byte(s)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	
	if (err != nil) {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(input string, current string) (bool) {
	inputPassword := []byte(input)
	currentPassword := []byte(current)
	err := bcrypt.CompareHashAndPassword(currentPassword, inputPassword)

	return err == nil
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}