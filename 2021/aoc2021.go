package aoc2021

import (
	. "aoc"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1(part2 bool) int {
	buf := Readfile("day1.txt")
	// buf = []string{"199",
	// 	"200",
	// 	"208",
	// 	"210",
	// 	"200",
	// 	"207",
	// 	"240",
	// 	"269",
	// 	"260",
	// 	"263"}
	countinc := func(a []int) (count int) {
		last := a[0]
		for _, n := range a[1:] {
			if n > last {
				count++
			}
			last = n
		}
		return
	}
	if !part2 {
		a := []int{}
		for _, s := range buf {
			n, _ := strconv.Atoi(s)
			a = append(a, n)
		}
		return countinc(a)
	} else {
		// part2
		// triplet := 0
		a := []int{}
		a3 := []int{}
		for _, s := range buf {
			n, _ := strconv.Atoi(s)
			a = append(a, n)
		}
		a3 = append(a3, a[0]+a[1]+a[2])
		n := a3[0]
		for i := 3; i < len(a); i++ {
			n = a3[i-3] - a[i-3] + a[i]
			a3 = append(a3, n)
		}
		// fmt.Println(a)
		// fmt.Println(a3)
		return countinc(a3)
	}
	return 0
}

func Day2(part2 bool) int {
	buf := Readfile("day2.txt")
	pos, depth, aim := 0, 0, 0
	if !part2 {
		for _, s := range buf {
			s := strings.Fields(s)
			n, _ := strconv.Atoi(s[1])
			switch s[0] {
			case "forward":
				pos += n
			case "down":
				depth += n
			case "up":
				depth -= n
			default:
				fmt.Errorf("unknown %s", s[0])
			}
		}
	} else {
		// part2
		for _, s := range buf {
			s := strings.Fields(s)
			n, _ := strconv.Atoi(s[1])
			switch s[0] {
			case "down":
				aim += n
			case "up":
				aim -= n
			case "forward":
				pos += n
				depth += aim * n
			}
		}
	}
	return pos * depth
}

func Day3(part2 bool) int {
	buf := Readfile("day3.txt")
	// buf = strings.Fields(`00100
	// 11110
	// 10110
	// 10111
	// 10101
	// 01111
	// 00111
	// 11100
	// 10000
	// 11001
	// 00010
	// 01010`)
	// returns the count of '1' at each bit-position
	getmode := func(buf []string, pos int) []int {
		var counts []int
		if pos < 0 {
			counts = make([]int, len(buf[0]))
			for _, s := range buf {
				for i := range s {
					if s[i] == '1' {
						counts[i]++
					}
				}
			}
		} else {
			counts = []int{0}
			for _, s := range buf {
				if s[pos] == '1' {
					counts[0]++
				}
			}
		}
		return counts
	}
	atoi := func(a string) (res int) {
		for i := range a {
			res <<= 1
			if a[i] == '1' {
				res |= 1
			}
		}
		return res
	}
	if !part2 {
		// gamma is bitmask of 1-dominant
		// epsilon is bitmask of 0-dominant
		gamma, epsilon := 0, 0
		counts := getmode(buf, -1)
		for i := range counts {
			gamma <<= 1
			epsilon <<= 1
			if counts[i] > len(buf)/2 {
				gamma |= 1
			} else {
				epsilon |= 1
			}
		}
		// println(gamma, epsilon)
		return gamma * epsilon
	} else {
		// part2
		// the index of the partition point between 'eliminated' and 'keep'
		lifesupport := func(buf []string, majority bool) (res int) {
			maj, min := byte('0'), byte('1')
			if majority {
				maj, min = byte('1'), byte('0')
			}
			part := 0
			for i := range buf[0] {
				count1 := getmode(buf[part:], i)[0]
				count0 := len(buf) - part - count1
				for j := part; j < len(buf); j++ {
					s := buf[j]
					if (count1 >= count0 && s[i] == min) ||
						(count1 < count0 && s[i] == maj) {
						buf[j], buf[part] = buf[part], buf[j]
						part++
						if part == len(buf)-1 {
							// fmt.Println("oxy", buf[part])
							res = atoi(buf[part])
							break
						}
					}
				}
			}
			return res
		}
		oxy := lifesupport(buf, true)
		co2 := lifesupport(buf, false)
		// fmt.Println(oxy, co2)
		return oxy * co2
	}
}

func Day4(part2 bool) int {
	buf := Readfile("day4.txt")
	// 	buf := strings.Split(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

	// 22 13 17 11  0
	//  8  2 23  4 24
	// 21  9 14 16  7
	//  6 10  3 18  5
	//  1 12 20 15 19

	//  3 15  0  2 22
	//  9 18 13 17  5
	// 19  8  7 25 23
	// 20 11 10 24  4
	// 14 21 16 12  6

	// 14 21 17 24  4
	// 10 16 15  9 19
	// 18  8 23 26 20
	// 22 11 13  6  5
	//  2  0 12  3  7
	// `, "\n")
	type card [5][5]int
	mark := func(c *card, n int) {
		for i := range c {
			for j := range c[0] {
				if c[i][j] == n {
					c[i][j] = -1
					return
				}
			}
		}
	}
	test := func(c *card) bool {
		for i := range c {
			row := 0
			col := 0
			for j := range c[0] {
				row += c[i][j]
				col += c[j][i]
			}
			if row == -5 || col == -5 {
				return true
			}
		}
		return false
	}
	sumUnmarked := func(c *card) (res int) {
		for i := range c {
			for j := range c[0] {
				if c[i][j] >= 0 {
					res += c[i][j]
				}
			}
		}
		return res
	}
	turns := strings.Split(buf[0], ",")
	cards := make([]card, (len(buf)-2)/6)
	for i := range cards {
		for r := 0; r < 5; r++ {
			s := strings.Fields(buf[2+(i*6)+r])
			for j := range s {
				cards[i][r][j], _ = strconv.Atoi(s[j])
			}
		}
	}
	lastscore := 0
	cardswon := make(map[int]bool)
	for t := range turns {
		t, _ := strconv.Atoi(turns[t])
		for c := range cards {
			mark(&cards[c], t)
			if test(&cards[c]) {
				cardswon[c] = true
				u := sumUnmarked(&cards[c])
				// fmt.Println("bingo", t, "card", c, "unmarked", u)
				lastscore = u * t
				if !part2 {
					return lastscore
				} else {
					if len(cardswon) == len(cards) {
						return lastscore
					}
				}
			}
		}
	}
	return lastscore
}

func Day5(part2 bool) int {
	buf := Readfile("day5.txt")
	// 	buf := strings.Split(`0,9 -> 5,9
	// 8,0 -> 0,8
	// 9,4 -> 3,4
	// 2,2 -> 2,1
	// 7,0 -> 7,4
	// 6,4 -> 2,0
	// 0,9 -> 2,9
	// 3,4 -> 1,4
	// 0,0 -> 8,8
	// 5,5 -> 8,2`, "\n")
	K := func(x, y int) string {
		return fmt.Sprintf("%d,%d", x, y)
	}
	m := make(map[string]int)
	for _, s := range buf {
		s := strings.Fields(s)
		var x1, y1, x2, y2 int
		if len(s) != 3 {
			fmt.Println("err not 3")
			fmt.Println(s)
		}
		fmt.Sscanf(s[0], "%d,%d", &x1, &y1)
		fmt.Sscanf(s[2], "%d,%d", &x2, &y2)
		if x1 == x2 {
			// horiz line
			if y1 > y2 {
				y2, y1 = y1, y2
			}
			for y1 <= y2 {
				m[K(x1, y1)]++
				y1++
			}
		} else if y1 == y2 {
			// vert line
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x1 <= x2 {
				m[K(x1, y1)]++
				x1++
			}
		} else {
			if part2 {
				// diagonal lines
				inc := 1
				if x1 > x2 {
					x1, x2, y1, y2 = x2, x1, y2, y1
				}
				if y1 > y2 {
					inc = -1
				}
				for x1 <= x2 {
					m[K(x1, y1)]++
					x1++
					y1 += inc
				}
			}
		}
	}
	count := 0
	for _, v := range m {
		if v > 1 {
			count++
		}
	}
	return count
}

func Day6(part2 bool) int {
	buf := Readfile("day6.txt")
	// buf := []string{`3,4,3,1,2`}
	s := strings.Split(buf[0], ",")
	fish := make([]int, 10)
	for i := range s {
		i, _ := strconv.Atoi(s[i])
		fish[i]++
	}
	//main loop
	limit := 80
	if part2 {
		limit = 256
	}
	for i := 0; i < limit; i++ {
		// if j == 0
		fish[7] += fish[0]
		fish[9] = fish[0] // to spawn
		for j := 1; j < 9; j++ {
			fish[j-1] = fish[j]
		}
		fish[8] = fish[9] // spawn
	}
	count := 0
	for _, x := range fish[:9] {
		count += x
	}
	return count
}

func Day7(part2 bool) int {
	buf := Readfile("day7.txt")
	// buf := []string{`16,1,2,0,4,2,7,1,2,14`}
	s := strings.Split(buf[0], ",")
	// min, mean, max
	min, max, sum := 999999, 0, 0
	input := make([]int, len(s))
	for i := range s {
		n, _ := strconv.Atoi(s[i])
		input[i] = n
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}
	cost := make([]int, max-min+1)
	if !part2 {
		for i := range cost {
			cost[i] = i
		}
	} else {
		// incremental fuel burn
		cumsum := 0
		for i := range cost {
			cumsum += i
			cost[i] = cumsum
		}
	}
	sumdev := func(s []int, ref int) (sum int) {
		for i := range s {
			d := s[i] - ref
			if d < 0 {
				d = -d
			}
			sum += cost[d]
		}
		return sum
	}
	minfuel := int(9e9)
	for i := min; i <= max; i++ {
		x := sumdev(input, i)
		if x < minfuel {
			minfuel = x
		}
	}
	return minfuel
}

func Day8(part2 bool) int {
	buf := Readfile("day8.txt")
	// `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`
	// buf := strings.Split(
	// 	`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
	// edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
	// fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
	// fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
	// aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
	// fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
	// dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
	// bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
	// egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
	// gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`, "\n")
	output := make([]string, len(buf))
	signal := make([]string, len(buf))
	for i, s := range buf {
		s := strings.Split(s, " | ")
		signal[i] = s[0]
		output[i] = s[1]
	}
	res := 0
	if !part2 {
		for i := range output {
			s := strings.Fields(output[i])
			for _, j := range s {
				l := len(j)
				if l == 2 || l == 3 || l == 4 || l == 7 {
					res++
				}
			}
		}
	} else {
		// assume the lists are sorted
		matchdigit := func(digits [][]byte, s []byte) byte {
			for i := range digits {
				if string(s) == string(digits[i]) {
					return byte('0' + i)
				}
			}
			return 0
		}
		// unscramble and decode the signal into the 4-digit value
		decode := func(signal, output string) int {
			s := strings.Fields(signal)
			sort.Slice(s, func(i, j int) bool { return len(s[i]) < len(s[j]) })
			digits := make([][]byte, 10)
			savedigit := func(s []byte, i int) {
				digits[i] = make([]byte, len(s))
				copy(digits[i], s)
			}
			res := make([]byte, 4)
			a := []byte{}
			cf := []byte{}
			bd := []byte{}
			for _, x := range s {
				j := []byte(x)
				sort.Slice(j, func(a, b int) bool { return j[a] < j[b] })
				switch len(j) {
				case 2: // '1'
					savedigit(j, 1)
					cf = j
				case 3: // '7'
					savedigit(j, 7)
					a = Setdiff(j, cf)
				case 4: // '4'
					savedigit(j, 4)
					bd = Setdiff(j, cf)
				case 5: // '235'
					abdcf := Setunion(a, Setunion(bd, cf))
					x := Setdiff(j, abdcf)
					if len(x) == 1 {
						// g = x
						if len(Setdiff(cf, j)) == 1 {
							// '5'
							savedigit(j, 5)
						} else {
							// '3'
							savedigit(j, 3)
						}
					} else if len(x) == 2 {
						// '2'
						savedigit(j, 2)
					}
				case 6: // '0' 6 9
					if len(Setdiff(bd, j)) > 0 {
						savedigit(j, 0)
					} else {
						if len(Setdiff(cf, j)) == 1 {
							savedigit(j, 6)
						} else {
							savedigit(j, 9)
						}
					}
				case 7:
					savedigit(j, 8)
				}
			}
			for i, s := range strings.Fields(output) {
				s := []byte(s)
				sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
				res[i] = matchdigit(digits, s)
			}
			n, _ := strconv.Atoi(string(res))
			return n
		}
		for i := range buf {
			n := decode(signal[i], output[i])
			res += n
		}
	}
	return res
}

func Day9(part2 bool) int {
	buf := Readfile("day9.txt")
	// 	buf := strings.Split(`2199943210
	// 3987894921
	// 9856789892
	// 8767896789
	// 9899965678`, "\n")
	res := 0
	height := len(buf)
	width := len(buf[0])
	a := make([][]uint8, height)
	type point struct {
		x, y       int
		size       int
		neighbours []*point
	}
	lowpoints := []point{}
	neighbours := func(a [][]uint8, p *point) (res []uint8) {
		i, j := p.x, p.y
		nb := []*point{}
		if i > 0 {
			nb = append(nb, &point{i - 1, j, 0, nil})
			res = append(res, a[i-1][j])
		}
		if i < height-1 {
			nb = append(nb, &point{i + 1, j, 0, nil})
			res = append(res, a[i+1][j])
		}
		if j > 0 {
			nb = append(nb, &point{i, j - 1, 0, nil})
			res = append(res, a[i][j-1])
		}
		if j < width-1 {
			nb = append(nb, &point{i, j + 1, 0, nil})
			res = append(res, a[i][j+1])
		}
		p.neighbours = nb
		return res
	}
	islowpoint := func(a [][]uint8, p *point) bool {
		i, j := p.x, p.y
		h := a[i][j]
		for _, x := range neighbours(a, p) {
			if h >= x {
				return false
			}
		}
		return true
	}
	in := func(y []point, x point) bool {
		for _, i := range y {
			if x.x == i.x && x.y == i.y {
				return true
			}
		}
		return false
	}
	// returns size of basin around p
	basin := func(a [][]uint8, p point) int {
		q := []point{p}
		visited := []point{}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if in(visited, p) {
				continue
			}
			visited = append(visited, p)
			neighbours(a, &p)
			for _, n := range p.neighbours {
				if a[n.x][n.y] < 9 && !in(visited, *n) {
					q = append(q, *n)
				}
			}
		}
		return len(visited)
	}
	// find lowpoints
	for i := range buf {
		d := make([]uint8, width)
		a[i] = d
		for j := range buf[i] {
			n, _ := strconv.Atoi(string(buf[i][j]))
			d[j] = uint8(n)
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			p := point{i, j, 0, nil}
			if islowpoint(a, &p) {
				lowpoints = append(lowpoints, p)
				res += int(a[i][j] + 1)
			}
		}
	}
	if !part2 {
		return res
	} else {
		// find basins
		for i := range lowpoints {
			size := basin(a, lowpoints[i])
			lowpoints[i].size = size
		}
		// sort descending
		sort.Slice(lowpoints, func(i, j int) bool { return lowpoints[i].size > lowpoints[j].size })
		res = lowpoints[0].size * lowpoints[1].size * lowpoints[2].size
	}
	return res
}

