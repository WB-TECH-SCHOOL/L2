package filesort

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strconv"
	"strings"
)

type line struct {
	Original string
	Columns  []string
}

func run(cmd *cobra.Command, args []string) {
	input := args[0]
	output := args[1]
	fmt.Printf("Input file: %s\n", input)
	fmt.Printf("Output file: %s\n", output)

	// Получение флагов
	columnIndex, _ := cmd.Flags().GetInt("column")
	numericSort, _ := cmd.Flags().GetBool("numeric")
	reverseSort, _ := cmd.Flags().GetBool("reverse")
	unique, _ := cmd.Flags().GetBool("unique")

	// Чтение входного файла
	file, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Считывание и обработка строк
	scanner := bufio.NewScanner(file)
	var lines []line
	for scanner.Scan() {
		text := scanner.Text()
		columns := strings.Fields(text)
		lines = append(lines, line{Original: text, Columns: columns})
	}

	// Сортировка
	sort.Slice(lines, func(i, j int) bool {
		if columnIndex >= len(lines[i].Columns) || columnIndex >= len(lines[j].Columns) {
			return false
		}
		if numericSort {
			iv, ierr := strconv.Atoi(lines[i].Columns[columnIndex])
			jv, jerr := strconv.Atoi(lines[j].Columns[columnIndex])
			if ierr == nil && jerr == nil {
				if reverseSort {
					return iv > jv
				}
				return iv < jv
			}
		}
		if reverseSort {
			return lines[i].Columns[columnIndex] > lines[j].Columns[columnIndex]
		}
		return lines[i].Columns[columnIndex] < lines[j].Columns[columnIndex]
	})

	// Удаление дубликатов
	if unique {
		tempLines := lines[:0]
		seen := make(map[string]bool)
		for _, line := range lines {
			if !seen[line.Original] {
				seen[line.Original] = true
				tempLines = append(tempLines, line)
			}
		}
		lines = tempLines
	}

	// Запись в выходной файл
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, l := range lines {
		fmt.Fprintln(writer, l.Original)
	}
	writer.Flush()
}
