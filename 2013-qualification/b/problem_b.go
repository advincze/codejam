package main

import (
	// "bufio"
	"flag"
	"fmt"
	// "os"
	"io/ioutil"
	"strconv"
	"strings"
)

var fname = flag.String("f", "", "lawnmower file")

func main() {
	flag.Parse()
	b, _ := ioutil.ReadFile(*fname)
	lines := strings.Split(string(b), "\n")

	// read count
	count, _ := strconv.Atoi(lines[0])
	//fmt.Println("count", count)
	// count = 4
	for i, k := 0, 1; i < count; i++ {
		arr := strings.Split(lines[k], " ")
		m, _ := strconv.Atoi(arr[0])
		n, _ := strconv.Atoi(arr[1])

		block := parseBlock(lines[k+1:k+m+1], m, n)
		maxx, maxy := calcRowMax(block)
		fmt.Printf("Case #%v: %v\n", i+1, YN(possible(block, maxx, maxy)))
		k += m + 1
	}
}

func YN(b bool) string {
	if b {
		return "YES"
	} else {
		return "NO"
	}
	panic("")
}

func parseBlock(block []string, m, n int) [][]int {
	bb := make([][]int, m)
	for i := 0; i < m; i++ {
		arr := strings.Split(block[i], " ")
		bb[i] = make([]int, n)
		for j := 0; j < n; j++ {
			bb[i][j], _ = strconv.Atoi(arr[j])
		}
	}
	return bb
}

func calcRowMax(bb [][]int) (maxx, maxy []int) {
	m := len(bb)
	n := len(bb[0])
	maxx = make([]int, m)
	maxy = make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := bb[i][j]
			if val > maxx[i] {
				maxx[i] = val
			}
			if val > maxy[j] {
				maxy[j] = val
			}
		}
	}
	return
}

func possible(bb [][]int, maxx, maxy []int) bool {
	m := len(bb)
	n := len(bb[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := bb[i][j]
			if val < maxx[i] && val < maxy[j] {
				return false
			}

		}
	}
	return true
}
