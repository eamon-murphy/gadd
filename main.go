package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gadd <file>")
		return
	}

	query := normalizePath(os.Args[1])

	candidates := getGitFiles()

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

func getGitFiles() []string {
	cmd := exec.Command("git", "ls-files")

	var out bytes.Buffer
	var errout bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &errout

	error := cmd.Run()
	if error != nil {
		fmt.Println("git error:", errout.String())
		return nil
	}

	output := strings.TrimSpace(out.String())

	if output == "" {
		return []string{}
	}

	fileList := strings.Split(output, "\n")

	var files []string
	for _, file := range fileList {
		files = append(files, normalizePath(file))
	}

	return files
}
