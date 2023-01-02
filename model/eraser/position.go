package eraser

import (
	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/pos"
)

type position struct {
	emptySet map[pos.Absolute]struct{}
}

func (p position) IsEmpty(x int, y int) bool {
	_, ok := p.emptySet[pos.Absolute{X: x, Y: y}]
	return ok
}

func Position(positions ...pos.Absolute) model.Eraser {
	p := new(position)
	hashMap := make(map[pos.Absolute]struct{})

	for _, v := range positions {
		hashMap[v] = struct{}{}
	}

	p.emptySet = hashMap
	return p
}

func Rectangle(leftTop pos.Absolute, rightDown pos.Absolute) model.Eraser {
	p := new(position)
	hashMap := make(map[pos.Absolute]struct{})

	for x := leftTop.X; x <= rightDown.X; x++ {
		for y := leftTop.Y; y <= rightDown.Y; y++ {
			hashMap[pos.Absolute{X: x, Y: y}] = struct{}{}
		}
	}

	p.emptySet = hashMap
	return p
}