func Day10(part2 bool) int {
	buf := strings.Split(`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`, "\n")
	buf = Readfile("day10.txt")
	push := func(st []byte, b byte) []byte {
		return append([]byte{b}, st...)
	}
	pop := func(st []byte) ([]byte, byte, bool) {
		if len(st) > 0 {
			x := st[0]
			st = st[1:]
			return st, x, true
		} else {
			return st, 0, false
		}
	}
	type chunk struct {
		c   byte
		err int
	}
	tbl := map[byte]chunk{
		')': {'(', 3},
		']': {'[', 57},
		'}': {'{', 1197},
		'>': {'<', 25137},
	}
	tbl2 := map[byte]chunk{
		'(': {')', 1},
		'[': {']', 2},
		'{': {'}', 3},
		'<': {'>', 4},
	}
	err := 0
	scores := []int{}
	for _, s := range buf {
		stack := []byte{}
		errflag := false
		for _, b := range s {
			b := byte(b)
			expected, ok := tbl[b]
			if !ok {
				// opening
				stack = push(stack, b)
			} else {
				// closing
				stck, x, ok := pop(stack)
				stack = stck
				if ok {
					if expected.c != x {
						err += expected.err
						errflag = true
						break
					}
				} else {
					fmt.Println("empty stack")
					errflag = true
				}
			}
		}
		if part2 && !errflag {
			// if stack is not empty => incomplete lines
			score := 0
			for _, x := range stack {
				expected := tbl2[x]
				score = score*5 + expected.err
			}
			scores = append(scores, score)
		}
	}
	if part2 {
		sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
		return scores[len(scores)/2]
	} else {
		return err
	}
}

