package aoc2021

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	days := []func(bool) int{
		Day1, Day2, Day3, Day4, Day5, Day6, Day7, Day8, Day9, Day10, Day11, Day12, Day13, Day14, Day15, // Day16, Day17, Day18,
	}
	for day, today := range days {
		fmt.Printf("Day %d.1 %v\n", day+1, today(false))
		fmt.Printf("Day %d.2 %v\n", day+1, today(true))
	}
}

func TestToday(t *testing.T) {
	today := Day15
	fmt.Printf("Today.1 %v\n", today(false))
	// fmt.Printf("Today.2 %v\n", today(true))
}
