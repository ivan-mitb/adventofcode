// package aoc2015 implements solutions for Advent of Code 2015
package aoc2015

import (
	"aoc"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/1/input

func Day1(part2 bool) int {
	buf := aoc.Readfile("day1.txt")
	i := 0
	for n, c := range buf[0] {
		if part2 {
			if i < 0 {
				return n
			}
		}
		switch c {
		case '(':
			i += 1
		case ')':
			i -= 1
		}
	}
	return i
}

func Day2(part2 bool) int {
	buf := aoc.Readfile("day2.txt")
	var area, peri int
	for _, s := range buf {
		l := []int{0, 0, 0}
		fmt.Sscanf(s, "%dx%dx%d", &l[0], &l[1], &l[2])
		sort.Ints(l)
		area += 3*l[0]*l[1] + 2*(l[1]*l[2]+l[2]*l[0])
		peri += 2*(l[0]+l[1]) + l[0]*l[1]*l[2]
	}
	if !part2 {
		return area
	} else {
		return peri
	}
}

func Day3(part2 bool) int {
	buf := aoc.Readfile("day3.txt")[0]
	type point [2]int
	var location [2]point
	var who int // 0 santa 1 robo
	acc := make(map[point]int)
	acc[location[who]] = 2
	for _, move := range buf {
		switch move {
		case '^':
			location[who][1] += 1
		case 'v':
			location[who][1] -= 1
		case '>':
			location[who][0] += 1
		case '<':
			location[who][0] -= 1
		}
		acc[location[who]] += 1
		if part2 {
			who = 1 - who
		}
	}
	houses := 0
	for _, c := range acc {
		if c > 1 {
			houses++
		}
	}
	return len(acc)
}

func Day4(part2 bool) int {
	PREFIX := "00000"
	if part2 {
		PREFIX = "000000"
	}
	buf := "ckczppom"
	// buf := "abcdef"
	n := 0
	for {
		s := fmt.Sprintf("%s%d", buf, n)
		sum := fmt.Sprintf("%x", md5.Sum([]byte(s)))
		if strings.HasPrefix(sum, PREFIX) {
			return n
		}
		n += 1
	}
	return 0
}

func Day5(part2 bool) int {
	buf := aoc.Readfile("day5.txt")
	res := 0
	if !part2 {
		// buf := []string{"ugknbfddgicrmopn", "aaa", "jchzalrnumimnmhp"}
		vowels := "aeiou"
		forbid := []string{"ab", "cd", "pq", "xy"}
		for _, s := range buf {
			// 3 or more vowels
			vowelcount := 0
			for i := 0; i < len(vowels); i++ {
				vowelcount += strings.Count(s, string(vowels[i]))
			}
			if vowelcount < 3 {
				continue
			}
			// twice in a row
			twice := false
			c := s[0]
			for i := 1; i < len(s); i++ {
				if c == s[i] {
					twice = true
					break
				}
				c = s[i]
			}
			// no forbidden
			noforbidden := true
			for _, f := range forbid {
				if strings.Contains(s, f) {
					noforbidden = false
					break
				}
			}
			if vowelcount >= 3 && twice && noforbidden {
				res++
			}
		}
	} else {
		// part 2
		// buf := []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}
		for _, s := range buf {
			haspair := false
			for i := 0; i < len(s)-1; i++ {
				pair := s[i : i+2]
				if strings.Contains(s[i+2:], pair) {
					haspair = true
					break
				}
			}
			hastriplet := false
			for i := 0; i < len(s)-2; i++ {
				if s[i] == s[i+2] {
					hastriplet = true
					break
				}
			}
			if haspair && hastriplet {
				res++
			}
		}
	}
	return res
}

func Day6(part2 bool) int {
	buf := aoc.Readfile("day6.txt")
	var arr [1000][1000]int8
	for _, s := range buf {
		var x1, y1, x2, y2 int
		re, _ := regexp.Compile(`(turn on|turn off|toggle) ([0-9]+,[0-9]+ through [0-9]+,[0-9]+)`)
		m := re.FindStringSubmatch(s)
		fmt.Sscanf(m[2], "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
		if !part2 {
			switch m[1] {
			case "toggle":
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						arr[i][j] = 1 - arr[i][j]
					}
				}
			case "turn on":
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						arr[i][j] = 1
					}
				}
			case "turn off":
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						arr[i][j] = 0
					}
				}
			}
		} else {
			// part2
			switch m[1] {
			case "turn on":
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						arr[i][j] += 1
					}
				}
			case "turn off":
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						var x int8
						if x = arr[i][j] - 1; x < 0 {
							x = 0
						}
						arr[i][j] = x
					}
				}
			case "toggle":
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						arr[i][j] += 2
					}
				}
			}
		}
	}
	sum := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			sum += int(arr[i][j])
		}
	}
	return sum
}