func Day11(part2 bool) int {
	buf := strings.Split(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`, "\n")
	/*
	   	buf = strings.Split(`11111
	   19991
	   19191
	   19991
	   11111`, "\n")
	*/
	buf = Readfile("day11.txt")
	height := len(buf)
	width := len(buf[0])
	type octo struct {
		energy uint8
		flash  bool
	}
	p := make([][]octo, height)
	for i, s := range buf {
		x := make([]octo, width)
		for j, c := range s {
			n, _ := strconv.Atoi(string(c))
			x[j] = octo{uint8(n), false}
		}
		p[i] = x
	}
	// main loop
	var boost func([][]octo, int, int)
	boost_neighbours := func(p [][]octo, i, j int) {
		x1 := Max(0, i-1)
		x2 := Min(height-1, i+1)
		y1 := Max(0, j-1)
		y2 := Min(width-1, j+1)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				boost(p, x, y)
			}
		}
	}
	boost = func(p [][]octo, i, j int) {
		p[i][j].energy++
		if p[i][j].energy > 9 && !p[i][j].flash {
			p[i][j].flash = true
			boost_neighbours(p, i, j)
		}
	}
	res := 0
	limit := 100
	// update and flag flashers
	if part2 {
		limit = math.MaxInt
	}
	for step := 0; step < limit; step++ {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				boost(p, i, j)
			}
		}
		if part2 {
			res = 0
		}
		// reset flashers
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if p[i][j].flash {
					res++
					p[i][j].energy = 0
					p[i][j].flash = false
				}
			}
		}
		if part2 {
			if res == height*width {
				return step + 1
			}
		}
	}
	return res
}

func Day12(part2 bool) int {
	ex2 := strings.Split(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, "\n")
	ex3 := strings.Split(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`, "\n")
	buf := strings.Split(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`, "\n")
	in := func(y map[string]int, x string) bool {
		for n, _ := range y {
			if x == n {
				return true
			}
		}
		return false
	}
	type node struct {
		name string
		prev *node
	}
	/*
		debug := func(n node) {
			for n.prev != nil {
				fmt.Print(n.name, ",")
				n = *n.prev
			}
			fmt.Println(n.name)
		}
	*/
	g := map[string][]string{}
	small := map[string]int{}
	// returns true if n is found in p's ancestry
	trace := func(n string, p node) bool {
		if !in(small, n) {
			return false
		}
		count := map[string]int{n: 1}
		flag := false
		for p.prev != nil {
			if part2 {
				// ignore uppercase
				if in(small, p.name) {
					count[p.name]++
					if count[p.name] > 1 {
						if flag { // someone else is already doubled
							// fmt.Println("trace FAIL", n, count)
							return true
						}
						flag = true
					}
				}
			} else {
				// part1
				if p.name == n {
					return true
				}
			}
			p = *p.prev
		}
		// fmt.Println("trace OK", n, count)
		return false
	}
	// ----start----
	// buf = ex3
	buf = Readfile("day12.txt")
	x := len(ex2) + len(ex3)
	for _, s := range buf {
		s := strings.Split(s, "-")
		if s[1] != "start" {
			g[s[0]] = append(g[s[0]], s[1])
		}
		if s[0] != "start" {
			g[s[1]] = append(g[s[1]], s[0])
		}
		if strings.ToLower(s[0]) == s[0] {
			small[s[0]]++
		}
		if strings.ToLower(s[1]) == s[1] {
			small[s[1]]++
		}
	}
	res := 0
	x++
	frontier := []node{{"start", nil}}
	for len(frontier) > 0 {
		n := frontier[0]
		frontier = frontier[1:]
		if n.name == "end" {
			// debug(n)
			res++
			continue
		}
		for _, i := range g[n.name] {
			if !trace(i, n) {
				frontier = append(frontier, node{i, &n})
			}
		}
	}
	return res
}

// bonus: output to PNG !
func Day13(part2 bool) int {
	buf := Readstring(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`)
	buf = Readfile("day13.txt")
	type point [2]int
	width, height := 0, 0
	dots := []point{}
	fold := []point{}
	// updates dots with new positions
	newpos := func(i, line int) int {
		if i > line {
			i = line - (i - line)
		}
		return i
	}
	makefold := func(p point) {
		if p[0] == 0 {
			// fold up along y=...
			for i, d := range dots {
				dots[i][1] = newpos(d[1], p[1])
			}
		} else {
			// fold left along x=...
			for i, d := range dots {
				dots[i][0] = newpos(d[0], p[0])
			}
		}
	}
	count := func() int {
		m := map[int]int{}
		for _, d := range dots {
			m[d[0]*0x100000000+d[1]]++
		}
		return len(m)
	}
	image := func(l, t, r, b int, data map[int]point) image.Image {
		img := image.NewGray(image.Rectangle{image.Point{l, t}, image.Point{r + 1, b + 1}})
		for _, p := range data {
			img.Set(p[0], p[1], color.Gray{0xff})
		}

		f, err := os.Create("day13image.png")
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(f, img); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		return img
	}
	// ingest
	for _, s := range buf {
		if len(s) == 0 {
			continue
		}
		if strings.Contains(s, "fold") {
			s := strings.TrimPrefix(s, "fold along ")
			var axis int
			fmt.Sscanf(s[2:], "%d", &axis)
			p := point{}
			if s[0] == 'x' {
				p = point{axis, 0}
			} else {
				p = point{0, axis}
			}
			fold = append(fold, p)
		} else {
			x, y := 0, 0
			fmt.Sscanf(s, "%d,%d", &x, &y)
			dots = append(dots, point{x, y})
			if x > width {
				width = x
			}
			if y > height {
				height = y
			}
		}
	}
	// fold & update pos
	if !part2 {
		makefold(fold[0])
		return count()
	} else {
		for _, f := range fold {
			makefold(f)
		}
		m := map[int]point{}
		l, r, t, b := width, 0, height, 0
		// unique points & new extents
		for _, d := range dots {
			m[d[0]*0x100000000+d[1]] = d
			if d[0] > r {
				r = d[0]
			}
			if d[0] < l {
				l = d[0]
			}
			if d[1] > b {
				b = d[1]
			}
			if d[1] < t {
				t = d[1]
			}
		}
		grid := make([][]byte, b-t+1)
		for i := range grid {
			row := bytes.Repeat([]byte{' '}, r-l+1)
			grid[i] = row
		}
		for _, p := range m {
			grid[p[1]][p[0]] = '#'
		}
		for i := range grid {
			fmt.Println(string(grid[i]))
		}
		if false {
			image(l, t, r, b, m)
		}
	}
	return 0
}

