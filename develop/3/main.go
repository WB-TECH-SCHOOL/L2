package main

import "filesort/filesort"

func main() {
	err := filesort.StartSortCmd()
	if err != nil {
		panic(err)
	}
}