func Day7new(part2 bool) int {
	type Node struct {
		input    []string
		output   []string
		operator []string
		value    uint16
		text     string
	}

	var valuenodes []string // list of nodes with value given
	buf := aoc.Readfile(`day7.txt`)
	var graph map[string]Node
	re := regexp.MustCompile(`([a-z]+|[0-9]+)* *([A-Z]+)* *([a-z0-9]+)`)

	updateNode := func(dest string, new Node) {
		// abort if dest is an immed value
		_, err := strconv.Atoi(dest)
		if err == nil {
			return
		}
		if node, ok := graph[dest]; ok {
			// update existing
			if new.input != nil {
				if node.input != nil {
					fmt.Println("overwriting existing input!")
				}
				node.input = new.input
			}
			if new.output != nil {
				node.output = append(node.output, new.output...)
			}
			if new.operator != nil {
				node.operator = new.operator
			}
			if new.value != 0 {
				node.value = new.value
			}
			if new.text != "" {
				node.text = new.text
			}
			graph[dest] = node
		} else {
			// create and add new node
			graph[dest] = new
		}
	}

	buildgraph := func(buf []string) map[string]Node {
		graph = make(map[string]Node)
		for _, s := range buf {
			tok := strings.Split(s, " -> ")
			m := re.FindStringSubmatch(tok[0])
			w2 := m[3]
			gate := m[2]
			w1 := m[1]
			dest := tok[1]
			switch gate {
			case "AND":
				fallthrough
			case "OR":
				updateNode(dest, Node{[]string{w1, w2}, nil, []string{gate}, 0, s})
				updateNode(w1, Node{nil, []string{dest}, nil, 0, ""})
				updateNode(w2, Node{nil, []string{dest}, nil, 0, ""})
			case "LSHIFT":
				fallthrough
			case "RSHIFT":
				updateNode(dest, Node{[]string{w1}, nil, []string{gate, w2}, 0, s})
				updateNode(w1, Node{nil, []string{dest}, nil, 0, ""})
			case "NOT":
				updateNode(dest, Node{[]string{w2}, nil, []string{gate}, 0, s})
				updateNode(w2, Node{nil, []string{dest}, nil, 0, ""})
			default:
				// whole value?
				n, err := strconv.Atoi(tok[0])
				if err == nil {
					// fmt.Println("value", n)
					updateNode(dest, Node{nil, nil, []string{"PASS"}, uint16(n), s})
					valuenodes = append(valuenodes, dest)
				} else {
					w2 = tok[0]
					// fmt.Println("pass", w2, s)
					updateNode(dest, Node{[]string{w2}, nil, []string{"PASS"}, 0, s})
					updateNode(w2, Node{nil, []string{dest}, nil, 0, ""})
				}
			}
		}
		return graph
	}

	// return the value of the given Node
	var recurse func(string) uint16
	recurse = func(n string) uint16 {
		node := graph[n]
		if node.value == 0 {
			switch node.operator[0] {
			case "NOT":
				node.value = recurse(node.input[0]) ^ 0xffff
			case "PASS":
				inp := node.input
				if inp != nil {
					node.value = recurse(inp[0])
				}
			case "LSHIFT":
				sw, _ := strconv.Atoi(node.operator[1])
				node.value = recurse(node.input[0]) << sw
			case "RSHIFT":
				sw, _ := strconv.Atoi(node.operator[1])
				node.value = recurse(node.input[0]) >> sw
			case "AND":
				if len(node.input) != 2 {
					fmt.Println("AND node doesn't have 2 inputs!")
					break
				}
				inp := node.input[0]
				n, err := strconv.Atoi(inp)
				if err == nil {
					node.value = uint16(n) & recurse(node.input[1])
				} else {
					node.value = recurse(inp) & recurse(node.input[1])
				}
			case "OR":
				if len(node.input) != 2 {
					fmt.Println("OR node doesn't have 2 inputs!")
					break
				}
				inp := node.input[0]
				n, err := strconv.Atoi(inp)
				if err == nil {
					node.value = uint16(n) | recurse(node.input[1])
				} else {
					node.value = recurse(inp) | recurse(node.input[1])
				}
			}
			graph[n] = node
		}
		return node.value
	}

	/*
		buf := []string{
			"123 -> x",
			"456 -> y",
			"x AND y -> d",
			"x OR y -> e",
			"x LSHIFT 2 -> f",
			"y RSHIFT 2 -> g",
			"NOT x -> h",
			"NOT y -> i",
			"d AND e -> j",
			"x OR j -> k",
		}
	*/
	graph = buildgraph(buf)
	if part2 {
		n := graph["b"]
		n.value = 956
		graph["b"] = n
	}
	recurse("a")
	// fmt.Println(graph)
	/*
		for k, v := range graph {
			fmt.Printf("%v: %v\n", k, v.value)
		}
	*/
	return int(graph["a"].value)
}

