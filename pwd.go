package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPassword(password string) error {
	if len(password) >= 8 && len(password) <= 16 {
		rex := "~!@#$%^&*()_+\\={}\"[]:;?,./'"
		var a, b, c, d int = 0, 0, 0, 0
		for _, r := range password {
			if r == ' ' {
				return errors.New("The password cannot contain Spaces")
			}
			if unicode.IsUpper(r) {
				a = 1
			}
			if unicode.IsLower(r) {
				b = 1
			}
			if unicode.IsDigit(r) {
				c = 1
			}
			if strings.ContainsRune(rex, r) {
				d = 1
			}
		}
		if a+b+c+d >= 3 {
			return nil
		} else {
			return errors.New("The password strength is too low")
		}
	} else {
		return errors.New("The password length does not meet the rules")
	}

}
