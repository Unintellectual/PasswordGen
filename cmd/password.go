package cmd

import (
	"github.com/spf13/cobra"
	"math/rand"
	"os"
)

var generateCmd = &cobra.Command{
	Use:   "Generate",
	Short: "Generate Random Passwords",
	Long: `A Random Password Generator that makes customizable length, numbers and special characters. 
	
	For example:
	
	PasswordGen generate -l 12 -d -s`,

		Run: GeneratePassword,
	
}

func init ()  {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("lenght", "l" , 8 , "Lenght of generated password")
	generateCmd.Flags().BoolP("Digits", "d" , false , "Addition of Digits")
	generateCmd.Flags().BoolP("SpecialChar", "s" , false,"Addition of Special Charaters")
}
func GeneratePassword(cmd *cobra.Command, args []string)  {
		length, _ := cmd.Flags().GetInt("Length")
		isDigits, _ := cmd.Flags().GetBool("Digits")
		isSpecialChar, _ := cmd.Flags().GetBool("Characters")

		charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghjklmnopqrstuvwxyz"

		if isDigits{
			charSet += "0123456789"
		}

		if isSpecialChar {
			charSet += "~!@#$%^&*()_+[]{}|;:',<.>/?`"
		}

		password := make([]byte, length)

		for i := range password{
			password[i] = charSet[rand.Intn(len(charSet))]
		}
}