func Day14(part2 bool) int {
	buf := Readstring(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`)
	buf = Readfile("day14.txt")
	template := buf[0]
	rules := map[string][]string{}
	for _, s := range buf[2:] {
		s := strings.Split(s, " -> ")
		rules[s[0]] = []string{s[0][:1] + s[1], s[1] + s[0][1:2]}
	}
	ngrams := map[string]int{}
	// chop template into 2-grams
	for i := 0; i < len(template)-1; i++ {
		ngrams[template[i:i+2]]++
	}
	tail := template[len(template)-2:]
	limit := 10
	if part2 {
		limit = 40
	}
	for cycle := 0; cycle < limit; cycle++ {
		res := map[string]int{}
		// fmt.Println(ngrams, tail)
		for i, n := range ngrams {
			for _, j := range rules[i] {
				res[j] += n
			}
		}
		tail = rules[tail][1]
		ngrams = res
	}
	// element counts
	freq := map[string]int{}
	for i, n := range ngrams {
		freq[i[0:1]] += n
	}
	freq[tail[1:]]++
	min, max := math.MaxInt, 0
	for _, n := range freq {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max - min
}

// TODO: this would benefit from using a heap as frontier
func Day15(part2 bool) int {
	buf := Readstring(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`)
	buf = Readfile("day15.txt")
	height := len(buf)
	width := len(buf[0])
	grid := make([][]uint8, height)
	for i := range grid {
		row := make([]uint8, width)
		for j := range buf[i] {
			n, _ := strconv.Atoi(buf[i][j : j+1])
			row[j] = uint8(n)
		}
		grid[i] = row
	}
	type point [2]int
	type node struct {
		p    point
		path []*node
		cost int
	}
	fcompact := func(f []node) []node {
		m := make(map[point]node)
		for _, n := range f {
			if mem, ok := m[n.p]; ok {
				if n.cost < mem.cost {
					m[n.p] = n
				}
			} else {
				m[n.p] = n
			}
		}
		s := make([]node, len(m))
		i := 0
		for _, n := range m {
			s[i] = n
			i++
		}
		sort.Slice(s, func(i, j int) bool { return s[i].cost < s[j].cost })
		return s
	}
	getgrid := func(p point) int {
		if !part2 {
			return int(grid[p[0]][p[1]])
		} else {
			tiley, offy := p[0]/height, p[0]%height
			tilex, offx := p[1]/width, p[1]%width
			n := int(grid[offy][offx])
			n = ((n - 1 + tilex + tiley) % 9) + 1
			return n
		}
	}
	pos := point{0, 0}
	end := point{height - 1, width - 1}
	if part2 {
		end = point{height*5 - 1, width*5 - 1}
	}
	frontier := []node{{pos, nil, 0}}
	visited := make(map[point]bool)
	bestcost := math.MaxInt
	// UCS
	for len(frontier) > 0 {
		frontier = fcompact(frontier)
		n := frontier[0]
		frontier = frontier[1:]
		p := n.p
		if p == end {
			// fmt.Println("found", n.cost)
			if n.cost < bestcost {
				bestcost = n.cost
			}
			// in UCS the first is the shortest path :)
			break
		}
		if visited[p] {
			// abandon already visited
			continue
		}
		visited[p] = true
		// newpath := append(n.path, &n)
		enqueue := func(np point) {
			cost := n.cost + getgrid(np)
			if !visited[np] {
				// frontier = append(frontier, node{np, newpath, cost})
				frontier = append(frontier, node{np, nil, cost})
			}
		}
		if p[0] > 0 {
			npoint := point{p[0] - 1, p[1]}
			enqueue(npoint)
		}
		if p[0] < end[0] {
			npoint := point{p[0] + 1, p[1]}
			enqueue(npoint)
		}
		if p[1] > 0 {
			npoint := point{p[0], p[1] - 1}
			enqueue(npoint)
		}
		if p[1] < end[1] {
			npoint := point{p[0], p[1] + 1}
			enqueue(npoint)
		}
	}
	return bestcost
}

