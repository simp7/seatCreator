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
	nameFormatter := nameformatter.Standard()

	rect := eraser.Rectangle(pos.Absolute{X: 8, Y: 10}, pos.Absolute{X: 10, Y: 11})
	specific := eraser.Position(pos.Absolute{X: 7, Y: 11})
	integrated := eraser.Integrated(rect, specific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
		Reverse:       true,
	}

	rect2 := eraser.Rectangle(pos.Absolute{X: 28, Y: 4}, pos.Absolute{X: 28, Y: 5})
	rect3 := eraser.Rectangle(pos.Absolute{X: 27,Y: 10},pos.Absolute{X: 28, Y: 10})
	position1 := eraser.Position(pos.Absolute{X: 12,Y: 10})
	position2 := eraser.Position(pos.Absolute{X: 28, Y: 11})
	integrated2 := eraser.Integrated(rect2, position1, rect3, position2)
	base = model.NewSeatBase(12, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         17,
		YSize:         8,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	rect4 := eraser.Rectangle(pos.Absolute{X: 30, Y: 10}, pos.Absolute{X: 32, Y: 11})
	position3 := eraser.Position(pos.Absolute{X: 33,Y: 11})
	integrated3 := eraser.Integrated(rect4, position3)
	base = model.NewSeatBase(30, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         8,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}


	integrated4 := eraser.Integrated()
	base = model.NewSeatBase(8, 10, "휠체어석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}

	base = model.NewSeatBase(32, 10, "휠체어석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	return group.Mixed(block1, block2, block3, block4, block5)
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
		Key:             "2F",
		Seats:           seats,
		XSize:           34,
		YSize:           8,
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