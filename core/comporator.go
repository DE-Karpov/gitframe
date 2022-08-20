package core

import (
	"fmt"
	"sort"
)

const (
	LINES   = "lines"
	COMMITS = "commits"
	FILES   = "files"
)

func Compare(data map[string]*record, orderBy string) []record {

	values := make([]record, 0, len(data))

	for _, val := range data {
        values = append(values, *val)
    }

	fmt.Println()

	switch orderBy {
	case COMMITS:
		sort.SliceStable(values, func(i, j int) bool{
			return len(values[i].commits) < len(values[j].commits)
		})
	case FILES:
		sort.SliceStable(values, func(i, j int) bool{
			return len(values[i].files) < len(values[j].files)
		})
	default:
		sort.SliceStable(values, func(i, j int) bool{
			return values[i].lines < values[j].lines
		})
	}

	return values
}
