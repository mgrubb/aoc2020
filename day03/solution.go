package main

import (
	"fmt"
	"io"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type sol struct{}

// InitSolution initializes the solution plugin
func InitSolution(pm plugin.Manager) {
	pm.RegisterSolution(&sol{})
}

func (s *sol) Name() string {
	return "day03"
}

func (s *sol) ProperName() string {
	return "Day 03"
}

func checkSlope(rise, run int, course []string) int {
	courseWidth := len(course[0])
	courseHeight := len(course)
	numTrees := 0
	for row, col := 0, 0; row < courseHeight; row, col = row+rise, (col+run)%courseWidth {
		if course[row][col] == '#' {
			numTrees++
		}
	}
	return numTrees
}

func (s *sol) solvePart1(course []string) error {
	numTrees := checkSlope(1, 3, course)
	fmt.Printf("\tPart 1 Solution: %d\n", numTrees)
	return nil
}

func (s *sol) solvePart2(course []string) error {
	var ans uint64 = 1
	slopes := [5][2]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	for _, m := range slopes {
		ans *= uint64(checkSlope(m[0], m[1], course))
	}

	fmt.Printf("\tPart 2 Solution: %d\n", ans)
	return nil
}

func (s *sol) Solve(input io.Reader) error {
	course, err := scanner.ReadLines(input)
	if err != nil {
		return err
	}
	err = s.solvePart1(course)
	if err != nil {
		return err
	}
	err = s.solvePart2(course)
	if err != nil {
		return err
	}
	return nil
}
