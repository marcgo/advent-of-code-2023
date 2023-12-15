package day15

import (
	_ "embed"
	"strings"

	"AdventOfCode2023/internal/util"
)

//go:embed input.data
var Input []byte

type (
	stepsT struct {
		steps []string
		//h     hasherT
	}
)

func loadData(dataB []byte) (steps stepsT) {
	lines := util.LoadLines(dataB)
	for _, line := range lines {
		steps.steps = append(steps.steps, strings.Split(line, ",")...)
	}
	return steps
}

func (st *stepsT) hash() (sum int64) {
	for _, s := range st.steps {
		sum += hash(s)
	}
	return sum
}

func hash(s string) (cv int64) {
	for _, c := range s {
		cv += int64(c)
		cv *= 17
		cv %= 256
	}
	return cv
}
