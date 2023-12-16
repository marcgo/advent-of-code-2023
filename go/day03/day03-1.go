package day03

import (
	_ "embed"
)

func Part1(dataB []byte) (sum int64) {
	engine := loadData(dataB)
	return engine.sumPartNumbers()
}
