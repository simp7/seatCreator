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

	integrated := eraser.Integrated()
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(4, 10, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 7),
		XSize:         7,
		YSize:         2,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(3, 12, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 9),
		XSize:         8,
		YSize:         7,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(4, 19, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 16),
		XSize:         7,
		YSize:         2,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(3, 21, "A석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 18),
		XSize:         8,
		YSize:         4,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specific := eraser.Position(pos.Absolute{X: 27, Y: 5},
								pos.Absolute{X: 27, Y: 7},
								pos.Absolute{X: 27, Y: 9},
								pos.Absolute{X: 27, Y: 11},
								pos.Absolute{X: 27, Y: 13},
	)
	rect1 := eraser.Rectangle(pos.Absolute{X: 28, Y: 4}, pos.Absolute{X: 28, Y: 24})
	integrated2 := eraser.Integrated(specific, rect1)
	base = model.NewSeatBase(12, 4, "B석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         17,
		YSize:         23,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}


	specific1 := eraser.Position(pos.Absolute{X: 37,Y: 4},
								 pos.Absolute{X: 37,Y: 10},
								 pos.Absolute{X: 37,Y: 11},
								 pos.Absolute{X: 37,Y: 19},
								 pos.Absolute{X: 37,Y: 20},
	)
	integrated3 := eraser.Integrated(specific1)
	base = model.NewSeatBase(30, 4, "C석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         21,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	rect2 := eraser.Rectangle(pos.Absolute{X: 5, Y: 1}, pos.Absolute{X: 6, Y: 1})
	rect3 := eraser.Rectangle(pos.Absolute{X: 34, Y: 1}, pos.Absolute{X: 35, Y: 1})
	integrated4 := eraser.Integrated(rect2, rect3)
	base = model.NewSeatBase(5, 1, "OP석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         31,
		YSize:         2,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	rect4 := eraser.Rectangle(pos.Absolute{X: 11, Y: 25}, pos.Absolute{X: 29, Y: 25})
	integrated5 := eraser.Integrated(rect4)
	base = model.NewSeatBase(7, 25, "휠체어석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         27,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	block8 := group.HorizontalBlock(blockInput8)
	block9 := group.HorizontalBlock(blockInput9)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8, block9)
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
		XSize:           34,
		YSize:           23,
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