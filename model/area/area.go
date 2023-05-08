package area

import (
	"fmt"
	"strings"

	"github.com/simp7/seatCreator/model"
)

type Area struct {
	Key             string      `json:"area_key"`
	XSize           int32       `json:"x_size"`
	YSize           int32       `json:"y_size"`
	BackgroundImage string      `json:"background_image"`
	Color           string      `json:"color"`
	Seats           model.Group `json:"seats"`
}

func (a Area) String() string {
	areaKey := fmt.Sprintf("area_key: \"%s\"", a.Key)
	xSize := fmt.Sprintf("x_size: %d", a.XSize)
	ySize := fmt.Sprintf("y_size: %d", a.YSize)
	backgroundImage := fmt.Sprintf("background_image: \"%s\"", a.BackgroundImage)
	color := fmt.Sprintf("color: \"%s\"", a.Color)
	seats := fmt.Sprintf("seats: [%s]", a.Seats)
	result := []string{areaKey, xSize, ySize, backgroundImage, color, seats}
	return "{\n" + strings.Join(result, ",\n") + "\n}"
}

func (a Area) Html() string {
	return a.Seats.Html()
}
