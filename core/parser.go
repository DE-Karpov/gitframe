package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	recordPattern    = `^[a-fA-F0-9]{40}(.|\s)*filename .*`
	authorPattern    = `author [A-Za-z0-9- ]*`
	commitPattern    = `^[a-fA-F0-9]{40}`
	fileNamePattern  = `filename .*`
	firstLinePattern = ` [0-9]* [0-9]* [0-9]*`
)

type record struct {
	author  string
	files   map[string]struct{}
	commits map[string]struct{}
	lines   int
}

func (r record) String() string {
	return fmt.Sprintf("Author: %s, files: %d, commits: %d lines: %d", r.author, len(r.files), len(r.commits), r.lines)
}

type records map[string]*record

func Parse(rawBlame string) map[string]*record {

	recordReg := regexp.MustCompile(recordPattern)
	authorReg := regexp.MustCompile(authorPattern)
	commitReg := regexp.MustCompile(commitPattern)
	lineReg := regexp.MustCompile(firstLinePattern)
	fileNameReg := regexp.MustCompile(fileNamePattern)
	matches := recordReg.FindAllStringSubmatch(rawBlame, -1)

	for _, v := range matches {
		author := authorReg.FindStringSubmatch(v[0])[0][6:]
		commitNumber := commitReg.FindStringSubmatch(v[0])[0]
		fileName := strings.Split(fileNameReg.FindStringSubmatch(v[0])[0], " ")[1]
		lines := lineReg.FindStringSubmatch(v[0])[0]

		if info, found := recordsSlice[author]; found {

			if _, found := info.files[fileName]; !found {
				info.files[fileName] = struct{}{}
			}

			if _, found := info.commits[commitNumber]; !found {
				info.commits[commitNumber] = struct{}{}
			}

			linesNum := getLinesNumber(lines)

			info.lines += linesNum
		} else {
			commit := make(map[string]struct{})
			filename := make(map[string]struct{})
			commit[commitNumber] = struct{}{}
			filename[fileName] = struct{}{}

			linesNum := getLinesNumber(lines)

			info = &record{author: author, commits: commit, files: filename, lines: linesNum}
			recordsSlice[author] = info
		}

	}
	return recordsSlice
}

func getLinesNumber(lines string) int {
	num := strings.Split(strings.TrimSpace(lines), " ")[2]
	linesNum, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return linesNum
}
