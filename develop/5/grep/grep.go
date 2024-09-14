package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type grepOptions struct {
	after       int
	before      int
	context     int
	count       bool
	ignoreCase  bool
	invertMatch bool
	fixedMatch  bool
	lineNumber  bool
	pattern     string
}

var opts grepOptions

var grepCmd = &cobra.Command{
	Use:   "go run main.go [pattern] [files...]",
	Short: "Search for patterns in files",
	Long: `Search for patterns in files and print lines where patterns are found.
This command mimics some of the behaviors of Unix grep.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		opts.pattern = args[0]
		files := args[1:]
		grep(files, opts)
	},
}

func init() {
	grepCmd.Flags().IntVarP(&opts.after, "after", "A", 0, "print +N lines after match")
	grepCmd.Flags().IntVarP(&opts.before, "before", "B", 0, "print +N lines before match")
	grepCmd.Flags().IntVarP(&opts.context, "context", "C", 0, "print ±N lines around match")
	grepCmd.Flags().BoolVarP(&opts.count, "count", "c", false, "count of matching lines")
	grepCmd.Flags().BoolVarP(&opts.ignoreCase, "ignore-case", "i", false, "ignore case distinctions")
	grepCmd.Flags().BoolVarP(&opts.invertMatch, "invert-match", "v", false, "select non-matching lines")
	grepCmd.Flags().BoolVarP(&opts.fixedMatch, "fixed", "F", false, "interpret pattern as a fixed string")
	grepCmd.Flags().BoolVarP(&opts.lineNumber, "line-number", "n", false, "print line number with output lines")
}

func grep(files []string, opts grepOptions) {
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0
		var results []string
		var counts int

		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			match := strings.Contains(line, opts.pattern)

			if opts.ignoreCase {
				line = strings.ToLower(line)
				match = strings.Contains(line, strings.ToLower(opts.pattern))
			}

			if opts.fixedMatch {
				match = line == opts.pattern
			}

			if opts.invertMatch {
				match = !match
			}

			if match {
				if opts.count {
					counts++
				} else {
					if opts.lineNumber {
						line = fmt.Sprintf("%d:%s", lineNum, line)
					}
					results = append(results, line)
				}
			}
		}

		if opts.count {
			fmt.Println(counts)
		} else {
			for _, result := range results {
				fmt.Println(result)
			}
		}
	}
}

func Execute() error {
	return grepCmd.Execute()
}
