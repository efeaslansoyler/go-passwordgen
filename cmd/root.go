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

	"github.com/efeaslansoyler/go-passwordgen/internal/generator"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-passwordgen",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := generator.PasswordOptions{
			Length:          length,
			UseSpecialChars: useSpecialChars,
			UseNumbers:      useNumbers,
			UseUpper:        useUpper,
			UseLower:        useLower,
			Count:           count,
		}
		passwords, err := generator.GeneratePassword(opts)
		if err != nil {
			return err
		}
		for i, password := range passwords {
			fmt.Printf("Password %d: %s\n", i+1, password)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
)

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	rootCmd.Flags().BoolVarP(&useSpecialChars, "special", "s", true, "Use special characters")
	rootCmd.Flags().BoolVarP(&useNumbers, "numbers", "n", true, "Use numbers")
	rootCmd.Flags().BoolVarP(&useUpper, "upper", "u", true, "Use uppercase letters")
	rootCmd.Flags().BoolVarP(&useLower, "lower", "o", true, "Use lowercase letters")
	rootCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of passwords to generate")
}
