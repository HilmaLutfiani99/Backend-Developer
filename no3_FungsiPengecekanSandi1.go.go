package main

import (
	"fmt"
	"unicode"
)

func checkPassword3a(password string) error {
	// Check minimum length
	if len(password) < 8 {
		return fmt.Errorf("Kata sandi harus memiliki minimal 8 karakter")
	}

	// Check maximum length
	if len(password) > 32 {
		return fmt.Errorf("Kata sandi harus memiliki maksimal 32 karakter")
	}

	// Check first character is not a digit
	if unicode.IsDigit(rune(password[0])) {
		return fmt.Errorf("Karakter awal tidak boleh angka")
	}

	// Check for at least one digit
	hasDigit := false
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return fmt.Errorf("Kata sandi harus memiliki setidaknya satu angka")
	}

	// Check for at least one uppercase letter and one lowercase letter
	hasUpper := false
	hasLower := false
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		}
	}
	if !hasUpper || !hasLower {
		return fmt.Errorf("Kata sandi harus memiliki setidaknya satu huruf kapital dan satu huruf kecil")
	}

	return nil
}

func main() {
	password := "5andiwara"

	err := checkPassword3a(password)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Kata sandi valid")
	}
}
