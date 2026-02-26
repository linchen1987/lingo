package cmd

import (
	"fmt"

	"github.com/linchen1987/lingo/internal/tools"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "base64 [encode|decode] [input]",
	Short: "Base64 encode or decode",
	Long: `Base64 encode or decode a string.

Examples:
  lin base64 encode hello
  lin base64 decode aGVsbG8=`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]
		input := args[1]

		switch action {
		case "encode":
			fmt.Println(tools.Base64Encode(input))
		case "decode":
			result, err := tools.Base64Decode(input)
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
