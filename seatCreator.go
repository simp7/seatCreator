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

	hall1 := eraser.HorizontalHallway(14)
	hall2 := eraser.VerticalHallway(15)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 9},pos.Absolute{X: 15, Y: 16})
	rect2 := eraser.Rectangle(pos.Absolute{X: 26,Y: 9},pos.Absolute{X: 36, Y: 15})
	rect3 := eraser.Rectangle(pos.Absolute{X: 27,Y: 16},pos.Absolute{X: 36, Y: 16})


	rect4 := eraser.Rectangle(pos.Absolute{X: 5,Y: 8},pos.Absolute{X: 9, Y: 8})
	rect15 := eraser.Rectangle(pos.Absolute{X: 12,Y: 8},pos.Absolute{X: 14, Y: 8})
	rect5 := eraser.Rectangle(pos.Absolute{X: 26,Y: 8},pos.Absolute{X: 29, Y: 8})
	rect16 := eraser.Rectangle(pos.Absolute{X: 32,Y: 8},pos.Absolute{X: 36, Y: 8})


	rect6 := eraser.Rectangle(pos.Absolute{X: 5,Y: 3},pos.Absolute{X: 10, Y: 3})
	rect7 := eraser.Rectangle(pos.Absolute{X: 5,Y: 4},pos.Absolute{X: 9, Y: 4})
	rect8 := eraser.Rectangle(pos.Absolute{X: 5,Y: 6},pos.Absolute{X: 5, Y: 6})
	rect9 := eraser.Rectangle(pos.Absolute{X: 5,Y: 7},pos.Absolute{X: 6, Y: 7})
	rect10 := eraser.Rectangle(pos.Absolute{X: 26,Y: 3},pos.Absolute{X: 26, Y: 7})
	rect11 := eraser.Rectangle(pos.Absolute{X: 32,Y: 3},pos.Absolute{X: 36, Y: 3})
	rect12 := eraser.Rectangle(pos.Absolute{X: 33,Y: 4},pos.Absolute{X: 36, Y: 4})
	rect13 := eraser.Rectangle(pos.Absolute{X: 36,Y: 6},pos.Absolute{X: 36, Y: 6})
	rect14 := eraser.Rectangle(pos.Absolute{X: 35,Y: 7},pos.Absolute{X: 36, Y: 7})

	rect17 := eraser.Rectangle(pos.Absolute{X: 16, Y: 3}, pos.Absolute{X: 16, Y: 7})
	rect18 := eraser.Rectangle(pos.Absolute{X: 25, Y: 3}, pos.Absolute{X: 25, Y: 6})
	rect19 := eraser.Rectangle(pos.Absolute{X: 17, Y: 3}, pos.Absolute{X: 17, Y: 5})
	rect20 := eraser.Rectangle(pos.Absolute{X: 24, Y: 3}, pos.Absolute{X: 24, Y: 4})
	rect21 := eraser.Rectangle(pos.Absolute{X: 18, Y: 3}, pos.Absolute{X: 18, Y: 3})

	integrated1 := eraser.Integrated(hall1,hall2, rect1, rect2, rect3,rect4,rect5, rect6, rect7, rect8, rect9,
								rect10,rect11, rect12, rect13, rect14, rect15, rect16, rect17,rect18,rect19,rect20,
								rect21)
	base := model.NewSeatBase(5, 3, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         32,
		YSize:         14,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	Wrect1 := eraser.Rectangle(pos.Absolute{X: 13, Y: 8}, pos.Absolute{X: 28, Y: 8})
	integrated2 := eraser.Integrated(Wrect1)
	base = model.NewSeatBase(12, 8, "휠체어석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         18,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput1)
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
		XSize:           32,
		YSize:           14,
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
