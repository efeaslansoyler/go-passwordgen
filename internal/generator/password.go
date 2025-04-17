package generator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

const (
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	numbers      = "0123456789"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
)

type PasswordOptions struct {
	Length          int
	UseSpecialChars bool
	UseNumbers      bool
	UseUpper        bool
	UseLower        bool
	Count           int
}

type GeneratedPassword struct {
	Value    string
	Strength string
	Entropy  float64
}

func validateOptions(opt PasswordOptions) error {
	var minLength int
	if opt.UseSpecialChars {
		minLength++
	}
	if opt.UseNumbers {
		minLength++
	}
	if opt.UseUpper {
		minLength++
	}
	if opt.UseLower {
		minLength++
	}

	if opt.Length < minLength {
		return errors.New("length is too short for the selected character sets")
	}
	if opt.Count < 1 {
		return errors.New("count must be greater than 0")
	}
	if !opt.UseUpper && !opt.UseLower && !opt.UseNumbers && !opt.UseSpecialChars {
		return errors.New("at least one character set must be selected")
	}
	return nil
}

func shuffle(runes []rune) error {
	N := len(runes)
	for i := range N - 1 {
		r, err := secureRandomInt(N - i)
		if err != nil {
			return err
		}
		r = r + i
		runes[i], runes[r] = runes[r], runes[i]
	}
	return nil
}

func buildCharset(opt PasswordOptions) string {
	var charset strings.Builder
	if opt.UseUpper {
		charset.WriteString(uppercase)
	}
	if opt.UseLower {
		charset.WriteString(lowercase)
	}
	if opt.UseNumbers {
		charset.WriteString(numbers)
	}
	if opt.UseSpecialChars {
		charset.WriteString(specialChars)
	}
	return charset.String()
}

func secureRandomInt(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %w", err)
	}
	return int(n.Int64()), nil
}

func PasswordEntropy(password string) (float64, string, error) {
	if len(password) == 0 {
		return 0, "", errors.New("password is empty")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, r := range password {
		switch {
		case 'A' <= r && r <= 'Z':
			hasUpper = true
		case 'a' <= r && r <= 'z':
			hasLower = true
		case '0' <= r && r <= '9':
			hasNumber = true
		case strings.ContainsRune(specialChars, r):
			hasSpecial = true
		}
	}

	charsetSize := 0
	if hasUpper {
		charsetSize += len(uppercase)
	}
	if hasLower {
		charsetSize += len(lowercase)
	}
	if hasNumber {
		charsetSize += len(numbers)
	}
	if hasSpecial {
		charsetSize += len(specialChars)
	}
	if charsetSize == 0 {
		return 0, "", errors.New("password contains no recognized character types")
	}

	entropy := float64(len([]rune(password))) * math.Log2(float64(charsetSize))

	var strength string
	switch {
	case entropy >= 80:
		strength = "Excellent"
	case entropy >= 60:
		strength = "Strong"
	case entropy >= 40:
		strength = "Moderate"
	default:
		strength = "Weak"
	}

	return entropy, strength, nil
}

func GeneratePassword(opt PasswordOptions) ([]GeneratedPassword, error) {
	if err := validateOptions(opt); err != nil {
		return nil, err
	}

	charset := buildCharset(opt)
	charsetRunes := []rune(charset)
	passwords := make([]GeneratedPassword, opt.Count)

	for i := range passwords {
		password := make([]rune, opt.Length)
		position := 0

		if opt.UseUpper {
			n, err := secureRandomInt(len(uppercase))
			if err != nil {
				return nil, err
			}
			password[position] = rune(uppercase[n])
			position++
		}
		if opt.UseLower {
			n, err := secureRandomInt(len(lowercase))
			if err != nil {
				return nil, err
			}
			password[position] = rune(lowercase[n])
			position++
		}
		if opt.UseNumbers {
			n, err := secureRandomInt(len(numbers))
			if err != nil {
				return nil, err
			}
			password[position] = rune(numbers[n])
			position++
		}
		if opt.UseSpecialChars {
			n, err := secureRandomInt(len(specialChars))
			if err != nil {
				return nil, err
			}
			password[position] = rune(specialChars[n])
			position++
		}

		for j := position; j < opt.Length; j++ {
			n, err := secureRandomInt(len(charsetRunes))
			if err != nil {
				return nil, err
			}
			password[j] = charsetRunes[n]
		}

		if err := shuffle(password); err != nil {
			return nil, err
		}

		pwdStr := string(password)
		entropy, strength, err := PasswordEntropy(pwdStr)
		if err != nil {
			return nil, err
		}

		passwords[i] = GeneratedPassword{
			Value:    pwdStr,
			Strength: strength,
			Entropy:  entropy,
		}
	}

	return passwords, nil
}
