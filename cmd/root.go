package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "passwordGen",
	Short: "Use this tool to generate passwords",
	Long: `This tool can help you generate passwords
Examples:

passwordGen generate -l 25 -d -s

-l: length of password
-d: add digits to password (default true)
-s: add special characters to password (default true)
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