func Day7(part2 bool) int {
	/*
		url := "https://adventofcode.com/2015/day/7/input"
		resp, err := http.Get(url)
		fmt.Println(err, resp)
	*/
	buf := aoc.Readfile(`day7.txt`)
	soln := make(map[string]int)
	dict := make(map[string][]string, len(buf))
	re := regexp.MustCompile(`([a-z]+|[0-9]+)* *([A-Z]+)* *([a-z0-9]+)`)
	for _, s := range buf {
		tok := strings.Split(s, " -> ")
		m := re.FindStringSubmatch(tok[0])
		w2 := m[3]
		gate := m[2]
		w1 := m[1]
		dest := tok[1]
		if _, ok := dict[dest]; ok {
			fmt.Println(dest, "exists")
		}
		switch gate {
		case "AND":
			fallthrough
		case "OR":
			dict[dest] = []string{gate, w1, w2}
		case "LSHIFT":
			fallthrough
		case "RSHIFT":
			dict[dest] = []string{gate, w1, w2}
		case "NOT":
			dict[dest] = []string{gate, w2}
		default:
			// whole value?
			n, err := strconv.Atoi(tok[0])
			if err == nil {
				// fmt.Println("value", n)
				soln[dest] = n
			} else {
				w2 = tok[0]
				// fmt.Println("wire", w2, s)
				dict[dest] = []string{w2}
			}
		}
	}
	// back substitute
	part2solved := false
	for {
		if a, ok := soln["a"]; ok {
			// "a" has been solved
			if !part2 {
				for k, v := range soln {
					fmt.Printf("%v: %v\n", k, v)
				}
				return a
			} else {
				if !part2solved {
					// part2
					soln["b"] = a
					delete(soln, "a")
					b := dict["b"]
					fmt.Println(b)
					delete(soln, b[1])
					delete(soln, b[2])
					part2solved = true
				} else {
					return a
				}
			}
		}
		for k, v := range dict {
			if _, ok := soln[k]; ok {
				// skip if already solved
				continue
			}
			switch v[0] {
			case "AND":
				i, ok := strconv.Atoi(v[1])
				m, ok1 := soln[v[1]]
				n, ok2 := soln[v[2]]
				if ok == nil {
					m = i
					ok1 = true
				}
				if ok1 && ok2 {
					soln[k] = m & n
					// delete(dict, k)
				}
			case "OR":
				m, ok1 := soln[v[1]]
				n, ok2 := soln[v[2]]
				if ok1 && ok2 {
					soln[k] = m | n
					// delete(dict, k)
				}
			case "LSHIFT":
				if n, ok := soln[v[1]]; ok {
					shift, _ := strconv.Atoi(v[2])
					soln[k] = n << shift
					// delete(dict, k)
				}
			case "RSHIFT":
				if n, ok := soln[v[1]]; ok {
					shift, _ := strconv.Atoi(v[2])
					soln[k] = n >> shift
					// delete(dict, k)
				}
			case "NOT":
				if n, ok := soln[v[1]]; ok {
					soln[k] = n ^ 0xffff
					// delete(dict, k)
				}
			default:
				if n, ok := soln[v[0]]; ok {
					soln[k] = n
					// delete(dict, k)
				}
			}
		}
	}
	return 0
}

func Day8(part2 bool) int {
	buf := aoc.Readfile("day8.txt")
	// buf = []string{`""`, `"abc"`, `"aaa\"aaa"`, `"\x27"`}
	// buf = []string{`"\"\\\x27"`}
	rePar, _ := regexp.Compile(`\\x[0-9a-f]{2}`)
	reEnc, _ := regexp.Compile(`\\x[0-9a-f]{2}|\\"|\\\\`)

	parse := func(s string) int {
		s = s[1 : len(s)-1]
		s = strings.Replace(s, `\\`, `.`, -1)
		s = strings.Replace(s, `\"`, `.`, -1)
		m := rePar.FindAllString(s, -1)
		return len(s) - 3*len(m)
	}

	encode := func(s string) int {
		// s = strings.Trim(s, `"`)
		s = s[1 : len(s)-1]
		m := reEnc.FindAllString(s, -1)
		count := 0
		for _, i := range m {
			switch i[:2] {
			case `\"`:
				count += 4 - 2
			case `\\`:
				count += 4 - 2
			case `\x`:
				count += 5 - 4
			}
		}
		// println(s, 6+len(s)+count)
		return 6 + len(s) + count
	}

	nliterals := 0
	nmemory := 0
	nnewstr := 0
	for _, s := range buf {
		nliterals += len(s)
		nmemory += parse(s)
		nnewstr += encode(s)
	}
	if !part2 {
		return nliterals - nmemory
	} else {
		return nnewstr - nliterals
	}
}

func Day9(part2 bool) int {
	buf := aoc.Readfile("day9.txt")
	// buf = []string{"London to Dublin = 464", "London to Belfast = 518", "Dublin to Belfast = 141"}
	dist := make(map[string]int)
	locations := make(map[string]int)
	key := func(i, j string) string {
		return i + "/" + j
	}
	parse := func(s string) (string, string, int) {
		m := strings.Fields(s)
		n, _ := strconv.Atoi(m[4])
		locations[m[0]]++
		locations[m[2]]++
		dist[key(m[0], m[2])] = n
		dist[key(m[2], m[0])] = n
		return m[0], m[2], n
	}
	// parse into dist and locations
	for _, s := range buf {
		parse(s)
	}
	var loc []string
	for x := range locations {
		loc = append(loc, x)
	}
	// sort.Strings(loc)

	routedist := func(route []int) int {
		totaldist := 0
		for i := 0; i < len(route)-1; i++ {
			k := key(loc[route[i]], loc[route[i+1]])
			// println(k)
			totaldist += dist[k]
		}
		return totaldist
	}

	indices := make([]int, len(loc))
	for i := 0; i < len(indices); i++ {
		indices[i] = i
	}
	// locperms := permute(nil, indices)
	locperms := aoc.Permute2(indices)
	// check no duplicates
	// checkdup(locperms)
	mindist := int(1e6)
	maxdist := 0
	for _, route := range locperms {
		d := routedist(route)
		if part2 {
			if d > maxdist {
				maxdist = d
				// fmt.Println("max", maxdist, route)
			}
		} else {
			if d < mindist {
				mindist = d
				// fmt.Println("min", mindist, route)
			}
		}
	}
	if !part2 {
		return mindist
	} else {
		return maxdist
	}
}

