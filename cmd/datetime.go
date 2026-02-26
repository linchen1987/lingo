package cmd

import (
	"fmt"

	"lin/internal/tools"

	"github.com/spf13/cobra"
)

var datetimeCmd = &cobra.Command{
	Use:   "datetime [input]",
	Short: "Convert between timestamp and datetime string",
	Long: `Convert between Unix timestamp and datetime string.

Examples:
  lin datetime 1772103423
  lin datetime '2026.02.26 18:57:35 GMT+8'
  lin datetime              # show current timestamp`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(tools.GetCurrentTimestamp())
			return
		}

		result, err := tools.ParseDatetime(args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}
