/*
Copyright © 2025 Efe Aslan Söyler efeaslan1703@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/efeaslansoyler/go-passwordgen/internal/generator"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-passwordgen",
	Short: "Generate secure random passwords from the command line",
	Long: `go-passwordgen generates secure, random passwords with customizable
length and character sets. Supports special characters, numbers, upper and
lowercase letters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := generator.PasswordOptions{
			Length:          length,
			UseSpecialChars: useSpecialChars,
			UseNumbers:      useNumbers,
			UseUpper:        useUpper,
			UseLower:        useLower,
			Count:           count,
		}
		start := time.Now()
		passwords, err := generator.GeneratePassword(opts)
		if err != nil {
			return err
		}
		elapsed := time.Since(start)
		if quiet {
			for _, p := range passwords {
				fmt.Println(p.Value)
			}
		} else {
			for i, p := range passwords {
				fmt.Printf("Password %d: %s (Strength: %s, Entropy: %.2f)\n", i+1, p.Value, colorStrength(p.Strength), p.Entropy)
			}
			fmt.Printf("Generation time: %s\n", elapsed)

		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	length          int
	useSpecialChars bool
	useNumbers      bool
	useUpper        bool
	useLower        bool
	count           int
	quiet           bool
)

var Version = "dev"

func init() {
	rootCmd.Version = Version
	rootCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	rootCmd.Flags().BoolVarP(&useSpecialChars, "special", "s", true, "Use special characters")
	rootCmd.Flags().BoolVarP(&useNumbers, "numbers", "n", true, "Use numbers")
	rootCmd.Flags().BoolVarP(&useUpper, "upper", "u", true, "Use uppercase letters")
	rootCmd.Flags().BoolVarP(&useLower, "lower", "o", true, "Use lowercase letters")
	rootCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of passwords to generate")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Suppress output(Print only password(s)")
}

func colorStrength(strength string) string {
	switch strength {
	case "Excellent":
		return color.New(color.FgHiGreen, color.Bold).Sprint(strength)
	case "Strong":
		return color.New(color.FgGreen).Sprint(strength)
	case "Moderate":
		return color.New(color.FgYellow).Sprint(strength)
	case "Weak":
		return color.New(color.FgRed).Sprint(strength)
	default:
		return strength
	}
}
