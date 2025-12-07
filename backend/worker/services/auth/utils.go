package authService

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(plaintext *string) error {
	//Convert password string to byte slice
	var passwordBytes = []byte(*plaintext)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return errors.New("error hashing password")
	}

	hashedPassword := string(hashedPasswordBytes)

	*plaintext = hashedPassword

	return nil
}

func doPasswordsMatch(hashedPassword string, plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(plaintext))
	return true, err
}
