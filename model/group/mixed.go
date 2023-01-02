package group

import (
	"strings"

	"github.com/simp7/seatCreator/model"
)

type mixedGroup struct {
	groups []model.Group
}

func (g mixedGroup) String() string {
	result := make([]string, 0)
	for _, v := range g.groups {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

func (g mixedGroup) Html() string {
	result := make([]string, 0)
	for _, v := range g.groups {
		result = append(result, v.Html())
	}
	return strings.Join(result, "\n")
}

func Mixed(groups ...model.Group) model.Group {
	return mixedGroup{groups: groups}
}
