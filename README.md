# Advent of Code 2020

Thought I would warm up for this year by doing last years (2020) AOC challenges.
This is all done in Go Lang (https://golang.org).

Thought I would experiment with writing a plugin system in go, so each day's solutions are written
as a shared library type plugin.  The Makefile will build the main program which is responsible for
providing the puzzle input to the plugins, and the daily plugins themselves which are responsible for
solving the puzzles.
