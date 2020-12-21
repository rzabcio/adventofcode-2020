package main

import (
	"fmt"
)

func Day17_1(filename string) int {
	fmt.Printf("")
	space := NewCubeSpaceFromFile(filename)
	space.NextGen()
	space.NextGen()
	space.NextGen()
	space.NextGen()
	space.NextGen()
	space.NextGen()
	return space.ActiveCount
}

func Day17_2(filename string) int {
	space := NewCubeSpaceFromFile(filename)
	fmt.Println("space: ", space)
	return 0
}

type CubeSpace struct {
	points       cubeSpacePoints
	z_min, z_max int
	y_min, y_max int
	x_min, x_max int
	ActiveCount  int
}

type cubeSpacePoints map[int](map[int](map[int]*Cube))

func NewCubeSpace() CubeSpace {
	space := new(CubeSpace)
	space.points = make(cubeSpacePoints)
	return *space
}

func NewCubeSpaceFromFile(filename string) CubeSpace {
	space := NewCubeSpace()
	z := 0
	for y, line := range inputSl(filename) {
		for x, char := range line {
			space.Set(x, y, z, string(char))
		}
	}
	return space
}

func (space *CubeSpace) Set(x, y, z int, st string) {
	if _, ok := space.points[z]; !ok {
		space.points[z] = make(map[int](map[int]*Cube))
	}
	if _, ok := space.points[z][y]; !ok {
		space.points[z][y] = make(map[int]*Cube)
	}
	if _, ok := space.points[z][y][x]; !ok {
		cube := &Cube{x, y, z, st}
		space.points[z][y][x] = cube
		if st == "#" {
			space.z_min, space.z_max = minMax([]int{space.z_min, space.z_max, z})
			space.y_min, space.y_max = minMax([]int{space.y_min, space.y_max, y})
			space.x_min, space.x_max = minMax([]int{space.x_min, space.x_max, x})
			space.ActiveCount++
		}
	} else {
		if st == "#" && space.points[z][y][x].st == "." {
			space.ActiveCount--
		} else if st == "." && space.points[z][y][x].st == "#" {
			space.ActiveCount++
		}
		space.points[z][y][x].st = st
	}
}

func (space *CubeSpace) Get(x, y, z int) Cube {
	if _, ok := space.points[z]; !ok {
		return Cube{x, y, z, "."}
	}
	if _, ok := space.points[z][y]; !ok {
		return Cube{x, y, z, "."}
	}
	if _, ok := space.points[z][y][x]; !ok {
		return Cube{x, y, z, "."}
	}
	cube := space.points[z][y][x]
	return *cube
}

//func (space *CubeSpace) CountActive() int {
//	result := 0
//	for x := space.x_min; x <= space.x_max; x++ {
//		for y := space.y_min; y <= space.y_max; y++ {
//			for z := space.z_min; z <= space.z_max; z++ {
//				if space.points[z][y][x].st == "#" {
//					result++
//				}
//			}
//		}
//	}
//	return result
//}

func (space *CubeSpace) GetNeigh(xs, ys, zs int) CubeSpace {
	neigh := NewCubeSpace()
	for x := xs - 1; x <= xs+1; x++ {
		for y := ys - 1; y <= ys+1; y++ {
			for z := zs - 1; z <= zs+1; z++ {
				//fmt.Printf("(%d,%d,%d)\n", x, y, z)
				neigh.Set(x, y, z, space.Get(x, y, z).st)
			}
		}
	}
	return neigh
}

func (space *CubeSpace) NextGen() {
	newSpace := NewCubeSpace()
	for z := space.z_min - 1; z <= space.z_max+1; z++ {
		for y := space.y_min - 1; y <= space.y_max+1; y++ {
			for x := space.x_min - 1; x <= space.x_max+1; x++ {
				cube := space.Get(x, y, z)
				neigh := space.GetNeigh(x, y, z)
				if cube.st == "#" && (neigh.ActiveCount-1 < 2 || neigh.ActiveCount-1 > 3) {
					newSpace.Set(x, y, z, ".")
				} else if cube.st == "." && neigh.ActiveCount == 3 {
					newSpace.Set(x, y, z, "#")
				} else {
					newSpace.Set(x, y, z, cube.st)
				}
			}
		}
	}
	*space = newSpace
}

func (space *CubeSpace) Print() string {
	result := ""
	for y := space.y_min; y <= space.y_max; y++ {
		for z := space.z_min; z <= space.z_max; z++ {
			for x := space.x_min; x <= space.x_max; x++ {
				result += space.Get(x, y, z).st
			}
			result += " "
		}
		result += "\n"
	}
	return result
}

type Cube struct {
	x  int
	y  int
	z  int
	st string
}
