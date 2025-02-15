package main

import (
	"strings"
	"strconv"
	"sort"
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

var (
	debugFlg	bool
	sc		= bufio.NewScanner(os.Stdin)
	wr		= bufio.NewWriter(os.Stdout)
)

func bfs(start int, neighbors [][]int) (int, int) {
	queue := []int{start}
	visited := make([]bool, len(neighbors))
	visited[start] = true
	distance := make([]int, len(neighbors))
	farthestNode := start
	for len(queue) > 0 {
		currVertex := queue[0]
		queue = queue[1:]
		for _, next := range neighbors[currVertex] {
			if visited[next] {
				continue
			}
			visited[next] = true
			distance[next] = distance[currVertex] + 1
			queue = append(queue, next)
			if distance[next] > distance[farthestNode] {
				farthestNode = next
			}
		}
	}
	return farthestNode, distance[farthestNode]
}
func outis(sl []int) {
	r := make([]string, len(sl))
	for i, v := range sl {
		r[i] = itoa(v)
	}
	out(strings.Join(r, " "))
}
func outYN(v bool) {
	if v {
		out("Yes")
	} else {
		out("No")
	}
}
func flush() {
	e := wr.Flush()
	if e != nil {
		panic(e)
	}
}
func nftoi(decimalLen int) int {
	sc.Scan()
	s := sc.Text()
	r := 0
	minus := strings.Split(s, "-")
	isMinus := false
	if len(minus) == 2 {
		s = minus[1]
		isMinus = true
	}
	t := strings.Split(s, ".")
	i := atoi(t[0])
	r += i * pow(10, decimalLen)
	if len(t) > 1 {
		i = atoi(t[1])
		i *= pow(10, decimalLen-len(t[1]))
		r += i
	}
	if isMinus {
		return -r
	}
	return r
}
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}
func itoa(i int) string {
	return strconv.Itoa(i)
}
func btoi(b byte) int {
	return atoi(string(b))
}
func ni2() (int, int) {
	return ni(), ni()
}
func ni3() (int, int, int) {
	return ni(), ni(), ni()
}
func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}
func ni2s(n int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = ni2()
	}
	return a, b
}
func ni3s(n int) ([]int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i] = ni3()
	}
	return a, b, c
}
func ni4s(n int) ([]int, []int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i], d[i] = ni4()
	}
	return a, b, c, d
}
func ni2a(n int) [][2]int {
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1] = ni2()
	}
	return a
}
func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}
func ns() string {
	sc.Scan()
	return sc.Text()
}
func out(v ...interface{}) {
	_, e := fmt.Fprintln(wr, v...)
	if e != nil {
		panic(e)
	}
}
func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func pow2(a int) int {
	return int(math.Pow(2, float64(a)))
}
func dbg(v ...interface{}) {
	if !debugFlg {
		return
	}
	out(v...)
}
func dbgi2s(sl [][]int) {
	if !debugFlg {
		return
	}
	for _, v := range sl {
		outis(v)
	}
	out("")
}
func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
}
func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}
func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}
func getS() string {
	sc.Scan()
	return sc.Text()
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}
func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}
func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
}
func main() {
	defer flush()
	N := ni()
	c1score := make([]int, N+1)
	c2score := make([]int, N+1)
	for i := 1; i <= N; i++ {
		c, s := ni2()
		if c == 1 {
			c1score[i] = c1score[i-1] + s
			c2score[i] = c2score[i-1]
		} else {
			c1score[i] = c1score[i-1]
			c2score[i] = c2score[i-1] + s
		}
	}
	dbg("c1score: ", c1score, "c2score: ", c2score)
	answer := ""
	Q := ni()
	for i := 0; i < Q; i++ {
		l, r := ni2()
		answer += fmt.Sprintf("%d %d", c1score[r]-c1score[l-1], c2score[r]-c2score[l-1])
		if i < Q-1 {
			answer += "\n"
		}
	}
	out(answer)
}
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input.txt")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
		debugFlg = true
	}
}
