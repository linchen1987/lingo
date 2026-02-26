package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lin",
	Short: "A CLI tool collection",
	Long:  `lin is a collection of useful CLI tools for encoding, decoding, and data conversion.`,
	Run: func(cmd *cobra.Command, args []string) {
		runTUI()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(datetimeCmd)
	rootCmd.AddCommand(base58Cmd)
	rootCmd.AddCommand(base64Cmd)
}
