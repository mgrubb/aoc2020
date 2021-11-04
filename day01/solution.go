package main

import (
	"fmt"
	"io"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type sol struct{}

// InitSolution registers the solution with the plugin manager
func InitSolution(mgr plugin.Manager) {
	mgr.RegisterSolution(&sol{})
}

func (s sol) Name() string {
	return "day01"
}

func (s sol) ProperName() string {
	return "Day 01"
}

func (s sol) solvePart1(nums []int64) error {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				fmt.Printf("\tPart 1 Solution: %d\n", nums[i]*nums[j])
				return nil
			}
		}
	}
	fmt.Printf("\tPart 1 No Solution Found\n")
	return nil
}

func (s sol) solvePart2(nums []int64) error {
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					fmt.Printf("\tPart 2 Solution: %d\n", nums[i]*nums[j]*nums[k])
					return nil
				}
			}
		}
	}
	return nil
}

func (s sol) Solve(input io.Reader) error {
	nums, err := scanner.LinesOfInts(input)
	if err != nil {
		return err
	}

	s.solvePart1(nums)
	s.solvePart2(nums)
	return nil
}