func Day16(part2 bool) int {
	// version, ID
	// ID=4 literal int, padded then expanded to 5-bit groups (leading 1 except last)
	// ID!=4 operator: length bit 0: 15 bit len; 1: 11 bit subpkt count
	//					subpackets
	ex := []string{`D2FE28`,
		`38006F45291200`,
		`EE00D40C823060`,
		`8A004A801A8002F478`,
		`620080001611562C8802118E34`,
		`C0015000016115A2E0802F182340`,
		`A0016C880162017C3686B18A3D4780`}
	ex2 := []string{
		`C200B40A82`,
		`04005AC33890`,
		`880086C3E88112`,
		`CE00C43D881120`,
		`D8005AC2A8F0`,
		`F600BC2D8F`,
		`9C005AC2F8F0`,
		`9C0141080250320F1802104A08`}
	extract := func(b []byte, start, leng int) []byte {
		off := start / 8
		shift := start % 8
		mask := []byte{0xff, 0x80, 0xc0, 0xe0, 0xf0, 0xf8, 0xfc, 0xfe}
		res := make([]byte, (leng+7)/8)
		for i := 0; i < len(res)-1; i++ {
			if shift > 0 {
				u16 := uint16(b[off+i])<<8 | uint16(b[off+i+1])
				res[i] = uint8(u16 >> (8 - shift))
				// res[i] = b[off+i]<<shift | (b[off+i+1]>>(8-shift))&mask[shift]
			} else {
				res[i] = b[off+i]
			}
		}
		// last byte
		if shift > 0 {
			x := uint16(0)
			if off+len(res) < len(b) {
				x = uint16(b[off+len(res)])
			}
			u16 := uint16(b[off+len(res)-1])<<8 | x
			res[len(res)-1] = uint8(u16>>(8-shift)) & mask[leng%8]
		} else {
			res[len(res)-1] = b[off+len(res)-1] & mask[leng%8]
		}
		return res
	}
	var evalpacket func([]byte) (int, int, uint)
	// recursively evaluates a packet and its subpackets
	evalpacket = func(packet []byte) (bits, version int, value uint) {
		version = (int(packet[0]) >> 5) & 7
		id := (packet[0] >> 2) & 7
		if id == 4 {
			// literal
			done := false
			off := 6
			value = uint(0)
			for !done {
				b := extract(packet, off, 5)[0] >> 3
				off += 5
				if b&0x10 == 0 {
					done = true
				}
				b &= 0x0f
				value = (value << 4) | uint(b)
			}
			return off, version, value
		}
		ltype := packet[0] >> 1 & 1
		start := 0
		vals := []uint{}
		if ltype == 0 {
			// 15 bit subpacket bitlen
			n := uint16(packet[1])<<8 | uint16(packet[2])
			n >>= 2
			if packet[0]&1 > 0 {
				n |= 0x4000
			}
			// fmt.Println(version, id, "type", ltype, "subpacket bitlen", n)
			start = 15 + 7
			for start < int(n)+15+7 {
				b, v, vl := evalpacket(extract(packet, start, len(packet)*8-start))
				start += b
				version += v
				vals = append(vals, vl)
			}
		} else {
			// 11 bit subpacket count
			n := uint16(packet[1])<<8 | uint16(packet[2])
			n >>= 6
			if packet[0]&1 > 0 {
				n |= 0x0400
			}
			// fmt.Println(version, id, "type", ltype, "subpacket count", n)
			start = 11 + 7
			for i := 0; i < int(n); i++ {
				b, v, vl := evalpacket(extract(packet, start, len(packet)*8-start))
				start += b
				version += v
				vals = append(vals, vl)
			}
		}
		if !part2 {
			return start, version, 0
		}
		switch id {
		case 0: // sum
			for _, v := range vals {
				value += v
			}
		case 1: // product
			value = 1
			for _, v := range vals {
				value *= v
			}
		case 2: // minimum
			min := uint(math.MaxUint)
			for _, v := range vals {
				if v < min {
					min = v
				}
			}
			value = min
		case 3: // maximum
			max := uint(0)
			for _, v := range vals {
				if v > max {
					max = v
				}
			}
			value = max
		case 5: // greater than
			if vals[0] > vals[1] {
				value = 1
			}
		case 6: // less than
			if vals[0] < vals[1] {
				value = 1
			}
		case 7: // equal to
			if vals[0] == vals[1] {
				value = 1
			}
		}
		return start, version, value
	}
	buf := ex[4]
	buf = ex2[0]
	buf = Readfile("day16.txt")[0]
	work := func(buf string) int {
		// ingest
		input := make([]byte, len(buf)/2)
		for i := 0; i < len(input); i++ {
			n := buf[i<<1] - byte('0')
			if n > 9 {
				n -= 7
			}
			m := buf[i<<1+1] - byte('0')
			if m > 9 {
				m -= 7
			}
			input[i] = n<<4 | m
		}
		_, versionsum, value := evalpacket(extract(input, 0, len(input)*8))
		if !part2 {
			return versionsum
		} else {
			return int(value)
		}
	}
	return work(buf)
}

