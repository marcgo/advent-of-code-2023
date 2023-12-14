package day06

import (
	_ "embed"
)

//go:embed test-2.data
var Test_2 []byte
var TestResult_2 = int64(71503)

//go:embed input-2.data
var Input_2 []byte

func Day06_2(dataB []byte) (product int64) {
	races := loadRaces(dataB)
	product = 1
	for _, race := range races.races {
		product *= race.WinningPossibilities()
	}
	return product
}
