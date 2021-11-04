#!/bin/bash

if [ -z "$1" ]
then
  echo "Usage: $(basename $0) <day>" >&2
  exit 1
fi

day="$(printf "day%02d\n" $1)"
DAY="$(printf "Day %02d\n" $1)"

if [ -d "${day}" ]
then
  echo "Day $1 already exists." >&2
  exit 1
fi

mkdir ${day} inputs/${day}

cat << EOF > ${day}/solution.go
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
	return "${day}"
}

func (s *sol) ProperName() string {
	return "${DAY}"
}

func (s *sol) solvePart1(lines []string) error {
	// fmt.Printf("\tPart 1 Solution: \n")
	return nil
}

func (s *sol) solvePart2(lines []string) error {
	// fmt.Printf("\tPart 2 Solution: \n")
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
EOF