func Day17(part2 bool) int {
	buf := Readstring(`target area: x=20..30, y=-10..-5`)[0]
	buf = Readfile("day17.txt")[0]
	type point [2]int
	x1, x2, y1, y2 := 0, 0, 0, 0
	if n, err := fmt.Sscanf(buf, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2); n != 4 {
		log.Fatal(err)
	}
	update := func(pos point, vector point) (point, point) {
		// each step:
		// x += vx, y += vy
		// vx = max(vx-1, 0)
		// vy--
		pos[0] += vector[0]
		pos[1] += vector[1]
		vector[0] = Max(vector[0]-1, 0)
		vector[1]--
		return pos, vector
	}
	intersect := func(pos point) bool {
		res := 0
		if pos[0] >= x1 {
			res |= 1
		}
		if pos[0] <= x2 {
			res |= 2
		}
		if pos[1] >= y1 {
			res |= 4
		}
		if pos[1] <= y2 {
			res |= 8
		}
		return res == 0xf
	}
	// return yhigh for this run, or -1 if no intersect
	simulate := func(vec point) int {
		yhigh := 0
		pos := point{0, 0}
		for !intersect(pos) {
			// fmt.Printf("pos %v vec %v\n", pos, vec)
			pos, vec = update(pos, vec)
			if pos[1] > yhigh {
				yhigh = pos[1]
			}
			if pos[0] > x2 || pos[1] < y1 {
				// left the arena
				return -1
			}
		}
		return yhigh
	}
	// vec := point{6, 9}
	vec := point{10, 0}
	// find a horizontal vector that is a solution
	yhigh := 0
	count := 0
	for x := 00; x < 600; x++ {
		for y := -300; y < 600; y++ {
			vec[0], vec[1] = x, y
			n := simulate(vec)
			if n >= 0 {
				count++
				// fmt.Println(vec, n)
			}
			if n > yhigh {
				yhigh = n
			}
		}
	}
	if !part2 {
		return yhigh
	} else {
		return count
	}
}

func Day18(part2 bool) int {
	type node struct {
		left, right    int
		lchild, rchild *node
	}
	mknode := func(x, y int) node {
		return node{left: x, right: y}
	}
	addchild := func(n node, l, r *node) node {
		n.lchild, n.rchild = l, r
		return n
	}
	add := func(a, b node) node {
		return node{lchild: &a, rchild: &b}
	}
	/*
			buf := Readstring(`[1,2]
		[[1,2],3]
		[9,[8,7]]
		[[1,9],[8,5]]
		[[[[1,2],[3,4]],[[5,6],[7,8]]],9]
		[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]
		[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]`)
			// ingest
			parse := func(s string) {

			}
	*/
	btree := mknode(1, 2)
	n := mknode(3, 4)
	fmt.Println(add(btree, n))
	btree = addchild(btree, &n, nil)
	fmt.Println(btree)
	return 0
}

