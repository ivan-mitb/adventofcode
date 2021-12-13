package aoc2015

import (
	"aoc"
	"fmt"
	"testing"
)

// var days []func(bool) int

func BenchmarkPopCount(b *testing.B) {
	var x uint64 = 0xfedcba9876543210
	b.Run("base", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PopCount(x)
		}
	})
	b.Run("loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PopCountLoop(x)
		}
	})
	b.Run("64-bits", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PopCount64(x)
		}
	})
	b.Run("clear", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PopCountClear(x)
		}
	})
}

func TestPermute2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	a = a[:10]
	// fmt.Println(permute2(a))
	aoc.Permute2(a)
}

func TestPermute3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	a = a[:10]
	// fmt.Println(permute3(a))
	aoc.Permute3(a)
}

func TestCombinate(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	a = a[:]
	fmt.Println(aoc.Combinate(a, 3))
}

func BenchmarkPermute(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = a[:8]
	var res, res2, res3 [][]int
	b.Run("permute", func(b *testing.B) {
		res = aoc.Permute(nil, a)
	})
	b.Run("permute2", func(b *testing.B) {
		res2 = aoc.Permute2(a)
	})
	b.Run("permute3", func(b *testing.B) {
		res3 = aoc.Permute3(a)
	})
	println(len(res) == len(res2))
	println(len(res) == len(res3))
}

func BenchmarkPrime(b *testing.B) {
	b.Run("findprime", func(b *testing.B) {
		aoc.Findprimes(100000)
	})
	b.Run("sieve", func(b *testing.B) {
		aoc.Sieve(100000)
	})
}

func TestAll(t *testing.T) {
	days := []func(bool) int{
		Day1, Day2, Day3, Day4, Day5, Day6, Day7new, Day8, Day9, Day10, Day11, Day12, Day13, Day14, Day15, Day16, Day17, Day18, Day19, Day20, Day21, Day22,
	}
	for day, today := range days {
		fmt.Printf("Day %d.1 %v\n", day+1, today(false))
		fmt.Printf("Day %d.2 %v\n", day+1, today(true))
	}
}

func TestToday(t *testing.T) {
	today := Day22
	fmt.Printf("Today.1 %v\n", today(false))
	// fmt.Printf("Today.2 %v\n", today(true))
}

func TestScratch(t *testing.T) {
	fmt.Println("scratch", Scratch())
}
