package solutions

import (
	"fmt"
)

func GenerateGraphAndReturnNumConnectedComponents(n int64, p float64) (int, int) {
	return 0, 0
}

func Sol2() {
	n := int64(100)
	p := 0.0
	step := 0.001

	for {
		numComps, maxCompSize := GenerateGraphAndReturnNumConnectedComponents(n, p)
		fmt.Printf("p: %.3f, numComps: %d, maxCompSize: %d\n", p, numComps, maxCompSize)
		p += step

		if numComps == 1 {
			break
		}
	}
}
