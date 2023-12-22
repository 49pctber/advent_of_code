package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type LutRange struct {
	Low    uint32
	High   uint32
	Offset uint32
}

func (lr *LutRange) Range() uint32 {
	return lr.High - lr.Low + 1
}

func (l *LutRange) String() string {
	return fmt.Sprintf("%d-%d: %d", l.Low, l.High, l.Offset)
}

type Lut struct {
	Ranges []LutRange
}

func (l *Lut) Print() {
	for i, lr := range l.Ranges {
		fmt.Println(i, lr.String())
	}
}

func (a Lut) Len() int           { return len(a.Ranges) }
func (a Lut) Swap(i, j int)      { a.Ranges[i], a.Ranges[j] = a.Ranges[j], a.Ranges[i] }
func (a Lut) Less(i, j int) bool { return a.Ranges[i].Low < a.Ranges[j].Low }

func (lut *Lut) Lookup(index uint32) uint32 {
	for _, lr := range lut.Ranges {
		if index > lr.High {
			continue
		} else if index <= lr.High && index >= lr.Low {
			return lr.Offset + (index - lr.Low)
		} else {
			return index
		}
	}

	panic("unexpected error in lookup table")
}

func (lut *Lut) ProcessIntervals(intervals *Intervals) Intervals {
	var new Intervals

	for _, interval := range *intervals {
		var index uint32 = interval.Low
		var remaining uint32 = interval.Range

		for remaining > 0 {
			for _, lr := range lut.Ranges {
				if index > lr.High {
					continue
				}

				ni := Interval{Low: lut.Lookup(index)}

				var high uint32
				if index <= lr.High && index >= lr.Low {
					high = lr.High
				} else {
					high = lr.Low - 1
				}

				ni.Range = high - index + 1
				if ni.Range > remaining {
					ni.Range = remaining
				}
				new = append(new, ni)
				remaining -= ni.Range
				index += ni.Range
				break
			}
		}

	}

	return new
}

type Interval struct {
	Low   uint32
	Range uint32
}

func (i *Interval) High() uint32 {
	return i.Low + i.Range - 1
}

func (i *Interval) String() string {
	return fmt.Sprintf("%d-%d (%d)", i.Low, i.High(), i.Range)
}

type Intervals []Interval

func (is *Intervals) Print() {
	for i, interval := range *is {
		fmt.Printf("%d: %s\n", i, interval.String())
	}
}

func (is *Intervals) NSeeds() uint32 {
	var sum uint32 = 0
	for _, interval := range *is {
		sum += interval.Range
	}
	return sum
}

func (is Intervals) Len() int           { return len(is) }
func (is Intervals) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is Intervals) Less(i, j int) bool { return is[i].Low < is[j].Low }

func loadSeeds() []uint32 {
	file, err := os.Open(`input\input5\seeds.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seeds []uint32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), " ") {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			seeds = append(seeds, uint32(v))
		}
	}

	return seeds
}

func loadLut(fname string) Lut {
	var lut Lut

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(\d+) (\d+) (\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())
		dest_start, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal(err)
		}
		src_start, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}
		range_length, err := strconv.Atoi(m[3])
		if err != nil {
			log.Fatal(err)
		}

		lut.Ranges = append(lut.Ranges, LutRange{Low: uint32(src_start), High: uint32(src_start + range_length - 1), Offset: uint32(dest_start)})
	}

	sort.Sort(lut)

	return lut
}

var seed_soil Lut
var soil_fertilizer Lut
var fertilizer_water Lut
var water_light Lut
var light_temp Lut
var temp_humidity Lut
var humidity_location Lut

func init() {
	seed_soil = loadLut(`input\input5\s2s.txt`)
	soil_fertilizer = loadLut(`input\input5\s2f.txt`)
	fertilizer_water = loadLut(`input\input5\f2w.txt`)
	water_light = loadLut(`input\input5\w2l.txt`)
	light_temp = loadLut(`input\input5\l2t.txt`)
	temp_humidity = loadLut(`input\input5\t2h.txt`)
	humidity_location = loadLut(`input\input5\h2l.txt`)
}

func ComputeLocation(seed uint32) uint32 {
	return humidity_location.Lookup(temp_humidity.Lookup(light_temp.Lookup(water_light.Lookup(fertilizer_water.Lookup(soil_fertilizer.Lookup(seed_soil.Lookup(seed)))))))
}

func main() {
	fmt.Println("Day 5")

	seeds := loadSeeds()

	var lowest_location uint32

	for i, seed := range seeds {
		location := ComputeLocation(seed)
		if i == 0 || location < lowest_location {
			lowest_location = location
		}
	}

	fmt.Println("Location:", lowest_location) // 251346198

	var intervals Intervals

	for i, v := range seeds {
		if i%2 == 0 {
			intervals = append(intervals, Interval{Low: v})
		} else {
			intervals[i/2].Range = v
		}
	}

	for _, lut := range []Lut{seed_soil, soil_fertilizer, fertilizer_water, water_light, light_temp, temp_humidity, humidity_location} {
		var new_intervals Intervals = lut.ProcessIntervals(&intervals)
		if intervals.NSeeds() != new_intervals.NSeeds() {
			log.Fatal("number of seeds changed")
		}
		intervals = new_intervals
	}

	sort.Sort(intervals)

	fmt.Println("Lowest:", intervals[0].Low) // 72263011
}
