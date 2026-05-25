package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "passwordGen",
	Short: "Generate random passwords",
	Long: `Generate random passwords with customisable options.

For example:

passwordGen generate -l 25 -d -s
`,
	Run: generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialCharacters, _ := cmd.Flags().GetBool("special")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialCharacters {
		charset += "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(string(password))
}


func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 12, "Length of the generated password")
	generateCmd.Flags().BoolP("digits", "d", true, "Include digits in generated password")
	generateCmd.Flags().BoolP("special", "s", true, "Include special characters in generated password")
}
