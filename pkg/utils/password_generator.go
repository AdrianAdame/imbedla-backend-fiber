package utils

import "golang.org/x/crypto/bcrypt"

// NormalizePassword func for a returning the users input as a byte slice.
func NormalizePassword(p string) []byte {
	return []byte(p)
}

// GeneratePassword func for a making hash & salt with user password.
func GeneratePassword(p string) string {
	// Normalize password from string to []byte.
	bytePwd := NormalizePassword(p)

	// MinCost is just an integer constant provided by the bcrypt package
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(hash)
}

// ComparePasswords func for a comparing password.
func ComparePasswords(hashedPwd, inputPwd string) bool {

	byteHash := NormalizePassword(hashedPwd)
	byteInput := NormalizePassword(inputPwd)

	// Return result.
	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
