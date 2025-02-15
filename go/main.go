package main

import (
	"bufio"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

var debugFlg bool

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

func main() {
	defer flush()

	// 10

	N, Q := ni2()
	As := nis(N)

	querys := make([][]int, Q)
	for i := 0; i < Q; i++ {
		querys[i] = nis(3)
	}

	currentIndex := 0
	for i := 0; i < Q; i++ {
		query := querys[i]
		dbg("query: ", query[0])
		if query[0] == 1 {
			fromI := ((query[1] - 1 - currentIndex) + N) % N
			toI := ((query[2] - 1 - currentIndex) + N) % N
			dbg("original index", query[1], query[2])
			from := As[fromI]
			to := As[toI]
			dbg("ft", fromI, from, toI, to, As)
			As[fromI] = to
			As[toI] = from
			dbg("As", As[fromI], As[toI])
			dbg("2out", As, currentIndex)
		} else if query[0] == 2 {
			currentIndex = (currentIndex + 1) % N
		} else {
			targetIndex := (query[1] - 1 - currentIndex) + N
			dbg("3out", As, currentIndex)
			out(As[targetIndex%N])
		}
	}

	// dbg("c1score: ", c1score, "c2score: ", c2score)

	// answer := ""
	// Q := ni()
	// for i := 0; i < Q; i++ {
	// 	l, r := ni2()
	// 	answer += fmt.Sprintf("%d %d", c1score[r]-c1score[l-1], c2score[r]-c2score[l-1])
	// 	if i < Q-1 {
	// 		answer += "\n"
	// 	}
	// }
	//
	// out(answer)

}
