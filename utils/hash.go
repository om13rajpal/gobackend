package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(inputPassword string, databasePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(inputPassword))
	return err == nil
}