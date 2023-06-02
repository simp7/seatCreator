package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/area"
	"github.com/simp7/seatCreator/model/eraser"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
)
func ArtriumSmall() model.Group {
	nameFormatter := nameformatter.Standard()

	hall := eraser.VerticalHallway(5, 17)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 16}, pos.Absolute{X: 4,Y: 21})
	rect2 := eraser.Rectangle(pos.Absolute{X: 18, Y: 16}, pos.Absolute{X: 19, Y: 21})
	specific := eraser.Position(pos.Absolute{X: 4,Y: 15}, pos.Absolute{X: 18, Y: 15})

	integrated := eraser.Integrated(hall, rect1, rect2, specific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         17,
		YSize:         18,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput)
	return group.Mixed(block1)
}

func copy(target model.Group) {
	err := clipboard.WriteAll(target.String())
	if err != nil {
		fmt.Println("Error occured when copying")
		return
	}
	fmt.Println("Copied!")
}
func handler(c *gin.Context) {
	seats := ArtriumSmall() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           17,
		YSize:           18,
		BackgroundImage: "",
		Color:           "#ff9f00",
    
  
	}
	fmt.Fprintf(c.Writer, "<h1>Seats Map</h1><body>\n%s\n</body>", target.Html())
	copy(target)
}
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", handler)
	r.Run()
}