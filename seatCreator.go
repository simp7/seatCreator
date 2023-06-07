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
		XSize:         6,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specific := eraser.Position(pos.Absolute{X: 24,Y: 5},
								pos.Absolute{X: 24,Y: 7},
								pos.Absolute{X: 24,Y: 9},
								pos.Absolute{X: 24,Y: 11},
								pos.Absolute{X: 24,Y: 13},
	)
	rect := eraser.Rectangle(pos.Absolute{X: 22,Y: 14}, pos.Absolute{X: 24, Y: 14})
	integrated2 := eraser.Integrated(specific, rect)
	base = model.NewSeatBase(10, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         15,
		YSize:         11,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(26, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         8,
		EmptyChecker:  integrated,
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
	seats := ConcertHall1F() // Put Seating Here
	target := area.Area{
		Key:             "2F",
		Seats:           seats,
		XSize:           29,
		YSize:           11,
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