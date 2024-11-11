package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type coordinate_t struct {
	x, y, z int
}

type brick_t struct {
	c0, c1 coordinate_t
}

func (brick brick_t) String() string {
	return fmt.Sprintf("%d,%d,%d~%d,%d,%d", brick.c0.x, brick.c0.y, brick.c0.z, brick.c1.x, brick.c1.y, brick.c1.z)
}

func (brick *brick_t) Fall(d int) {
	brick.c0.z -= d
	brick.c1.z -= d
}

func (brick brick_t) XExtrema() (int, int) {
	return min(brick.c0.x, brick.c1.x), max(brick.c0.x, brick.c1.x)
}

func (brick brick_t) YExtrema() (int, int) {
	return min(brick.c0.y, brick.c1.y), max(brick.c0.y, brick.c1.y)
}

func (brick brick_t) ZExtrema() (int, int) {
	return min(brick.c0.z, brick.c1.z), max(brick.c0.z, brick.c1.z)
}

type bricks_t []brick_t

func (bricks bricks_t) String() string {
	rows := []string{}
	for _, brick := range bricks {
		rows = append(rows, brick.String())
	}
	return strings.Join(rows, "\n")
}
func (bricks bricks_t) Len() int      { return len(bricks) }
func (bricks bricks_t) Swap(i, j int) { bricks[i], bricks[j] = bricks[j], bricks[i] }
func (bricks bricks_t) Less(i, j int) bool {
	iz := min(bricks[i].c0.z, bricks[i].c1.z)
	jz := min(bricks[j].c0.z, bricks[j].c1.z)
	if iz != jz {
		return iz < jz
	}

	iy := min(bricks[i].c0.y, bricks[i].c1.y)
	jy := min(bricks[j].c0.y, bricks[j].c1.y)
	if iy != jy {
		return iy < jy
	}

	ix := min(bricks[i].c0.x, bricks[i].c1.x)
	jx := min(bricks[j].c0.x, bricks[j].c1.x)
	if ix != jx {
		return ix < jx
	}
	// log.Fatalf("shouldn't be here (%d,%d,%d)~(%d,%d,%d)", ix, jx, iy, jy, iz, jz)
	return true
}

type volume_t struct {
	xmax    int
	ymax    int
	heights [][]int
}

func (bricks *bricks_t) Fall() int {
	sort.Sort(bricks)

	n_falling_bricks := 0

	xmax, ymax := 0, 0
	for _, brick := range *bricks {
		xmax = max(xmax, brick.c0.x, brick.c1.x)
		ymax = max(ymax, brick.c0.y, brick.c1.y)
	}

	volume := volume_t{xmax: xmax, ymax: ymax, heights: make([][]int, xmax+1)}
	for i := range volume.heights {
		volume.heights[i] = make([]int, ymax+1)
	}

	for i := range *bricks {
		xmin, xmax := (*bricks)[i].XExtrema()
		ymin, ymax := (*bricks)[i].YExtrema()
		zmin, _ := (*bricks)[i].ZExtrema()

		collision_height := 0

		for x := xmin; x <= xmax; x++ {
			for y := ymin; y <= ymax; y++ {
				collision_height = max(collision_height, volume.heights[x][y])
			}
		}

		fall_distance := zmin - collision_height - 1

		if fall_distance > 0 {
			n_falling_bricks++
		}

		(*bricks)[i].Fall(fall_distance)

		zmin, zmax := (*bricks)[i].ZExtrema()
		if zmin != collision_height+1 {
			log.Fatal("this is an error here")
		}
		for x := xmin; x <= xmax; x++ {
			for y := ymin; y <= ymax; y++ {
				volume.heights[x][y] = zmax
			}
		}
	}

	// fmt.Printf("volume.heights: %v\n", volume.heights)

	return n_falling_bricks
}

func (bricks *bricks_t) TestDisintegration(i int) int {
	copy_bricks := make(bricks_t, len(*bricks))
	copy(copy_bricks, *bricks)
	remaining_bricks := append(copy_bricks[:i], copy_bricks[i+1:]...)
	return remaining_bricks.Fall()
}

func (bricks *bricks_t) CountCandidateBricks() int {

	bricks.Fall()
	// sort.Sort(bricks)

	safe_bricks := 0
	fall_sum := 0
	for i := range *bricks {
		n := bricks.TestDisintegration(i)
		fall_sum += n
		if n == 0 {
			safe_bricks++
		}

	}

	return safe_bricks

}

func (bricks *bricks_t) CountFallingBricks() int {
	bricks.Fall()
	fall_sum := 0
	for i := range *bricks {
		fall_sum += bricks.TestDisintegration(i)
	}
	return fall_sum
}

func ParseDay22Input(s string) bricks_t {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`^(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)$`)

	var bricks bricks_t

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		var c0, c1 coordinate_t
		c0.x, _ = strconv.Atoi(r[1])
		c0.y, _ = strconv.Atoi(r[2])
		c0.z, _ = strconv.Atoi(r[3])
		c1.x, _ = strconv.Atoi(r[4])
		c1.y, _ = strconv.Atoi(r[5])
		c1.z, _ = strconv.Atoi(r[6])
		bricks = append(bricks, brick_t{c0: c0, c1: c1})
	}

	sort.Sort(bricks)

	return bricks
}

func main() {
	bricks := ParseDay22Input(filepath.Join("input", "input22.txt"))
	part1 := bricks.CountCandidateBricks()
	fmt.Printf("part1: %v\n", part1) // 497
	part2 := bricks.CountFallingBricks()
	fmt.Printf("part2: %v\n", part2) // 67468
}
