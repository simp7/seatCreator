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
 	
	rect := eraser.Rectangle(pos.Absolute{X: 7,Y: 4}, pos.Absolute{X: 16, Y: 4})
	rect3 := eraser.Rectangle(pos.Absolute{X: 11,Y: 5}, pos.Absolute{X: 16, Y: 5})
	rect2 := eraser.Rectangle(pos.Absolute{X: 12,Y: 6}, pos.Absolute{X: 16, Y: 6})
	rect4 := eraser.Rectangle(pos.Absolute{X: 13,Y: 7}, pos.Absolute{X: 16, Y: 8})
	integrated := eraser.Integrated(rect, rect2,rect3,rect4)
	base := model.NewSeatBase(3, 4, "마석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	rect5 := eraser.Rectangle(pos.Absolute{X: 28, Y: 5}, pos.Absolute{X: 30, Y: 8})
	integrated2 := eraser.Integrated(rect5)
	base = model.NewSeatBase(18, 5, "바석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         19,
		YSize:         5,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}
	
	rect6 := eraser.Rectangle(pos.Absolute{X: 42, Y: 4}, pos.Absolute{X: 42, Y: 6})
	rect7 := eraser.Rectangle(pos.Absolute{X: 38, Y: 5}, pos.Absolute{X: 38, Y: 6})
	integrated3 := eraser.Integrated(rect6, rect7)
	base = model.NewSeatBase(38, 4, "사석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         4,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}



	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1, block2, block3)
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
		Key:             "2F",
		Seats:           seats,
		XSize:           37,
		YSize:           6,
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
