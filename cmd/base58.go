package cmd

import (
	"fmt"

	"github.com/linchen1987/lingo/internal/tools"

	"github.com/spf13/cobra"
)

var base58Cmd = &cobra.Command{
	Use:   "base58 [encode|decode] [input]",
	Short: "Base58 encode or decode",
	Long: `Base58 encode or decode a string.

Examples:
  lin base58 encode hello
  lin base58 decode StV1DL6CwTryKyV`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]
		input := args[1]

		switch action {
		case "encode":
			fmt.Println(tools.Base58Encode(input))
		case "decode":
			result, err := tools.Base58Decode(input)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Println(result)
		default:
			fmt.Printf("Unknown action: %s. Use 'encode' or 'decode'\n", action)
		}
	},
}
