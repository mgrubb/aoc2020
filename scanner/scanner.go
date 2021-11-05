package scanner

import (
	"bufio"
	"io"
	"strconv"
)

// ReadLines reads entirety of io.Reader and returns a slice of lines
func ReadLines(r io.Reader) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// LinesOfInts reads io.Reader line by line and converts each line into an int64, then returns a slice of those ints.
func LinesOfInts(r io.Reader) ([]int64, error) {
	lines := make([]int64, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, nil
}

// ScanGroups scans lines as groups of items separated by an empty line and returns a slice of string slices
func ScanGroups(lines []string) [][]string {
	recs := [][]string{}
	start := 0
	for i, l := range lines {
		if len(l) == 0 {
			recs = append(recs, lines[start:i])
			start = i + 1
		} else if i == len(lines)-1 {
			recs = append(recs, lines[start:])
		}
	}
	return recs
}

// ReadGroups reads lines as groups of items separated by an empty line and returns a slice of string slices
func ReadGroups(r io.Reader) ([][]string, error) {
	lines, err := ReadLines(r)
	if err != nil {
		return nil, err
	}
	return ScanGroups(lines), nil
}
