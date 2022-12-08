package emptychecker

import "github.com/simp7/seatCreator/model"

type positionChecker struct {
	emptySet map[model.Pos]struct{}
}

func (p positionChecker) IsEmpty(x int, y int) bool {
	_, ok := p.emptySet[model.Pos{X: x, Y: y}]
	return ok
}

func NewPositionChecker(positions ...model.Pos) EmptyChecker {
	p := new(positionChecker)
	hashMap := make(map[model.Pos]struct{})

	for _, v := range positions {
		hashMap[v] = struct{}{}
	}

	p.emptySet = hashMap
	return p
}

func NewRectangleChecker(leftTop model.Pos, rightDown model.Pos) EmptyChecker {
	p := new(positionChecker)
	hashMap := make(map[model.Pos]struct{})

	for x := leftTop.X; x <= rightDown.X; x++ {
		for y := leftTop.Y; y <= rightDown.Y; y++ {
			hashMap[model.Pos{X: x, Y: y}] = struct{}{}
		}
	}

	p.emptySet = hashMap
	return p
}
