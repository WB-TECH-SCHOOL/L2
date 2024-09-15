package cut

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type cutOptions struct {
	fields    []int
	delimiter string
	separated bool
}

var opts cutOptions

var cutCmd = &cobra.Command{
	Use:   "type [file] | go run main.go cut [flags] OR go run main.go cut [flags] + STDIN",
	Short: "Cut out selected portions of each line of a file",
	Long: `Cut out selected portions of each line of a file or standard input and print them to standard output.
This command mimics some of the behaviors of Unix cut.`,
	Run: func(cmd *cobra.Command, args []string) {
		cut(opts)
	},
}

func init() {
	cutCmd.Flags().IntSliceVarP(&opts.fields, "fields", "f", []int{}, "select only these fields")
	cutCmd.Flags().StringVarP(&opts.delimiter, "delimiter", "d", "\t", "use DELIM instead of TAB for field delimiter")
	cutCmd.Flags().BoolVarP(&opts.separated, "separated", "s", false, "do not print lines not containing delimiters")
}

func cut(opts cutOptions) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, opts.delimiter)

		if opts.separated && len(columns) < 2 {
			continue
		}

		var output []string
		for _, field := range opts.fields {
			if field > 0 && field <= len(columns) {
				output = append(output, columns[field-1])
			}
		}

		fmt.Println(strings.Join(output, opts.delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func Execute() error {
	return cutCmd.Execute()
}
