package aoc

// common routines for use by AoC code in subdirectories

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Readfile(fn string) (buf []string) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		buf = append(buf, s.Text())
	}
	f.Close()
	return buf
}

// returns example string split into []string
func Readstring(s string) (buf []string) {
	// buf = strings.Split(s, "\n")
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		buf = append(buf, sc.Text())
	}
	return buf
}

func Setdiff(a, b []byte) (res []byte) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
			continue
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
			continue
		}
		if a[i] >= b[j] {
			j++
		}
	}
	if i < len(a) {
		res = append(res, a[i:]...)
	}
	return res
}

func Setunion(a, b []byte) []byte {
	res := []byte{}
	i, j := 0, 0
	for {
		if i < len(a) && j < len(b) {
			if a[i] == b[j] {
				res = append(res, a[i])
				i++
				j++
			} else if a[i] < b[j] {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		} else {
			if i < len(a) {
				res = append(res, a[i:]...)
			} else if j < len(b) {
				res = append(res, b[j:]...)
			}
			break
		}
	}
	return res
}

func Abs(i int) int {
	return Max(i, -i)
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// var permute func([]int, []int) [][]int

// recursive function that returns all permutations of input
func Permute(prefix []int, input []int) [][]int {
	if len(input) == 0 {
		return [][]int{prefix}
	}
	var res [][]int
	for i := range input {
		cc := make([]int, len(input))
		copy(cc, input) // builtin
		if i > 0 {
			cc[0], cc[i] = cc[i], cc[0]
		}
		// newprefix := append(prefix, cc[0])
		newprefix := make([]int, len(prefix)+1)
		copy(newprefix, prefix)
		newprefix[len(prefix)] = cc[0]
		p := Permute(newprefix, cc[1:])
		res = append(res, p...)
		// fmt.Println("  append", prefix, p)
	}
	return res
}

// using minimal copying; swap and unswap items in-place
func Permute2(a []int) (res [][]int) {
	var r func([]int, int)
	r = func(a []int, i int) {
		if i == len(a)-1 {
			x := make([]int, len(a))
			copy(x, a)
			res = append(res, x)
			return
		}
		for x := i; x < len(a); x++ {
			if x > i {
				a[x], a[i] = a[i], a[x]
			}
			r(a, i+1)
			if x > i {
				a[x], a[i] = a[i], a[x]
			}
		}
	}
	r(a, 0)
	return res
}

// using goroutines is 10x slower
func Permute3(a []int) (res [][]int) {
	ch := make(chan []int, 10)
	var r func([]int, int)
	r = func(a []int, i int) {
		if i == len(a)-1 {
			ch <- a
			return
		}
		for x := i; x < len(a); x++ {
			if x > i {
				a[x], a[i] = a[i], a[x]
			}
			b := make([]int, len(a))
			copy(b, a)
			if i < 2 {
				go r(b, i+1)
			} else {
				r(b, i+1)
			}
			if x > i {
				a[x], a[i] = a[i], a[x]
			}
		}
	}
	go r(a, 0)
	for i := 0; i < Factorial(len(a)); i++ {
		a := <-ch
		res = append(res, a)
	}
	return res
}

func Factorial(n int) int {
	i := n
	for i > 1 {
		i--
		n = n * i
	}
	return n
}

func Combinate(a []int, n int) (res [][]int) {
	var r func(int, int)
	buf := make([]int, n)
	r = func(i, j int) {
		if j == 0 {
			x := make([]int, n)
			copy(x, buf)
			res = append(res, x)
			return
		}
		for x := i; x < len(a)-j+1; x++ {
			buf[n-j] = a[x]
			r(x+1, j-1)
		}
	}
	r(0, n)
	return
}

func Findprimes(extent int) []int {
	res := []int{}
	for i := 3; i < extent; i += 2 {
		div := false
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				div = true
				break
			}
		}
		if !div {
			res = append(res, i)
		}
	}
	return append([]int{2}, res...)
}

// of Eratosthenes B-)
func Sieve(extent int) []int {
	n := make([]bool, extent)
	res := []int{}
	for i := 2; i < extent; i++ {
		if !n[i] {
			res = append(res, i)
			for j := i; j < extent; j += i {
				n[j] = true
			}
		}
	}
	return res
}
