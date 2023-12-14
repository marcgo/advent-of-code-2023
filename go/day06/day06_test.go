package day06

import (
	"testing"

	"AdventOfCode2023/internal/util"
)

func TestDay06_1(t *testing.T) {
	util.CheckI64(TestResult_1, Day06_1(Test), "day 6 1 test", t)
}
