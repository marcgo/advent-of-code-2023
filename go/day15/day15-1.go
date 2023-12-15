package day15

import (
	_ "embed"
)

func Part1(dataB []byte) (sum int64) {
	steps := loadData(dataB)
	return steps.hash()
}
