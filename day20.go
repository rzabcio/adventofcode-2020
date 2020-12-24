package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day20_1(filename string) int {
	fmt.Printf("")
	cam := NewCamArray(filename)
	cam.Assembly()
	fmt.Println(cam.PrintPic())
	return cam.MulCornerIds()
}

func Day20_2(filename string) int {
	fmt.Printf("")
	cam := NewCamArray(filename)
	cam.Assembly()
	//fmt.Println(cam.PrintPic())
	return cam.MulCornerIds()
}

type CamArray struct {
	tiles        []*CamTile
	pic          map[int](map[int]*CamTile)
	x_min, x_max int
	y_min, y_max int
}

func NewCamArray(filename string) CamArray {
	cam := new(CamArray)
	cam.tiles = make([]*CamTile, 0)
	tile := NewCamTile()
	r_newTile := regexp.MustCompile(`^Tile (\d*):$`)
	for line := range inputCh(filename) {
		//new tile
		if r_newTile.MatchString(line) {
			tileId, _ := strconv.Atoi(r_newTile.FindStringSubmatch(line)[1])
			tile.id = tileId
			//fmt.Println("     new tile: ", tile.id)
		} else if len(line) == 0 {
			tile.GenerateBorders()
			cam.tiles = append(cam.tiles, tile)
			//fmt.Printf("     finishing %d tile: %d\n", len(cam.tiles), tile.id)
			tile = NewCamTile()
		} else {
			tile.pics = append(tile.pics, line)
			//fmt.Printf("    adding line %d to tile: %d\n", len(tile.pics), tile.id)
		}
	}
	tile.GenerateBorders()
	cam.tiles = append(cam.tiles, tile)
	//fmt.Printf("   - finishing %d tile: %d\n", len(cam.tiles), tile.id)
	return *cam
}

func (cam *CamArray) Assembly() {
	cam.pic = make(map[int](map[int]*CamTile))
	cam.SetTile(0, 0, cam.tiles[0])
	noChange := false
	for !cam.Assembled() && !noChange {
		noChange = true
		for i := 1; i < len(cam.tiles); i++ {
			newTile := cam.tiles[i]
			if newTile.assembled {
				continue
			}
			for y, subm := range cam.pic {
				for x, tile := range subm {
					i1, i2 := cam.FindBorders(tile, newTile)
					sideCount := (2 - i2 + i1) % 4
					newTile.Rotate(sideCount)
					if i1 == 0 {
						//fmt.Printf("   + (%d,%d) %d ↑ %d (%d,%d)\n", x, y, tile.id, newTile.id, x, y-1)
						if cam.SetTile(x, y-1, newTile) {
							noChange = false
							break
						}
					} else if i1 == 1 {
						//fmt.Printf("   + (%d,%d) %d → %d (%d,%d)\n", x, y, tile.id, newTile.id, x+1, y)
						if cam.SetTile(x+1, y, newTile) {
							noChange = false
							break
						}
					} else if i1 == 2 {
						//fmt.Printf("   + (%d,%d) %d ↓ %d (%d,%d)\n", x, y, tile.id, newTile.id, x, y+1)
						if cam.SetTile(x, y+1, newTile) {
							noChange = false
							break
						}
					} else if i1 == 3 {
						//fmt.Printf("   + (%d,%d) %d ← %d (%d,%d)\n", x, y, tile.id, newTile.id, x-1, y)
						if cam.SetTile(x-1, y, newTile) {
							noChange = false
							break
						}
					}
				}
				if newTile.assembled {
					break
				}
			}
		}
	}
}

func (cam *CamArray) Assembled() bool {
	assembledCount := 0
	for _, subm := range cam.pic {
		assembledCount += len(subm)
	}
	//fmt.Printf("assembled? %t, (all tiles: %s, assembled tiles: %d, tiles to assembly: %d)\n", (assembledCount >= len(cam.tiles)), len(cam.tiles), assembledCount, FilterCamTiles(cam.tiles, func(tile *CamTile) bool { return !tile.assembled }))
	fmt.Printf("--> assembled? %t, (all tiles: %s, assembled tiles: %d)\n", (assembledCount >= len(cam.tiles)), len(cam.tiles), assembledCount)
	return assembledCount >= len(cam.tiles)
}

func (cam *CamArray) FindBorders(tile1, tile2 *CamTile) (int, int) {
	for i1, b1 := range tile1.borders {
		for i2, b2 := range tile2.borders {
			if b1 == b2 {
				return i1, i2
			}
		}
	}
	return -1, -1
}

