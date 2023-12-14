package day13

import (
	"AdventOfCode2023/internal/util"
)

type (
	fieldT struct {
		horL []string
		verL []string
	}
)

func loadData(dataB []byte) (fieldsL []fieldT) {
	lines := util.LoadLines(dataB)
	field := fieldT{}
	for _, line := range lines {
		if line == "" {
			addField(&field, &fieldsL)
			//field.verL = make([]string, len(field.horL[0]))
			//for i := range field.horL {
			//	for j := range field.horL[i] {
			//		field.verL[j] += string(field.horL[i][j])
			//	}
			//}
			//fieldsL = append(fieldsL, field)
			//field = fieldT{}
			continue
		}
		field.horL = append(field.horL, line)
	}
	addField(&field, &fieldsL)
	return fieldsL
}

func addField(field *fieldT, fieldsL *[]fieldT) {
	field.verL = make([]string, len(field.horL[0]))
	for i := range field.horL {
		for j := range field.horL[i] {
			field.verL[j] += string(field.horL[i][j])
		}
	}
	*fieldsL = append(*fieldsL, *field)
	*field = fieldT{}
}
