package core

import (
	"fmt"
	"os/exec"
	"strings"
	"golang.org/x/exp/maps"
)

var recordsSlice records

func Process(repo, orderBy, format string, languages, extensions, include, exclude []string) {
	filenames, err := Search(repo, orderBy, extensions, include, exclude)
	if err != nil {
		fmt.Printf("Couldn't find files in %s: %v", repo, err)
		return
	}
	data := blame(filenames)

	orderedResult := Compare(data, orderBy)

	for _, result := range orderedResult {
		fmt.Println(result)
	}
}

func blame(filenames []string) map[string]*record {
	recordsSlice = map[string]*record{}
	for _, file := range filenames {

		dir := file[:strings.LastIndex(file, "/")+1]
		filename := file[strings.LastIndex(file, "/")+1:]

		if strings.Contains(file, "/.") {
			continue
		}

		cmd := exec.Command("git", "blame", "--porcelain", filename)
		cmd.Dir = dir

		result, err := cmd.CombinedOutput()
		if err != nil {
			continue
		}

		maps.Copy(recordsSlice, Parse(string(result)))
	}
	return recordsSlice
}
