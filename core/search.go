package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func Search(root, orderBy string, exts, incl, excl []string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if len(exts) > 0 {
			if !info.IsDir() && contains(exts, filepath.Ext(path)) {
				files = append(files, path)
			}
		} else {
			if !info.IsDir() {
				files = append(files, path)
			}
		}
		return nil
	})

	if len(excl) > 0 {
		for _, e := range excl {
			files = exclude(files, e)
		}
	}

	if len(incl) > 0 {
		for _, i := range incl {
			files = include(files, i)
		}
	}

	return files, err
}

func exclude(files []string, pattern string) []string {
	exclfiles, err := filepath.Glob(pattern)
	fmt.Println(exclfiles)
	if err != nil {
		fmt.Println(err)
	}

	return difference(files, exclfiles)
}

func include(files []string, pattern string) []string {
	inclfiles, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err)
	}

	return inclfiles
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
