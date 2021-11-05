package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type sol struct{}

// InitSolution initializes the solution plugin
func InitSolution(pm plugin.Manager) {
	pm.RegisterSolution(&sol{})
}

func (s *sol) Name() string {
	return "day05"
}

func (s *sol) ProperName() string {
	return "Day 05"
}

func parseBoardingPass(input string) uint16 {
	r := strings.NewReplacer("F", "0", "L", "0", "B", "1", "R", "1")
	i, _ := strconv.ParseUint(r.Replace(input), 2, 16)
	return uint16(i)
}

func (s *sol) solvePart1(lines []string) error {
	var max uint16
	for _, line := range lines {
		bpass := parseBoardingPass(line)
		if bpass > max {
			max = bpass
		}
	}
	fmt.Printf("\tPart 1 Solution: %d\n", max)
	return nil
}

func (s *sol) solvePart2(lines []string) error {
	seats := []uint16{}
	for _, line := range lines {
		seat := parseBoardingPass(line)
		seats = append(seats, seat)
	}
	sort.SliceStable(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	var mySeat uint16
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1] != seats[i]+1 {
			mySeat = seats[i] + 1
			break
		}
	}
	fmt.Printf("\tPart 2 Solution: %d\n", mySeat)

	return nil
}

func (s *sol) Solve(input io.Reader) error {
	lines, err := scanner.ReadLines(input)
	if err != nil {
		return err
	}
	err = s.solvePart1(lines)
	if err != nil {
		return err
	}
	err = s.solvePart2(lines)
	if err != nil {
		return err
	}
	return nil
}
