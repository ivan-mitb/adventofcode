package aoc

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Readfile(fn string) (buf []string) {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		buf = append(buf, s.Text())
	}
	f.Close()
	return
}

func setdiff(a, b []byte) (res []byte) {
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

func setunion(a, b []byte) []byte {
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
					a = setdiff(j, cf)
				case 4: // '4'
					savedigit(j, 4)
					bd = setdiff(j, cf)
				case 5: // '235'
					abdcf := setunion(a, setunion(bd, cf))
					x := setdiff(j, abdcf)
					if len(x) == 1 {
						// g = x
						if len(setdiff(cf, j)) == 1 {
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
					if len(setdiff(bd, j)) > 0 {
						savedigit(j, 0)
					} else {
						if len(setdiff(cf, j)) == 1 {
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
	// buf := Readfile("day9.txt")
	return 0
}