func Day10(part2 bool) int {
	buf := "1321131112"
	// buf = "1"
	runlen := func(s string) string {
		c := s[0]
		var res strings.Builder
		var i, j int
		for i = range s {
			if s[i] != c {
				res.WriteString(fmt.Sprintf("%d%c", i-j, c))
				j = i
				c = s[i]
			}
		}
		res.WriteString(fmt.Sprintf("%d%c", i+1-j, c))
		// fmt.Println(res.String())
		return res.String()
	}
	t := buf
	var iterations int
	if !part2 {
		iterations = 40
	} else {
		iterations = 50
	}
	for i := 0; i < iterations; i++ {
		t = runlen(t)
	}
	return len(t)
}

func Day11(part2 bool) int {
	buf := "hxbxwxba"
	checkstr := func(s string) bool {
		d := make([]byte, len(s)-1)
		// diff and find a 11 sequence
		for i := 0; i < len(buf)-1; i++ {
			d[i] = buf[i+1] - buf[i]
			if i > 0 && d[i-1] == 1 && d[i] == 1 {
				return true
			}
		}
		return false
	}
	checkIOL := func(s string) bool {
		return !strings.ContainsAny(s, "iol")
	}
	checkpairs := func(s string) bool {
		// 2 or more different pairs
		var p []byte
		count := 0
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				p = append(p, s[i])
				count++
				i++
			}
		}
		return count > 1 && p[0] != p[1]
	}
	nextpasswd := func(s string) string {
		incr := func(c *byte) bool {
			*c++
			if *c > 'z' {
				*c = 'a'
				return true
			}
			if *c == 'i' || *c == 'o' || *c == 'l' {
				*c++
			}
			return false
		}
		res := make([]byte, len(s))
		copy(res, s)
		l := len(s) - 1
		for incr(&res[l]) {
			l--
			if l < 0 {
				break
			}
		}
		// fmt.Println(s, string(res))
		return string(res)
	}
	// buf = "hijklmmn"
	// buf = "abbceffg"
	// buf = "abbcegjk"
	// buf = "abcdffaa"
	// buf = "ghjaabcc"
	// println(buf, checkstr(buf), checkIOL(buf), checkpairs(buf))
	// buf = "abcdefgh"
	// buf = "ghijklmn"
	parts := 0
	for true {
		buf = nextpasswd(buf)
		// println(buf)
		if checkIOL(buf) && checkstr(buf) && checkpairs(buf) {
			if !part2 {
				fmt.Println("found", buf)
				return 0
			} else {
				parts++
				if parts == 2 {
					fmt.Println("found", buf)
					return 0
				}
			}
		}
	}
	return -1
}

func Day12(part2 bool) int {
	buf := aoc.Readfile("day12.txt")
	/*
		buf := []string{"[]", "{}", "[1,2,3]",
			`{"a":2,"b":4}`,
			"[[[3]]]", `{"a":{"b":4},"c":-1}`,
			`{"a":[-1,1]}`, `[-1,{"a":1}]`,
			`[1,{"c":"red","b":2},3]`,
			`{"d":"red","e":[1,2,3,4],"f":5}`,
			`[1,"red",5]`}
	*/
	var rsum func(interface{}) int
	rsum = func(m interface{}) int {
		res := 0
		switch m.(type) {
		case []interface{}:
			m := m.([]interface{})
			for _, i := range m {
				switch i.(type) {
				case float64:
					res += int(i.(float64))
				case string:
					if i.(string) == "red" {
						// println("red is ok")
					}
				default:
					res += rsum(i)
				}
			}
		case map[string]interface{}:
			m := m.(map[string]interface{})
			for _, i := range m {
				switch i.(type) {
				case float64:
					res += int(i.(float64))
				case string:
					if part2 && i.(string) == "red" {
						// println("red")
						return 0 //, true
					}
				default:
					res += rsum(i)
				}
			}
		}
		return res
	}
	var m interface{}
	res := 0
	for _, buf := range buf {
		json.Unmarshal([]byte(buf), &m)
		// fmt.Printf("%T %v ", m, m)
		// fmt.Println(rsum(m))
		res += rsum(m)
	}
	return res
}

