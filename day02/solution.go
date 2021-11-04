package main

import (
	"fmt"
	"io"

	"github.com/mgrubb/aoc2020/plugin"
)

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

func (s *sol) Solve(input io.Reader) error {
	fmt.Println("Day 02")
	return nil
}