func Day19(part2 bool) int {
	buf := Readstring(`--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`)
	buf = Readfile("day19.txt")
	type point [3]int
	type scanner struct {
		id         int
		pings      []point
		dist       map[int][]point
		relativeto *scanner
		rot, off   point
		parent     *scanner
		children   []*scanner
	}
	ingest := func() []scanner {
		scanners := []scanner{}
		pings := []point{}
		i := 0
		for _, s := range buf {
			n, _ := fmt.Sscanf(s, "--- scanner %d ---", &i)
			if n == 1 {
				if i > 0 {
					scanners = append(scanners,
						scanner{id: i - 1, pings: pings, dist: nil})
					pings = []point{}
				}
				continue
			}
			x, y, z := 0, 0, 0
			n, _ = fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z)
			if n == 3 {
				pings = append(pings, point{x, y, z})
			}
		}
		scanners = append(scanners, scanner{id: i, pings: pings, dist: nil})
		return scanners
	}
	// dist matrix
	distance := func(p []point, method string) map[int][]point {
		m := map[int][]point{}
		sq := func(x int) int { return x * x }
		if method == "manhattan" {
			sq = func(x int) int {
				if x < 0 {
					return -x
				}
				return x
			}
		}
		for i := 0; i < len(p)-1; i++ {
			for j := i + 1; j < len(p); j++ {
				d := sq(p[i][0]-p[j][0]) + sq(p[i][1]-p[j][1]) + sq(p[i][2]-p[j][2])
				// crazy non-unique dist exist
				m[d] = append(m[d], []point{p[i], p[j]}...)
			}
		}
		return m
	}
	Z := point{0, 0, 0}
	vecdiff := func(p, q point) point { return point{p[0] - q[0], p[1] - q[1], p[2] - q[2]} }
	vecmul := func(p, q point) point { return point{p[0] * q[0], p[1] * q[1], p[2] * q[2]} }
	// vecadd := func(p, q point) point { return point{p[0] + q[0], p[1] + q[1], p[2] + q[2]} }
	xmat := []point{
		{1, 2, 3}, {2, -1, 3}, {-1, -2, 3}, {-2, 1, 3},
		{-1, 2, -3}, {-2, -1, -3}, {1, -2, -3}, {2, 1, -3},
		{-3, 2, 1}, {-3, -1, 2}, {-3, -2, -1}, {-3, 1, -2},
		{3, 2, -1}, {3, -1, -2}, {3, -2, 1}, {3, 1, 2},
		{1, -3, 2}, {2, -3, -1}, {-1, -3, -2}, {-2, -3, 1},
		{1, 3, -2}, {2, 3, 1}, {-1, 3, 2}, {-2, 3, -1},
	}
	transform := func(p, xmat point) point {
		if xmat == Z {
			return p
		}
		m := point{}
		sign := point{1, 1, 1}
		for i := range m {
			m[i] = xmat[i]
			if m[i] < 0 {
				sign[i] = -1
				m[i] = -m[i]
			}
			m[i]--
		}
		r := point{p[m[0]], p[m[1]], p[m[2]]}
		r = vecmul(r, sign)
		return r
	}

	originalScanners := ingest()
	scanners := originalScanners
	autoinc := 40

	for len(scanners) > 1 {
		for i := range scanners {
			scanners[i].dist = distance(scanners[i].pings, "manhattan")
		}

		// given a list of common distances, look up the pointpairs and
		// work out the orientation vector and offset (from j to i)
		findOrientation := func(i, j int, I []int) (rot, off point, ok bool) {
			// d := I[0]
			try := func(v1, v2 point) (res point) {
				// transform v2 to match v1
				for j := range xmat {
					r := transform(v2, xmat[j])
					rv1 := point{-v1[0], -v1[1], -v1[2]} // reversed
					if vecdiff(v1, r) == Z || vecdiff(rv1, r) == Z {
						// fmt.Println("rot match", xmat[j], v1, v2)
						return xmat[j]
					}
				}
				// fmt.Println("no rot match")
				return res
			}
			// returns modal candidate
			best := func(candidates map[point]int) point {
				n := 0
				rot := point{}
				for k, v := range candidates {
					if v > n {
						rot = k
						n = v
					}
				}
				return rot
			}
			// returns the offset between the scanners' centres
			checkAlignment := func(rot point) point {
				count := 0
				off := map[point]int{}
				for _, d := range I {
					d1 := scanners[i].dist[d]
					d2 := scanners[j].dist[d]
					v1 := vecdiff(d1[0], d1[1])
					rv1 := vecdiff(d1[1], d1[0])
					d2[0] = transform(d2[0], rot)
					d2[1] = transform(d2[1], rot)
					v2 := vecdiff(d2[0], d2[1])
					if vecdiff(v2, v1) == Z {
						off[vecdiff(d2[0], d1[0])]++
						count++
					} else if vecdiff(v2, rv1) == Z {
						off[vecdiff(d2[1], d1[0])]++
						count++
					}
				}
				// fmt.Printf("aligned %d/%d\n  %v\n", count, len(I), off)
				for k := range off {
					return k
				}
				return point{}
			}
			// --------
			candidates := map[point]int{}
			for _, d := range I {
				d1 := scanners[i].dist[d]
				d2 := scanners[j].dist[d]
				// v1 := vecdiff(d1[0], d1[1])
				// v2 := vecdiff(d2[0], d2[1])
				p := 0
				for p < len(d1) {
					v1 := vecdiff(d1[p], d1[p+1])
					p += 2
					q := 0
					for q < len(d2) {
						v2 := vecdiff(d2[q], d2[q+1])
						q += 2
						if t := try(v1, v2); t != Z {
							candidates[t]++
						}
					}
				}
			}
			rot = best(candidates)
			if rot == Z {
				// fmt.Printf("%d-%d no candidates!\n", i, j)
				// fmt.Println(I)
				// fmt.Println(candidates)
				return rot, off, false
			}
			// check that all of j's pointpairs match i's after rot & translation
			off = checkAlignment(rot)
			// fmt.Println(i, j, "chose", rot, off, "from", candidates)
			return rot, off, true
		}

		// identify scanner-pairs that have common pointpairs, and figure out the transform
		findGroups := func() [][2]int {
			res := [][2]int{}
			for i := range scanners {
				x := []int{}
				for k := range scanners[i].dist {
					x = append(x, k)
				}
				for j := i + 1; j < len(scanners); j++ {
					if scanners[j].relativeto != nil {
						continue
					}
					y := []int{}
					for k := range scanners[j].dist {
						y = append(y, k)
					}
					// intersect = union - setdiff(x,y) - setdiff(y,x)
					U := SetunionInt(x, y)
					Dleft := SetdiffInt(x, y)
					Dright := SetdiffInt(y, x)
					I := SetdiffInt(U, SetunionInt(Dleft, Dright))
					// find the orientation vector & offset using pointpairs common to both
					if len(I) > 11 {
						if rot, off, ok := findOrientation(i, j, I); ok {
							res = append(res, [2]int{i, j})
							scanners[j].relativeto = &scanners[i]
							scanners[j].rot = rot
							scanners[j].off = off
						}
					}
				}
			}
			return res
		}
		findGroups()

		/*
			println("scanner children rot off")
			for _, s := range scanners {
				fmt.Println(s.id, len(s.children), s.rot, s.off)
			}
			println()
		*/
		// transform each scanner's pings
		g := map[*scanner][]*scanner{} // use a graph to collect the pieces
		for i := range scanners {
			s := &scanners[i]
			relto := s.relativeto
			rot, off := s.rot, s.off
			if relto == nil {
				g[s] = []*scanner{s}
				continue
			}
			var last *scanner = nil
			for relto != nil {
				for j, p := range s.pings {
					s.pings[j] = vecdiff(transform(p, rot), off)
				}
				rot = relto.rot
				off = relto.off
				last = relto
				relto = relto.relativeto
			}
			g[last] = append(g[last], s)
		}
		// fmt.Println("graph", g)
		// fuse scanners together, then rinse and repeat
		// build up the scanner tree for part2
		fuse := func(g map[*scanner][]*scanner) []scanner {
			x := make([]*scanner, len(g))
			i := 0
			for k := range g {
				x[i] = k
				i++
			}
			sort.Slice(x, func(i, j int) bool { return x[i].id < x[j].id })
			scan2 := []scanner{}
			for _, v := range x {
				pings := map[point]bool{}
				children := []*scanner{}
				// fmt.Printf("fuse %d - ", autoinc)
				newscanner := scanner{id: autoinc, dist: map[int][]point{}}
				autoinc++
				for _, s := range g[v] {
					// fmt.Printf("%d ", s.id)
					s.parent = &newscanner
					children = append(children, s)
					for _, p := range s.pings {
						pings[p] = true
					}
				}
				// println()
				pslice := make([]point, len(pings))
				i := 0
				for p := range pings {
					pslice[i] = p
					i++
				}
				newscanner.children = children
				newscanner.pings = pslice
				scan2 = append(scan2, newscanner)
			}
			return scan2
		}
		scanners = fuse(g)
	}

	if !part2 {
		return len(scanners[0].pings)
	} else {
		scannerpos := make([]point, len(originalScanners))
		i := 0
		type node struct {
			scanner
			xform [][2]point
		}
		frontier := []node{{scanners[0], [][2]point{}}}
		for len(frontier) > 0 {
			n := frontier[len(frontier)-1]
			frontier = frontier[:len(frontier)-1]
			if len(n.children) == 0 {
				relto := &n.scanner
				p := point{}
				for {
					p = vecdiff(transform(p, relto.rot), relto.off)
					relto = relto.relativeto
					if relto == nil {
						break
					}
				}
				for _, x := range n.xform {
					p = vecdiff(transform(p, x[0]), x[1])
				}
				scannerpos[n.id] = p
				i++
				continue
			}
			x := make([][2]point, len(n.xform)+1)
			copy(x[1:], n.xform)
			x[0] = [2]point{n.rot, n.off}
			for _, c := range n.children {
				frontier = append(frontier, node{*c, x})
			}
		}
		// manhattan distance
		max := 0
		for d := range distance(scannerpos, "manhattan") {
			if d > max {
				max = d
			}
		}
		return max
	}
	return 0
}

func Day20(part2 bool) int {
	return 0
}