func Day13(part2 bool) int {
	buf := aoc.Readfile("day13.txt")
	// 	buf = strings.Split(`Alice would gain 54 happiness units by sitting next to Bob.
	// Alice would lose 79 happiness units by sitting next to Carol.
	// Alice would lose 2 happiness units by sitting next to David.
	// Bob would gain 83 happiness units by sitting next to Alice.
	// Bob would lose 7 happiness units by sitting next to Carol.
	// Bob would lose 63 happiness units by sitting next to David.
	// Carol would lose 62 happiness units by sitting next to Alice.
	// Carol would gain 60 happiness units by sitting next to Bob.
	// Carol would gain 55 happiness units by sitting next to David.
	// David would gain 46 happiness units by sitting next to Alice.
	// David would lose 7 happiness units by sitting next to Bob.
	// David would gain 41 happiness units by sitting next to Carol.`, "\n")
	M := make(map[string]map[string]int)
	// process input
	for _, s := range buf {
		s := strings.Fields(strings.TrimSuffix(s, "."))
		who, sign, partner := s[0], s[2], s[10]
		bias, _ := strconv.Atoi(s[3])
		if sign == "lose" {
			bias = -bias
		}
		// key := who + partner
		if _, ok := M[who]; !ok {
			M[who] = map[string]int{partner: bias}
		} else {
			M[who][partner] = bias
		}
	}
	// unique persons
	var names []string
	for n := range M {
		names = append(names, n)
	}
	// fmt.Println(M, names)
	l := len(names)
	if part2 {
		l++ // myself
		names = append(names, "myself")
	}
	// make permutations
	n := make([]int, l)
	for i := range n {
		n[i] = i
	}
	bestscore := int(-1e6)
	// for _, p := range permute(nil, n) {
	for _, p := range aoc.Permute2(n) {
		score := 0
		// compute each diner's score
		for i := range names {
			who := names[p[i]]
			left := names[p[(i+1)%l]]
			right := names[p[(i+l-1)%l]]
			score += M[who][left] + M[who][right]
		}
		if score > bestscore {
			bestscore = score
		}
	}
	return bestscore
}

func Day14(part2 bool) int {
	buf := aoc.Readfile("day14.txt")
	// buf = strings.Split(`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	// Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`, "\n")
	// parse
	if !part2 {
		maxdist := int(-1e6)
		for _, s := range buf {
			s := strings.Fields(s)
			// who, _ := strconv.Atoi(s[0])
			spd, _ := strconv.Atoi(s[3])
			dur, _ := strconv.Atoi(s[6])
			rest, _ := strconv.Atoi(s[13])
			mul := 2503 / (dur + rest)
			frac := 2503 % (dur + rest)
			dist := mul * dur * spd
			if frac < dur {
				dist += frac * spd
			} else {
				dist += dur * spd
			}
			if dist > maxdist {
				maxdist = dist
			}
		}
		// after 2503 sec
		return maxdist
	} else {
		// part2
		type reindeer struct {
			speed, dur, rest, dist int
			isflying               bool
			tdur, trest, score     int
		}
		max, imax := 0, -1
		reindeers := []*reindeer{}
		for _, s := range buf {
			s := strings.Fields(s)
			// who := s[0]
			spd, _ := strconv.Atoi(s[3])
			dur, _ := strconv.Atoi(s[6])
			rest, _ := strconv.Atoi(s[13])
			reindeers = append(reindeers, &reindeer{spd, dur, rest, 0, true, 0, 0, 0})
		}
		for time := 0; time < 2503; time++ {
			// update dist
			for _, r := range reindeers {
				if r.isflying {
					if r.tdur < r.dur {
						r.tdur++
						r.dist += r.speed
					} else {
						r.isflying = false
						r.tdur = 0
						r.trest++
					}
				} else {
					if r.trest < r.rest {
						r.trest++
					} else {
						r.isflying = true
						r.trest = 0
						r.tdur++
						r.dist += r.speed
					}
				}
			}
			// who is leading?
			max, imax = -1, -1
			for i := range reindeers {
				if reindeers[i].dist > max {
					max = reindeers[i].dist
					imax = i
				}
			}
			reindeers[imax].score++
		}
		max, imax = 0, -1
		for i, r := range reindeers {
			if r.score > max {
				max = r.score
				imax = i
			}
		}
		return reindeers[imax].score
	}
}

