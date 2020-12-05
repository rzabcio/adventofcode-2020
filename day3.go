package main


func Day3_1(filename string) int {
	terrain := inputSl(filename)
	return countTrees(terrain, 3, 1)
}

func Day3_2(filename string) int {
	terrain := inputSl(filename)
	slopes := [5]int{11, 31, 51, 71, 12}
	treeCountMul := 1
	for _, slope := range slopes {
		dx, dy := int(slope/10), slope%10
		treeCountMul *= countTrees(terrain, dx, dy)
	}
	return treeCountMul
}

func countTrees(terrain []string, dx, dy int) int {
	treeCount := 0
	for x, y := 0, 0; y < len(terrain); x, y = (x+dx)%len(terrain[0]), y+dy {
		if '#' == terrain[y][x] {
			treeCount++
		}
	}
	return treeCount
}
