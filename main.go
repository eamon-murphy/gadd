package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gadd <file>")
		return
	}

	query := normalizePath(os.Args[1])

	candidates := []string{
		"this/directory/ok/file.go",
		"this/directory/file.go",
		"some/directory/file.go",
		"README.md",
	}

	matches := findMatches(query, candidates)

	printResult(query, matches)
}

func normalizePath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func findMatches(query string, candidates []string) []string {
	var matches []string

	for _, candidate := range candidates {
		candidate = normalizePath(candidate)

		if candidate == query || strings.HasSuffix(candidate, "/"+query) {
			matches = append(matches, candidate)
		}
	}
	return matches
}

func printResult(query string, matches []string) {
	switch len(matches) {
	case 0:
		fmt.Println(query + " not found.")

	case 1:
		fmt.Println("unique match")
		fmt.Println(matches[0])

	default:
		fmt.Println("ambiguous match")
		for _, match := range matches {
			fmt.Println(" -", match)
		}
	}
}