func Day15(part2 bool) int {
	buf := aoc.Readfile("day15.txt")
	buf = strings.Split(`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`, "\n")
	type ing struct {
		name                                            string
		capacity, durability, flavor, texture, calories int
	}
	ingredients := []*ing{}
	for _, s := range buf {
		ss := strings.ReplaceAll(s, ",", "")
		s := strings.Fields(ss)
		cap, _ := strconv.Atoi(s[2])
		dur, _ := strconv.Atoi(s[4])
		fla, _ := strconv.Atoi(s[6])
		tex, _ := strconv.Atoi(s[8])
		cal, _ := strconv.Atoi(s[10])
		ingredients = append(ingredients, &ing{
			strings.Trim(s[0], ":"), cap, dur, fla, tex, cal})
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	// returns ing, score
	compute := func(quant []int) (ing, int) {
		total := ing{name: "total"}
		for n, i := range ingredients {
			q := quant[n]
			total.capacity += q * i.capacity
			total.durability += q * i.durability
			total.flavor += q * i.flavor
			total.texture += q * i.texture
			total.calories += q * i.calories
		}
		score := max(total.capacity, 0) * max(total.durability, 0) * max(total.flavor, 0) * max(total.texture, 0)
		return total, score
	}

	best := 0
	// quant := [...]int{44, 56}
	quant := make([]int, len(ingredients))
	var r func([]int, int, int) int
	r = func(a []int, j int, sum int) int {
		var score int
		if j < len(a)-1 {
			for i := 0; i <= 100-sum; i++ {
				a[j] = i
				score = r(a, j+1, sum+i)
			}
		} else {
			a[j] = 100 - sum
			var total ing
			total, score = compute(a)
			if score > best {
				if !part2 || (part2 && total.calories == 500) {
					// fmt.Println("best", a, score)
					best = score
				}
			}
		}
		return best
	}
	best = r(quant, 0, 0)
	return best
}

func Day16(part2 bool) int {
	buf := aoc.Readfile("day16.txt")
	// the real aunt
	/*
		type aunt struct {
			children, cats, samoyeds, pomeranians, akitas int
			vizslas, goldfish, trees, cars, perfumes      int
		}
		sue := aunt{3, 7, 2, 3, 0, 0, 5, 3, 2, 1}
	*/
	type aunt struct {
		id  int
		m   map[string]int
		err int
	}
	sue := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1}
	var candidates []aunt
	cand := -1
	for _, a := range buf {
		aa := strings.ReplaceAll(a, ",", "")
		a := strings.Split(strings.ReplaceAll(aa, ":", ""), " ")
		id, _ := strconv.Atoi(a[1])
		m := make(map[string]int)
		for i := 2; i < len(a); i += 2 {
			n, _ := strconv.Atoi(a[i+1])
			m[a[i]] = n
		}
		err := 0
		for k, v := range m {
			if part2 {
				if k == "cats" || k == "trees" {
					if sue[k] < v {
						continue
					}
				}
				if k == "pomeranians" || k == "goldfish" {
					if sue[k] > v {
						continue
					}
				}
				if sue[k] == v {
					continue
				}
				err++
			} else {
				// part1
				d := sue[k] - v
				err += d * d // squared error
			}
		}
		candidates = append(candidates, aunt{id, m, err})
		if err == 0 {
			fmt.Println(id, m)
			cand = id
		}
	}
	return cand
}

func Day17(part2 bool) int {
	buf := aoc.Readfile("day17.txt")
	// buf = []string{"20", "15", "10", "5", "5"}
	containers := []int{}
	for _, s := range buf {
		n, _ := strconv.Atoi(s)
		containers = append(containers, n)
	}
	p := make([]int, len(containers))
	for i := range p {
		p[i] = i
	}
	mincontainers := make(map[int]int)
	score := 0
	for n := 0; n < len(containers); n++ {
		for _, p := range aoc.Combinate(p, n) {
			sum := 0
			for _, i := range p {
				sum += containers[i]
			}
			// if sum == 25 {
			if sum == 150 {
				// fmt.Println(p)
				score++
				mincontainers[n]++
			}
		}
	}
	if !part2 {
		return score
	} else {
		kmin := 999
		for k := range mincontainers {
			if k < kmin {
				kmin = k
			}
		}
		return mincontainers[kmin]
	}
}

func Day18(part2 bool) int {
	buf := aoc.Readfile("day18.txt")
	// buf = []string{".#.#.#",
	// 	"...##.",
	// 	"#....#",
	// 	"..#...",
	// 	"#.#..#",
	// 	"####.."}
	dim := len(buf)
	grid := make([][]bool, dim+2)
	grid2 := make([][]bool, dim+2)
	// top n bottom blank rows
	b := make([]bool, dim+2)
	grid[0] = b
	grid2[0] = b
	grid[dim+1] = b
	grid2[dim+1] = b
	// initialise
	for j, s := range buf {
		b := make([]bool, dim+2)
		grid2[j+1] = b
		b = make([]bool, dim+2)
		for i := range s {
			switch s[i] {
			case '#':
				b[i+1] = true
			case '.':
				b[i+1] = false
			}
		}
		grid[j+1] = b
	}
	litecorners := func(grid [][]bool) {
		grid[1][1] = true
		grid[dim][1] = true
		grid[1][dim] = true
		grid[dim][dim] = true
	}
	if part2 {
		litecorners(grid)
	}
	// turn a block into a linear slice
	serialise := func(grid [][]bool, up, down, left, right int) []bool {
		size := (down - up + 1) * (right - left + 1)
		buf := make([]bool, size)
		i := 0
		for row := up; row < down+1; row++ {
			for col := left; col < right+1; col++ {
				buf[i] = grid[row][col]
				i++
			}
		}
		return buf
	}
	litcount := func(grid [][]bool) (sum int) {
		for _, x := range serialise(grid, 1, dim, 1, dim) {
			if x {
				sum++
			}
		}
		return
	}
	surrcount := func(row, col int, grid [][]bool) (sum int) {
		x := serialise(grid, row-1, row+1, col-1, col+1)
		x[4] = false
		for _, x := range x {
			if x {
				sum++
			}
		}
		return
	}
	update := func(row, col int, from, to [][]bool) (res bool) {
		c := surrcount(row+1, col+1, from)
		// fmt.Printf("%d,%d %d\n", row, col, c)
		res = from[row+1][col+1]
		if res {
			if c < 2 || c > 3 {
				res = false
			}
		} else {
			if c == 3 {
				res = true
			}
		}
		to[row+1][col+1] = res
		return
	}
	toString := func(grid [][]bool) {
		for r := 1; r < dim+1; r++ {
			x := []byte{}
			for c := 1; c < dim+1; c++ {
				z := byte('.')
				if grid[r][c] {
					z = '#'
				}
				x = append(x, z)
			}
			fmt.Println(string(x))
		}
		fmt.Println()
	}
	// do work
	g1 := grid
	g2 := grid2
	for i := 0; i < 100; i++ {
		for r := 0; r < dim; r++ {
			for c := 0; c < dim; c++ {
				update(r, c, g1, g2)
			}
		}
		g1, g2 = g2, g1
		if part2 {
			litecorners(g1)
		}
		if false {
			toString(g1)
		}
	}
	return litcount(grid)
}

