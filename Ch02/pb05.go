package ch02

import "fmt"

func FindLargestRegion(m [][]int) int {
	max := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			tmp := bfsArea(m, i, j)
			if max < tmp {
				max = tmp
			}
		}
	}

	return max
}

func bfsArea(m [][]int, i int, j int) int {
	visited := make(map[string]bool)

	qi := []int{i}
	qj := []int{j}

	// moving direction from (i,j)
	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	dist := 0
	for len(qi) > 0 && len(qj) > 0 {
		level := len(qi)

		for it := 0; it < level; it++ {

			i := qi[0]
			j := qj[0]

			// pop queue
			qi := qi[1:]
			qj := qj[1:]

			// traverse all 8 directions
			for k := 0; k < 8; k++ {
				nxtI := i + di[k]
				nxtJ := j + dj[k]

				coor := fmt.Sprintf("%d,%d", nxtI, nxtJ)

				if nxtI >= 0 && nxtI < len(m) &&
					nxtJ >= 0 && nxtJ < len(m[0]) &&
					m[nxtI][nxtJ] == 1 && !visited[coor] {
					qi = append(qi, nxtI)
					qj = append(qj, nxtJ)

					// marked as visited
					visited[coor] = true
				}
			}
		}

		dist++
	}
	return dist
}
