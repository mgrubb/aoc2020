package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type sol struct{}

const summation rune = '∑'

// InitSolution initializes the solution plugin
func InitSolution(pm plugin.Manager) {
	pm.RegisterSolution(&sol{})
}

func (s *sol) Name() string {
	return "day06"
}

func (s *sol) ProperName() string {
	return "Day 06"
}

func tallyGroup1(group string) int {
	res := make(map[rune]int)
	for _, c := range group {
		res[c]++
	}
	return len(res)
}

func (s *sol) solvePart1(groups [][]string) error {
	total := 0
	for _, g := range groups {
		tally := tallyGroup1(strings.Join(g, ""))
		// fmt.Printf("\tGroup %d: %d\n", i+1, len(group)-1)
		total += tally
	}
	fmt.Printf("\tPart 1 Solution: %d\n", total)
	return nil
}

func tallyGroup2(group []string) int {
	res := make(map[rune]int)
	for _, c := range strings.Join(group, "") {
		res[c]++
	}
	count := 0
	for _, v := range res {
		if v == len(group) {
			count++
		}
	}
	return count
}

func (s *sol) solvePart2(groups [][]string) error {
	total := 0
	for _, g := range groups {
		total += tallyGroup2(g)
	}
	fmt.Printf("\tPart 2 Solution: %d\n", total)
	return nil
}

func (s *sol) Solve(input io.Reader) error {
	groups, err := scanner.ReadGroups(input)
	if err != nil {
		return err
	}
	err = s.solvePart1(groups)
	if err != nil {
		return err
	}
	err = s.solvePart2(groups)
	if err != nil {
		return err
	}
	return nil
}
