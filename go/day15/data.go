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
	hasherT struct {
		c int64
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
		h := hasherT{}
		h.hash(s)
		sum += h.c
	}
	return sum
}

func (h *hasherT) hash(s string) {
	for _, c := range s {
		h.c += int64(c)
		h.c *= 17
		h.c %= 256
	}
}