func (cam *CamArray) SetTile(x, y int, tile *CamTile) bool {
	if _, ok := cam.pic[y]; !ok {
		cam.pic[y] = make(map[int]*CamTile)
	}
	if _, ok := cam.pic[y][x]; ok {
		fmt.Printf("!!! cannot set %d to (%2d,%2d) - already occupied by %d\n", tile.id, x, y, cam.pic[y][x].id)
		return false
	}
	cam.pic[y][x] = tile
	tile.assembled = true
	cam.y_min, cam.y_max = minMax([]int{cam.y_min, cam.y_max, y})
	cam.x_min, cam.x_max = minMax([]int{cam.x_min, cam.x_max, x})
	return true
}

func (cam *CamArray) GetTile(x, y int) *CamTile {
	if _, ok := cam.pic[y]; !ok {
		emptyTile := new(CamTile)
		emptyTile.id = 0
		return emptyTile
	}
	if _, ok := cam.pic[y][x]; !ok {
		emptyTile := new(CamTile)
		emptyTile.id = 0
		return emptyTile
	}
	tile := cam.pic[y][x]
	return tile
}

func (cam *CamArray) GetTileById(id int) *CamTile {
	for _, tile := range cam.tiles {
		if tile.id == id {
			return tile
		}
	}
	return NewCamTile()
}

func (cam *CamArray) MulCornerIds() int {
	result := 1
	result *= cam.GetTile(cam.x_min, cam.y_min).id
	result *= cam.GetTile(cam.x_max, cam.y_min).id
	result *= cam.GetTile(cam.x_min, cam.y_max).id
	result *= cam.GetTile(cam.x_max, cam.y_max).id
	return result
}

func (cam *CamArray) String() string {
	s := ""
	for _, tile := range cam.tiles {
		s += tile.String() + "\n\n"
	}
	return s
}

func (cam *CamArray) PrintPic() string {
	s := ""
	for y := cam.y_min; y <= cam.y_max; y++ {
		s += fmt.Sprintf("|%10d ", y)
	}
	s += "\n"
	for y := cam.y_min; y <= cam.y_max; y++ {
		for x := cam.x_min; x <= cam.x_max; x++ {
			s += fmt.Sprintf("%2d,%2d: %-4d ", x, y, cam.GetTile(x, y).id)
			//s += fmt.Sprintf("%5d ", cam.GetTile(x, y).id)
		}
		s += "\n"
	}
	return s
}

type CamTile struct {
	id        int
	pics      []string
	borders   []string
	assembled bool
}

func NewCamTile() *CamTile {
	tile := new(CamTile)
	tile.id = 0
	tile.assembled = false
	return tile
}

func (tile *CamTile) GenerateBorders() {
	tile.borders = make([]string, 4)
	tile.borders[0] = tile.pics[0]
	tile.borders[2] = tile.pics[len(tile.pics)-1]
	tile.borders[1], tile.borders[3] = "", ""
	for y := 0; y < len(tile.pics); y++ {
		tile.borders[3] += string(tile.pics[y][0])
		tile.borders[1] += string(tile.pics[y][len(tile.pics[y])-1])
	}
}

func (tile *CamTile) Rotate(sideCount int) {
	newPics := make([]string, 0)
	if sideCount == 1 {
		for x := 0; x <= len(tile.pics)-1; x++ {
			newPic := ""
			for y := len(tile.pics) - 1; y >= 0; y-- {
				newPic += string(tile.pics[y][x])
			}
			newPics = append(newPics, newPic)
		}
	} else if sideCount == 2 {
		for y := len(tile.pics) - 1; y >= 0; y-- {
			newPic := ""
			for x := len(tile.pics[y]) - 1; x >= 0; x-- {
				newPic += string(tile.pics[y][x])
			}
			newPics = append(newPics, newPic)
		}
	} else if sideCount == 3 {
		for x := len(tile.pics[0]) - 1; x >= 0; x-- {
			newPic := ""
			for y := 0; y <= len(tile.pics)-1; y++ {
				newPic += string(tile.pics[y][x])
			}
			newPics = append(newPics, newPic)
		}
	} else {
		return
	}
	tile.pics = newPics
	tile.GenerateBorders()
}

func (tile *CamTile) Print() string {
	s := fmt.Sprintf("Tile %d:\n", tile.id)
	s += strings.Join(tile.pics, "\n")
	return s
}

func (tile *CamTile) String() string {
	return fmt.Sprintf("%d", tile.id)
}

//TOOLS
func FilterCamTiles(arr []*CamTile, cond func(*CamTile) bool) []*CamTile {
	result := []*CamTile{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