// amazing A* search
func Day19(part2 bool) int {
	buf := aoc.Readfile("day19.txt")
	// 	buf = strings.Split(`e => H
	// e => O
	// H => HO
	// H => OH
	// O => HH

	// HOH`, "\n")
	rules := [][]string{}
	revrules := [][]string{}
	var medicine string
	for _, s := range buf {
		s := strings.Fields(s)
		if len(s) == 0 {
			continue
		}
		if len(s) == 1 {
			medicine = s[0]
		} else {
			rules = append(rules, []string{s[0], s[2]})
			revrules = append(revrules, []string{s[2], s[0]})
		}
	}
	if !part2 {
		res := make(map[string]bool)
		for _, r := range rules {
			re, _ := regexp.Compile(r[0])
			m := re.FindAllStringSubmatchIndex(medicine, -1)
			// fmt.Println("rule", r[0], r[1])
			for _, i := range m {
				x := medicine[:i[0]] + r[1] + medicine[i[1]:]
				// fmt.Println(x)
				res[x] = true
			}
		}
		return len(res)
	} else {
		// part2
		type Node struct {
			str   string
			depth int
			dist  int
			prev  *Node
		}
		// returns all unique single substitutions of string s
		substitutions := func(n Node) []Node {
			nodeset := make(map[string]bool)
			s := n.str
			for _, r := range revrules {
				re, _ := regexp.Compile(r[0])
				m := re.FindAllStringSubmatchIndex(s, -1)
				// fmt.Printf("subs %s -> %s %s (%d)\n", s, r[0], r[1], len(m))
				for _, i := range m {
					x := s[:i[0]] + r[1] + s[i[1]:]
					// if len(x) <= len(medicine) {
					nodeset[x] = true
					// }
				}
			}
			res := make([]Node, len(nodeset))
			i := 0
			for k := range nodeset {
				res[i] = Node{k, n.depth + 1, len(k), &n}
				i++
			}
			return res
		}
		// A* search B-)
		frontier := []Node{Node{medicine, 0, 0, nil}}
		visited := []string{}
		isVisited := func(s string, v []string) bool {
			for _, x := range v {
				if x == s {
					return true
				}
			}
			return false
		}
		backtrace := func(n *Node) []string {
			chain := []string{}
			for n != nil {
				chain = append(chain, n.str)
				n = n.prev
			}
			return chain
		}
		for len(frontier) > 0 {
			// fmt.Println("frontier", len(frontier), "visited", len(visited))
			// sort frontier, pick the node with shortest dist
			sort.Slice(frontier, func(i, j int) bool {
				return frontier[i].dist < frontier[j].dist
			})
			n := frontier[0]
			frontier = frontier[1:]
			// fmt.Println(n)
			if isVisited(n.str, visited) {
				continue
			}
			visited = append(visited, n.str)
			if n.str == "e" {
				backtrace(&n)
				return n.depth
			}
			// enqueue all substitutions
			subs := substitutions(n)
			frontier = append(frontier, subs...)
		}
		return 0
	}
}

