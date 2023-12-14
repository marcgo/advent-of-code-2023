package day06

import (
	_ "embed"
	"fmt"
	"strings"

	"AdventOfCode2023/internal/util"
)

type (
	race struct {
		time     int
		distance int
	}
	races struct {
		races []race
	}
)

//go:embed test.data
var Test []byte
var TestResult_1 = int64(288)

//go:embed input.data
var Input []byte

func Day06_1(dataB []byte) int64 {
	races := loadData(dataB)
	fmt.Printf("races: %+v\n", races)
	return -1
}

func loadData(dataB []byte) (races races) {
	lines := util.LoadLines(dataB)
	if len(lines) != 2 {
		panic("oops")
	}
	if !strings.HasPrefix(lines[0], "Time:") {
		panic("oops")
	}
	if !strings.HasPrefix(lines[1], "Distance:") {
		panic("oops")
	}

	for _, timeS := range strings.Split(lines[0], " ") {
		if timeS == "Time:" || timeS == "" {
			continue
		}
		race := race{
			time:     int(util.MustAtoi(timeS)),
			distance: 0,
		}
		races.races = append(races.races, race)
	}

	i := 0
	for _, distS := range strings.Split(lines[1], " ") {
		if distS == "Distance:" || distS == "" {
			continue
		}
		dist := util.MustAtoi(distS)
		races.races[i].distance = int(dist)
		i++
	}
	return races
}
