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

func ConcertHall1F() model.Group {
	vHall := eraser.VerticalHallway(7,27)
	hHall := eraser.HorizontalHallway(16)

	leftTopRect1 := eraser.Rectangle(pos.Absolute{X: 2, Y: 8}, pos.Absolute{X: 4, Y: 10})
	leftTopRect2 := eraser.Rectangle(pos.Absolute{X: 2, Y: 11}, pos.Absolute{X: 3, Y: 11})
	leftTopRect3 := eraser.Rectangle(pos.Absolute{X: 2, Y: 12}, pos.Absolute{X: 2, Y: 13})

	middleSpecific1 := eraser.Position(pos.Absolute{X: 8, Y: 8})
	middleSpecific2 := eraser.Position(pos.Absolute{X: 26, Y: 8})

	rightTopRect1 := eraser.Rectangle(pos.Absolute{X: 30, Y: 8}, pos.Absolute{X: 32, Y: 10})
	rightTopRect2 := eraser.Rectangle(pos.Absolute{X: 31, Y: 11}, pos.Absolute{X: 32, Y: 11})
	rightTopRect3 := eraser.Rectangle(pos.Absolute{X: 32, Y: 12}, pos.Absolute{X: 32, Y: 13})

	rightBottomRect1 := eraser.Rectangle(pos.Absolute{X: 28, Y: 17}, pos.Absolute{X: 32, Y: 17})
	rightBottomRect2 := eraser.Rectangle(pos.Absolute{X: 29, Y: 18}, pos.Absolute{X: 31, Y: 18})
	integrated := eraser.Integrated(vHall, hHall, leftTopRect1, leftTopRect2, leftTopRect3, middleSpecific1, middleSpecific2, rightTopRect1, rightTopRect2, rightTopRect3, rightBottomRect1, rightBottomRect2)

	endLineLeck1 := eraser.Rectangle(pos.Absolute{X: 5, Y: 25}, pos.Absolute{X: 9, Y: 25})
	endLineLeck2 := eraser.Rectangle(pos.Absolute{X: 11, Y: 25}, pos.Absolute{X: 23, Y: 25})
	endLineLeck3 := eraser.Rectangle(pos.Absolute{X: 25, Y: 25}, pos.Absolute{X: 27, Y: 25})
	endLineLeck4 := eraser.Rectangle(pos.Absolute{X: 29, Y: 25}, pos.Absolute{X: 31, Y: 25})
	integrated2 := eraser.Integrated(endLineLeck1, endLineLeck2, endLineLeck3, endLineLeck4)

	wheelSeatLeck1 := eraser.Rectangle(pos.Absolute{X: 2, Y: 25}, pos.Absolute{X: 4, Y: 25})
	wheelSeatLeck2 := eraser.Rectangle(pos.Absolute{X: 6, Y: 25}, pos.Absolute{X: 8, Y: 25})
	wheelSeatLeck3 := eraser.Rectangle(pos.Absolute{X: 10, Y: 25}, pos.Absolute{X: 24, Y: 25})
	wheelSeatLeck4 := eraser.Rectangle(pos.Absolute{X: 26, Y: 25}, pos.Absolute{X: 28, Y: 25})
	wheelSeatLeck5 := eraser.Rectangle(pos.Absolute{X: 31, Y: 25}, pos.Absolute{X: 32, Y: 25})
	integrated3 := eraser.Integrated(wheelSeatLeck1, wheelSeatLeck2,wheelSeatLeck3,wheelSeatLeck4,wheelSeatLeck5)

	integrated4 := eraser.Integrated()

	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(2, 8, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         31,
		YSize:         17,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(2, 25, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 17),
		XSize:         31,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(2, 25, "휠체어석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         31,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('W'),
	}

	base = model.NewSeatBase(29, 18, "휠체어석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 6, 1),
		XSize:         2,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	return group.Mixed(block1, block2, block3, block4)
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
	seats := ConcertHall1F() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           31,
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