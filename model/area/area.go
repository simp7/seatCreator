package area

import (
	"fmt"
	"strings"

	"github.com/simp7/seatCreator/model"
)

type Area struct {
	Key             string
	XSize           int32
	YSize           int32
	BackgroundImage string
	Color           string
	Seats           model.Group
}

func (a Area) String() string {
	areaKey := "area_key: " + a.Key
	xSize := fmt.Sprintf("x_size: %d", a.XSize)
	ySize := fmt.Sprintf("x_size: %d", a.YSize)
	backgroundImage := "background_image: " + a.BackgroundImage
	color := "color: " + a.Color
	seats := "seats: " + a.Seats.String()
	result := []string{areaKey, xSize, ySize, backgroundImage, color, seats}
	return strings.Join(result, ",\n")
}

func (a Area) Html() string {
	return a.Seats.Html()
}
