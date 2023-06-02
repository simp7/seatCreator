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
	hall := eraser.VerticalHallway(13)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 6, Y: 5})
	rect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 15}, pos.Absolute{X: 7, Y: 15})
	specific1 := eraser.Position(pos.Absolute{X: 3, Y: 6})
	integrated := eraser.Integrated(hall, rect1,rect2, specific1)
	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         12,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 3,Y: 6},pos.Absolute{X: 22,Y: 14})
	rect4 := eraser.Rectangle(pos.Absolute{X: 6,Y: 4},pos.Absolute{X: 22,Y: 15})
	rect5 := eraser.Rectangle(pos.Absolute{X: 6,Y: 15},pos.Absolute{X: 7,Y: 15})
	rect6 := eraser.Rectangle(pos.Absolute{X: 3,Y: 4},pos.Absolute{X: 3,Y: 15})
	specific2 := eraser.Position(pos.Absolute{X: 4, Y: 4})
	specific3 := eraser.Position(pos.Absolute{X: 4, Y: 5})
	specific4 := eraser.Position(pos.Absolute{X: 5, Y: 4})
	integrated2 := eraser.Integrated(rect3,rect4,rect5,rect6, specific2,specific3,specific4)

	base = model.NewSeatBase(3, 4, "휠체어석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         12,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
	}


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	return group.Mixed(block1, block2)
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
		XSize:           20,
		YSize:           12,
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