func Day21(part2 bool) int {
	buf := Readstring(`Player 1 starting position: 4
Player 2 starting position: 8`)
	// buf = Readfile("day21.txt")
	pos := [2]int{}
	pos[0], _ = strconv.Atoi(strings.Fields(buf[0])[4])
	pos[1], _ = strconv.Atoi(strings.Fields(buf[1])[4])
	rollcount := 0
	score := [2]int{}
	die := 1
	for {
		// p1 roll 3x
		x := 0
		for i := 0; i < 3; i++ {
			x += die
			die = (die % 100) + 1
			rollcount++
		}
		pos[0] = ((pos[0] - 1 + x) % 10) + 1
		score[0] += pos[0]
		if score[0] >= 1000 {
			fmt.Println(score[1], rollcount)
			return score[1] * rollcount
		}
		// p2 roll 3x
		x = 0
		for i := 0; i < 3; i++ {
			x += die
			die = (die % 100) + 1
			rollcount++
		}
		pos[1] = ((pos[1] - 1 + x) % 10) + 1
		score[1] += pos[1]
		if score[1] >= 1000 {
			fmt.Println(score[0], rollcount)
			return score[0] * rollcount
		}
	}
	return 0
}

func Day22(part2 bool) int {
	return 0
}

func Day23(part2 bool) int {
	return 0
}

func Day24x(part2 bool) int {
	buf := Readstring(`inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`)
	buf = Readfile("day24.txt")
	text := [][]string{}
	blocks := [][2]int{}
	start, end := 0, -1
	for i, s := range buf {
		s := strings.Fields(s)
		if s[0] == "inp" {
			end = i
			if end > 0 {
				blocks = append(blocks, [2]int{start, end})
			}
			start = i
		}
		text = append(text, s)
	}
	blocks = append(blocks, [2]int{start, len(buf)})
	// reg := [4]int{}
	memo := map[string][4]int{}
	runblock := func(input byte, n int, reg [4]int) [4]int {
		argreg := func(s string) uint8 {
			return s[0] - byte('w')
		}
		arg := func(s string) int {
			n, err := strconv.Atoi(s)
			if err != nil {
				n = reg[argreg(s)]
			}
			return n
		}
		start := blocks[n][0]
		end := blocks[n][1]
		for _, s := range text[start:end] {
			switch s[0] {
			case "inp":
				x := argreg(s[1])
				reg[x] = int(input - byte('0'))
			case "add":
				x, y := argreg(s[1]), arg(s[2])
				reg[x] += y
			case "mul":
				x, y := argreg(s[1]), arg(s[2])
				reg[x] *= y
			case "div":
				x, y := argreg(s[1]), arg(s[2])
				reg[x] /= y
			case "mod":
				x, y := argreg(s[1]), arg(s[2])
				reg[x] %= y
			case "eql":
				x, y := argreg(s[1]), arg(s[2])
				if reg[x] == y {
					reg[x] = 1
				} else {
					reg[x] = 0
				}
			default:
				fmt.Println("error:", s[0])
			}
		}
		return reg
	}
	var monad func([]byte) [4]int
	monad = func(input []byte) [4]int {
		inputp := len(input) - 1
		reg, ok := memo[string(input[:inputp])]
		if !ok {
			if inputp > 1 {
				reg = monad(input[:inputp])
			} else {
				reg = runblock(input[0], 0, reg)
			}
			memo[string(input[:inputp])] = reg
		}
		reg = runblock(input[inputp], inputp, reg)
		return reg
	}
	var decr func([]byte, int) []byte
	decr = func(s []byte, p int) []byte {
		n := int(s[p]-byte('0')) - 1
		if n == 0 {
			n = 9
			if p > 0 {
				decr(s, p-1)
			}
		}
		s[p] = byte(n) + byte('0')
		return s
	}
	/* parallel listing
	for i := 0; i < blocks[0][1]; i++ {
		for j := range blocks {
			fmt.Printf("%10s", buf[blocks[j][0]+i])
		}
		println()
	}
	return 0
	*/
	// exhaustive search
	input := []byte("11119227949959")
	reps := 0
	for {
		r := monad(input)
		if r[3] == 0 {
			fmt.Println("found", input, r)
			break
		}
		input = decr(input, 13)
		reps++
		if reps%1000 == 0 {
			print(string(input), "\r")
			// fmt.Println(memo)
			// break
		}
	}
	return 0
}

func Day24(part2 bool) int {
	zmod := []int{14, 14, 14, 12, 15, -12, -12, 12, -7, 13, -8, -5, -10, -7}
	zdiv := []int{1, 1, 1, 1, 1, 26, 26, 1, 26, 1, 26, 26, 26, 26}
	wadd := []int{14, 2, 1, 13, 5, 5, 5, 9, 3, 13, 2, 1, 11, 8}
	runone := func(input byte, step int, z int) int {
		w := int(input - byte('0'))
		if w != (z%26)+zmod[step] {
			z = z / zdiv[step]
			z = z*26 + (w + wadd[step])
		} else {
			z = z / zdiv[step]
		}
		println("z after step", step, z, z%26)
		return z
	}
	monad := func(input []byte) int {
		z := 0
		for i := 0; i < 14; i++ {
			w := int(input[i] - byte('0'))
			if w != (z%26)+zmod[i] {
				z = z / zdiv[i]
				z = z*26 + (w + wadd[i])
			} else {
				z = z / zdiv[i]
			}
		}
		return z
	}
	// work backwards
	// for step := 13; step >= 0; step-- {
	step := 12
	for i := 1; i < 10; i++ {
		for z := 0; z < 26; z++ {
			res := runone('0'+byte(i), step, 10220+z)
			fmt.Printf("step %d: i %d z %d -> %d %d ", step, i, 10220+z, res, res%26)
		}
	}
	// }

	// z 0, 15, 393, ... 6909090 after step4
	z := 0
	z = runone('1', 0, z)
	z = runone('1', 1, z)
	z = runone('1', 2, z)
	z = runone('1', 3, z)
	z = runone('9', 4, z)
	z = runone('2', 5, z) // z 13-21
	z = runone('2', 6, z) // z 13-21
	z = runone('7', 7, z) // i 1-7 z 0
	z = runone('9', 8, z) // z 8-16
	z = runone('4', 9, z) // z must be 0
	z = runone('9', 10, z)
	z = runone('1', 11, z)
	z = runone('5', 12, z)
	z = runone('9', 13, z)
	println(z)
	input := []byte("11119227949959")
	println(monad(input))
	return 0
	reps := 1
	for {
		print(string(input), "\r")
		if monad(input) == 0 {
			fmt.Println("found", input)
			break
		}
		reps--
		if reps == 0 {
			println()
			break
		}
	}
	return 0
}

func Day25(part2 bool) int {
	return 0
}
