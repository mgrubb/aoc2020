package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type policy struct {
	lower int64
	upper int64
	char  string
}

type sol struct{}

// InitSolution initializes the solution plugin
func InitSolution(pm plugin.Manager) {
	pm.RegisterSolution(&sol{})
}

func (s *sol) Name() string {
	return "day02"
}

func (s *sol) ProperName() string {
	return "Day 02"
}

func (s *sol) solvePart1(lines []string) error {
	numValid := 0
	for _, line := range lines {
		pol, passwd, err := parseLine(line)
		if err != nil {
			return err
		}
		num := int64(strings.Count(passwd, pol.char))
		if num >= pol.lower && num <= pol.upper {
			numValid++
		}
	}
	fmt.Printf("\tPart 1 Solution: %d\n", numValid)
	return nil
}

func (s *sol) solvePart2(lines []string) error {
	numValid := 0
	for _, line := range lines {
		pol, passwd, err := parseLine(line)
		if err != nil {
			return err
		}
		hasLower := passwd[pol.lower-1] == pol.char[0]
		hasUpper := passwd[pol.upper-1] == pol.char[0]
		// xor
		if (hasLower || hasUpper) && !(hasLower && hasUpper) {
			numValid++
		}

	}
	fmt.Printf("\tPart 2 Solution: %d\n", numValid)
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

func parseLine(line string) (*policy, string, error) {
	parts := strings.Split(line, ": ")
	policyParts := strings.Split(parts[0], " ")
	rangeParts := strings.Split(policyParts[0], "-")
	il, err := strconv.ParseInt(rangeParts[0], 10, 32)
	if err != nil {
		return nil, "", err
	}
	iu, err := strconv.ParseInt(rangeParts[1], 10, 64)
	if err != nil {
		return nil, "", err
	}
	pol := &policy{
		lower: il,
		upper: iu,
		char:  policyParts[1],
	}
	return pol, parts[1], nil
}
