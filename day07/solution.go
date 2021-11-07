package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type sol struct{}

type (
	container map[string]int
	ruleMap   map[string]container
)

// InitSolution initializes the solution plugin
func InitSolution(pm plugin.Manager) {
	pm.RegisterSolution(&sol{})
}

func (s *sol) Name() string {
	return "day07"
}

func (s *sol) ProperName() string {
	return "Day 07"
}

func parseRule(line string, rules ruleMap) error {
	rootParts := strings.Split(line, " bags contain ")
	contained := strings.Split(rootParts[1], ", ")
	newContainer := make(container)
	for _, cont := range contained {
		contParts := strings.SplitN(cont, " ", 2)
		bagPart := strings.Split(contParts[1], " bag")
		if contParts[0] == "no" {
			break
		}
		num, err := strconv.Atoi(contParts[0])
		if err != nil {
			return err
		}
		newContainer[bagPart[0]] = num
	}
	rules[rootParts[0]] = newContainer
	return nil
}

func outputDot(rules ruleMap) {
	tmpl := `
	digraph day07 {
		{{range $k, $conts := . }}
			{{range $n, $wt := $conts }}
			{{ $k }} -> {{ $n }} [label="$wt"];
			{{end}}
		{{end}}
	}`
	tt := template.New("solution")
	template.Must(tt.Parse(tmpl))
	tt.Execute(os.Stdout, rules)
}

func (s *sol) solvePart1(lines []string) error {
	fmt.Printf("\tPart 1 Solution: \n")
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
