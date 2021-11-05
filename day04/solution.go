package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/mgrubb/aoc2020/plugin"
	"github.com/mgrubb/aoc2020/scanner"
)

type sol struct{}

type passport map[string]interface{}

// InitSolution initializes the solution plugin
func InitSolution(pm plugin.Manager) {
	pm.RegisterSolution(&sol{})
}

func (s *sol) Name() string {
	return "day04"
}

func (s *sol) ProperName() string {
	return "Day 04"
}

func scanRecord(record string) (passport, error) {
	pp := make(passport)
	var err error
	fields := strings.Split(record, " ")
	for _, field := range fields {
		kv := strings.Split(field, ":")
		switch kv[0] {
		case "byr", "iyr", "eyr", "cid":
			pp[kv[0]], err = strconv.ParseUint(kv[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("Error parsing field %s: %w", kv[0], err)
			}
		case "pid", "hgt", "hcl", "ecl":
			pp[kv[0]] = kv[1]
		default:
			return nil, fmt.Errorf("Unknown key: %s", kv[0])
		}
	}
	return pp, nil
}

func checkRange(n, l, u uint64) bool {
	return l <= n && n <= u
}

func checkHeight(val string) bool {
	heightNum, err := strconv.ParseUint(val[0:len(val)-2], 10, 64)
	if err != nil {
		return false
	}
	if strings.HasSuffix(val, "cm") {
		return checkRange(heightNum, 150, 193)
	} else if strings.HasSuffix(val, "in") {
		return checkRange(heightNum, 59, 76)
	}
	return false
}

func checkHairColor(val string) bool {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return re.MatchString(val)
}

func checkEyeColor(val string) bool {
	validColors := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
	_, ok := validColors[val]
	return ok
}

func checkPassportID(val string) bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.MatchString(val)
}

func part1IsValid(pp passport) bool {
	reqKeys := map[string]struct{}{
		"byr": {},
		"iyr": {},
		"eyr": {},
		"pid": {},
		"hgt": {},
		"hcl": {},
		"ecl": {},
	}

	if len(pp) < 7 || len(pp) > 8 {
		return false
	}

	missingKeys := []string{}
	for k := range reqKeys {
		_, ok := pp[k]
		if !ok {
			missingKeys = append(missingKeys, k)
		}
	}
	if len(missingKeys) != 0 {
		return false
	}
	return true
}

func part2IsValid(pp passport) bool {
	return checkRange(pp["byr"].(uint64), 1920, 2002) &&
		checkRange(pp["iyr"].(uint64), 2010, 2020) &&
		checkRange(pp["eyr"].(uint64), 2020, 2030) &&
		checkHeight(pp["hgt"].(string)) &&
		checkHairColor(pp["hcl"].(string)) &&
		checkEyeColor(pp["ecl"].(string)) &&
		checkPassportID(pp["pid"].(string))
}

func collapseRecs(lines []string) []string {
	recs := []string{}
	for _, group := range scanner.ScanGroups(lines) {
		recs = append(recs, strings.Join(group, " "))
	}
	return recs
}

func (s *sol) solvePart1(recs []string) error {
	validPassports := 0
	for _, rec := range recs {
		pp, err := scanRecord(rec)
		if err != nil {
			return err
		}
		if part1IsValid(pp) {
			validPassports++
		}
	}
	fmt.Printf("\tPart 1 Solution: %d\n", validPassports)
	return nil
}

func (s *sol) solvePart2(recs []string) error {
	validPassports := 0
	for _, rec := range recs {
		pp, err := scanRecord(rec)
		if err != nil {
			return err
		}
		if !part1IsValid(pp) {
			continue
		}

		if part2IsValid(pp) {
			validPassports++
		}
	}
	fmt.Printf("\tPart 2 Solution: %d\n", validPassports)
	return nil
}

func (s *sol) Solve(input io.Reader) error {
	lines, err := scanner.ReadLines(input)
	if err != nil {
		return err
	}
	recs := collapseRecs(lines)
	err = s.solvePart1(recs)
	if err != nil {
		return err
	}
	err = s.solvePart2(recs)
	if err != nil {
		return err
	}
	return nil
}
