package aoc2021

import (
	"aoc"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1(part2 bool) int {
	buf := aoc.Readfile("day1.txt")
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
	buf := aoc.Readfile("day2.txt")
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
	buf := aoc.Readfile("day3.txt")
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
	buf := aoc.Readfile("day4.txt")
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
	buf := aoc.Readfile("day5.txt")
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
	buf := aoc.Readfile("day6.txt")
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
	buf := aoc.Readfile("day7.txt")
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
	buf := aoc.Readfile("day8.txt")
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
					a = aoc.Setdiff(j, cf)
				case 4: // '4'
					savedigit(j, 4)
					bd = aoc.Setdiff(j, cf)
				case 5: // '235'
					abdcf := aoc.Setunion(a, aoc.Setunion(bd, cf))
					x := aoc.Setdiff(j, abdcf)
					if len(x) == 1 {
						// g = x
						if len(aoc.Setdiff(cf, j)) == 1 {
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
					if len(aoc.Setdiff(bd, j)) > 0 {
						savedigit(j, 0)
					} else {
						if len(aoc.Setdiff(cf, j)) == 1 {
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
	buf := aoc.Readfile("day9.txt")
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
	buf = aoc.Readfile("day10.txt")
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
	buf = aoc.Readfile("day11.txt")
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
		x1 := aoc.Max(0, i-1)
		x2 := aoc.Min(height-1, i+1)
		y1 := aoc.Max(0, j-1)
		y2 := aoc.Min(width-1, j+1)
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
		limit = 1e6
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
	buf = aoc.Readfile("day12.txt")
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
	buf := aoc.Readstring(`6,10
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
	buf = aoc.Readfile("day13.txt")
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
	buf := aoc.Readstring(`NNCB

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
	// buf = aoc.Readfile("day14.txt")
	template := buf[0]
	rules := map[string]string{}
	for _, s := range buf[2:] {
		s := strings.Split(s, " -> ")
		rules[s[0]] = s[1]
	}
	for cycle := 0; cycle < 10; cycle++ {
		res := bytes.NewBufferString(template[:1])
		for i := 0; i < len(template)-1; i++ {
			res.WriteString(rules[template[i:i+2]])
			res.WriteByte(byte(template[i+1]))
			// fmt.Print(res, " ")
		}
		template = res.String()
		fmt.Println(len(template))
	}
	// element counts
	freq := map[string]int{}
	for _, c := range template {
		freq[string(c)]++
	}
	min, max := int(1e6), 0
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