func Day20(part2 bool) int {
	primes := aoc.Sieve(10000)
	memo := make(map[int][]int)
	factorise := func(estimate int) []int {
		n := estimate
		f := make(map[int]bool)
		f[n] = true
		g := []int{n}
		gp := 0
		for gp < len(g) {
			n = g[gp]
			gp++
			// use memoisation
			if factors, ok := memo[n]; ok {
				g = append(g, factors...)
			} else {
				factors := []int{}
				for _, p := range primes {
					if n%p == 0 {
						q := n / p
						f[q] = true
						factors = append(factors, q)
					}
				}
				g = append(g, factors...)
				memo[n] = factors
			}
		}
		res := []int{}
		for i := range f {
			if part2 {
				if estimate < i*50 {
					res = append(res, i)
				}
			} else {
				res = append(res, i)
			}
		}
		return res
	}
	// given a list of factors, return sum of presents delivered
	houseprez := func(f []int) int {
		var perhouse int
		if !part2 {
			perhouse = 10
		} else {
			perhouse = 11
		}

		sum := 0
		for i := range f {
			sum += f[i]
		}
		return sum * perhouse
	}
	threshold := 33100000
	// answer : 776160 | 786240
	estimate := 776100
	search := func(min, max int) int {
		estimate := min
		for estimate < max {
			h := houseprez(factorise(estimate))
			// println(estimate, h, "\r")
			if h > threshold {
				println(estimate)
				return estimate
			} else {
				estimate++
			}
		}
		return 0
	}
	return search(estimate, estimate+100000)
}

func Day21(part2 bool) int {
	buf := aoc.Readstring(`Hit Points: 12
Damage: 7
Armor: 2`)
	buf = aoc.Readfile("day21.txt")
	type player struct {
		hp, damage, armor int
	}
	makeplayer := func(buf []string) player {
		hp, dmg, arm := 0, 0, 0
		fmt.Sscanf(buf[0], "Hit Points: %d", &hp)
		fmt.Sscanf(buf[1], "Damage: %d", &dmg)
		fmt.Sscanf(buf[2], "Armor: %d", &arm)
		return player{hp, dmg, arm}
	}
	// returns true if me wins
	play := func(me, boss player) bool {
		for {
			boss.hp -= aoc.Max(1, me.damage-boss.armor)
			if boss.hp <= 0 {
				return true
			}
			me.hp -= aoc.Max(1, boss.damage-me.armor)
			if me.hp <= 0 {
				return false
			}
		}
	}
	type item struct {
		name, class         string
		cost, damage, armor int
	}
	itemstats := strings.Split(`
Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`, "\n")
	getitems := func(buf []string) []item {
		items := []item{}
		class := ""
		for _, s := range itemstats {
			if strings.Contains(s, ":") {
				class = strings.Split(s, ":")[0]
				continue
			}
			if len(s) < 20 {
				continue
			}
			name := strings.TrimSpace(s[:12])
			cost, _ := strconv.Atoi(strings.TrimSpace(s[12:15]))
			dmge, _ := strconv.Atoi(strings.TrimSpace(s[17:22]))
			armr, _ := strconv.Atoi(strings.TrimSpace(s[25:]))
			items = append(items, item{name, class, cost, dmge, armr})
		}
		return items
	}

	items := getitems(itemstats)
	boss := makeplayer(buf)
	me := player{100, 0, 0}
	// test
	// boss = player{12, 7, 2}
	// me = player{8, 5, 5}

	nextweapon := func(items []item) func() item {
		w := []item{}
		for _, i := range items {
			if i.class == "Weapons" {
				w = append(w, i)
			}
		}
		i := 0
		return func() item {
			i++
			if i <= len(w) {
				return w[i-1]
			}
			return item{}
		}
	}(items)
	iterarmor := func(items []item) func() item {
		w := []item{{"none", "Armor", 0, 0, 0}}
		for _, i := range items {
			if i.class == "Armor" {
				w = append(w, i)
			}
		}
		i := 0
		return func() item {
			i++
			if i <= len(w) {
				return w[i-1]
			}
			return item{}
		}
	}
	iterring := func(items []item) func() item {
		r := []item{{"none", "Rings", 0, 0, 0}}
		r = append(r, r[0]) //duplicate
		for _, i := range items {
			if i.class == "Rings" {
				r = append(r, i)
			}
		}
		n := make([]int, len(r))
		for i := range n {
			n[i] = i
		}
		c := aoc.Combinate(n, 2)
		i := 0
		return func() item {
			i++
			if i <= len(c) {
				t1, t2 := r[c[i-1][0]], r[c[i-1][1]]
				t := item{t1.name + "/" + t2.name, "Rings",
					t1.cost + t2.cost, t1.damage + t2.damage, t1.armor + t2.armor}
				return t
			}
			return item{}
		}
	}
	// combinate items
	bestcost := 0
	if !part2 {
		bestcost = int(1e6)
	}
	for w := nextweapon(); w.class != ""; w = nextweapon() {
		nextarmor := iterarmor(items)
		for a := nextarmor(); a.class != ""; a = nextarmor() {
			nextring := iterring(items)
			for r := nextring(); r.class != ""; r = nextring() {
				cost := w.cost + a.cost + r.cost
				me.hp = 100
				me.damage = w.damage + r.damage
				me.armor = a.armor + r.armor
				// fmt.Println(w.name, a.name, r.name, cost, me)
				boss := player{boss.hp, boss.damage, boss.armor}
				if !part2 {
					if play(me, boss) && cost < bestcost {
						bestcost = cost
					}
				} else {
					if !play(me, boss) && cost > bestcost {
						bestcost = cost
					}
				}
			}
		}
	}
	return bestcost
}

func Day22(part2 bool) int {
	// buf := aoc.Readfile("day22.txt")
	return 0
}

func Scratch() int {
	return 0
}
