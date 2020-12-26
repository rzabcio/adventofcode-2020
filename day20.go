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
	//fmt.Println(cam.PrintNeighs())
	cam.Assemble()
	//fmt.Println(cam.PrintPic())
	return cam.MulCornerIds()
}

func Day20_2(filename string) int {
	fmt.Printf("")
	cam := NewCamArray(filename)
	cam.Assemble()
	cam.CreateImage()
	return cam.image.CountRoughWaters()
}

type CamArray struct {
	tiles        []*CamTile
	pic          map[int](map[int]*CamTile)
	image        *CamTile
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

func (cam *CamArray) Assembled() bool {
	assembledCount := 0
	for _, subm := range cam.pic {
		assembledCount += len(subm)
	}
	return assembledCount >= len(cam.tiles)
}

func (cam *CamArray) Assemble() {
	cam.pic = make(map[int](map[int]*CamTile))
	corner := cam.GetCornerTiles()[0]
	cam.AssembleTile(0, 0, Transposition{corner.id, 0, 0, 0, 0}, 1)
}

func (cam *CamArray) AssembleTile(x, y int, trans Transposition, lvl int) {
	log := strings.Repeat(" ", lvl)
	tile := cam.GetTileById(trans.Id)
	log += fmt.Sprintf("%d -> (%3d,%3d)", tile.id, x, y)
	if tile.assembled || cam.GetTile(x, y).id != 0 {
		log += fmt.Sprintf(" - already added")
		//fmt.Println(log)
		return
	}
	log += fmt.Sprintf(" -> rot: %d, flip: %d", trans.Rot, trans.Flip)
	tile.Transpose(trans)
	cam.GetMatchingTiles(tile)
	cam.SetTile(x, y, tile)
	tile.assembled = true

	log += fmt.Sprintf(" -> %v", tile.neighs)
	//fmt.Println(log)

	for _, ntrans := range tile.neighs {
		nx, ny := -1, -1
		if ntrans.Side == 0 {
			nx, ny = x, y-1
		} else if ntrans.Side == 1 {
			nx, ny = x+1, y
		} else if ntrans.Side == 2 {
			nx, ny = x, y+1
		} else if ntrans.Side == 3 {
			nx, ny = x-1, y
		}
		cam.AssembleTile(nx, ny, ntrans, lvl+1)
	}
}

func (cam *CamArray) FindTransposition(tile1, tile2 *CamTile) Transposition {
	if tile1.id == tile2.id {
		return Transposition{-1, -1, -1, -1, -1}
	}
	for i1, b1 := range tile1.borders[0] {
		for i2 := 0; i2 < len(tile2.borders[0]); i2++ {
			if b1 == tile2.borders[0][i2] {
				return Transposition{tile2.id, i1, i2, (6 - i2 + i1) % 4, 1}
			}
			if b1 == tile2.borders[1][i2] {
				return Transposition{tile2.id, i1, i2, (6 - i2 + i1) % 4, 0}
			}
		}
	}
	return Transposition{-1, -1, -1, -1, -1}
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
	for _, tile := range cam.GetCornerTiles() {
		result *= tile.id
	}
	return result
}

func (cam *CamArray) GetCornerTiles() []*CamTile {
	cornerTiles := make([]*CamTile, 0)
	for _, tile := range cam.tiles {
		matchingTiles := cam.GetMatchingTiles(tile)
		if len(matchingTiles) == 2 {
			cornerTiles = append(cornerTiles, tile)
		}
		if len(cornerTiles) == 4 {
			break
		}
	}
	return cornerTiles
}

func (cam *CamArray) GetMatchingTiles(tile1 *CamTile) []*CamTile {
	if len(tile1.neighs) == 0 {
		cam.CalculateNeighs(tile1)
	}

	matching := make([]*CamTile, 0)
	for _, trans := range tile1.neighs {
		matching = append(matching, cam.GetTileById(trans.Id))
	}
	return matching
}

func (cam *CamArray) CalculateNeighs(tile1 *CamTile) {
	tile1.neighs = make(map[int]Transposition, 0)
	for _, tile2 := range cam.tiles {
		trans := cam.FindTransposition(tile1, tile2)
		if trans.Side >= 0 {
			tile1.neighs[trans.Id] = trans
		}
		if len(tile1.neighs) == 4 {
			continue
		}
	}
}

func (cam *CamArray) CreateImage() {
	cam.image = NewCamTile()
	cam.image.pics = make([]string, 0)
	lineLen := len(cam.GetTile(0, 0).pics[0])
	for py := cam.y_min; py <= cam.y_max; py++ {
		for lineNo := 1; lineNo < len(cam.GetTile(0, py).pics)-1; lineNo++ {
			imageLine := ""
			for px := cam.x_min; px <= cam.x_max; px++ {
				imageLine += cam.GetTile(px, py).pics[lineNo][1 : lineLen-1]
				//imageLine += " "
			}
			cam.image.pics = append(cam.image.pics, imageLine)
		}
		//cam.image.pics = append(cam.image.pics, " ")
	}
}

func (cam *CamArray) PrintImage() string {
	s := strings.Repeat("=", len(cam.image.pics[0])) + "\n"
	s += cam.image.Print() + "\n"
	s += strings.Repeat("=", len(cam.image.pics[0])) + "\n"
	return s
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
	for x := cam.x_min; x <= cam.x_max; x++ {
		s += fmt.Sprintf("|%12d ", x)
	}
	s += "\n"
	for y := cam.y_min; y <= cam.y_max; y++ {
		for x := cam.x_min; x <= cam.x_max; x++ {
			s += fmt.Sprintf("%3d,%3d: %-4d ", x, y, cam.GetTile(x, y).id)
			//s += fmt.Sprintf("%5d ", cam.GetTile(x, y).id)
		}
		s += "\n"
	}
	return s
}

func (cam *CamArray) PrintNeighs() string {
	s := "\n=== tiles ===\n"
	for _, tile1 := range cam.tiles {
		s += fmt.Sprintf("     %d: %d neighs: %v\n", tile1.id, len(tile1.neighs), tile1.neighs)
	}
	return s
}

type CamTile struct {
	id        int
	pics      []string
	borders   [][]string
	assembled bool
	neighs    map[int]Transposition
}

func NewCamTile() *CamTile {
	tile := new(CamTile)
	tile.id = 0
	tile.assembled = false
	return tile
}

func (tile *CamTile) GenerateBorders() {
	tile.borders = make([][]string, 2)
	tile.borders[0] = make([]string, 4)
	tile.borders[0][0] = tile.pics[0]
	tile.borders[0][2] = ""
	for x := len(tile.pics[0]) - 1; x >= 0; x-- {
		tile.borders[0][2] += string(tile.pics[len(tile.pics)-1][x])
	}
	tile.borders[0][1], tile.borders[0][3] = "", ""
	for y := 0; y < len(tile.pics); y++ {
		tile.borders[0][1] += string(tile.pics[y][len(tile.pics[y])-1])
	}
	for y := len(tile.pics) - 1; y >= 0; y-- {
		tile.borders[0][3] += string(tile.pics[y][0])
	}
	tile.borders[1] = make([]string, 4)
	for i := 0; i < 4; i++ {
		tile.borders[1][i] = ReverseStr(tile.borders[0][i])
	}
}

func (tile *CamTile) Transpose(trans Transposition) {
	newPics := make([]string, 0)
	y_first := true
	x_asc := true
	y_asc := true

	if trans.Rot == 0 {
		y_first, x_asc, y_asc = true, true, true
		if trans.Flip == 1 && (trans.Side == 0 || trans.Side == 2) {
			x_asc = !x_asc // !x_asc
		} else if trans.Flip == 1 {
			y_asc = !y_asc // !y_asc
		}
	} else if trans.Rot == 1 {
		y_first, x_asc, y_asc = false, true, false
		if trans.Flip == 1 && (trans.Side == 0 || trans.Side == 2) {
			y_asc = !y_asc // !y_asc
		} else if trans.Flip == 1 {
			x_asc = !x_asc // !x_asc
		}
	} else if trans.Rot == 2 {
		y_first, x_asc, y_asc = true, false, false
		if trans.Flip == 1 && (trans.Side == 0 || trans.Side == 2) {
			x_asc = !x_asc // !x_asc
		} else if trans.Flip == 1 {
			y_asc = !y_asc // !y_asc
		}
	} else if trans.Rot == 3 {
		y_first, x_asc, y_asc = false, false, true
		if trans.Flip == 1 && (trans.Side == 0 || trans.Side == 2) {
			y_asc = !y_asc // !y_asc
		} else if trans.Flip == 1 {
			x_asc = !x_asc // !x_asc
		}
	}
	//fmt.Printf("trans: %v -> %t, %t, %t\n", trans, y_first, x_asc, y_asc)

	ys, ye, yd := 0, 0, 0
	xs, xe, xd := 0, 0, 0
	if x_asc {
		xs, xe, xd = 0, len(tile.pics[0]), 1
	} else {
		xs, xe, xd = len(tile.pics[0])-1, -1, -1
	}
	if y_asc {
		ys, ye, yd = 0, len(tile.pics), 1
	} else {
		ys, ye, yd = len(tile.pics)-1, -1, -1
	}
	if y_first {
		for y := ys; y != ye; y += yd {
			newPic := ""
			for x := xs; x != xe; x += xd {
				newPic += string(tile.pics[y][x])
			}
			newPics = append(newPics, newPic)
		}
	} else {
		for x := xs; x != xe; x += xd {
			newPic := ""
			for y := ys; y != ye; y += yd {
				newPic += string(tile.pics[y][x])
			}
			newPics = append(newPics, newPic)
		}
	}

	tile.pics = newPics
	tile.GenerateBorders()
	tile.neighs = make(map[int]Transposition)
}

func (tile *CamTile) CountMonsters() (int, int) {
	monster := []string{
		"..................#.",
		"#....##....##....###",
		".#..#..#..#..#..#...",
	}
	//joiner := strings.Repeat(".", len(tile.pics[0])-len(monster[0]))
	joiner := fmt.Sprintf(".{%d}", len(tile.pics[0])-len(monster[0]))
	monNorAsc := strings.Join(monster, joiner)
	monRevAsc := strings.Join(ReverseStrArr(monster), joiner)
	monHashCount := strings.Count(monNorAsc, "#")

	r_monster := []*regexp.Regexp{
		regexp.MustCompile(monNorAsc),
		regexp.MustCompile(monRevAsc),
		regexp.MustCompile(ReverseStr(monNorAsc)),
		regexp.MustCompile(ReverseStr(monRevAsc)),
	}

	res := 0
	picJoined := strings.Join(tile.pics, "")
	for _, r := range r_monster {
		start := 0
		for {
			match := r.FindStringIndex(picJoined[start:])
			if match == nil {
				break
			}
			start = match[0] + 1
			res++
		}
	}
	tile.Transpose(Transposition{0, 0, 0, 1, 0})
	picJoined = strings.Join(tile.pics, "")
	for _, r := range r_monster {
		start := 0
		for {
			match := r.FindStringIndex(picJoined[start:])
			if match == nil {
				break
			}
			start += match[0] + 1
			res++
		}
	}

	return res, monHashCount
}

func (tile *CamTile) CountRoughWaters() int {
	monCount, monHashCount := tile.CountMonsters()
	hashCount := tile.CountHashes()
	return hashCount - monCount*monHashCount
}

func (tile *CamTile) CountHashes() int {
	result := 0
	for _, pic := range tile.pics {
		result += strings.Count(pic, "#")
	}
	return result
}

func (tile *CamTile) PrintFull() string {
	s := fmt.Sprintf("==== %4d ====\n", tile.id)
	s += tile.Print() + "\n"
	s += fmt.Sprintf("     b[0]: %s\n", tile.borders[0])
	s += fmt.Sprintf("     b[1]: %s\n", tile.borders[1])
	s += fmt.Sprintf("   neighs: %v\n", tile.neighs)
	return s
}

func (tile *CamTile) Print() string {
	return strings.Join(tile.pics, "\n")
}

func (tile *CamTile) String() string {
	return fmt.Sprintf("%d", tile.id)
}

type Transposition struct {
	Id    int
	Side  int
	Side2 int
	Rot   int
	Flip  int
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

func (cam *CamArray) compareTiles(id1, id2 int) {
	tile1, tile2 := cam.GetTileById(id1), cam.GetTileById(id2)
	cam.CalculateNeighs(tile1)
	cam.CalculateNeighs(tile2)
	fmt.Println(tile1.PrintFull())
	fmt.Println(tile2.PrintFull())
	trans := cam.FindTransposition(tile1, tile2)
	fmt.Printf("--> %v\n", trans)
	tile2.Transpose(trans)
	fmt.Println(tile2.PrintFull())
}
