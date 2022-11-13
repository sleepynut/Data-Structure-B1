package ch02

import "fmt"

func FindLargestRegion(m [][]int) int {
	max := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			tmp := bfsArea(m, i, j)
			if max < tmp {
				fmt.Printf("(%d,%d): %d\n", i, j, max)
				max = tmp
			}
		}
	}

	return max
}

func bfsArea(m [][]int, i int, j int) int {
	type Coor struct {
		i    int
		j    int
		dist int
	}

	visited := make(map[string]bool)

	q := []Coor{
		{i: i, j: j, dist: 1},
	}

	// moving direction from (i,j)
	// up, left, right, down
	di := []int{-1, 0, 0, 1}
	dj := []int{0, -1, 1, 0}

	max := 0
	for len(q) > 0 {
		level := len(q)

		for it := 0; it < level; it++ {

			i := q[0].i
			j := q[0].j
			dist := q[0].dist

			// pop queue
			q = q[1:]

			// mark current position visited
			coor := fmt.Sprintf("%d,%d", i, j)
			if visited[coor] {
				continue
			}

			visited[coor] = true
			if dist > max {
				max = dist
			}

			// traverse all 8 directions
			for k := 0; k < 4; k++ {
				nxtI := i + di[k]
				nxtJ := j + dj[k]

				coor = fmt.Sprintf("%d,%d", nxtI, nxtJ)

				if nxtI >= 0 && nxtI < len(m) &&
					nxtJ >= 0 && nxtJ < len(m[0]) &&
					m[nxtI][nxtJ] == 1 && !visited[coor] {
					q = append(q, Coor{i: nxtI, j: nxtJ, dist: dist + 1})

				}
			}
			// fmt.Println("end")
		}

	}
	return max
}
