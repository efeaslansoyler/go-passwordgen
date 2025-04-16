package generator

import (
	"strings"
	"testing"
	"unicode"
)

func containsAny(s string, set string) bool {
	for _, r := range s {
		if strings.ContainsRune(set, r) {
			return true
		}
	}
	return false
}

func TestGeneratePassword_Basic(t *testing.T) {
	opt := PasswordOptions{
		Length:          12,
		UseSpecialChars: true,
		UseNumbers:      true,
		UseUpper:        true,
		UseLower:        true,
		Count:           5,
	}
	passwords, err := GeneratePassword(opt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(passwords) != 5 {
		t.Fatalf("expected 5 passwords, got %d", len(passwords))
	}
	for _, pwd := range passwords {
		if len([]rune(pwd)) != 12 {
			t.Errorf("expected password length 12, got %d", len([]rune(pwd)))
		}
		if !containsAny(pwd, specialChars) {
			t.Errorf("password missing special char: %s", pwd)
		}
		if !containsAny(pwd, numbers) {
			t.Errorf("password missing number: %s", pwd)
		}
		if !containsAny(pwd, uppercase) {
			t.Errorf("password missing uppercase: %s", pwd)
		}
		if !containsAny(pwd, lowercase) {
			t.Errorf("password missing lowercase: %s", pwd)
		}
	}
}

func TestGeneratePassword_OnlyNumbers(t *testing.T) {
	opt := PasswordOptions{
		Length:          8,
		UseSpecialChars: false,
		UseNumbers:      true,
		UseUpper:        false,
		UseLower:        false,
		Count:           2,
	}
	passwords, err := GeneratePassword(opt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, pwd := range passwords {
		for _, r := range pwd {
			if !unicode.IsDigit(r) {
				t.Errorf("expected only digits, got: %q", r)
			}
		}
	}
}

func TestGeneratePassword_OnlyUpper(t *testing.T) {
	opt := PasswordOptions{
		Length:          10,
		UseSpecialChars: false,
		UseNumbers:      false,
		UseUpper:        true,
		UseLower:        false,
		Count:           1,
	}
	passwords, err := GeneratePassword(opt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, pwd := range passwords {
		for _, r := range pwd {
			if !unicode.IsUpper(r) {
				t.Errorf("expected only uppercase, got: %q", r)
			}
		}
	}
}

func TestGeneratePassword_InvalidOptions(t *testing.T) {
	opt := PasswordOptions{
		Length:          8,
		UseSpecialChars: false,
		UseNumbers:      false,
		UseUpper:        false,
		UseLower:        false,
		Count:           1,
	}
	_, err := GeneratePassword(opt)
	if err == nil {
		t.Error("expected error for no charset selected")
	}

	opt = PasswordOptions{
		Length:          1,
		UseSpecialChars: true,
		UseNumbers:      true,
		UseUpper:        false,
		UseLower:        false,
		Count:           1,
	}
	_, err = GeneratePassword(opt)
	if err == nil {
		t.Error("expected error for length too short")
	}

	opt = PasswordOptions{
		Length:          8,
		UseSpecialChars: true,
		UseNumbers:      true,
		UseUpper:        true,
		UseLower:        true,
		Count:           0,
	}
	_, err = GeneratePassword(opt)
	if err == nil {
		t.Error("expected error for count < 1")
	}
}

func TestGeneratePassword_MinLength(t *testing.T) {
	opt := PasswordOptions{
		Length:          4,
		UseSpecialChars: true,
		UseNumbers:      true,
		UseUpper:        true,
		UseLower:        true,
		Count:           1,
	}
	passwords, err := GeneratePassword(opt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(passwords) != 1 {
		t.Fatalf("expected 1 password, got %d", len(passwords))
	}
	pwd := passwords[0]
	if len([]rune(pwd)) != 4 {
		t.Errorf("expected password length 4, got %d", len([]rune(pwd)))
	}
}

func TestGeneratePassword_Uniqueness(t *testing.T) {
	opt := PasswordOptions{
		Length:          16,
		UseSpecialChars: true,
		UseNumbers:      true,
		UseUpper:        true,
		UseLower:        true,
		Count:           10,
	}
	passwords, err := GeneratePassword(opt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	seen := make(map[string]struct{})
	for _, pwd := range passwords {
		if _, exists := seen[pwd]; exists {
			t.Errorf("duplicate password generated: %s", pwd)
		}
		seen[pwd] = struct{}{}
	}
}
