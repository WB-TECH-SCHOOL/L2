package filesort

import (
	"github.com/spf13/cobra"
	"os"
)

var sortCmd = &cobra.Command{
	Use:   "go run main.go [input] [output]",
	Short: "Script for sorting strings in file with provided flags",
	Args:  cobra.ExactArgs(2),
	Run:   run,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
		}
	},
}

func StartSortCmd() error {
	initFlags()

	err := sortCmd.Execute()
	if err != nil {
		os.Exit(2)
	}

	return nil
}

func initFlags() {
	sortCmd.Flags().IntP("column", "c", 0, "sort by column")
	sortCmd.Flags().BoolP("numeric", "n", false, "sort by numeric")
	sortCmd.Flags().BoolP("reverse", "r", false, "reverse filesort")
	sortCmd.Flags().BoolP("unique", "u", false, "unique filesort")
}
