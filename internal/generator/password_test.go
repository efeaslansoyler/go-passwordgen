// Package generator contains tests for the password generation and analysis logic.
package generator

import (
	"strings"
	"testing"
	"unicode"
)

// containsAny returns true if any rune in s is present in set.
func containsAny(s string, set string) bool {
	for _, r := range s {
		if strings.ContainsRune(set, r) {
			return true
		}
	}
	return false
}

// TestGeneratePassword_Basic checks that generated passwords meet all option requirements
// and contain at least one character from each selected set.
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
	for _, gp := range passwords {
		pwd := gp.Value
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

// TestGeneratePassword_OnlyNumbers checks that generated passwords contain only digits
// when only numbers are enabled.
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
	for _, gp := range passwords {
		pwd := gp.Value
		for _, r := range pwd {
			if !unicode.IsDigit(r) {
				t.Errorf("expected only digits, got: %q", r)
			}
		}
	}
}

// TestGeneratePassword_OnlyUpper checks that generated passwords contain only uppercase letters
// when only uppercase is enabled.
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
	for _, gp := range passwords {
		pwd := gp.Value
		for _, r := range pwd {
			if !unicode.IsUpper(r) {
				t.Errorf("expected only uppercase, got: %q", r)
			}
		}
	}
}

// TestGeneratePassword_InvalidOptions checks that invalid options (e.g., no charset, too short, count < 1)
// return an error.
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

// TestGeneratePassword_MinLength checks that the minimum length for all enabled charsets is enforced.
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
	pwd := passwords[0].Value
	if len([]rune(pwd)) != 4 {
		t.Errorf("expected password length 4, got %d", len([]rune(pwd)))
	}
}

// TestGeneratePassword_Uniqueness checks that all generated passwords are unique in a batch.
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
	for _, gp := range passwords {
		pwd := gp.Value
		if _, exists := seen[pwd]; exists {
			t.Errorf("duplicate password generated: %s", pwd)
		}
		seen[pwd] = struct{}{}
	}
}
