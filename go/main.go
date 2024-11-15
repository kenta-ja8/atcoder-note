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

	_, K := ni2()
	dbg("K: ", K)
	S := ns()
	dbg("S: ", S)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetNext := make([][26]int, len(S)+1)
	// for i := 0; i < len(S); i++ {
	// 	ss := make([]int, len(alphabet))
	// 	for _i := 0; _i < len(alphabet); _i++ {
	// 		ss[_i] = -1
	// 	}
	// 	alphabetNext[i] = ss
	// }
	for i := 0; i < 26; i++ {
		alphabetNext[len(S)][i] = -1
	}
	// 次出現位置テーブルの構築
	for i := len(S) - 1; i >= 0; i-- {
		copy(alphabetNext[i][:], alphabetNext[i+1][:])
		c := S[i] - 'a'
		alphabetNext[i][c] = i
	}
	// for index, s := range S {
	// 	s := string(s)
	// 	alphaNum := strings.Index(alphabet, s)
	// 	for ni := 0; ni < len(S); ni++ {
	// 		if ni > index {
	// 			break
	// 		}
	// 		nn := alphabetNext[ni][alphaNum]
	// 		if nn != -1 {
	// 			continue
	// 		}
	// 		alphabetNext[ni][alphaNum] = index
	// 	}
	// }

	// fmt.Printf("%+v", alphabetNext)
	// for rowIndex := 0; rowIndex < 26; rowIndex++ {
	// 	fmt.Printf("%s ", string(alphabet[rowIndex]))
	// 	for columnIndex := 0; columnIndex < len(S); columnIndex++ {
	// 		fmt.Printf("%d ", alphabetNext[columnIndex][rowIndex])
	// 	}
	// 	fmt.Println("")
	// }

	answer := ""
	for sindex := 0; sindex < len(S); sindex++ {
		for _, alpha := range alphabet {
			alphaNum := strings.Index(alphabet, string(alpha))
			nextPos := alphabetNext[sindex][alphaNum]
			if nextPos == -1 {
				continue
			}
			maxLength := nextPos + K - (len(answer) + 1) + 1
			if maxLength <= len(S) {
				sindex = nextPos
				answer += string(alpha)
				break
			}
		}
		if len(answer) == K {
			break
		}
	}

	out(answer)

	// all := make([][]int, 0, H)
	// for i := 0; i < H; i++ {
	// 	rowSlice := nis(W)
	// 	dbg("rowSlice: ", rowSlice)
	// 	all = append(all, rowSlice)
	// }
	// rowAggregate := make([]int, H)
	// columnAggregate := make([]int, W)
	// for rowIndex, row := range all {
	// 	for columnIndex, v := range row {
	//      dbg("v: ", v, "index: ", columnIndex ,"rowindex: ", rowIndex, "row: ", row)
	// 		rowAggregate[rowIndex] += v
	// 		columnAggregate[columnIndex] += v
	// 	}
	// }
	//
	// for i := 0; i < H; i++ {
	//    texts := make([]int, W)
	// 	for j := 0; j < W; j++ {
	// 		// out(rowAggregate[i] + columnAggregate[j] - all[i][j])
	//      // text += itoa(rowAggregate[i] + columnAggregate[j] - all[i][j]) + " "
	//      texts[j] = rowAggregate[i] + columnAggregate[j] - all[i][j]
	// 	}
	//    outis(texts)
	// }
}
